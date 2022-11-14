package chrome

import (
	"time"

	"github.com/tebeka/selenium"
)

func (c *ChromeClient) InspectForms() error {
	inputs, err := c.driver.FindElements(selenium.ByCSSSelector, "input")
	if err != nil {
		return err
	}

	buttons, err := c.driver.FindElements(selenium.ByCSSSelector, "button")
	if err != nil {
		return err
	}

	clickables, err := c.driver.FindElements(selenium.ByCSSSelector, "[onclick]:not([onclick=''])")
	if err != nil {
		return err
	}

	for _, input := range inputs {
		input.SendKeys(c.Payload)
	}

	for _, button := range buttons {
		button.Click()
	}

	for _, clickable := range clickables {
		clickable.Click()
	}

	return nil
}

func (c *ChromeClient) Inspect(url string) (ChromeInspectResult, error) {
	defer c.service.Stop()

	if err := c.driver.Get(url); err != nil {
		return c.result, err
	}

	time.Sleep(1 * time.Second)

	if err := c.TakeScreenshot(url); err != nil {
		return c.result, err
	}

	if err := c.InspectForms(); err != nil {
		return c.result, err
	}

	return c.result, nil
}
