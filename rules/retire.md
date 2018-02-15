# Utilisation de Retire

Afin d'assurer que toutes les dépendances du projet ne possede pas de vulnérabilités, il est conseillé d'utiliser le module Retire

```json
{
  "name": "nodejs",
  "version": "0.0.0",
  "private": true,
  "scripts": {
    "start": "node ./app/index.js",
    "retire": "retire -n -p"
  },
  "dependencies": {
    "body-parser": "~1.18.2",
    "express": "~4.15.5",
    "morgan": "^1.9.0"
  },
  "devDependencies": {
    "retire": "^1.5.1"
  }
}
```
