# Sécurité

Version anglaise : `SECURITY.md`.

## Signaler une vulnérabilité

Par l'onglet **Security** du dépôt, **Report a vulnerability**. Le signalement
reste privé jusqu'à la publication d'un correctif.

Pas d'issue publique : une évasion du bac à sable décrite en clair avant
correctif vaut mode d'emploi.

Le projet n'a pas d'équipe dédiée. La réponse vient dès que possible, sans délai
garanti.

## Versions concernées

Aucune version n'est publiée à ce jour. Seule la branche `master` est concernée.

## Ce qui relève d'une vulnérabilité

Le modèle de menace complet est dans `docs/securite.md`. L'hypothèse de départ
est que le HTML déposé est hostile, y compris venant d'un tenant authentifié :
toute évasion des garanties du bac à sable en relève.

- Accès réseau sortant depuis un document rendu.
- Lecture ou écriture hors du répertoire de travail du travail en cours.
- Résolution d'une ressource hors du lot.
- Franchissement du cloisonnement entre tenants, disque ou parallélisme.
- Contenu de document ou donnée personnelle apparaissant dans les journaux.

## Ce qui n'en relève pas

`docs/securite.md` énumère ce que le projet ne garantit pas. Un signalement qui
porte là-dessus est traité comme une demande d'évolution, pas comme une
vulnérabilité :

- L'isolation entre documents d'un même lot. Le lot est l'unité de confiance.
- Le mot de passe d'archive vu comme protection d'un destinataire vis-à-vis d'un
  autre : il protège le transport, et c'est le chiffrement par PDF qui protège
  le destinataire.
- Le déni de service par document pathologique, au-delà des quotas et des délais
  configurés.

Une contribution qui ajoute une option relâchant une garantie du bac à sable est
refusée par principe, quelle que soit sa qualité par ailleurs.
