// Copyright 2026 Stéphane Primault <sprimault@users.noreply.github.com>
// SPDX-License-Identifier: Apache-2.0

// Package queue ordonnance les travaux entre tenants.
//
// L'ordonnancement est équitable, jamais FIFO : un lot de 6000 documents ne
// doit pas bloquer le lot de 3 déposé par un autre tenant.
package queue
