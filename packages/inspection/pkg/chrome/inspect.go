package chrome

import (
	"time"

	"github.com/tebeka/selenium"
)

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

	time.Sleep(1 * time.Second)

	if err := c.TakeScreenshot(); err != nil {
		return c.result, err
	}

	if err := c.InspectForms(); err != nil {
		return c.result, err
	}

	return c.result, nil
}
