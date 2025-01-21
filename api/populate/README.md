# Populate the database

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