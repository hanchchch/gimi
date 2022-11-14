package urls

import "strings"

func EnsureProtocol(url string) string {
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		return "http://" + url
	}
	return url
}

func TrimProtocol(url string) string {
	return strings.Split(url, "://")[1]
}

func GetHost(url string) string {
	return strings.Split(TrimProtocol(url), "/")[0]
}
