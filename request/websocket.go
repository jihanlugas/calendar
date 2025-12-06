package request

import "github.com/jihanlugas/calendar/constant"

type WsRead struct {
	Type    constant.WsType `json:"type"`
	Payload interface{}     `json:"payload"`
}
