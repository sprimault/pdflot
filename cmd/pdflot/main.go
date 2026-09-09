// Copyright 2026 Stéphane Primault <sprimault@users.noreply.github.com>
// SPDX-License-Identifier: Apache-2.0

// Commande pdflot : transformation de lots HTML en PDF.
package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err := run(ctx, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// run dispatche vers les sous-commandes. Le code de sortie distingue le succès
// complet, le succès partiel et l'échec global : un lot dont quelques documents
// échouent reste un lot livré.
func run(ctx context.Context, args []string) error {
	return fmt.Errorf("non implémenté")
}
