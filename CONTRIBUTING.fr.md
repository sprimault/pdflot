# Contribuer

Version anglaise : `CONTRIBUTING.md`.

## Comment ce dépôt est écrit

Le code est écrit avec l'assistance d'un modèle de langage, sous relecture. Les
règles de projet qui en découlent sont dans `docs/` et s'appliquent à tout
contributeur, humain ou non.

## Branche et publication

Une branche par lot, nommée `<type>/<sujet>`. Un commit. Retour par pull
request, fusionnée en squash.

Un lot qui touche à `internal/sandbox` ou au schéma du manifeste part seul : ce
sont les deux endroits où une régression ne se voit pas au diff.

## Avant de pousser

1. `make controles` passe — formatage, `go vet`, sommes de `go.sum`, `gosec`,
   `govulncheck` et tests.
2. `make rendu-verif` passe, dans l'image de référence.
3. La documentation touchée par le lot part dans le même commit.

Le premier point se lance tout seul : `.githooks/pre-push` l'exécute avant chaque
poussée, une fois le hook activé par `git config core.hooksPath .githooks` — à
faire une fois par clone. `git push --no-verify` le contourne, pour pousser une
branche de travail qu'on sait cassée.

Une seule cible, appelée par le hook comme par l'intégration continue. Ce n'est
pas une commodité : tant que les contrôles étaient énumérés de part et d'autre,
les deux listes ont divergé sans bruit, et le job refusait un formatage que le
contrôle local acceptait.

CodeQL est le seul contrôle qu'on ne rejoue pas en local, son analyse demandant
un bundle de près d'un gigaoctet. Il tourne à chaque proposition sur GitHub, et
son absence de `make controles` est délibérée.

Le point 3 n'est pas une politesse : `docs/conception.md` fait foi, et un
document qui prend du retard sur le code devient un vieux papier que plus
personne ne lit.

## Messages

Types conventionnels — `fix`, `feat`, `docs`, `refactor`, `chore` — non
traduits. `chore` couvre ce qui ne touche ni au comportement ni à la
documentation : outillage, construction, intégration continue. Le
scope suit le paquet : `fix(render)`, `feat(manifest)`. Un document français se
traduit dans la moitié anglaise : `docs(conception)` devient `docs(design)`.

Messages bilingues, les deux moitiés séparées par `***`. Quatre lignes de corps
par langue au maximum : au-delà, le contenu appartient à un document du dépôt et
on est en train de l'y doubler.

Un message porte ce qui change et pourquoi — jamais la façon dont le travail a
été mené.

## Documents de test

Aucun extrait de document rendu dans un message, une PR ou une issue. Les lots
d'exemple sont fictifs, mais l'habitude se prend sur eux et se perd sur un lot
réel. Un document se désigne par son identifiant.

## Dépendances

Bibliothèque standard par défaut. Toute dépendance nouvelle se justifie dans la
PR, avec sa licence. Aucune dépendance sous AGPL : elle contaminerait les
utilisateurs du service.

Les versions se figent dans `go.mod` ; les dépendances externes hors Go —
Chromium, polices — s'épinglent dans l'image et se reportent dans le manifeste
de sortie.

## Sécurité

Voir `docs/securite.md` pour le modèle de menace, et `SECURITY.fr.md` pour le
signalement d'une vulnérabilité.

Une contribution qui ajoute une option relâchant une garantie du bac à sable est
refusée par principe, quelle que soit sa qualité par ailleurs.
