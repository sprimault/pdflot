// Copyright 2026 Stéphane Primault <sprimault@users.noreply.github.com>
// SPDX-License-Identifier: Apache-2.0

// Package sandbox porte les invariants de sécurité du rendu.
//
// Les garanties de ce package ne sont pas des options : rien dans la
// configuration ne doit permettre de les relâcher. Réseau sortant coupé, aucun
// accès hors du répertoire de travail, pas de résolution de chemin absolu ni de
// remontée d'arborescence, quota disque par tenant.
package sandbox
