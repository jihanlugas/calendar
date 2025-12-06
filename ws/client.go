package ws

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/jihanlugas/calendar/app/event"
	"github.com/jihanlugas/calendar/constant"
	"github.com/jihanlugas/calendar/db"
	"github.com/jihanlugas/calendar/request"
	"github.com/jihanlugas/calendar/response"
	"github.com/jihanlugas/calendar/utils"
)

type EchoValidator interface {
	Validate(i interface{}) error
}

type Client struct {
	Hub       *PropertyHub
	Conn      *websocket.Conn
	Send      chan []byte
	Validator EchoValidator
}

func bindPayload[T any](payload interface{}, out *T) error {
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, out)
}

func (c *Client) ReadPump() {
	defer func() {
		c.Hub.unregister <- c
		c.Conn.Close()
	}()

	conn, closeConn := db.GetConnection()
	defer closeConn()

	eventRepo := event.NewRepository()

	for {

		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			break
		}

		var data request.WsRead

		err = json.Unmarshal(msg, &data)
		if err != nil {
			break
		}

		switch data.Type {
		case constant.WS_TYPE_GET_EVENT:
			req := new(request.TimelineEvent)
			if err := bindPayload(data.Payload, req); err != nil {
				fmt.Println("error bind payload:", err)
				break
			}

			utils.TrimWhitespace(req)

			err = c.Validator.Validate(req)
			if err != nil {
				fmt.Println("error validation:", err)
				break
			}

			vEvents, err := eventRepo.Timeline(conn, *req)
			if err != nil {
				break
			}

			res := response.WsMessage{
				Type:    constant.WS_TYPE_DATA_EVENT,
				Payload: vEvents,
			}

			resByte, err := json.Marshal(res)
			if err != nil {
				break
			}

			c.Send <- resByte
		}
	}
}

func (c *Client) WritePump() {
	defer c.Conn.Close()

	for msg := range c.Send {
		err := c.Conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			return
		}
	}
}
