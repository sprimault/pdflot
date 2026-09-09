# Modèle de menace

## Hypothèse de départ

pdflot exécute du HTML fourni par un tiers dans un moteur de navigateur
complet. Le déposant d'un lot est un attaquant potentiel, y compris s'il s'agit
d'un tenant authentifié.

C'est précisément le scénario contre lequel le mainteneur de wkhtmltopdf mettait
en garde, et que ses CVE de SSRF et de traversée de répertoires illustrent.

## Ce qui est garanti

- Pas d'accès réseau sortant depuis le document rendu, donc pas de SSRF vers un
  service interne ou un point de métadonnées d'instance.
- Pas de lecture de fichier hors du répertoire de travail du travail en cours.
- Ressources résolues uniquement à l'intérieur du lot.
- Cloisonnement disque et parallélisme par tenant.

Ces garanties sont des invariants : aucune option de configuration ne les
relâche. Chacune est couverte par un test négatif.

## Ce qui n'est pas garanti

- L'isolation entre documents d'un même lot. Le lot est l'unité de confiance.
- La protection d'un destinataire vis-à-vis d'un autre par le seul mot de passe
  d'archive : celui-ci protège le transport. Il faut le chiffrement AES par PDF.
- La résistance à un déni de service par document pathologique au-delà des
  quotas et délais configurés.

## Exposition

Le service n'est pas conçu pour être exposé directement sur un réseau public. Il
attend une passerelle d'authentification devant lui.
