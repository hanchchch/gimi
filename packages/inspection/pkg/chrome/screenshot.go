package chrome

func (c *ChromeClient) TakeScreenshot(url string) error {
	b, err := c.driver.Screenshot()
	if err != nil {
		return err
	}

	c.result.Screenshot = b

	return nil
}
