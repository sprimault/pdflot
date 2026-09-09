# Construction

## Principe

**Le rendu est défini par l'image, pas par le poste.** Chromium ne se comporte
pas pareil sous Windows et sous Linux — polices, moteur, métriques. Développer et
valider dans le conteneur dès la première étape, sinon les écarts au pixel se
découvrent en production.

Le binaire multiplateforme reste un objectif ; la cible de référence reste
l'image.

## Ce que le livrable contient, et ce qu'il ne contient pas

Chromium fait plusieurs centaines de mégaoctets de code natif : il n'est pas
embarqué dans le binaire et ne peut pas l'être. Contrairement à wkhtmltopdf, qui
embarquait son moteur, le livrable réel est :

- un binaire Go statique, plus un Chromium présent sur la machine ;
- une image Docker avec Chromium et polices épinglés, qui est la référence.

## Épinglage

L'image épingle une version exacte de Chromium et une liste exacte de polices.
Ces deux valeurs sont exposées en labels d'image et reportées dans le manifeste
de sortie de chaque lot.

Un lot rejoué six mois plus tard doit pouvoir être comparé au précédent. Sans cet
épinglage, aucune promesse de reproductibilité n'est tenable.

## Cibles

- `make build` — binaire local, pour le développement du code hors rendu.
- `make test` — tests unitaires, hors rendu.
- `make lint` — `gofmt -l` vide et `go vet` propre.
- `make sec` — `gosec` sur l'ensemble des paquets.
- `make vulncheck` — `govulncheck` sur les dépendances.
- `make tools` — installe `gosec` et `govulncheck`, qui ne viennent pas avec la
  distribution Go.
- `make image` — image de référence.
- `make rendu-verif` — rendus comparés aux empreintes de `testdata/golden`, dans
  l'image. **Doit passer.**

L'intégration continue n'exécute pas autre chose : elle appelle ces mêmes
cibles. Une commande recopiée dans le workflow serait une seconde définition du
contrôle, et c'est ainsi que le job a fini par refuser un formatage que
`make lint` acceptait.

## Réglages propres à une machine

Le Makefile fait un `-include makefile.local`, ignoré par git. C'est là que se
déclarent les chemins de cache d'un poste et les prérequis qu'il ajoute aux
cibles de contrôle — typiquement, réinstaller les outils avant chaque
exécution. Rien de ce qui relève du projet n'y a sa place.

## Empreintes de référence

`testdata/golden` ne se régénère que dans l'image de référence. Un diff après
régénération sans changement de code signale une version de Chromium ou un jeu
de polices différent : c'est un défaut d'environnement, on corrige l'image,
jamais le fichier de référence.

L'écriture dans `testdata/golden` est refusée par la configuration de session
pour cette raison.

## Dépendances externes à venir

PDF/A n'est pas couvert par pdfcpu. Si le besoin d'archivage légal se confirme,
Ghostscript ou veraPDF entre dans l'image, épinglé au même titre que Chromium, et
sa version rejoint le manifeste de sortie.
