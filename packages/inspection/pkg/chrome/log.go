package chrome

import (
	"encoding/json"

	"github.com/tebeka/selenium/log"
)

type NetworkLogMessage struct {
	Method string `json:"method"`
	Params struct {
		// Common
		RequestId    string  `json:"requestId"`
		FrameId      string  `json:"frameId"`
		HasExtraInfo bool    `json:"hasExtraInfo"`
		Timestamp    float64 `json:"timestamp"`
		Type         string  `json:"type"`

		// Request Network.requestWillBeSent
		DocumentURL    string `json:"documentURL"`
		HasUserGesture bool   `json:"hasUserGesture"`
		Initiator      struct {
			Stack struct {
				CallFrames []struct {
					ColumnNumber int    `json:"columnNumber"`
					FunctionName string `json:"functionName"`
					LineNumber   int    `json:"lineNumber"`
					URL          string `json:"url"`
					ScriptId     string `json:"scriptId"`
				} `json:"callFrames"`
			} `json:"stack"`
			Type string `json:"type"`
		} `json:"initiator"`
		LoaderId             string `json:"loaderId"`
		RedirectHasExtraInfo bool   `json:"redirectHasExtraInfo"`
		Request              struct {
			HasPostData    bool              `json:"hasPostData"`
			Headers        map[string]string `json:"headers"`
			PostData       string            `json:"postData"`
			IsSameSite     bool              `json:"isSameSite"`
			Method         string            `json:"method"`
			ReferrerPolicy string            `json:"referrerPolicy"`
			URL            string            `json:"url"`
		} `json:"request"`

		// Response Network.responseReceived
		Response struct {
			ConnectionID      int                `json:"connectionId"`
			ConnectionReused  bool               `json:"connectionReused"`
			EncodedDataLength float64            `json:"encodedDataLength"`
			FromDiskCache     bool               `json:"fromDiskCache"`
			FromServiceWorker bool               `json:"fromServiceWorker"`
			Headers           map[string]string  `json:"headers"`
			MimeType          string             `json:"mimeType"`
			Protocol          string             `json:"protocol"`
			RemoteIPAddress   string             `json:"remoteIPAddress"`
			RemotePort        int                `json:"remotePort"`
			ResponseTime      float64            `json:"responseTime"`
			SecurityDetails   interface{}        `json:"securityDetails"`
			SecurityState     string             `json:"securityState"`
			Status            int                `json:"status"`
			StatusText        string             `json:"statusText"`
			Timing            map[string]float64 `json:"timing"`
			URL               string             `json:"url"`
		} `json:"response"`
	} `json:"params"`
}

type NetworkLog struct {
	Message NetworkLogMessage `json:"message"`
	WebView string            `json:"webview"`
}

func ParseLogFromJson(data []byte) (*NetworkLog, error) {
	var log *NetworkLog
	err := json.Unmarshal(data, &log)
	if err != nil {
		return nil, err
	}
	return log, nil
}

func (n *NetworkLog) GetRequestID() string {
	return n.Message.Params.RequestId
}

func (n *NetworkLog) GetFrameID() string {
	return n.Message.Params.FrameId
}

func (n *NetworkLog) GetRequestURL() string {
	return n.Message.Params.Request.URL
}

func (n *NetworkLog) IsRequestWillBeSent() bool {
	return n.Message.Method == "Network.requestWillBeSent"
}

func (n *NetworkLog) IsResponseReceived() bool {
	return n.Message.Method == "Network.responseReceived"
}

type NetworkLogEntries []*NetworkLog

func NewNetworkLogEntries() *NetworkLogEntries {
	return &NetworkLogEntries{}
}

func (n *NetworkLogEntries) ParseFromDriverLogs(l []log.Message) {
	for _, entry := range l {
		log, err := ParseLogFromJson([]byte(entry.Message))
		if err != nil {
			continue
		}
		n.Append(log)
	}
}

func (n *NetworkLogEntries) Len() int {
	return len(*n)
}

func (n *NetworkLogEntries) Append(l *NetworkLog) {
	*n = append(*n, l)
}

func (n *NetworkLogEntries) Filter(filter func(*NetworkLog) bool) *NetworkLogEntries {
	filtered := NewNetworkLogEntries()
	for _, entry := range *n {
		if filter(entry) {
			filtered.Append(entry)
		}
	}
	return filtered
}

func (n *NetworkLogEntries) FilterByRequestId(requestId string) *NetworkLogEntries {
	return n.Filter(func(entry *NetworkLog) bool {
		return entry.GetRequestID() == requestId
	})
}

func (n *NetworkLogEntries) Map(mapper func(*NetworkLog) *NetworkLog) *NetworkLogEntries {
	mapped := NewNetworkLogEntries()
	for _, entry := range *n {
		mapped.Append(mapper(entry))
	}
	return mapped
}
