# Initialiser la base de données

- [Français](#français)
- [English](#english)

## Français

Ce module contient le code pour initialiser la base de données avec des mots.
Il est utiliser par le docker compose dans le dossier parent.

Avant de l'utiliser, assurez-vous d'avoir que ce dossier contient un fichier `mots.txt` avec les mots à ajouter.

Ensuite, soit vous utilisez le fichier `docker-compose.yml` dans le dossier parent, soit vous pouvez exécuter le fichier `populate.go` directement.

### Docker compose (Recommandé)

Pour initialiser la base de données avec Docker, exécutez la commande suivante dans le dossier parent :

```bash
docker compose --profile init up dico-db populate
```

Le conteneur va s'éteindre une fois que les mots seront ajoutés à la base de données.

### Executer le fichier directement

> [!IMPORTANT]  
> Vous devez modifier les paramètres de connexions à la base de donnée.

Exécutez la commande suivante :

```bash
go run populate.go
```

Ou si vous souhaitez supprimer le précédent mots :

```bash
go run populate.go -clear
```

En cas d'erreur, assurez-vous que le serveur mongodb est en cours d'exécution.

## English

This module contains the code to initialize the database with words.
It is used by the docker compose in the parent folder.
Before using it, make sure that this folder contains a `mots.txt` file with the words to add.

Then, either use the `docker-compose.yml` file in the parent folder, or you can run the `populate.go` file directly.

### Docker compose

To initialize the database with Docker, run the following command in the parent folder:

```bash
docker compose --profile init up dico-db populate
```

The container will shut down once the words are added to the database.

### Run the file directly

> [!IMPORTANT]  
> You must modify the database connection settings.

Run the following command:

```bash
go run populate.go
```

Or if you want to remove the previous words:

```bash
go run populate.go -clear
```

In case of an error, make sure that the mongodb server is running.
