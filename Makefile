.PHONY: build test lint sec vulncheck tools image rendu-verif

# Réglages propres à une machine : chemins de cache, prérequis ajoutés aux
# cibles de contrôle. Le tiret parce que la plupart des machines n'en ont pas.
-include makefile.local

build:
	go build -trimpath -o bin/pdflot ./cmd/pdflot

test:
	go test ./...

# `gofmt -l` liste les fichiers mal formés sans jamais échouer. Sans le
# `test -z`, la cible passait au vert sur du code que la CI refusait.
#
# Les répertoires du module plutôt que `.` : un poste peut placer GOCACHE dans
# le dépôt pour échapper à son antivirus, et le cache contient du Go généré que
# personne n'a à formater. `go list` suit le même périmètre que `./...`, qui
# ignore déjà les répertoires en point.
lint:
	test -z "$$(gofmt -l $$(go list -f '{{.Dir}}' ./...))"
	go vet ./...

sec:
	gosec ./...

vulncheck:
	govulncheck ./...

# Ni gosec ni govulncheck ne viennent avec la distribution Go. La cible les
# dépose dans le GOPATH ; c'est à chaque machine de décider, dans son
# makefile.local, si les contrôles la rejouent avant de s'exécuter.
tools:
	go install github.com/securego/gosec/v2/cmd/gosec@latest
	go install golang.org/x/vuln/cmd/govulncheck@latest

# Les tests de rendu ne valent que dans l'image de référence.
image:
	docker build -t pdflot:dev .

# Compare les rendus aux empreintes de testdata/golden. Un diff sans changement
# de code signale un écart d'environnement, pas une référence à mettre à jour.
rendu-verif: image
	docker run --rm -v "$$PWD:/src:ro" pdflot:dev verif-rendu /src
