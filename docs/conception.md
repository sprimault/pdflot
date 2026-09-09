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

Le recto/verso le démontre. Si 6000 attestations sont fusionnées dans un PDF
destiné à l'imprimeur, chaque attestation doit commencer sur un recto : une
attestation de 3 pages impose une page blanche avant la suivante. La frontière
entre documents reste donc signifiante **à l'intérieur** de l'output ; elle ne
disparaît pas à la fusion.

Cette même frontière porte les signets, la numérotation par section et le
bordereau. Une implémentation qui fusionne des PDF opaques la perd.

### Plusieurs outputs, un seul rendu

Un même lot peut produire plusieurs outputs à partir des mêmes documents : un
PDF par destinataire pour l'envoi par courriel, et un PDF unique calé en
recto/verso et trié pour le routeur. Le rendu Chromium ne doit être payé qu'une
fois. C'est l'avantage principal sur une boucle appelant un convertisseur
unitaire.

## Rendu

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

### Concaténation plutôt que fusion

Pour un document composé de plusieurs HTML, on les concatène avec un saut de
page entre chaque et on fait **un seul** rendu. Trois gains : un appel navigateur
au lieu de N, une pagination continue native, des compteurs CSS cohérents.

Le prix : collisions de CSS entre fichiers, et explosion mémoire sur les groupes
trop gros. Au-delà d'un seuil configuré, on bascule sur rendu unitaire puis
fusion.

### Signets

Chromium n'en produit pas. Ils sont construits au post-traitement depuis la
structure de documents — information que pdflot possède, contrairement à qui
fusionne des PDF opaques.

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
