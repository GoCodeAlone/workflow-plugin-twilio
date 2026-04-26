package internal

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	twilio "github.com/twilio/twilio-go"
	twilioClient "github.com/twilio/twilio-go/client"
)

// urlRewriteTransport is an http.RoundTripper that rewrites the scheme+host of
// outbound requests to a configured base URL, enabling redirection of Twilio API
// calls to a local mock server.
type urlRewriteTransport struct {
	base      string // e.g. "http://localhost:19051"
	Transport http.RoundTripper
}

func (t *urlRewriteTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	clone := req.Clone(req.Context())
	// Replace scheme+host with the mock base, preserving the path+query.
	parts := strings.SplitN(t.base, "://", 2)
	if len(parts) == 2 {
		clone.URL.Scheme = parts[0]
		clone.URL.Host = parts[1]
	}
	tr := t.Transport
	if tr == nil {
		tr = http.DefaultTransport
	}
	return tr.RoundTrip(clone)
}

// twilioModule creates a Twilio REST client and registers it.
type twilioModule struct {
	name   string
	config map[string]any
}

func newTwilioModule(name string, config map[string]any) (*twilioModule, error) {
	return &twilioModule{name: name, config: config}, nil
}

// Init creates the Twilio REST client and registers it in the global registry.
func (m *twilioModule) Init() error {
	accountSid, _ := m.config["accountSid"].(string)
	authToken, _ := m.config["authToken"].(string)
	apiKey, _ := m.config["apiKey"].(string)
	apiSecret, _ := m.config["apiSecret"].(string)
	baseURL, _ := m.config["baseURL"].(string)
	if missingTwilioCredentials(accountSid, authToken, apiKey, apiSecret) {
		if configBool(m.config, "optional") {
			UnregisterClient(m.name)
			return nil
		}
		return fmt.Errorf("twilio.provider %q: accountSid and authToken, or accountSid, apiKey, and apiSecret are required", m.name)
	}
	// Build an HTTP client that optionally redirects to a mock server.
	httpClient := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	if baseURL != "" {
		httpClient.Transport = &urlRewriteTransport{base: baseURL}
	}

	var client *twilio.RestClient

	if apiKey != "" {
		tc := &twilioClient.Client{Credentials: twilioClient.NewCredentials(apiKey, apiSecret), HTTPClient: httpClient}
		tc.SetAccountSid(accountSid)
		params := twilio.ClientParams{
			Username:   apiKey,
			Password:   apiSecret,
			AccountSid: accountSid,
			Client:     tc,
		}
		client = twilio.NewRestClientWithParams(params)
	} else {
		tc := &twilioClient.Client{Credentials: twilioClient.NewCredentials(accountSid, authToken), HTTPClient: httpClient}
		tc.SetAccountSid(accountSid)
		params := twilio.ClientParams{
			Username:   accountSid,
			Password:   authToken,
			AccountSid: accountSid,
			Client:     tc,
		}
		client = twilio.NewRestClientWithParams(params)
	}

	if region, ok := m.config["region"].(string); ok && region != "" {
		client.SetRegion(region)
	}
	if edge, ok := m.config["edge"].(string); ok && edge != "" {
		client.SetEdge(edge)
	}

	RegisterClient(m.name, client)
	return nil
}

func missingTwilioCredentials(accountSid, authToken, apiKey, apiSecret string) bool {
	accountSid = strings.TrimSpace(accountSid)
	authToken = strings.TrimSpace(authToken)
	apiKey = strings.TrimSpace(apiKey)
	apiSecret = strings.TrimSpace(apiSecret)
	return accountSid == "" || (authToken == "" && (apiKey == "" || apiSecret == ""))
}

func configBool(config map[string]any, key string) bool {
	switch value := config[key].(type) {
	case bool:
		return value
	case string:
		return strings.EqualFold(strings.TrimSpace(value), "true")
	default:
		return false
	}
}

// Start is a no-op for this module.
func (m *twilioModule) Start(_ context.Context) error { return nil }

// Stop unregisters the Twilio client.
func (m *twilioModule) Stop(_ context.Context) error {
	UnregisterClient(m.name)
	return nil
}
