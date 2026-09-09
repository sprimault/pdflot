# Conventions Go

## Langue

Identifiants, noms de fichiers et de répertoires en anglais quand c'est le terme
naturel du domaine : `document`, `output`, `bundle`, `render`. Commentaires et
godoc en français.

Éviter dans le code et la documentation les termes français à double sens. Les
termes techniques sans ambiguïté dans leur contexte restent acceptables.

## Style

- Le code doit se lire comme du code écrit par un humain expérimenté.
- Pas de commentaire qui paraphrase la ligne suivante. Un commentaire explique
  pourquoi, jamais quoi.
- Pas de bannière, pas d'emoji, pas de séparateur décoratif.
- Pas de défensive inutile : ne pas vérifier ce que le type garantit déjà.
- Rule-of-three avant tout refactoring, sauf pour les utilitaires de formatage.
- `gofmt` et `go vet` en CI, non négociables.

## Bassin d'entités

Le vocabulaire du domaine est fermé. Une entité qui n'est ni un `document`, ni un
`output`, ni un `bundle`, ni un `job` doit se justifier avant d'exister.

- `Document` — unité logique, porte ses fichiers HTML et ses métadonnées de
  bordereau.
- `Output` — PDF produit, porte sa liste ordonnée de documents et sa règle de
  calage.
- `Bundle` — livraison, porte ses outputs, le bordereau et le rapport.
- `Job` — un lot en cours de traitement, porte son tenant, son répertoire de
  travail et son état.

## Gestion d'erreur

- Erreurs enveloppées avec `%w` et un contexte utile, jamais une chaîne nue.
- Une erreur de rendu unitaire n'interrompt jamais un lot : elle est collectée
  et reportée.
- Le code de sortie distingue succès complet, succès partiel et échec global.
- Les messages nomment le document fautif et la cause.
- Les fonctionnalités non écrites renvoient
  `errors.New("à implémenter : étape N")`, avec le numéro de l'étape du ROADMAP.
  C'est ce marqueur que compte `/initdoc`.

## Contexte et attentes

`context` propagé partout où il y a une attente : rendu, appel externe, écriture
disque. Aucune temporisation en dur en remplacement d'un signal — c'est
exactement l'erreur de `--javascript-delay`.

## Journalisation

`log/slog`, JSON en service, texte en ligne de commande. Jamais de contenu de
document ni de donnée personnelle : identifiant, tenant, durée, code d'erreur.
Le niveau debug peut mentionner des chemins, jamais du contenu.

## Doctrine de test

- **Rendus** : comparaison d'empreinte contre `testdata/golden`, exécutée dans
  l'image de référence uniquement. Une régénération depuis un poste de
  développement est à rejeter.
- **Sécurité** : un test négatif par garantie annoncée — accès réseau, accès
  fichier, chemin absolu, débordement de quota. Un invariant sans test négatif
  n'est pas un invariant.
- **Échec partiel** : le lot `examples/lots/degrade` doit produire les documents
  valides et un rapport exploitable.
- **Débit** : mesuré et consigné à chaque modification touchant au pool de
  rendu.

## Traces

Aucune mention d'assistant IA dans les commits, les PR, les issues ou les notes
de publication. Pas de trailer `Co-Authored-By` — l'attribution est vidée dans
`.claude/settings.json`.

Identité git configurée en local sur le dépôt (`--local`), adresse noreply
GitHub. `CLAUDE.md`, `.claude/` et `memory/` ne sont jamais suivis par git : ils
passent par `.git/info/exclude`, local au clone. Ce sont des outils de travail,
pas des fichiers du projet — les déclarer dans le dépôt reviendrait à y inscrire
l'existence de fichiers qu'un contributeur n'a aucune raison d'avoir.

Messages de commit en français, ligne de sujet courte à l'impératif.

## Dépendances

Bibliothèque standard par défaut. Une dépendance s'ajoute quand elle remplace du
code qu'on ne veut pas maintenir, pas quand elle fait gagner dix lignes.

Licences compatibles Apache 2.0. Aucune dépendance PDF sous AGPL : cela
contaminerait les utilisateurs du service, ce qui tuerait son adoption en
entreprise.

Socle retenu : `chromedp` pour le pilotage CDP, `pdfcpu` pour tout le
post-traitement PDF.
