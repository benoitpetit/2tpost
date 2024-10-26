package utils

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	TwitterApiKey       string `json:"twitter_api_key"`
	TwitterApiSecret    string `json:"twitter_api_secret"`
	TwitterAccessToken  string `json:"twitter_access_token"`
	TwitterAccessSecret string `json:"twitter_access_secret"`
	ThreadsAccessToken  string `json:"threads_access_token"`
	ThreadsUserID       string `json:"threads_user_id"`
}

var AppConfig Config

func LoadConfig() {
	configFilePath := "config.json"

	file, err := os.Open(configFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("Le fichier config.json n'existe pas. Un fichier vide sera créé.")
			AppConfig = Config{}
			return
		}
		log.Printf("Erreur lors de l'ouverture du fichier config: %v\n", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&AppConfig)
	if err != nil {
		log.Printf("Erreur lors du décodage du fichier config: %v\n", err)
		AppConfig = Config{}
	}
}

func SaveConfig() error {
	file, err := os.Create("config.json")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(AppConfig)
}
