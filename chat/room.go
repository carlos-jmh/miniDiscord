package chat

import (
	"time"

	goaway "github.com/TwiN/go-away"
	"github.com/olahol/melody"
)

type Room struct {
	Name      string
	Melody    *melody.Melody
	ParentHub *Hub
}

type Message struct {
	Timestamp time.Time
	Content   string
}

func (r *Room) SaveMessage(msg *Message) {
	r.ParentHub.DAO.Put(r, msg)
}

func (r *Room) GetMessages() []*Message {
	return r.ParentHub.DAO.Get(r)
}

// NewConnectionHandler is called when a new client connects to the room
func (r *Room) NewConnectionHandler(s *melody.Session) {
	r.RestoreMessages(s)
}

// NewMessageHandler is called when a new message is received
func (r *Room) NewMessageHandler(_ *melody.Session, data []byte) {
	if goaway.IsProfane(string(data)) {
		return
	}

	message := &Message{
		Timestamp: time.Now(),
		Content:   string(data),
	}

	r.SaveMessage(message)

	str := message.toString(customTimeFormat)

	// send message to all online clients
	r.Melody.Broadcast([]byte(str))
}

// RestoreMessages sends all messages in the room to the client
func (r *Room) RestoreMessages(s *melody.Session) {
	for _, msg := range r.GetMessages() {
		str := msg.toString(customTimeFormat)
		s.Write([]byte(str))
	}
}

func (m *Message) toString(format string) string {
	time := m.Timestamp.Format(format)
	content := m.Content

	return time + " " + content
}
