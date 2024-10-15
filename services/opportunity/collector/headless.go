package collector

import (
	"github.com/tebeka/selenium"
)

func Headless() {
	// Connect to the WebDriver instance running locally.
	caps := selenium.Capabilities{"browserName": "chrome"}
	wd, err := selenium.NewRemote(caps, "")
	if err != nil {
		panic(err)
	}
	defer wd.Quit()

	// Navigate to the page you want to interact with
	if err := wd.Get("http://example.com"); err != nil {
		panic(err)
	}

	// Find a button by its ID and click it
	btn, err := wd.FindElement(selenium.ByID, "myButton")
	if err != nil {
		panic(err)
	}
	if err := btn.Click(); err != nil {
		panic(err)
	}

	// You can now interact with the new elements or dialogs opened as a result of the click
	// ...

	// Don't forget to close the page and the WebDriver.
	wd.Quit()
}
