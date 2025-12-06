package ws

type PropertyHub struct {
	propertyID string
	clients    map[*Client]bool
	Register   chan *Client
	unregister chan *Client
	broadcast  chan []byte
}

func NewPropertyHub(propertyID string) *PropertyHub {
	return &PropertyHub{
		propertyID: propertyID,
		clients:    make(map[*Client]bool),
		Register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan []byte),
	}
}

func (h *PropertyHub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.clients[client] = true

		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.Send)
			}

		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.clients, client)
				}
			}
		}
	}
}
