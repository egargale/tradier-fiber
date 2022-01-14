package rest

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/gofiber/fiber/v2"
	"nhooyr.io/websocket"
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
	req.Header.Set("Authorization", tc.authHeader)
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
	log.Println(sessionResp.Stream.SessionId)
	go OpenStreamSocket(sessionResp.Stream.SessionId)
	return sessionResp, nil
}

func OpenStreamSocket(id string) {

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	c, _, err := websocket.Dial(ctx, "wss://ws.tradier.com/v1/markets/events", nil)
	if err != nil {
		//log.Printf("handshake failed with status %d", resp.StatusCode)
		log.Fatal("dial:", err)
	}
	defer c.Close(websocket.StatusInternalError, "the sky is falling")

	type streamPayload struct {
		Symbols   []string `json:"symbols"`
		SessionId string   `json:"sessionid"`
		Linebreak string   `json:"linebreak"`
	}

	defaultsymbol := []string{"SPY"}
	var initial_payload streamPayload
	initial_payload.SessionId = id
	initial_payload.Symbols = defaultsymbol
	initial_payload.Linebreak = "true"
	spew.Dump(initial_payload)

	// err = wsjson.Write(ctx, c, initial_payload)
	fakepayload := fmt.Sprintf("{\"symbols\": [\"SPY\"], \"sessionid\": \"%s\", \"linebreak\": true}", id)
	log.Println(fakepayload)
	err = c.Write(ctx, websocket.MessageText, []byte(fakepayload))
	if err != nil {
		log.Printf("Error in subscribing stream %s", err)
	}
	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			_, message, err := c.Read(ctx)
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	ticker := time.NewTicker(time.Second * 30)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			c.Close(websocket.StatusNormalClosure, "")
			return
		case t := <-ticker.C:
			log.Printf("time: %s", t.String())
			// err := c.Write(ctx, websocket.MessageText, []byte(t.String()))
			// if err != nil {
			// 	log.Println("write:", err)
			// 	return
			// }
		}
	}
}
