// Copyright 2026 Stéphane Primault <sprimault@users.noreply.github.com>
// SPDX-License-Identifier: Apache-2.0

// Package render pilote Chromium et produit les PDF bruts.
//
// Le rendu se fait en deux passes lorsque des variables de pagination sont
// employées : la première produit le corps et donne le nombre de pages, la
// seconde rend les en-têtes et pieds au format exact de la page. Sans variable
// de pagination, la seconde passe est évitée.
//
// Les documents composés de plusieurs fichiers HTML sont concaténés avec un
// saut de page entre chaque et rendus en une fois, pour conserver une
// pagination continue et des compteurs CSS cohérents. Au-delà du seuil
// configuré, on bascule sur rendu unitaire puis fusion.
package render
