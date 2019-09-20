package sendpulse

import (
	"context"
	"net/http"

	resty "github.com/go-resty/resty/v2"
	"golang.org/x/oauth2/clientcredentials"
)

type Sendpulse struct {
	client *resty.Client

	SMTP *apiSmtp
}

func New(clientId, clientSecret string) *Sendpulse {
	oauthClient := newOAuthClient(clientId, clientSecret)

	client := resty.NewWithClient(oauthClient)
	client.SetHostURL("https://api.sendpulse.com")

	return &Sendpulse{
		client: client,
		SMTP:   &apiSmtp{client},
	}
}

func newOAuthClient(clientId, clientSecret string) *http.Client {
	config := clientcredentials.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		TokenURL:     "https://api.sendpulse.com/oauth/access_token",
	}

	ctx := context.Background()
	return config.Client(ctx)
}
