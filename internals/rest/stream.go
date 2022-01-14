package rest

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/session"
	// "github.com/pkg/errors"
)

func (tc *Client) GetSessionID() (interface{}, error) {

	// First create a streaming session.
	createSessionUrl := tc.endpoint + "/v1/markets/events/session"

	a := fiber.AcquireAgent()
	defer fiber.ReleaseAgent(a)

	r := fiber.AcquireResponse()
	defer fiber.ReleaseResponse(r)

	a.Reuse()

	//Set request
	req := a.Request()
	req.Header.SetMethod(fiber.MethodPost)
	req.Header.Set("accept", "application/json")
	req.SetRequestURI(createSessionUrl)

	if err := a.Parse(); err != nil {
		fmt.Println(err.Error())
	}

	// var retCode int
	var retBody []byte
	var errs []error

	// Send out the HTTP request
	var sessionResp struct {
		Stream struct {
			SessionId string
			Url       string
		}
	}

	if _, retBody, errs = a.Struct(&sessionResp); len(errs) > 0 {
		log.Printf("received: %v", string(retBody))
		log.Printf("could not send HTTP request: %v", errs)
		return nil, errs[len(errs)-1]
	}
	return sessionResp, nil
}
