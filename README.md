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

## Utilisation

Le script va automatiquement récupérer les formations Udemy gratuites et les envoyer sur le canal Discord configuré. Il est recommandé de lancer le script régulièrement (par exemple, via un cron job) pour ne manquer aucune nouvelle formation.

## Contribution

Les contributions sont les bienvenues ! N'hésitez pas à ouvrir une issue ou à soumettre une pull request.

## Licence

Ce projet est sous licence MIT.

## Auteur

* Maxence Villet
