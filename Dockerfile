# Copyright 2026 Stéphane Primault <sprimault@users.noreply.github.com>
# SPDX-License-Identifier: Apache-2.0

# Image de référence : c'est elle qui définit le rendu, pas le poste de
# développement. Chromium et polices sont épinglés, leurs versions sont
# reportées dans le manifeste de sortie.

# syntax=docker/dockerfile:1

FROM golang:1.27 AS build
WORKDIR /src
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -trimpath -o /out/pdflot ./cmd/pdflot

FROM debian:bookworm-slim
# TODO étape 2 : épingler la version de Chromium et la liste des polices.
COPY --from=build /out/pdflot /usr/local/bin/pdflot
USER 65532:65532
ENTRYPOINT ["/usr/local/bin/pdflot"]
