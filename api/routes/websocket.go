package routes

import (
	"context"
	"encoding/json"
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/hosseinmirzapur/rthnk/client"
)

type WSMessage struct {
	Msg string `json:"msg"`
}

func registerWSRoutes(app *fiber.App) {
	ws := app.Group("/ws")

	ws.Get("/:secret", websocket.New(func(c *websocket.Conn) {

		secret := c.Params("secret")

		rc := client.NewRedisClient()

		pubsub := rc.Subscribe(context.Background(), secret)
		defer pubsub.Close()

		for {
			_, msg, err := c.ReadMessage()
			if err != nil {
				log.Println("r error:", err.Error())
				break
			}

			var wsMsg WSMessage
			err = json.Unmarshal(msg, &wsMsg)
			if err != nil {
				log.Println("r error:", err.Error())
				break
			}

			err = rc.Publish(context.Background(), secret, wsMsg.Msg).Err()
			if err != nil {
				log.Println("redis pub error:", err.Error())
				break
			}

			psMsg, err := pubsub.ReceiveMessage(context.Background())
			if err != nil {
				log.Println("redis read error:", err.Error())
				break
			}

			err = c.WriteMessage(websocket.TextMessage, []byte(psMsg.Payload))
			if err != nil {
				log.Println("w error:", err.Error())
				break
			}

		}
		log.Println("connection closed...")
	}))
}
