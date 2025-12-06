package websocket

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/jihanlugas/calendar/config"
	"github.com/jihanlugas/calendar/response"
	"github.com/jihanlugas/calendar/utils"
	"github.com/jihanlugas/calendar/ws"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	hubManager *ws.HubManager
}

func NewHandler(hubManager *ws.HubManager) Handler {
	return Handler{
		hubManager: hubManager,
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		origin := r.Header.Get("Origin")
		if len(config.WsAllowedOrigins) == 0 {
			return true
		} else {
			return utils.StringContains(config.WsAllowedOrigins, origin)
		}
	},
}

func (h Handler) Serve(c echo.Context) error {
	var err error

	propertyId := c.QueryParam("propertyId")
	if propertyId == "" {
		return response.Error(http.StatusBadRequest, response.ErrorHandlerGetUserInfo, err, nil).SendJSON(c)
	}

	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return response.Error(http.StatusBadRequest, response.ErrorHandlerGetUserInfo, err, nil).SendJSON(c)
	}

	hub := h.hubManager.GetHub(propertyId)

	client := &ws.Client{
		Hub:       hub,
		Conn:      conn,
		Send:      make(chan []byte),
		Validator: h.hubManager.Validator,
	}

	hub.Register <- client

	go client.WritePump()
	go client.ReadPump()

	return nil
}
