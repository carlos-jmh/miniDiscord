package chat

import (
	"context"
	"log"
	"time"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

// DAO (data access object) interacts with a database
type DAO struct {
	db *firestore.Client
}

func NewDAO() *DAO {
	return &DAO{
		db: firebaseInit(),
	}
}

// Put saves a message to the database
func (d *DAO) Put(room *Room, msg *Message) {
	ctx := context.Background()

	doc := d.db.Collection(room.ParentHub.Name).
		Doc(room.Name).
		Collection("messages")

	_, _, _ = doc.Add(ctx, map[string]interface{}{
		"timestamp": msg.Timestamp,
		"message":   msg.Content,
	})
}

// Get retrieves messages from the database
func (d *DAO) Get(room *Room) []*Message {
	ctx := context.Background()

	messagesCollection := d.db.Collection(room.ParentHub.Name).
		Doc(room.Name).
		Collection("messages")

	documents, err := messagesCollection.OrderBy("timestamp", firestore.Asc).
		Documents(ctx).
		GetAll()
	if err != nil {
		log.Println("error getting data from firestore: ", err)
		return nil
	}

	var allMessages []*Message

	for _, document := range documents {
		currMsg := &Message{
			Timestamp: document.Data()["timestamp"].(time.Time),
			Content:   document.Data()["message"].(string),
		}
		allMessages = append(allMessages, currMsg)
	}

	return allMessages
}

// firebaseInit initializes a connection to the database
func firebaseInit() *firestore.Client {
	ctx := context.Background()

	// Replace this with your own credentials
	opt := option.WithCredentialsFile("secret.json")

	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Println("error initializing app: ", err)
		return nil
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Println("error initializing Firestore client: ", err)
		return nil
	}

	return client
}
