package chrome

import (
	"runtime"
	"strings"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

const driverPath = "./drivers/chrome"

type ChromeInspectResult struct {
	Malicious bool
}
type ChromeClient struct {
	driver  selenium.WebDriver
	service *selenium.Service
	result  ChromeInspectResult
}

func NewChromeClient() (*ChromeClient, error) {
	service, err := selenium.NewChromeDriverService(strings.Join([]string{driverPath, runtime.GOARCH, runtime.GOOS}, "/"), 4444)
	if err != nil {
		return nil, err
	}

	caps := selenium.Capabilities{}
	caps.AddChrome(chrome.Capabilities{Args: []string{
		"window-size=1920x1080",
		"disable-gpu",
		// "--headless",
	}})

	driver, err := selenium.NewRemote(caps, "")
	if err != nil {
		return nil, err
	}

	return &ChromeClient{
		driver:  driver,
		service: service,
		result:  ChromeInspectResult{},
	}, nil
}

func (c *ChromeClient) InspectForms() error {
	forms, err := c.driver.FindElements(selenium.ByCSSSelector, "form")
	if err != nil {
		return err
	}

	for _, form := range forms {
		inputs, err := form.FindElements(selenium.ByCSSSelector, "input")
		if err != nil {
			continue
		}

		for _, input := range inputs {
			input.SendKeys("test")
		}

		form.Submit()
	}

	return nil
}

func (c *ChromeClient) Inspect(url string) (ChromeInspectResult, error) {
	defer c.service.Stop()

	if err := c.driver.Get(url); err != nil {
		return c.result, err
	}

	if err := c.InspectForms(); err != nil {
		return c.result, err
	}

	return c.result, nil
}
