package main

import (
	"context"
	"log"

	"os"
	"os/signal"
	"time"

	"nhooyr.io/websocket"
	
	// "nhooyr.io/websocket/wsjson"
)

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	c, _, err := websocket.Dial(ctx, "ws://localhost:3000/ws", nil)
	if err != nil {
		//log.Printf("handshake failed with status %d", resp.StatusCode)
		log.Fatal("dial:", err)
	}
	defer c.Close(websocket.StatusInternalError, "the sky is falling")
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

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			// err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
			err := c.Write(ctx, websocket.MessageText, []byte(t.String()))
			if err != nil {
				log.Println("write:", err)
				return
			}
		case <-interrupt:
			log.Println("interrupt")
			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.Close(websocket.StatusNormalClosure, "")
			if err != nil {
				log.Println("write close:", err)
				return
			}
		}
	}
}
