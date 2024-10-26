<template>
  <div
    id="app"
    class="flex flex-col items-center justify-start h-screen bg-gray-900 text-gray-100"
  >
    <div
      class="w-full bg-gray-800 text-white flex items-center justify-between p-4"
    >
      <h1 class="text-lg font-semibold">
        2tpost - Publiez en un clic sur X et Threads
      </h1>
      <button
        @click="toggleSettings"
        class="text-sm bg-gray-700 hover:bg-gray-600 px-4 py-2 rounded transition flex items-center"
      >
        <i class="fas fa-cog mr-2"></i> Paramètres
      </button>
    </div>

    <div
      class="flex flex-col w-full max-w-2xl mt-6 p-4 bg-gray-800 rounded-md shadow-md"
    >
      <div class="relative">
        <a
          @click.prevent="openBuyMeACoffee"
          class="absolute top-2 right-2 text-gray-400 hover:text-gray-300 hover:cursor-pointer text-xl transform transition-transform duration-300 hover:scale-125"
        >
          <i class="fas fa-heart text-red-500 mx-1"></i>
        </a>
        <textarea
          v-model="content"
          placeholder="Entrez le contenu à poster..."
          class="w-full h-80 p-4 mb-4 border bg-gray-800 border-gray-600 rounded focus:outline-none focus:ring-2 focus:ring-blue-500 whitespace-pre-wrap"
        ></textarea>
        <div class="flex justify-between text-sm mb-2">
          <span
            :class="{
              'text-gray-500': !isExceeded('twitter'),
              'text-red-500': isExceeded('twitter'),
            }"
          >
            Twitter/X: {{ content.length }}/{{ characterLimits.twitter }}
          </span>
          <span
            :class="{
              'text-gray-500': !isExceeded('threads'),
              'text-red-500': isExceeded('threads'),
            }"
          >
            Threads: {{ content.length }}/{{ characterLimits.threads }}
          </span>
        </div>
      </div>

      <div class="flex justify-between space-x-4">
        <button
          @click="preparePost('twitter')"
          :disabled="
            isLoading ||
            !isConfigValid ||
            (!allowExceedLimits && content.length > characterLimits.twitter)
          "
          class="flex-grow bg-blue-600 hover:bg-blue-700 text-white py-2 px-4 rounded transition flex items-center justify-center"
        >
          <span v-if="isLoading && confirmationAction === 'twitter'">
            <i class="fas fa-spinner fa-spin mr-2"></i> Publication...
          </span>
          <span v-else>
            <i class="fa-brands fa-x-twitter mr-2"></i> Twitter X
          </span>
        </button>
        <button
          @click="preparePost('threads')"
          :disabled="
            isLoading ||
            !isConfigValid ||
            (!allowExceedLimits && content.length > characterLimits.threads)
          "
          class="flex-grow bg-green-600 hover:bg-green-700 text-white py-2 px-4 rounded transition flex items-center justify-center"
        >
          <span v-if="isLoading && confirmationAction === 'threads'">
            <i class="fas fa-spinner fa-spin mr-2"></i> Publication...
          </span>
          <span v-else>
            <i class="fa-brands fa-threads mr-2"></i> Threads
          </span>
        </button>
        <button
          @click="preparePost('both')"
          :disabled="
            isLoading ||
            !isConfigValid ||
            (!allowExceedLimits &&
              (content.length > characterLimits.twitter ||
                content.length > characterLimits.threads))
          "
          class="flex-grow bg-purple-600 hover:bg-purple-700 text-white py-2 px-4 rounded transition flex items-center justify-center"
        >
          <span v-if="isLoading && confirmationAction === 'both'">
            <i class="fas fa-spinner fa-spin mr-2"></i> Publication...
          </span>
          <span v-else
            ><i class="fa-regular fa-paper-plane mr-2"></i> Les deux</span
          >
        </button>
      </div>

      <p v-if="!isConfigValid" class="text-red-500 mt-2">
        Veuillez configurer les tokens avant de publier.
      </p>
    </div>

    <transition name="fade">
      <div
        v-if="showConfirmation"
        class="fixed inset-0 flex items-center justify-center bg-gray-900 bg-opacity-50 z-50"
      >
        <div
          class="bg-gray-800 rounded-lg shadow-lg p-6 w-2/3 max-h-80vh overflow-y-auto relative"
        >
          <button
            @click="closeConfirmation"
            class="absolute top-4 right-4 text-gray-500 hover:text-gray-300 focus:outline-none"
          >
            <i class="fas fa-times"></i>
          </button>
          <h2 class="text-2xl font-bold mb-4 text-gray-100">
            Confirmation de publication
          </h2>
          <p class="text-gray-300 mb-6">
            Êtes-vous sûr de vouloir publier ce contenu sur
            <span v-if="confirmationAction === 'twitter'">Twitter X</span>
            <span v-else-if="confirmationAction === 'threads'">Threads</span>
            <span v-else>les deux plateformes</span> ?
          </p>
          <div class="flex justify-end space-x-4">
            <button
              @click="confirmPost"
              class="bg-blue-500 hover:bg-blue-600 text-white font-semibold py-2 px-4 rounded transition flex items-center"
            >
              <i class="fas fa-check mr-2"></i> Confirmer
            </button>
            <button
              @click="closeConfirmation"
              class="bg-red-500 hover:bg-red-600 text-white font-semibold py-2 px-4 rounded transition flex items-center"
            >
              <i class="fas fa-times mr-2"></i> Annuler
            </button>
          </div>
        </div>
      </div>
    </transition>

    <transition name="fade">
      <div
        v-if="showSettings"
        class="fixed inset-0 flex items-center justify-center bg-gray-900 bg-opacity-50 z-50"
      >
        <div
          class="bg-gray-800 rounded-lg shadow-lg p-6 w-2/3 max-h-80vh overflow-y-auto relative"
        >
          <button
            @click="toggleSettings"
            class="absolute top-4 right-4 text-gray-500 hover:text-gray-300 focus:outline-none"
          >
            <i class="fas fa-times"></i>
          </button>
          <h2 class="text-2xl font-bold mb-4 text-gray-100">
            Paramètres des tokens
          </h2>

          <div class="mb-6">
            <h3
              class="text-lg font-semibold text-gray-300 mb-3 cursor-pointer flex items-center justify-between"
              @click="toggleTwitterTokens"
            >
              <span>Tokens API Twitter X</span>
              <span v-if="showTwitterTokens">⬆️</span>
              <span v-else>⬇️</span>
            </h3>

            <transition name="slide">
              <div v-if="showTwitterTokens" class="space-y-4">
                <input
                  v-model="twitterAPIKey"
                  placeholder="Clé API Twitter"
                  class="w-full p-3 border bg-gray-900 border-gray-600 rounded focus:outline-none focus:ring-2 focus:ring-blue-500 placeholder-gray-500"
                  aria-label="Clé API Twitter"
                />
                <input
                  v-model="twitterAPISecret"
                  placeholder="Secret API Twitter"
                  class="w-full p-3 border bg-gray-900 border-gray-600 rounded focus:outline-none focus:ring-2 focus:ring-blue-500 placeholder-gray-500"
                  aria-label="Secret API Twitter"
                />
                <input
                  v-model="twitterAccessToken"
                  placeholder="Jeton d'accès Twitter"
                  class="w-full p-3 border bg-gray-900 border-gray-600 rounded focus:outline-none focus:ring-2 focus:ring-blue-500 placeholder-gray-500"
                  aria-label="Jeton d'accès Twitter"
                />
                <input
                  v-model="twitterAccessSecret"
                  placeholder="Secret d'accès Twitter"
                  class="w-full p-3 border bg-gray-900 border-gray-600 rounded focus:outline-none focus:ring-2 focus:ring-blue-500 placeholder-gray-500"
                  aria-label="Secret d'accès Twitter"
                />
              </div>
            </transition>
          </div>

          <div class="mb-6">
            <h3
              class="text-lg font-semibold text-gray-300 mb-3 cursor-pointer flex items-center justify-between"
              @click="toggleThreadsTokens"
            >
              <span>Tokens API Threads</span>
              <span v-if="showThreadsTokens">⬆️</span>
              <span v-else>⬇️</span>
            </h3>

            <transition name="slide">
              <div v-if="showThreadsTokens" class="space-y-4">
                <input
                  v-model="threadsAccessToken"
                  placeholder="Jeton d'accès Threads"
                  class="w-full p-3 border bg-gray-900 border-gray-600 rounded focus:outline-none focus:ring-2 focus:ring-blue-500 placeholder-gray-500"
                  aria-label="Jeton d'accès Threads"
                />
                <input
                  v-model="threadsUserID"
                  placeholder="ID Utilisateur Threads"
                  class="w-full p-3 border bg-gray-900 border-gray-600 rounded focus:outline-none focus:ring-2 focus:ring-blue-500 placeholder-gray-500"
                  aria-label="ID Utilisateur Threads"
                />
              </div>
            </transition>
          </div>

          <div class="mb-6">
            <label class="flex items-center">
              <input
                type="checkbox"
                v-model="allowExceedLimits"
                class="form-checkbox h-5 w-5 text-blue-600"
              />
              <span class="ml-2 text-gray-300"
                >Autoriser le dépassement des limites de caractères
                (Premium)</span
              >
            </label>
          </div>

          <div class="flex justify-center mt-6 mb-4">
            <p class="text-sm text-gray-400 text-center">
              Pour obtenir les tokens d'accès, veuillez vous référer à la
              <a
                @click.prevent="openDocumentation"
                class="text-blue-500 hover:underline cursor-pointer"
                >documentation</a
              >.
            </p>
          </div>

          <div class="flex justify-between space-x-3 mt-6">
            <button
              @click="saveConfig"
              class="flex items-center bg-blue-500 hover:bg-blue-600 text-white font-semibold py-2 px-6 rounded transition"
              :disabled="isSaving"
            >
              <span v-if="isSaving">
                <svg
                  class="animate-spin h-5 w-5 mr-2"
                  xmlns="http://www.w3.org/2000/svg"
                  fill="none"
                  viewBox="0 0 24 24"
                >
                  <circle
                    class="opacity-25"
                    cx="12"
                    cy="12"
                    r="10"
                    stroke="currentColor"
                    stroke-width="4"
                  ></circle>
                  <path
                    class="opacity-75"
                    fill="currentColor"
                    d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z"
                  ></path>
                </svg>
                Sauvegarde...
              </span>
              <span v-else> <i class="fas fa-save mr-2"></i> Sauvegarder </span>
            </button>
            <button
              @click="toggleSettings"
              class="flex items-center bg-red-500 hover:bg-red-600 text-white font-semibold py-2 px-6 rounded transition"
            >
              <i class="fas fa-times mr-2"></i> Annuler
            </button>
          </div>
        </div>
      </div>
    </transition>

    <footer
      class="w-full bg-gray-800 text-gray-400 text-sm p-4 mt-auto flex flex-col items-center"
    >
      <div class="flex space-x-6">
        <a
          @click.prevent="openTwitterX"
          class="cursor-pointer text-blue-500 hover:text-blue-400 flex items-center"
          title="TwitterX"
        >
          <i class="fa-brands fa-x-twitter mr-2"></i> TwitterX
        </a>
        <a
          @click.prevent="openThreads"
          class="cursor-pointer text-blue-500 hover:text-blue-400 flex items-center"
          title="Threads"
        >
          <i class="fa-brands fa-threads mr-2"></i> Threads
        </a>
        <a
          @click.prevent="openBuyMeACoffee"
          class="cursor-pointer text-blue-500 hover:text-blue-400 flex items-center"
          title="Buy me a coffee"
        >
          <i class="fas fa-coffee mr-2"></i> Buy me a Coffee
        </a>
        <a
          @click.prevent="openGithub"
          class="cursor-pointer text-blue-500 hover:text-blue-400 flex items-center"
          title="Github"
        >
          <i class="fab fa-github mr-2"></i> Github
        </a>
        <a
          @click.prevent="openDocumentation"
          class="cursor-pointer text-blue-500 hover:text-blue-400 flex items-center"
          title="Documentation"
        >
          <i class="fas fa-book mr-2"></i> Documentation
        </a>
      </div>
    </footer>
  </div>
