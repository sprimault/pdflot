# Conception

**La conception fait foi.** Le code s'y conforme, pas l'inverse. Un désaccord
entre ce document et le code est un défaut du code.

## Problème

Transformer un lot de fichiers HTML — de l'ordre de 6000 — en PDF, en respectant
des règles de regroupement qui changent d'une demande à l'autre, avec calage
recto/verso, et livrer le tout empaqueté.

Le rendu unitaire HTML vers PDF est un problème résolu. Ce qui ne l'est pas,
c'est la couche au-dessus : Gotenberg est sans état, une requête vaut un
document. Le job de lot — gabarit et jeu de données vers N PDF, avec file,
reprise, progression, regroupement, empaquetage — chacun se le réécrit dans son
propre back. pdflot est cette couche.

## Modèle de données

Trois niveaux, jamais aplatis :

- **document** — 1 à N fichiers HTML formant une unité logique : une
  attestation, une facture, un bulletin.
- **output** — 1 à N documents assemblés dans un PDF, avec sa règle de calage.
- **bundle** — la livraison : les outputs empaquetés, éventuellement chiffrés,
  avec le bordereau et le rapport.

### Pourquoi trois niveaux et pas deux

Le cas courant le montre sans même faire appel au recto/verso : plusieurs
centaines d'attestations fusionnées dans un seul PDF, et un bordereau qui dit
combien de pages porte chacune — sans quoi personne, en aval, ne sait où l'une
s'arrête et où la suivante commence. La frontière entre documents reste donc
signifiante **à l'intérieur** de l'output ; elle ne disparaît pas à la fusion.

Le recto/verso durcit la démonstration. Quand l'output est calé, une attestation
de 3 pages impose une page blanche avant la suivante : la frontière ne sert plus
seulement à décrire le résultat, elle le modifie.

Cette même frontière porte les signets, la numérotation par section et le
bordereau. Une implémentation qui fusionne des PDF opaques la perd
définitivement.

### Plusieurs outputs, un seul rendu

Un même lot peut produire plusieurs outputs à partir des mêmes documents : un
PDF par destinataire pour l'envoi par courriel, et un PDF unique calé en
recto/verso et trié pour le routeur. Le rendu Chromium ne doit être payé qu'une
fois. C'est l'avantage principal sur une boucle appelant un convertisseur
unitaire.

## Rendu

### Origine des documents

Les documents sont servis par un serveur HTTP sur la boucle locale, tenu par
pdflot. Le schéma `file:` est interdit au navigateur.

Ce n'est pas un détail d'implémentation. En `file://`, c'est Chromium qui résout
les sous-ressources — `<img src>`, `<link>`, `<iframe src>`, `url()` — et le code
ne voit jamais ces requêtes : il ne lui reste que le chemin d'entrée, c'est-à-dire
pas ce qui est attaqué. Le PDF rendu, puis remis au déposant, devient alors le
canal de sortie d'un fichier local lu à l'insu du service. Servir en HTTP fait de
chaque sous-ressource une requête que le bac à sable résout, refuse, et éprouve
sans navigateur.

Un serveur par travail, sur un port éphémère, fermé avec lui. L'origine distincte
n'est pas une commodité : elle empêche un document d'atteindre un autre travail
par une requête que le bac à sable verrait comme légitime, puisqu'elle resterait
à l'intérieur de sa racine. Le navigateur devient ainsi une seconde barrière,
derrière la résolution.

HTTP/1.1 en clair. Aucun navigateur ne négocie HTTP/2 sans TLS, et la limite de
six connexions par origine qui en découle borne le chargement d'un document, pas
le nombre de travaux simultanés.

### Deux passes

Les variables de pagination — `[page]`, `[topage]`, `[section]`, `[document]` —
imposent deux passes : le nombre de pages n'est connu qu'après le premier rendu.
La passe 2 rend les en-têtes et pieds au format exact de la page et les appose.

C'est plus lourd que le gabarit natif de Chromium, mais celui-ci n'hérite ni des
styles du document ni de ses ressources. La seconde passe achète la fidélité CSS
complète.

Sans variable de pagination dans le lot, la seconde passe est évitée.

Le sommaire relève de la même mécanique : les numéros ne sont connus qu'après le
premier rendu.

### Ordre d'assemblage

L'enchaînement n'est pas libre : le calage recto/verso insère des pages, donc il
déplace tout ce qui a été numéroté avant lui.

1. Passe 1 sur les corps, une fois par document quel que soit le nombre
   d'outputs qui le reprennent.
2. Fusion des documents d'un output.
3. Insertion des pages blanches de calage, **en duplex seulement**.
4. Table des placements : quelle page appartient à quel document.
5. Passe 2, qui rend les bandeaux de cet output à partir de ces placements et les
   appose.

En duplex, les pages blanches sont numérotées comme les autres — `[page]` et
`[topage]` en tiennent compte — mais ne portent ni en-tête ni pied. La pagination
correspond alors aux feuilles remises à l'imprimeur, seule lecture qui vaille
pour qui reçoit le PDF.

### Simplex par défaut

Un output n'insère aucune page blanche tant que le manifeste ne déclare pas le
duplex. Le PDF livré est continu.

C'est le cas courant : le routeur ou l'imprimeur fait sa propre imposition, à
partir du bordereau qui donne le nombre de pages de chaque document, et selon sa
machine — plieuse, mise sous pli, format d'enveloppe. Lui remettre un PDF déjà
calé fausse son décompte ou double les blanches. Le calage est donc une
déclaration explicite, jamais un effet de bord.

