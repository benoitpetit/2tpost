package handlers

import (
	"context"
	"errors"
	"net/url"
	"time"

	"2tpost/utils"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

type Metadata struct {
	Title       string
	Description string
	Image       string
}

// extrait les métadonnées d'une URL et retourne une structure Metadata
func FetchMetadataFromURL(targetURL string) (*Metadata, error) {
	if targetURL == "" {
		return nil, errors.New("aucune URL fournie")
	}

	// Contexte avec délai d'attente
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Lancer le navigateur en mode headless
	u := launcher.New().Headless(true).MustLaunch()

	browser := rod.New().ControlURL(u).Context(ctx)
	if err := browser.Connect(); err != nil {
		utils.ErrorLogger.Println("Erreur lors de la connexion au navigateur :", err)
		return nil, errors.New("erreur lors de la connexion au navigateur")
	}
	defer browser.Close()

	// Créer une nouvelle page avec le contexte et naviguer vers l'URL
	page := browser.MustPage().Context(ctx)
	if err := page.Navigate(targetURL); err != nil {
		utils.ErrorLogger.Println("Erreur lors de la navigation :", err)
		return nil, errors.New("erreur lors de la navigation vers l'URL")
	}
	page.WaitLoad()

	// Extraire les métadonnées via JavaScript
	metadata, err := page.Evaluate(rod.Eval(`() => {
		const getMetaTagContent = (property) =>
			document.querySelector('meta[property="' + property + '"]')?.getAttribute("content") ||
			document.querySelector('meta[name="' + property + '"]')?.getAttribute("content");
		return {
			title: document.title,
			description: getMetaTagContent("og:description") || document.querySelector('meta[name="description"]')?.getAttribute("content"),
			image: getMetaTagContent("og:image") || getMetaTagContent("twitter:image"),
		};
	}`))
	if err != nil {
		utils.ErrorLogger.Println("Erreur lors de l'extraction des métadonnées :", err)
		return nil, errors.New("erreur lors de l'extraction des métadonnées")
	}

	// Extraire les valeurs des métadonnées
	var title, description, image string
	metaDataMap := metadata.Value.Map()

	if titleVal, ok := metaDataMap["title"]; ok {
		title = titleVal.Str()
	}

	if descVal, ok := metaDataMap["description"]; ok {
		description = descVal.Str()
	}

	if imageVal, ok := metaDataMap["image"]; ok {
		image = imageVal.Str()
	}

	// Si l'image est un chemin relatif, la convertir en URL absolue
	parsedURL, err := url.Parse(targetURL)
	if err == nil && image != "" {
		imageURL, err := url.Parse(image)
		if err == nil && !imageURL.IsAbs() {
			image = parsedURL.ResolveReference(imageURL).String()
		}
	}

	// Retourner les métadonnées
	return &Metadata{
		Title:       title,
		Description: description,
		Image:       image,
	}, nil
}
