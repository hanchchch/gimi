package chrome

import (
	"encoding/json"
	"log"
	"strings"
	"time"

	"github.com/tebeka/selenium"
	slog "github.com/tebeka/selenium/log"
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

func (c *ChromeClient) InspectNetwork() error {
	l, err := c.driver.Log(slog.Performance)
	if err != nil {
		return err
	}

	logs := NewNetworkLogEntries()
	logs.ParseFromDriverLogs(l)

	for _, payloadLog := range *logs.Filter(func(entry *NetworkLog) bool {
		return entry.IsRequestWillBeSent()
	}).Filter(func(entry *NetworkLog) bool {
		if b, err := json.Marshal(entry); err != nil {
			return false
		} else {
			return strings.Contains(string(b), c.Payload)
		}
	}) {
		c.result.SendingTo = append(c.result.SendingTo, payloadLog.GetRequestURL())
		for _, responseLog := range *logs.FilterByRequestId(payloadLog.GetRequestID()).Filter(func(entry *NetworkLog) bool {
			return entry.IsResponseReceived()
		}) {
			status := responseLog.Message.Params.Response.Status
			log.Printf("captured input data posting - %d", status)
			if 200 <= status && status < 400 {
				c.result.Reasons = append(c.result.Reasons, "Payload delivered successfully without validation.")
			}
		}
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

	if err := c.InspectNetwork(); err != nil {
		return c.result, err
	}

	return c.result, nil
}