</template>

<script>
import {
  PostToTwitter,
  PostToThreads,
  PostToBoth,
  SaveConfig,
  GetConfig,
} from "../wailsjs/go/main/App";
import { BrowserOpenURL } from "../wailsjs/runtime/runtime.js";
import Toastify from "toastify-js";
import "toastify-js/src/toastify.css";

export default {
  data() {
    return {
      content: "",
      showSettings: false,
      showTwitterTokens: false,
      showThreadsTokens: false,
      isLoading: false,
      isSaving: false,
      twitterAPIKey: "",
      twitterAPISecret: "",
      twitterAccessToken: "",
      twitterAccessSecret: "",
      threadsAccessToken: "",
      threadsUserID: "",
      showConfirmation: false,
      confirmationAction: null,
      characterLimits: {
        twitter: 280,
        threads: 500,
      },
      allowExceedLimits: false,
    };
  },
  computed: {
    isConfigValid() {
      return (
        this.twitterAPIKey &&
        this.twitterAPISecret &&
        this.twitterAccessToken &&
        this.twitterAccessSecret &&
        this.threadsAccessToken &&
        this.threadsUserID
      );
    },
    highlightedContent() {
      return this.content.replace(
        /(#[a-zA-Z0-9_]+)/g,
        '<span class="text-blue-500">$1</span>'
      );
    },
  },
  methods: {
    toggleSettings() {
      this.showSettings = !this.showSettings;
    },
    toggleTwitterTokens() {
      this.showTwitterTokens = !this.showTwitterTokens;
    },
    toggleThreadsTokens() {
      this.showThreadsTokens = !this.showThreadsTokens;
    },
    preparePost(action) {
      if (!this.content) {
        this.showError("Le contenu est vide.");
        return;
      }

      // Vérification des limites de caractères
      if (!this.allowExceedLimits) {
        if (
          action === "twitter" &&
          this.content.length > this.characterLimits.twitter
        ) {
          this.showError(
            `Le contenu dépasse la limite de ${this.characterLimits.twitter} caractères pour Twitter X.`
          );
          return;
        }
        if (
          action === "threads" &&
          this.content.length > this.characterLimits.threads
        ) {
          this.showError(
            `Le contenu dépasse la limite de ${this.characterLimits.threads} caractères pour Threads.`
          );
          return;
        }
        if (
          action === "both" &&
          (this.content.length > this.characterLimits.twitter ||
            this.content.length > this.characterLimits.threads)
        ) {
          this.showError(
            `Le contenu dépasse les limites de caractères pour Twitter X et/ou Threads.`
          );
          return;
        }
      }

      this.confirmationAction = action;
      this.showConfirmation = true;
    },
    confirmPost() {
      this.showConfirmation = false;
      switch (this.confirmationAction) {
        case "twitter":
          this.postToTwitter();
          break;
        case "threads":
          this.postToThreads();
          break;
        case "both":
          this.postToBoth();
          break;
        default:
          break;
      }
      this.confirmationAction = null;
    },
    closeConfirmation() {
      this.showConfirmation = false;
      this.confirmationAction = null;
    },
    postToTwitter() {
      this.isLoading = true;
      PostToTwitter(this.content)
        .then((response) => this.showSuccess(response))
        .catch((err) =>
          this.showError(
            err.message || "Erreur lors de la publication sur Twitter X."
          )
        )
        .finally(() => {
          this.isLoading = false;
        });
    },
    postToThreads() {
      this.isLoading = true;
      PostToThreads(this.content)
        .then((response) => this.showSuccess(response))
        .catch((err) =>
          this.showError(
            err.message || "Erreur lors de la publication sur Threads."
          )
        )
        .finally(() => {
          this.isLoading = false;
        });
    },
    postToBoth() {
      this.isLoading = true;
      PostToBoth(this.content)
        .then((response) => this.showSuccess(response))
        .catch((err) =>
          this.showError(
            err.message ||
              "Erreur lors de la publication sur les deux plateformes."
          )
        )
        .finally(() => {
          this.isLoading = false;
        });
    },
    saveConfig() {
      this.isSaving = true;
      SaveConfig(
        this.twitterAPIKey,
        this.twitterAPISecret,
        this.twitterAccessToken,
        this.twitterAccessSecret,
        this.threadsAccessToken,
        this.threadsUserID
      )
        .then((response) => {
          this.showSuccess(response);
          this.showSettings = false;
        })
        .catch((err) =>
          this.showError(
            err.message || "Erreur lors de la sauvegarde des configurations."
          )
        )
        .finally(() => {
          this.isSaving = false;
        });
    },
    getConfig() {
      GetConfig()
        .then((config) => {
          this.twitterAPIKey = config.twitterAPIKey;
          this.twitterAPISecret = config.twitterAPISecret;
          this.twitterAccessToken = config.twitterAccessToken;
          this.twitterAccessSecret = config.twitterAccessSecret;
          this.threadsAccessToken = config.threadsAccessToken;
          this.threadsUserID = config.threadsUserID;
        })
        .catch((err) => {
          this.showError(
            err.message || "Erreur lors du chargement de la configuration."
          );
        });
    },
    showSuccess(message) {
      Toastify({
        text: message,
        duration: 3000,
        backgroundColor: "linear-gradient(to right, #00b09b, #96c93d)",
      }).showToast();
    },
    showError(message) {
      Toastify({
        text: message,
        duration: 3000,
        backgroundColor: "linear-gradient(to right, #ff5f6d, #ffc371)",
      }).showToast();
    },
    openDocumentation() {
      BrowserOpenURL(
        "https://benoitpetit.notion.site/Documentation-pour-l-application-2tpost-1266e0e9df2b80f3afbfc98e3e08c7c6"
      );
    },
    openTwitterX() {
      BrowserOpenURL("https://twitter.com/devbyben");
    },
    openThreads() {
      BrowserOpenURL("https://threads.net/benoitpetit.dev");
    },
    openGithub() {
      BrowserOpenURL("https://github.com/benoitpetit");
    },
    openBuyMeACoffee() {
      BrowserOpenURL("https://www.buymeacoffee.com/benoitpetit");
    },
    isExceeded(platform) {
      if (platform === "twitter") {
        return (
          !this.allowExceedLimits &&
          this.content.length > this.characterLimits.twitter
        );
      }
      if (platform === "threads") {
        return (
          !this.allowExceedLimits &&
          this.content.length > this.characterLimits.threads
        );
      }
      return false;
    },
  },
  mounted() {
    this.getConfig();
  },
};
</script>