// Copyright 2026 Stéphane Primault <sprimault@users.noreply.github.com>
// SPDX-License-Identifier: Apache-2.0

// Package bundle empaquette la livraison.
//
// Deux protections distinctes et indépendantes : le mot de passe d'archive, qui
// protège le transport, et le chiffrement AES par PDF, seul à protéger un
// destinataire d'un autre. Les documents nominatifs relèvent du second.
package bundle
