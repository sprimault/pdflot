// Copyright 2026 Stéphane Primault <sprimault@users.noreply.github.com>
// SPDX-License-Identifier: Apache-2.0

// Package assemble construit les PDF de sortie à partir des rendus.
//
// L'assemblage conserve les frontières de documents à l'intérieur d'un output.
// C'est cette information qui permet le calage recto/verso (insertion de pages
// blanches pour qu'un document commence toujours sur un recto), les signets, la
// numérotation par section et le bordereau.
package assemble
