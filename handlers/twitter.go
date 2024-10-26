package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strings"

	"2tpost/utils"

	"github.com/dghubble/oauth1"
)

type TweetResponse struct {
	Data struct {
		ID   string `json:"id"`
		Text string `json:"text"`
	} `json:"data"`
	Errors []struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
	} `json:"errors"`
}

func PostToTwitter(content string) error {
	if strings.TrimSpace(content) == "" {
		return errors.New("le contenu du tweet est vide")
	}
	if len(content) > 280 {
		return errors.New("le contenu du tweet dépasse la limite de 280 caractères")
	}

	config := oauth1.NewConfig(utils.AppConfig.TwitterApiKey, utils.AppConfig.TwitterApiSecret)
	token := oauth1.NewToken(utils.AppConfig.TwitterAccessToken, utils.AppConfig.TwitterAccessSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	twitterAPIURL := "https://api.twitter.com/2/tweets"
	tweetData := map[string]interface{}{
		"text": content,
	}

	jsonData, err := json.Marshal(tweetData)
	if err != nil {
		log.Println("Erreur lors de l'encodage des données du tweet :", err)
		return err
	}

	req, err := http.NewRequest("POST", twitterAPIURL, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("Erreur lors de la création de la requête :", err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Println("Erreur lors de l'envoi de la requête :", err)
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Erreur lors de la lecture de la réponse :", err)
		return err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		log.Printf("Erreur lors de la réponse de l'API Twitter : %s", resp.Status)
		return errors.New(string(body))
	}

	var tweetResp TweetResponse
	if err := json.Unmarshal(body, &tweetResp); err != nil {
		log.Println("Erreur lors du parsing de la réponse JSON :", err)
		return err
	}

	if len(tweetResp.Errors) > 0 {
		log.Printf("Erreur API Twitter : %s", tweetResp.Errors[0].Message)
		return errors.New(tweetResp.Errors[0].Message)
	}

	return nil
}
