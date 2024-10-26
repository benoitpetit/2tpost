package main

import (
	"2tpost/handlers"
	"embed"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"context"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

// Config struct to store tokens in the config.json file
type Config struct {
	TwitterAPIKey       string `json:"twitterAPIKey"`
	TwitterAPISecret    string `json:"twitterAPISecret"`
	TwitterAccessToken  string `json:"twitterAccessToken"`
	TwitterAccessSecret string `json:"twitterAccessSecret"`
	ThreadsAccessToken  string `json:"threadsAccessToken"`
	ThreadsUserID       string `json:"threadsUserID"`
}

// getConfigPath returns the absolute path to config.json
func getConfigPath() (string, error) {
	configPath := "config.json"

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		exePath, err := os.Executable()
		if err != nil {
			return "", fmt.Errorf("error retrieving executable path: %v", err)
		}
		exeDir := filepath.Dir(exePath)
		configPath = filepath.Join(exeDir, "config.json")
	}

	return configPath, nil
}

// LoadConfig loads the configuration from the config.json file
func LoadConfig() (Config, error) {
	var config Config
	configPath, err := getConfigPath()
	if err != nil {
		return config, err
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return config, nil
	}

	file, err := os.Open(configPath)
	if err != nil {
		return config, fmt.Errorf("error opening config file: %v", err)
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return config, fmt.Errorf("error reading config file: %v", err)
	}

	err = json.Unmarshal(bytes, &config)
	if err != nil {
		return config, fmt.Errorf("error parsing config file: %v", err)
	}

	return config, nil
}

// SaveConfig saves the configuration to the config.json file
func (a *App) SaveConfig(twitterAPIKey, twitterAPISecret, twitterAccessToken, twitterAccessSecret, threadsAccessToken, threadsUserID string) (string, error) {
	config := Config{
		TwitterAPIKey:       twitterAPIKey,
		TwitterAPISecret:    twitterAPISecret,
		TwitterAccessToken:  twitterAccessToken,
		TwitterAccessSecret: twitterAccessSecret,
		ThreadsAccessToken:  threadsAccessToken,
		ThreadsUserID:       threadsUserID,
	}

	file, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return "", fmt.Errorf("error formatting config: %v", err)
	}

	configPath, err := getConfigPath()
	if err != nil {
		return "", err
	}

	err = ioutil.WriteFile(configPath, file, 0644)
	if err != nil {
		return "", fmt.Errorf("error writing config file: %v", err)
	}

	return "Configuration saved successfully", nil
}

// GetConfig retrieves the current configuration to send to the frontend
func (a *App) GetConfig() (Config, error) {
	config, err := LoadConfig()
	if err != nil {
		return Config{}, err
	}
	return config, nil
}

// PostToTwitter posts content to Twitter/X
func (a *App) PostToTwitter(content string) (string, error) {
	err := handlers.PostToTwitter(content)
	if err != nil {
		return "", fmt.Errorf("error posting to Twitter/X: %v", err)
	}
	return "Post successfully published on Twitter/X.", nil
}

// PostToThreads posts content to Threads
func (a *App) PostToThreads(content string) (string, error) {
	err := handlers.PostToThreads(content)
	if err != nil {
		return "", fmt.Errorf("error posting to Threads: %v", err)
	}
	return "Post successfully published on Threads.", nil
}

// PostToBoth posts content to both Twitter/X and Threads
func (a *App) PostToBoth(content string) (string, error) {
	errTwitter := handlers.PostToTwitter(content)
	errThreads := handlers.PostToThreads(content)

	if errTwitter != nil || errThreads != nil {
		return "", fmt.Errorf("error posting:\nTwitter/X: %v\nThreads: %v", errTwitter, errThreads)
	}
	return "Post successfully published on Twitter/X and Threads.", nil
}

func loadIcon(path string) []byte {
	icon, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("error loading icon: %v\n", err)
		return nil
	}
	return icon
}

func main() {
	app := &App{}

	err := wails.Run(&options.App{
		Title:         "2tpost",
		Width:         850,
		Height:        770,
		Fullscreen:    false,
		DisableResize: false,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		Windows: &windows.Options{
			WebviewIsTransparent: true,
			IsZoomControlEnabled: false,
			Theme:                windows.SystemDefault,
		},
		Mac: &mac.Options{
			TitleBar:            mac.TitleBarHiddenInset(),
			WindowIsTranslucent: true,
			About: &mac.AboutInfo{
				Title:   "2tpost",
				Message: "Post content to Twitter and Threads with one click.",
				Icon:    loadIcon("../appicon.png"),
			},
		},
		Linux: &linux.Options{
			ProgramName:         "2tpost",
			Icon:                loadIcon("../appicon.png"),
			WindowIsTranslucent: true,
		},
		OnStartup: func(ctx context.Context) {
			fmt.Println("Application started")
		},
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
