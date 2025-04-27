# game-api

Cette API est implementée en Go. 

Go 1.24 est requis pour compiler le projet.

Pour exécuter le serveur en local:

```bash
make run
```

L'API est disponible à `http://localhost:8080`.

Au lancement le binaire initialize la base de données et insère un ensemble de Platforms.

Pour créer un jeu:

```bash
curl -d '{"name":"", "release_date":"2023-09-10", "ratings":19,"platforms":[{"name":"PC"}], "studio":"studio2"}' -X POST -H "Content-Type: application/json" http://localhost:8080/game/
```

La relation entre les jeux et les plateformes présente actuellement un bug qui ne permet pas de satisfaire exactement le schema de requête/réponse attendu: il est necéssaire de fournir un objet avec le champ `name` au lieu du nom de la plateform.
