package chrome

import (
	"os"
	"runtime"
	"strings"

	"github.com/google/uuid"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"github.com/tebeka/selenium/log"
)

type ChromeInspectResult struct {
	Screenshot []byte
	SendingTo  []string
	Reasons    []string
}
type ChromeClient struct {
	enableNetwork bool
	driver        selenium.WebDriver
	service       *selenium.Service
	result        ChromeInspectResult
	Payload       string
}

type ChromeClientOptions struct {
	ChromeArgs []string
}

func NewChromeClient(options ChromeClientOptions) (*ChromeClient, error) {
	c := &ChromeClient{
		enableNetwork: true,
		result: ChromeInspectResult{
			Screenshot: []byte{},
			SendingTo:  []string{},
			Reasons:    []string{},
		},
		Payload: uuid.New().String(),
	}

	driverPath := os.Getenv("CHROME_DRIVER_PATH")
	if driverPath == "" {
		driverPath = strings.Join([]string{"./drivers/chrome", runtime.GOARCH, runtime.GOOS}, "/")
	}
	service, err := selenium.NewChromeDriverService(driverPath, 4444)
	if err != nil {
		return nil, err
	}

	c.service = service

	caps := selenium.Capabilities{}
	caps.AddChrome(chrome.Capabilities{
		Args: options.ChromeArgs,
		PerfLoggingPrefs: &chrome.PerfLoggingPreferences{
			EnableNetwork: &c.enableNetwork,
		},
	})
	caps.SetLogLevel(log.Performance, log.All)

	driver, err := selenium.NewRemote(caps, "")
	if err != nil {
		return nil, err
	}

	c.driver = driver

	return c, nil
}

func (c *ChromeClient) Run(url string) (ChromeInspectResult, error) {
	return c.Inspect(url)
}
