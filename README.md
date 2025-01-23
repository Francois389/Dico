# Dico

- [Français](#français)
- [English](#english)

## Français

Je n'ai pas trouvé d'API permettant de demander un mot de la langue française en ligne.

Donc j'ai décidé de la faire moi-même.

La liste de mot que j'utilise a été téléchargé sur le site de [3Z Software](http://www.3zsoftware.com/fr/listes.php), il
s'agit de la liste *Petit Larousse Illustré 2007* que l'on peut
télécharger [ici](http://www.3zsoftware.com/listes/pli07.zip).

Cette API utilise une base de donnée MongoDB qui est dans un conteneur Docker.

## Mise en place du serveur

Cloner le projet en local

Éxecuter le `docker-compose` avec la commande suivante afin de lancer le conteneur :

```bash
docker compose up -d
```

Dès que le conteneur est lancé, il faut remplir la base de donnée avec des mots.

Vouys avez deux options :

1. Le faire à la main ( la base de donnée est accesible en `localhost` sur le port `27027` )
2. Utiliser le fichier [`populate.go`](https://github.com/Francois389/Dico/blob/main/api/populate/populate.go), vous
   pouvez suivre les instructions du README qui se trouve dans le même dossier.

Une fois que la base de donnée est initialisé, vous pouvez executer le serveur.

## Utilisation de l'API

Voici la liste des routes qui sont actuellement accessibles :

Tous les résultats de mots sont renvoyé sous cette forme :

```json
{
  "word": "festivals",
  "length": 9,
  "first_letter": "f",
  "sorted_letter": "aefilsstv"
}
```

### 1. `words/{lettre}`

Renvoie tous les mots qui commencent par la lettre fournie

<details>
  <summary>Erreur possible</summary>

- Si la lettre fournie ne fait pas 1 caractère, une erreur sera renvoyé :

`400 Bad Request`

```json
{
  "error": "Invalid first letter. Expected one character."
}
```

Exemple : `words/abc` renverra cette erreur.

- Si aucun mot ne commence par la lettre fournie, une erreur sera renvoyé :

`404 Not Found`

```json
{
  "error": "No words start with a (╚)"
}
```

Exemple : `words/╚` renverra cette erreur.

</details>

### 2. `word/{lettre}`

Renvoie un mot choisi aléatoirement qui commence par la lettre fournie.

<details>
  <summary>Erreur possible</summary>

- Si la lettre fournie ne fait pas 1 caractère, une erreur sera renvoyé :

`400 Bad Request`

```json
{
  "error": "Invalid first letter. Expected one character."
}
```

Exemple : `word/abc` renverra cette erreur.

- Si aucun mot ne commence par la lettre fournie, une erreur sera renvoyé :

`404 Not Found`

```json
{
  "error": "No words start with a (╚)"
}
```

Exemple : `word/╚` renverra cette erreur.

</details>

### 3. `word/length/{length}`

Renvoie un mot choisi aléatoirement qui a la longueur demandée

<details>
  <summary>Erreur possible</summary>

- Si la longueur fournie n'est pas un nombre, une erreur sera renvoyé :

`400 Bad Request`

```json
{
  "error": "Please give a number"
}
```

Exemple : `word/length/a` renverra cette erreur.

- Si aucun mot avec la longueur fournie n'est trouvé, une erreur sera renvoyé :

`404 Not Found`

```json
{
  "error": "No words with length (111)"
}
```

Exemple : `word/length/111` renverra cette erreur.

</details>

### 4. `anagrams/{mot}`

Renvoie tous les anagrammes du mot fourni

<details>
  <summary>Erreur possible</summary>

- Si le mot fourni n'a pas d'anagramme, une erreur sera renvoyé :

`404 Not Found`

```json
{
  "error": "No anagram found for this word (aaaaaaaa)"
}
```

Exemple : `anagrams/aaaaaaaa` renverra cette erreur.

</details>

<hr>

## English

I couldn't find an API allowing me to request a word of the French language online.

So I decided to do it myself.

The word list that I use was downloaded from the [3Z Software](http://www.3zsoftware.com/fr/listes.php) website,
this is the list *Petit Larousse Illustré 2007*
which can be downloaded [here](http://www.3zsoftware.com/listes/pli07.zip).

This API uses a MongoDB database, which is in a Docker container.

## Setting up the server

Clone the project in local

Run the `docker-compose` file with the following command to launch the container:

```bash
docker compose up -d
```

As soon as the container is launched, you must fill the database with words.

You have two options:

1. Do it by hand (the database is accessible via `localhost` on port `27027`)
2. Use the file [`populate.go`](https://github.com/Francois389/Dico/blob/main/api/populate/populate.go), you can follow
   the instructions in the README which is located in the same folder.

Once the database is initialized, you can run the server.

## Using API

Here is the list of routes that are currently accessible:

All word results are returned in this form:

```json
{
  "word": "festivals",
  "length": 9,
  "first_letter": "f",
  "sorted_letter": "aefilsstv"
}
```

### 1. `words/{letter}`

Returns all words that start with the letter provided

<details>
  <summary>Error possible</summary>

- If the letter provided is not one character long, an error will be returned:

`400 Bad Request`

```json
{
  "error": "Invalid first letter. Expected one character."
}
```

Example: `words/abc` will return this error.

- If no word begins with the letter provided, an error will be returned:

`404 Not Found`

```json
{
  "error": "No words start with a (╚)"
}
```

Example : `words/╚` will return this error.

</details>

### 2. `word/{letter}`

Returns a randomly chosen word that begins with the letter provided.

<details>
  <summary>Possible error</summary>

- If the letter provided is not one character long, an error will be returned:

`400 Bad Request`

```json
{
  "error": "Invalid first letter. Expected one character."
}
```

Example: `word/abc` will return this error.

- If no word begins with the letter provided, an error will be returned:

`404 Not Found`

```json
{
  "error": "No words start with a (╚)"
}
```

Example : `word/╚` will return this error.

</details>

### 3. `word/length/{length}`

Returns a randomly chosen word that has the requested length

<details>
  <summary>Possible error</summary>

- If the length provided is not a number, an error will be returned:

`400 Bad Request`

```json
{
  "error": "Please give a number"
}
```

Example: `word/length/a` will return this error.

- If no word with the provided length is found, an error will be returned:

`404 Not Found`

```json
{
  "error": "No words with length (111)"
}
```

Example: `word/length/111` will return this error.

</details>

### 4. `anagrams/{word}`

Returns all anagrams of the word provided

<details>
  <summary>Possible error</summary>

- If the word provided does not have an anagram, an error will be returned:

`404 Not Found`

```json
{
  "error": "No anagram found for this word (aaaaaaaa)"
}
```

Example: `anagrams/aaaaaaaa` will return this error.

</details>
