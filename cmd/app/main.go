package main

import (
	"log"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"

	"linkedin-automation/config"
	"linkedin-automation/stealth"
	"linkedin-automation/utils"
)

func main() {
	cfg := config.Load()

	// Launch browser
	url := launcher.New().
		Headless(cfg.Headless).
		Set("user-agent", cfg.UserAgent).
		MustLaunch()

	browser := rod.New().ControlURL(url).MustConnect()
	defer browser.Close()

	page := browser.MustPage("https://example.com")

	// Apply stealth
    stealth.Apply(page);

	// Simulate human wait
	utils.HumanDelay(1200, 2500)

	log.Println("PoC executed successfully")
}
