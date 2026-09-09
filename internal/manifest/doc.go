// Copyright 2026 Stéphane Primault <sprimault@users.noreply.github.com>
// SPDX-License-Identifier: Apache-2.0

// Package manifest décrit le contrat d'entrée d'un lot.
//
// Le manifeste porte les trois niveaux du modèle : document (1 à N fichiers
// HTML formant une unité logique), output (1 à N documents assemblés dans un
// PDF, avec sa règle de calage) et bundle (la livraison empaquetée).
//
// En l'absence de manifeste, la convention implicite s'applique : un fichier
// vaut un document, un document vaut un output.
package manifest
