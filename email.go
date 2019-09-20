package sendpulse

import (
	"encoding/base64"
	"encoding/json"
	"errors"
)

type EmailFields struct {
	Html    string    `json:"html"`
	Text    string    `json:"text"`
	Subject string    `json:"subject"`
	From    Address   `json:"from"`
	To      []Address `json:"to"`
}

type Email struct {
	EmailFields
}

func NewEmail(from Address, to interface{}, subject, html, text string) (*Email, error) {
	var addresses []Address

	if value, ok := to.([]Address); ok {
		addresses = value
	} else {
		if value, ok := to.(Address); ok {
			addresses = []Address{value}
		} else {
			return nil, errors.New("\"to\" must be list of Address or one Address")
		}
	}

	return &Email{
		EmailFields{
			Html:    html,
			Text:    text,
			Subject: subject,
			From:    from,
			To:      addresses,
		},
	}, nil
}

func (e *Email) MarshalJSON() ([]byte, error) {
	var data struct {
		Email EmailFields `json:"email"`
	}

	data.Email = e.EmailFields
	data.Email.Html = base64.StdEncoding.EncodeToString([]byte(e.Html))

	return json.Marshal(data)
}
