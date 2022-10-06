package chat

import (
	"log"
	"time"

	goaway "github.com/TwiN/go-away"
	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
)

var (
	customTimeFormat = "2006-01-02 15:04:05"
)

type Hub struct {
	Name  string
	Rooms map[string]*Room
	DAO   *DAO
	Url   string
}

type Room struct {
	Name      string
	Melody    *melody.Melody
	ParentHub *Hub
}

func NewHub(name string, r *gin.Engine) *Hub {
	hub := &Hub{
		Name:  name,
		Rooms: make(map[string]*Room),
		DAO:   NewDAO(),
		Url:   "/" + name,
	}

	r.GET(hub.Url+"/:roomID", func(c *gin.Context) {
		roomID := c.Param("roomID")
		room, ok := hub.Rooms[roomID]
		if !ok {
			log.Println("Room does not exist")
			return
		}

		room.Melody.HandleRequest(c.Writer, c.Request)
	})

	return hub
}

func (h *Hub) AddRoom(name string) *Room {
	room := &Room{
		Name:      name,
		Melody:    melody.New(),
		ParentHub: h,
	}

	h.Rooms[name] = room

	room.Melody.HandleConnect(room.NewClientConnection)

	room.Melody.HandleMessage(room.NewMessage)

	return h.Rooms[name]
}

func (r *Room) AddMessage(msg string) {
	r.ParentHub.DAO.Put(r.Name, msg)
}

func (r *Room) GetMessages() []string {
	return r.ParentHub.DAO.Get(r.Name)
}

func (r *Room) NewClientConnection(s *melody.Session) {
	r.RestoreMessages(s)
}

func (r *Room) NewMessage(_ *melody.Session, msg []byte) {
	timestamp := time.Now().Format(customTimeFormat)

	if goaway.IsProfane(string(msg)) {
		return
	}

	finalMsg := timestamp + " " + string(msg)
	r.AddMessage(finalMsg)
	r.Melody.Broadcast([]byte(finalMsg))
}

func (r *Room) RestoreMessages(s *melody.Session) {
	for _, msg := range r.GetMessages() {
		s.Write([]byte(msg))
	}
}
