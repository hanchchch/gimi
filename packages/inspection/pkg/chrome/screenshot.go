package chrome

import "io/ioutil"

func (c *ChromeClient) TakeScreenshot() error {
	b, err := c.driver.Screenshot()
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile("screenshot.png", b, 0644); err != nil {
		return err
	}

	return nil
}
