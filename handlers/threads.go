package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"time"

	"2tpost/utils"
)

type ThreadsResponse struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

func PostToThreads(content string) error {
	mediaParams := url.Values{
		"media_type":   {"TEXT"},
		"text":         {content},
		"access_token": {utils.AppConfig.ThreadsAccessToken},
	}

	mediaResponse, err := http.PostForm("https://graph.threads.net/v1.0/"+utils.AppConfig.ThreadsUserID+"/threads", mediaParams)
	if err != nil {
		log.Println("Erreur lors de la création du conteneur de média :", err)
		return err
	}
	defer mediaResponse.Body.Close()

	var threadsResp ThreadsResponse
	if err := json.NewDecoder(mediaResponse.Body).Decode(&threadsResp); err != nil {
		log.Println("Erreur lors du décodage de la réponse :", err)
		return err
	}

	time.Sleep(5 * time.Second)

	publishParams := url.Values{
		"creation_id":  {threadsResp.ID},
		"access_token": {utils.AppConfig.ThreadsAccessToken},
	}

	publishResponse, err := http.PostForm("https://graph.threads.net/v1.0/"+utils.AppConfig.ThreadsUserID+"/threads_publish", publishParams)
	if err != nil {
		log.Println("Erreur lors de la publication sur Threads :", err)
		return err
	}
	defer publishResponse.Body.Close()

	var publishResp ThreadsResponse
	if err := json.NewDecoder(publishResponse.Body).Decode(&publishResp); err != nil {
		log.Println("Erreur lors du décodage de la réponse de la publication :", err)
		return err
	}

	return nil
}
