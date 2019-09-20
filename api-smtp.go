package sendpulse

import (
	"fmt"

	resty "github.com/go-resty/resty/v2"
)

type apiSmtp struct {
	client *resty.Client
}

type SuccessResponse struct {
	Result bool `json:"result"`
}

type ErrorResponse struct {
	Code    int    `json:"error_code"`
	Message string `json:"message"`
}

func (a *apiSmtp) Send(email *Email) error {
	var (
		successResp SuccessResponse
		errResp     ErrorResponse
	)

	resp, err := a.client.R().
		SetBody(email).
		SetResult(&successResp).
		SetError(&errResp).
		Post("/smtp/emails")

	if err != nil {
		return fmt.Errorf("Error during request: http_code = %d  sendpulse_code = %d  message = \"%s\", err = \"%s\"",
			resp.StatusCode(),
			errResp.Code,
			errResp.Message,
			err,
		)
	}

	if !successResp.Result {
		return fmt.Errorf("Sendpulse respond with success = false")
	}

	return nil
}
