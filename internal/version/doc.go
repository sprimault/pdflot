// Copyright 2026 Stéphane Primault <sprimault@users.noreply.github.com>
// SPDX-License-Identifier: Apache-2.0

// Package version expose l'empreinte d'environnement.
//
// Version du binaire, de Chromium et du rastériseur, et liste des polices
// installées. Ces valeurs sont reportées dans le manifeste de sortie : un lot
// rejoué six mois plus tard doit pouvoir être comparé au précédent.
//
// Elles se lisent dans le fichier que l'image génère à sa construction, et non
// dans ses labels : un label n'est pas lisible depuis le conteneur, c'est une
// métadonnée que seul le démon expose.
package version
