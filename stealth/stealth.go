package stealth

import "github.com/go-rod/rod"

func Apply(page *rod.Page) {
	page.MustEvaluate(rod.Eval(`() => {
		Object.defineProperty(navigator, 'webdriver', {
			get: () => undefined
		});
	}`))
}

