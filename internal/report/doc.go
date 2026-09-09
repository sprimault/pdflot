// Copyright 2026 Stéphane Primault <sprimault@users.noreply.github.com>
// SPDX-License-Identifier: Apache-2.0

// Package report produit le rapport d'exécution et le bordereau.
//
// Le rapport recense les échecs unitaires : un lot ne s'interrompt jamais parce
// qu'un document casse. Le bordereau liste, par document, l'identifiant, le
// nombre de pages, le nombre de feuilles en recto/verso et le destinataire —
// de quoi alimenter un routeur ou calculer un affranchissement.
package report
