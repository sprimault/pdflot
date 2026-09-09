# Feuille de route

Chaque étape produit quelque chose d'exécutable et de démontrable. On ne passe
pas à la suivante tant que la précédente n'est pas utilisable en l'état.

L'étape N se publie en `0.N.0` et lui donne son titre : c'est ce que reprend
`CHANGELOG.md`, puis la note de publication. Les correctifs livrés à l'intérieur
d'une étape incrémentent le dernier chiffre et portent leur propre titre.

## Étape 1 — Squelette et rendu unitaire

Binaire Go, commande `pdflot render`, un fichier HTML en entrée, un PDF en
sortie. Chromium piloté par chromedp, lancé et arrêté proprement.

Fin d'étape : `pdflot render facture.html -o facture.pdf` fonctionne.

## Étape 2 — Image de référence

Image avec Chromium et polices épinglés, exposés en labels. Les rendus de
référence s'exécutent dans le conteneur et jamais sur le poste ; c'est là, et
seulement là, que naissent les empreintes de `testdata/golden`.

Fin d'étape : `make rendu-verif` compare un rendu aux empreintes et échoue sur
un écart d'un pixel.

## Étape 3 — Bac à sable

Réseau coupé, aucun accès hors du répertoire de travail, résolution des
ressources en chemins relatifs uniquement. Tests négatifs : une iframe vers une
adresse interne, un `file://` remontant, un chemin absolu.

Fin d'étape : les trois tests d'évasion échouent.

## Étape 4 — Lot et manifeste

Archive en entrée, `manifeste.json` optionnel à la racine, convention implicite
si absent (un fichier égale un document, un document égale une sortie). Modèle à
trois niveaux en place. Archive de PDF en sortie.

Fin d'étape : un ZIP de 50 HTML donne un ZIP de 50 PDF.

## Étape 5 — Rapport et échec partiel

`rapport.json` dans le paquet, poursuite du lot en cas d'échec unitaire, code de
sortie distinguant succès complet, succès partiel et échec global.

Fin d'étape : un lot avec 3 HTML volontairement cassés sur 50 produit 47 PDF et
un rapport exploitable.

## Étape 6 — Regroupement

Plusieurs HTML vers un même document par concaténation avec saut de page,
plusieurs documents vers une même sortie, bascule sur rendu unitaire puis fusion
au-delà du seuil. Frontières de documents conservées.

Fin d'étape : un manifeste décrivant trois regroupements différents produit
trois sorties correctes.

## Étape 7 — Pagination et seconde passe

Rendu en deux passes. Variables `[page]`, `[topage]`, `[section]`, `[document]`.
En-têtes et pieds en HTML rendus au format exact puis apposés. Sommaire.

Fin d'étape : une sortie fusionnée affiche une pagination continue juste et une
numérotation par section juste, et un lot sans variable de pagination ne paie
pas la seconde passe.

## Étape 8 — Calage recto/verso

Insertion de pages blanches pour qu'un document commence toujours sur un recto,
via `pages insert`. Modes simplex et duplex par sortie, **simplex par défaut** :
une sortie n'insère rien tant que le manifeste ne déclare pas le duplex, le
routeur faisant sa propre imposition dans le cas courant.

Fin d'étape : une sortie duplex de documents de 1, 2 et 3 pages est correctement
calée, et la même sortie en simplex ne gagne pas une page.

## Étape 9 — Signets et métadonnées

Signets construits à partir de la structure de documents. Métadonnées de
document renseignées depuis le manifeste. Manifeste de sortie portant
l'empreinte d'environnement — version de Chromium, du rastériseur et liste des
polices — lue dans le fichier que l'image génère à sa construction. Pas dans ses
labels : ceux-ci ne sont pas lisibles depuis le conteneur, ce sont des
métadonnées destinées au démon.

Fin d'étape : le PDF fusionné est navigable par signets, et le manifeste de
sortie suffit à savoir dans quelle image le lot a été rendu.

## Étape 10 — Pool de rendu et montée en charge

Pool de contextes Chromium recyclés tous les N rendus, parallélisme
configurable, cache des ressources partagées entre documents d'un même lot.

