# 2tpost - Publiez en un clic sur X et Threads

**2tpost** est une application multiplateforme qui permet de publier du contenu simultanément sur **Twitter X** et **Threads** en un seul clic. Conçue pour être simple d’utilisation, elle vous aide à gérer efficacement vos publications sur ces deux plateformes à partir d'une interface intuitive.

### Fonctionnalités

- **Publication Multiplateforme** : Publiez simultanément sur Twitter X et Threads.
- **Interface Intuitive** : Une interface épurée pour une prise en main rapide.
- **Gestion des Tokens API** : Enregistrez et gérez vos tokens API directement dans l'application.
- **Vérification des Limites de Caractères** : Assurez-vous que vos publications respectent les limites imposées par chaque plateforme.
- **Option Premium** : Dépassez les limites de caractères pour les utilisateurs premium.
- **Notifications** : Recevez des alertes de succès ou d’erreur après chaque publication.

---

## Installation

Vous pouvez télécharger l'exécutable pour votre système d'exploitation ou compiler l'application à partir du code source.

### Téléchargement des binaires

- **Windows** : [Télécharger pour Windows](https://github.com/benoitpetit/2tpost/releases/latest/download/2tpost_windows.exe)
- **Linux** : [Télécharger pour Linux](https://github.com/benoitpetit/2tpost/releases/latest/download/2tpost_linux)
- **Mac** : (Disponible bientôt)

### Prérequis

Avant de commencer, assurez-vous que les éléments suivants sont installés sur votre système :

- **Node.js** (version 14 ou supérieure) : [Télécharger Node.js](https://nodejs.org/)
- **Go** (version 1.16 ou supérieure) : [Télécharger Go](https://golang.org/dl/)
- **Wails** (version 2) : [Installer Wails](https://wails.io/)
- **Git** : [Télécharger Git](https://git-scm.com/)

### Compilation à partir du code source

1. **Cloner le Répertoire**

   ```bash
   git clone https://github.com/benoitpetit/2tpost.git
   cd 2tpost
   ```

2. **Installer les Dépendances**

   ```bash
   go mod tidy
   ```

3. **Construire l'Application**

   Pour compiler l'application :

   ```bash
   wails build
   ```

   Ou pour exécuter en mode développement :

   ```bash
   wails dev
   ```

---

## Configuration

### Obtention des Tokens API

Pour publier via **2tpost**, vous devrez configurer vos tokens API pour **Twitter X** et **Threads**. Vous pouvez suivre les guides spécifiques de chaque plateforme pour obtenir ces jetons. Consultez la [documentation complète ici](https://benoitpetit.notion.site/Documentation-pour-l-application-2tpost-1266e0e9df2b80f3afbfc98e3e08c7c6).

### Sauvegarde des Tokens

1. **Lancer l'Application** : Ouvrez **2tpost**.
2. **Accéder aux Paramètres** : Cliquez sur **Paramètres** dans le header.
3. **Entrer les Tokens** :
   - Dans la section **Tokens API Twitter X**, saisissez vos clés et jetons.
   - Dans la section **Tokens API Threads**, entrez votre jeton d’accès et ID utilisateur.
4. **Sauvegarder** : Cliquez sur **Sauvegarder** pour enregistrer vos tokens dans le fichier `config.json`.

---

## Utilisation

### Publier du Contenu

1. **Entrer le Contenu** : Saisissez le texte que vous souhaitez publier dans la zone de texte.
2. **Vérifier les Limites de Caractères** : Un compteur vous montre les limites spécifiques à Twitter X (280 caractères) et Threads (500 caractères).
3. **Publier** : Sélectionnez où publier entre **Twitter X**, **Threads**, ou **Les deux**.

---

## Sécurité

Le fichier `config.json` contient vos tokens API, il est donc crucial de le protéger et de ne pas le partager publiquement. `Il ce trouve a la racine de l'application.`

---

## Soutenez le Développement de **2tpost**

Si vous aimez l'application et souhaitez contribuer à son développement, vous pouvez me soutenir en m’offrant un café via **Buy Me a Coffee**. Votre soutien permet de maintenir le projet à jour et de continuer à ajouter de nouvelles fonctionnalités.

**Soutenez-moi ici :** [Buy Me a Coffee](https://www.buymeacoffee.com/benoitpetit)

---

## Informations supplémentaires

- **Mon site** : [benoitpetit.dev](https://benoitpetit.dev)
- **GitHub** : [benoitpetit](https://github.com/benoitpetit)
- **Twitter X** : [@devbyben](https://twitter.com/devbyben)
- **Threads** : [benoitpetit.dev](https://www.threads.net/@benoitpetit.dev)

---

**Made with ❤️ by Ben**