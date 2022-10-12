package chat

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
)

var (
	customTimeFormat = "2006-01-02 3:04 PM"
)

type Hub struct {
	Name  string
	Rooms map[string]*Room
	DAO   *DAO
	Url   string
}

func NewHub(name string, router *gin.Engine) *Hub {
	hub := &Hub{
		Name:  name,
		Rooms: make(map[string]*Room),
		DAO:   NewDAO(),
		Url:   "/" + name,
	}

	hub.Register(router)

	return hub
}

// Register registers the hub with the router
func (h *Hub) Register(router *gin.Engine) {
	router.GET(h.Url+"/:roomID", func(c *gin.Context) {
		roomID := c.Param("roomID")
		room, ok := h.Rooms[roomID]
		if !ok {
			log.Println("Room does not exist")
			return
		}

		room.Melody.HandleRequest(c.Writer, c.Request)
	})
}

// AddRoom adds a new room to the hub
func (h *Hub) AddRoom(name string) *Room {
	room := &Room{
		Name:      name,
		Melody:    melody.New(),
		ParentHub: h,
	}

	h.Rooms[name] = room

	room.Melody.HandleConnect(room.NewConnectionHandler)

	room.Melody.HandleMessage(room.NewMessageHandler)

	return h.Rooms[name]
}
