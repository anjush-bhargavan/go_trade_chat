package repo

// import (
// 	"github.com/anjush-bhargavan/go_trade_chat/pkg/models"
// )

// type Hub struct {
// 	Clients    map[uint]*models.Client
// 	Register   chan *models.Client
// 	Unregister chan *models.Client
// }

// func NewHub() *Hub {
// 	return &Hub{
// 		Clients:    make(map[uint]*models.Client),
// 		Register:   make(chan *models.Client),
// 		Unregister: make(chan *models.Client),
// 	}
// }


// func (h *Hub) Run() {
// 	for {
// 		select {
// 		case cl := <-h.Register:
	
// 			h.Clients[cl.ID] = cl

// 		case cl := <-h.Unregister:
// 			// Unregister the client directly from the hub
// 			if _, ok := h.Clients[cl.ID]; ok {
// 				delete(h.Clients, cl.ID)
// 				close(cl.Message)
// 			}

// 		}
// 	}
// }