Le risque de l'autre bord vaut d'être écrit, parce qu'il ne se voit qu'après
impression : un output non calé imprimé en recto/verso fait commencer un document
au verso de la dernière feuille du précédent dès que celui-ci a un nombre impair
de pages. Sur du courrier nominatif, deux destinataires se retrouvent sur la même
feuille, qui part dans une seule enveloppe. La responsabilité en revient à qui
imprime, mais un outil qui produit ce genre de document le dit plutôt que de le
laisser découvrir.

C'est la raison pour laquelle le bordereau donne le nombre de feuilles autant que
le nombre de pages, y compris en simplex : sans cette colonne, on retire les
blanches sans donner de quoi les recalculer.

Corollaire sur le rendu unique : c'est le corps qui n'est rendu qu'une fois. Le
bandeau est une propriété de l'output, et un document repris dans deux outputs a
deux bandeaux — il n'y a pas là de travail dupliqué à supprimer.

### Concaténation plutôt que fusion

Pour un document composé de plusieurs HTML, on les concatène avec un saut de
page entre chaque et on fait **un seul** rendu. Trois gains : un appel navigateur
au lieu de N, une pagination continue native, des compteurs CSS cohérents.

Le prix : collisions de CSS entre fichiers, et explosion mémoire sur les groupes
trop gros. Au-delà d'un seuil configuré, on bascule sur rendu unitaire puis
fusion.

### Signets

Le plan que Chromium sait produire est déduit des titres du HTML, pas de la
structure de documents : `generateDocumentOutline` reste à faux. Les signets sont
construits au post-traitement depuis le manifeste — information que pdflot
possède, contrairement à qui fusionne des PDF opaques.

Laisser ce paramètre à vrai ferait apporter à chaque document son propre plan, qui
viendrait concurrencer les signets dans un PDF fusionné où plus rien ne les
distingue.

### Page de garde

Ce n'est pas une option : c'est un document comme un autre dans l'output.

## Empaquetage

Deux protections distinctes et indépendantes :

- **Mot de passe d'archive** — protège le transport, pas les destinataires.
- **Chiffrement AES par PDF** — la seule protection d'un destinataire vis-à-vis
  d'un autre. C'est ce que réclament les documents nominatifs.

Les deux se déclarent au manifeste.

## Exploitation

### Échec partiel

Sur 6000 documents, il y en aura toujours quelques-uns qui cassent. Le lot ne
s'interrompt jamais : on poursuit, et le détail part dans `rapport.json` à
l'intérieur du paquet. C'est ce qui distingue un outil de production d'un jouet.

Le code de sortie porte le résultat du lot : `0` succès complet, `1` succès
partiel, `2` échec global. Deux autres valeurs ne décrivent pas un résultat de
lot et ne s'y substituent jamais — `3` quand le traitement a été interrompu, le
lot n'étant pas allé au bout et ses rendus annulés n'étant pas des échecs de
document ; `64` quand la commande est mal invoquée, auquel cas aucun lot n'a
commencé.

### Ce qui attend, et combien de temps

Un contexte de navigateur coûte trop cher pour qu'on en ouvre un par document en
attente : le parallélisme de rendu est borné, et une demande qui arrive alors que
tout est occupé attend. Ce qui se garantit n'est donc pas l'absence d'attente,
mais sa borne.

Le pool distribue par document, jamais par lot. Un travail de 6000 documents rend
six mille unités de travail qui s'entrelacent avec celles des autres, et la
demande qui patiente n'attend que la libération d'un contexte — la durée d'un
document. Distribuer par lot ferait attendre dix minutes le lot de trois déposé
derrière, même avec une file parfaitement équitable.

Le reste du chemin n'attend pas : le serveur sert chaque requête indépendamment,
et la résolution des ressources est sans état partagé. Un verrou à cet endroit
ramènerait tout le lot à un seul fil, et le défaut ne se verrait qu'à la montée en
charge.

Conséquence sur la surface : la fonction de rendu prend un document et rend un
PDF. Une fonction qui prendrait un lot rendrait l'entrelacement impossible à
ajouter ensuite sans reprendre tout ce qui l'appelle.

### Équité entre tenants

En FIFO, un lot de 6000 bloque dix minutes le client d'à côté qui en demande
trois. La file est équitable par tenant.

### Reproductibilité

Version de Chromium et polices épinglées dans l'image, reportées dans le
manifeste de sortie. Un lot rejoué six mois plus tard doit sortir au pixel près.
C'est le seul point sur lequel les besoins d'archivage légal n'ont aujourd'hui
aucune réponse propre.

## Héritage wkhtmltopdf

Le vocabulaire d'options est repris parce que des milliers d'équipes le
connaissent et cherchent à migrer. Le code ne l'est pas : wkhtmltopdf est en
LGPLv3 et écrit contre l'API QtWebKit, il n'y a pas de logique portable en
dessous.

Le cas unitaire reste le comportement de base : un lot déposé sans manifeste sort
comme il serait sorti d'une boucle sur wkhtmltopdf, un HTML donnant un PDF de
même nom. L'assemblage, le calage et l'empaquetage viennent en plus, jamais à la
place.

Deux options ne sont volontairement pas reprises :

- `--javascript-delay`, aveu d'impuissance. On attend un signal explicite du
  document, avec délai maximal.
- `--disable-local-file-access`, qui n'est pas une option mais le comportement
  par défaut.

## Hors périmètre

- Édition ou génération de HTML.
- Stockage durable des lots au-delà du TTL.
- Marques OMR de mise sous pli : chaque plieuse a sa spécification propre, et
  c'est invérifiable sans la machine en face. Architecturé comme point
  d'extension — post-traitement avec profil par tenant — pas implémenté.
- PDF/A : pdfcpu ne fait ni conversion ni validation. Passera par Ghostscript ou
  veraPDF si le besoin se confirme.
