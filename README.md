# Dico

Je n'ai pas trouvé d'API permettant de demander un mot de la langue française en ligne.

Donc j'ai décidé de la faire moi même.

La liste de mot que j'utilise a été téléchargé sur le site de [3Z Software](http://www.3zsoftware.com/fr/listes.php), il s'agit de la liste *Petit Larousse Illustré 2007* que l'on peut télécharger [ici](http://www.3zsoftware.com/listes/pli07.zip).

Cette API utilise une base de donnée MongoDB qui est dans un conteneur Docker.

## Mise en place du serveur

Cloner le projet en local

Éxecuter le `docker-compose` avec la commande suivante afin de lancer le conteneur :

```bash
docker compose up -d
```

Dés que le conteneur est lancé, il faut remplir la base de donnée avec des mots.

Vouys avez duex options :

1. Le faire à la main ( la base de donnée est accesible en `localhost` sur le port `27027` )
2. Utiliser le fichier [`populate.go`](https://github.com/Francois389/Dico/blob/main/api/populate/populate.go), vous pouvez suivre les instructions du README qui ce trouve dans le même dossier.

Une fois que la base de donnée est initialisé, vous pouvez executer le serveur.

## Utilisation de l'API

Voici la liste des routes qui sont actuellement accessibles :

Tous les résultats de mots sont renvoyé sous cette forme :

```json
{
    "word":"festivals",
    "length":9,
    "first_letter":"f",
    "sorted_letter":"aefilsstv"
}
```

### 1. `mots/{lettre}`

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

Exemple : `mots/abc` renverra cette erreur.

- Si aucun mot ne commence par la lettre fournie, une erreur sera renvoyé :

`404 Not Found`

```json
{
  "error": "No words start with a (╚)"
}
```

Exemple : `mots/╚` renverra cette erreur.

</details>

### 2. `mot/{lettre}`

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

Exemple : `mot/abc` renverra cette erreur.

- Si aucun mot ne commence par la lettre fournie, une erreur sera renvoyé :

`404 Not Found`

```json
{
  "error": "No words start with a (╚)"
}
```

Exemple : `mot/╚` renverra cette erreur.

</details>

### 3. `mot/length/{length}`

Renvoie un mot choisi aléatoirement qui a la longueur demandé

<details>
  <summary>Erreur possible</summary>

- Si la longueur fournie n'est pas un nombre, une erreur sera renvoyé :

`400 Bad Request`

```json
{
  "error":"Please give a number"
}
```

Exemple : `mot/length/a` renverra cette erreur.

- Si aucun mot avec la longueur fournie n'est trouvé, une erreur sera renvoyé :

`404 Not Found`

```json
{
  "error": "No words with length (111)"
}
```

Exemple : `mot/length/111` renverra cette erreur.

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

</details>
