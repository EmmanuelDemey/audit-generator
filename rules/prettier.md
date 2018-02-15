# Utilisation de Prettier

Afin d'uniformiser le code source d'un projet dans une équipe de développeurs, nous pouvons ajouter dans le workflow de développement, l'outil
Prettier, développé par Facebook.

Afin d'uniformiser l'utilisation de Prettier dans votre IDE et dans votre système d'intégration continue, il est prérable de configurer l'outil dans le
fichier package.json de votre project :

```json
{
  "name": "nodejs",
  "version": "0.0.0",
  "private": true,
  "scripts": {
    "start": "node ./app/index.js",
    "format": "prettier --write \"app/**/*.js\""
  },
  "dependencies": {
    "body-parser": "1.18.2",
    "express": "4.15.5",
    "morgan": "1.9.0"
  },
  "devDependencies": {
    "prettier": "1.10.2"
  },
  "prettier": {
    "singleQuote": true
  }
}
```
