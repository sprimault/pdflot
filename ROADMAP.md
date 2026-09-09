# Feuille de route

Chaque étape produit quelque chose d'exécutable et de démontrable. On ne passe
pas à la suivante tant que la précédente n'est pas utilisable en l'état.

## Étape 1 — Squelette et rendu unitaire

Binaire Go, commande `pdflot render`, un fichier HTML en entrée, un PDF en
sortie. Chromium piloté par chromedp, lancé et arrêté proprement.

Fin d'étape : `pdflot render facture.html -o facture.pdf` fonctionne.

## Étape 2 — Image Docker de référence

Image avec Chromium et polices épinglés, version reportée dans les métadonnées
du PDF produit. Tests de rendu exécutés dans le conteneur, pas sur le poste.

Fin d'étape : deux exécutions à un mois d'écart donnent un PDF identique.

## Étape 3 — Bac à sable

Réseau coupé, pas d'accès hors du répertoire temporaire, résolution des
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

## Étape 6 — Pool de rendu et montée en charge

Pool de contextes Chromium recyclés tous les N rendus, parallélisme
configurable, cache des ressources partagées entre documents d'un même lot.

Fin d'étape : 6000 documents traités sans fuite mémoire, débit mesuré et
consigné.

## Étape 7 — Regroupement

Plusieurs HTML vers un même document par concaténation avec saut de page,
plusieurs documents vers une même sortie, bascule sur rendu unitaire puis fusion
au-delà du seuil. Frontières de documents conservées.

Fin d'étape : un manifeste décrivant trois regroupements différents produit
trois sorties correctes.

## Étape 8 — Calage recto/verso

Insertion de pages blanches pour qu'un document commence toujours sur un recto,
via `pages insert`. Modes simplex et duplex par sortie.

Fin d'étape : une sortie duplex de documents de 1, 2 et 3 pages est correctement
calée.

## Étape 9 — Pagination et seconde passe

Rendu en deux passes. Variables `[page]`, `[topage]`, `[section]`, `[document]`.
En-têtes et pieds en HTML rendus au format exact puis apposés. Sommaire.

Fin d'étape : un PDF fusionné de 6000 documents affiche une pagination continue
juste et une numérotation par section juste.

## Étape 10 — Signets et métadonnées

Signets construits à partir de la structure de documents. Métadonnées de
document renseignées depuis le manifeste. Empreinte d'environnement (version de
Chromium, polices) dans le manifeste de sortie.

Fin d'étape : le PDF fusionné est navigable par signets.

## Étape 11 — Empaquetage et chiffrement

Archive protégée par mot de passe, et surtout chiffrement AES par PDF via
`encrypt`, avec mot de passe utilisateur et propriétaire distincts. Les deux
mécanismes sont indépendants et déclarés dans le manifeste.

Fin d'étape : un lot d'attestations sort avec un mot de passe différent par
destinataire.

## Étape 12 — Bordereau

Fichier CSV et XML listant, par document : identifiant, nombre de pages, nombre
de feuilles en recto/verso, adresse destinataire, nom du PDF. Suffisant pour un
routeur ou un service courrier, y compris pour le calcul d'affranchissement.

Fin d'étape : le bordereau d'un lot de 6000 est cohérent avec les PDF produits.

## Étape 13 — API asynchrone

Dépôt du lot, identifiant de travail, suivi de progression, récupération du
paquet, TTL et purge. Le binaire garde son mode ligne de commande hors ligne.

Fin d'étape : un lot complet passe par l'API de bout en bout.

## Étape 14 — Multi-tenant

Cloisonnement par tenant, file équitable, quotas disque et de parallélisme,
politique de purge par tenant.

Fin d'étape : un lot de 6000 ne retarde pas de plus de quelques secondes un lot
de 3 déposé par un autre tenant.

## Étape 15 — Compatibilité wkhtmltopdf

Table de correspondance des options historiques vers le manifeste, et mode de
compatibilité en ligne de commande pour les cas simples. Guide de migration.

Fin d'étape : une commande wkhtmltopdf courante s'exécute sans modification.

## Envisagé, non planifié

- PDF/A via Ghostscript ou veraPDF, si le besoin d'archivage légal se confirme.
- Marquage OMR ou datamatrix par page, en post-traitement avec profil par
  tenant.
- Interface web de suivi des travaux.
- Sortie vers un stockage objet plutôt qu'en réponse HTTP.
