# go-sendpulse
Golang client for sendpulse api. Currently you can only send emails with it.

This library use `golang.org/x/oauth2/clientcredentials` for sendpulse's oauth2 endpoint.

## Usage example

```go
package main

import (
	"log"

	sp "github.com/ont/go-sendpulse"
)

func main() {
	s := sp.New("... client-id ...", "... client-secret ...")

	email, err := sp.NewEmail(
		sp.Address{"from me", "sender@some.com"},
		sp.Address{"to this person", "person@test.com"},
		"test subject",
		"<b>test</b> html",
		"test text",
	)
	log.Println(err)

	err = s.SMTP.Send(email)
	log.Println(err)
}
```
