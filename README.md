<img src="Images/Logo.png" width="800" height="800" />

# Scrapper 

[![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/) [![Discord](https://img.shields.io/badge/Discord-%235865F2.svg?style=for-the-badge&logo=discord&logoColor=white)](https://discord.com/) [![Udemy](https://img.shields.io/badge/Udemy-%23EC5252.svg?style=for-the-badge&logo=udemy&logoColor=white)](https://www.udemy.com/)

Ce projet Go permet de récupérer automatiquement toutes les formations Udemy disponibles avec un coupon de réduction de 100% et de les envoyer directement sur un canal Discord.

## Fonctionnalités

* **Scraping Udemy :** Récupération des formations gratuites grâce à un coupon de réduction.
* **Intégration Discord :** Envoi des informations des formations sur un canal Discord spécifié.
* **Configuration facile :** Utilisation d'un fichier `.env` pour stocker les informations sensibles.

## Prérequis

* Compte Discord avec un bot créé et autorisé à envoyer des messages.
* Clé API Discord du bot.
* ID du canal Discord où les formations seront envoyées.
* Go installé sur votre machine.

## Installation

1.  Clonez le dépôt :

    ```bash
    git clone https://github.com/Maxence-villet/Scrapper.git
    cd Scrapper
    ```

2.  Créez un fichier `.env` à la racine du projet et remplissez-le avec vos informations :

    ```
    API_KEY=VOTRE_CLE_API_DISCORD
    CHANNEL_ID=ID_DU_CANAL_DISCORD
    ```

3.  Installez les dépendances :

    ```bash
    go mod tidy
    ```

4.  Compilez et exécutez le projet :

    ```bash
    go run main.go
    ```

### Options disponibles

- **`-k` ou `--key`** : Spécifie une liste d'éléments qui iront dans la whitelist.  
  Exemple : `-k "[project,app,build]"`  
  Cette option attend une chaîne de caractères représentant une liste de mot-clés.

- **`-b` ou `--blacklist`** : Spécifie une liste d'éléments à exclure.  
  Exemple : `-b "[salesforce,php]"`  
  Cette option attend une chaîne de caractères représentant une liste d'éléments à exclure.

- **`-i` ou `--input`** : Récupérer les éléments d'un fichier CSV.
  Exemple : `-i example.csv`
  Cette option attend une chaîne de caractères resprésentant le nom du fichier.

- **`-h` ou `--help`** : Affiche l'aide et la liste des options disponibles.  
  Exemple : `-h` ou `--help`

## Lecture de fichier CSV pour la blacklist et la whitelist

En plus des options de ligne de commande, vous pouvez également utiliser un fichier CSV pour définir la blacklist et la whitelist. Cela peut être utile si vous avez une liste importante d'éléments à inclure ou à exclure.

### Utilisation

Pour utiliser un fichier CSV, vous pouvez spécifier le fichier de sortie avec l'option `-i` ou `--input`. Le fichier CSV doit être structuré de la manière suivante :

```csv
build,horse
app,java
linux, 
```
**⚠️ Les mots  toujours en minuscules dans le fichier CSV ⚠️**

Par exemple :
```bash
go run main.go -i example.csv
  or
go run main.go --input example.csv
   ```

On obtiendra :
```
Whitelist = [build,app,linux]
Blacklist = [horse,java]
```



### Exemples d'utilisation

1. **Utilisation basique avec les options `-k` et `-b`** :

   ```bash
   go run main.go -k "[project,app,build]" -b "[salesforce,php]"
   ```

## Utilisation

Le script va automatiquement récupérer les formations Udemy gratuites et les envoyer sur le canal Discord configuré. Il est recommandé de lancer le script régulièrement (par exemple, via un cron job) pour ne manquer aucune nouvelle formation.

## Contribution

Les contributions sont les bienvenues ! N'hésitez pas à ouvrir une issue ou à soumettre une pull request.

## Licence

Ce projet est sous licence GPL-3.0 license.

## Auteur

* Maxence-villet
