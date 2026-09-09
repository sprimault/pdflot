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

1. `make lint` passe — `gofmt -l .` ne renvoie rien, `go vet ./...` est propre.
2. `make test` passe.
3. `make sec` et `make vulncheck` passent.
4. `make rendu-verif` passe, dans l'image de référence.
5. La documentation touchée par le lot part dans le même commit.

Les quatre premiers points se lancent par `make` et non à la main : c'est la
même définition que celle qu'exécute l'intégration continue, et deux listes
tenues en parallèle finissent par diverger sans que rien ne le dise.

Le point 5 n'est pas une politesse : `docs/conception.md` fait foi, et un
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
