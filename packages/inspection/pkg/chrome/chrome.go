package chrome

import (
	"os"
	"runtime"
	"strings"

	"github.com/google/uuid"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

type ChromeInspectResult struct {
	Malicious  bool
	Screenshot []byte
}
type ChromeClient struct {
	driver  selenium.WebDriver
	service *selenium.Service
	result  ChromeInspectResult
	Payload string
}

type ChromeClientOptions struct {
	ChromeArgs []string
}

func NewChromeClient(options ChromeClientOptions) (*ChromeClient, error) {
	driverPath := os.Getenv("CHROME_DRIVER_PATH")
	if driverPath == "" {
		driverPath = strings.Join([]string{"./drivers/chrome", runtime.GOARCH, runtime.GOOS}, "/")
	}
	service, err := selenium.NewChromeDriverService(driverPath, 4444)
	if err != nil {
		return nil, err
	}

	caps := selenium.Capabilities{}
	caps.AddChrome(chrome.Capabilities{Args: options.ChromeArgs})

	driver, err := selenium.NewRemote(caps, "")
	if err != nil {
		return nil, err
	}

	return &ChromeClient{
		driver:  driver,
		service: service,
		result:  ChromeInspectResult{},
		Payload: uuid.New().String(),
	}, nil
}

func (c *ChromeClient) Run(url string) (ChromeInspectResult, error) {
	return c.Inspect(url)
}