Le pipeline est complet à ce stade, et c'est la raison de cette place : le coût
d'un rendu n'est connu qu'une fois qu'un document peut être une concaténation de
plusieurs HTML et, le cas échéant, être rendu deux fois. Dimensionner le pool
avant, c'est le dimensionner sur une unité de travail qui n'existe plus après.

Fin d'étape : 6000 documents traités sans fuite mémoire, débit mesuré et
consigné.

## Étape 11 — Empaquetage et chiffrement

Archive protégée par mot de passe, et surtout chiffrement AES par PDF via
`encrypt`, avec mot de passe utilisateur et propriétaire distincts. Les deux
mécanismes sont indépendants et déclarés dans le manifeste.

Fin d'étape : un lot d'attestations sort avec un mot de passe différent par
destinataire.

## Étape 12 — Bordereau

Fichier CSV et XML listant, par output et par document : identifiant, nombre de
pages, nombre de feuilles, adresse destinataire, nom du PDF, position dans
l'output. Suffisant pour un routeur ou un service courrier, y compris pour le
calcul d'affranchissement.

La clé est bien le couple, pas le document seul : un même document repris dans
deux outputs y porte deux noms de PDF, deux positions et, si l'un est calé et
l'autre non, deux nombres de feuilles. Rattaché à un seul output, ou dupliqué
sans clé, l'affranchissement calculé dessus serait faux.

Fin d'étape : le bordereau d'un lot de 6000 est cohérent avec les PDF produits.

## Étape 13 — Cloisonnement et quotas par tenant

Cloisonnement des travaux par tenant, quotas disque et de parallélisme,
politique de purge par tenant.

Avant l'API, et non après : `docs/securite.md` range le cloisonnement dans ce qui
est garanti, et `.claude/critical-rules.md` en fait un invariant. Ouvrir le dépôt
de lots à plusieurs tenants avant de l'avoir posé publierait des versions dont le
document de sécurité serait faux, et un cloisonnement ajouté après coup se heurte
à un stockage déjà organisé sans lui.

Fin d'étape : un tenant qui atteint son quota est refusé sans que le voisin s'en
aperçoive.

## Étape 14 — API, dépôt et récupération

Dépôt du lot, identifiant de travail, récupération du paquet. Le binaire garde
son mode ligne de commande hors ligne.

Fin d'étape : un lot complet passe par l'API de bout en bout.

## Étape 15 — API, progression et purge

Suivi de progression pendant le traitement, TTL par travail, purge des lots
expirés.

Fin d'étape : un lot de 6000 rend compte de son avancement pendant qu'il tourne,
et un lot expiré ne laisse rien sur le disque.

## Étape 16 — Reprise d'un lot interrompu

Un lot interrompu — arrêt du service, redémarrage, défaillance de l'hôte —
repart de ce qui a déjà été rendu.

C'est ce qui distingue un traitement de lot d'une boucle : sur 6000 documents,
une reprise à 5000 coûte quelques minutes là où un nouveau départ coûte l'heure
entière. `docs/conception.md` la nomme parmi les raisons d'être du produit.

Fin d'étape : un lot interrompu à mi-course puis relancé produit le même paquet,
sans rejouer les rendus déjà faits.

## Étape 17 — File équitable

Ordonnancement équitable entre tenants. En FIFO, un lot de 6000 bloque dix
minutes le client d'à côté qui en demande trois.

Fin d'étape : un lot de 6000 ne retarde pas de plus de quelques secondes un lot
de 3 déposé par un autre tenant.

## Étape 18 — Compatibilité wkhtmltopdf

Table de correspondance des options historiques vers le manifeste, et mode de
compatibilité en ligne de commande pour les cas simples. Guide de migration.

Fin d'étape : une commande wkhtmltopdf courante s'exécute sans modification.

## Envisagé, non planifié

- PDF/A via Ghostscript ou veraPDF, si le besoin d'archivage légal se confirme.
- Marquage OMR ou datamatrix par page, en post-traitement avec profil par
  tenant.
- Interface web de suivi des travaux.
- Sortie vers un stockage objet plutôt qu'en réponse HTTP.
