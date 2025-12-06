package listener

import (
	"encoding/json"
	"log"
	"time"

	"github.com/jihanlugas/calendar/constant"
	"github.com/jihanlugas/calendar/db"
	"github.com/jihanlugas/calendar/model"
	"github.com/jihanlugas/calendar/response"
	"github.com/jihanlugas/calendar/ws"
	"github.com/lib/pq"
)

type EventChangePayload struct {
	Operation string       `json:"operation"`
	Old       *model.Event `json:"old"`
	New       *model.Event `json:"new"`
}

func StartEventListener(manager *ws.HubManager) {
	var err error
	listener := pq.NewListener(
		db.Dns,
		10*time.Second, // min reconnect
		time.Minute,    // max reconnect
		nil,
	)

	err = listener.Listen("event_changes")
	if err != nil {
		log.Println("Failed to LISTEN:", err)
		return
	}

	go func() {
		sendRefetch := func(propertyID string) {
			res := response.WsMessage{Type: constant.WS_TYPE_REFETCH}
			resByte, _ := json.Marshal(res)
			manager.SendToProperty(propertyID, resByte)
		}

		for {
			select {
			case notification := <-listener.Notify:
				if notification == nil {
					continue
				}

				log.Println("Received:", notification.Extra)

				var payload EventChangePayload
				err = json.Unmarshal([]byte(notification.Extra), &payload)
				if err != nil {
					log.Println("Unmarshal error:", err)
					continue
				}

				switch payload.Operation {

				case "INSERT", "UPDATE":
					if payload.New == nil {
						log.Println("ERROR: payload.New nil on", payload.Operation)
						continue
					}
					sendRefetch(payload.New.PropertyID)

				case "DELETE":
					if payload.Old == nil {
						log.Println("ERROR: payload.Old nil on DELETE")
						continue
					}
					sendRefetch(payload.Old.PropertyID)

				default:
					log.Println("Unknown operation:", payload.Operation)
					continue
				}

			case <-time.After(90 * time.Second):
				log.Println("Keepalive: Ping()")
				listener.Ping()
			}
		}
	}()
}
