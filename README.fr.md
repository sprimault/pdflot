# pdflot

Transformation de lots de fichiers HTML en PDF, avec règles de regroupement,
calage recto/verso et livraison en archive.

On dépose un lot — une archive de fichiers HTML, leurs ressources et un
manifeste optionnel. pdflot rend les documents avec Chromium, les assemble selon
les règles demandées, et renvoie une archive contenant les PDF, un bordereau et
un rapport d'exécution. L'ordre de grandeur visé est de 6000 documents par lot.

## Avertissement

pdflot exécute du HTML fourni par un tiers dans un moteur de navigateur. Tout
HTML entrant est traité comme hostile : réseau coupé, aucun accès au système de
fichiers hors du répertoire de travail, ressources résolues uniquement à
l'intérieur du lot. Ces garanties ne sont pas désactivables par configuration.

Ne pas exposer le service sur un réseau public sans passerelle
d'authentification devant.

## Ce que pdflot n'est pas

Ce n'est pas un convertisseur unitaire HTML vers PDF : Gotenberg occupe déjà
cette place et le fait bien. La valeur de pdflot est la couche d'orchestration
au-dessus — regroupement, calage, empaquetage, ordonnancement multi-tenant.

Ce n'est pas non plus un fork de wkhtmltopdf. Seul le vocabulaire d'options est
repris, pour que la migration soit immédiate.

## État

Projet en démarrage. Voir `ROADMAP.md`.

## Licence

Apache 2.0.
