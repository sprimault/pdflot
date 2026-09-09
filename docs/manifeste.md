# Manifeste de lot

Le manifeste est le contrat central de pdflot. Il se place à la racine du lot
sous le nom `manifeste.json`.

Il est optionnel : en son absence, la convention implicite s'applique — un
fichier HTML vaut un document, un document vaut un output, et le paquet contient
tous les PDF à plat.

## Structure

À rédiger sous forme de schéma JSON avant la première implémentation. Les trois
niveaux à couvrir :

- **documents** — identifiant, liste ordonnée de fichiers HTML, métadonnées
  destinées au bordereau (destinataire, référence).
- **outputs** — nom du PDF produit, liste ordonnée de documents, mode de calage
  (`simplex` ou `duplex`), en-tête et pied, page de garde, sommaire, signets,
  chiffrement propre à ce PDF.
- **bundle** — format d'archive, mot de passe éventuel, formats de bordereau
  demandés.

## Variables de pagination

Reprises de wkhtmltopdf pour que la migration soit immédiate : `[page]`,
`[topage]`, `[section]`, `[document]`. Leur emploi déclenche le rendu en deux
passes ; sans elles, la seconde passe est évitée.

## Options volontairement absentes

- `javascript-delay` : on attend un signal explicite du document, avec délai
  maximal.
- `disable-local-file-access` : comportement par défaut, non désactivable.

La page de garde n'est pas une option : c'est un document comme un autre dans
l'output.
