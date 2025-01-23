# Initialiser la base de données

- [Français](#français)
- [English](#english)

## Français

Pour ajouter des mots à la base de données, copiez un fichier nommé `mots.txt` qui contient sur chaque ligne un mot à ajouter. 

Ensuite, exécutez la commande suivante :

```bash 
go run populate.go
```

Ou si vous souhaitez supprimer le précédent mots :

```bash
go run populate.go -clear
```

En cas d'erreur, assurez-vous que le serveur mongodb est en cours d'exécution.

## English

To add some words to the database, copy a file named `mots.txt` who contain in each line a word to add.

Then, run the following command:

```bash
go run populate.go
```

or if you want to remove the previous words:

```bash
go run populate.go -clear
```

If any error, be sure that the mongodb server is running.