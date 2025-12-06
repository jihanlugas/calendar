package response

import "github.com/jihanlugas/calendar/constant"

type WsMessage struct {
	Type    constant.WsType `json:"type"`
	Payload interface{}     `json:"payload"`
}
