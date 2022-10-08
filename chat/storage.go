package chat

import (
	"context"
	"fmt"
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

func (d *DAO) Put(roomID string, msg string) {
	ctx := context.Background()

	_, _, _ = d.db.Collection("hub").Doc(roomID).Collection("messages").Add(ctx, map[string]interface{}{
        "message": msg,
		"timestamp": time.Now().Format(time.RFC3339),
	})
}

func (d *DAO) Get(roomID string) []string {
	allMessages := []string{}
	ctx := context.Background()

	// get all "messages" documents in order from oldest to newest
	it, err := d.db.Collection("hub").Doc(roomID).Collection("messages").OrderBy("timestamp", firestore.Asc).Documents(ctx).GetAll()
	if err != nil {
		fmt.Println("error getting data from firestore: ", err)
		return nil
	}

	for _, document := range it {
		allMessages = append(allMessages, document.Data()["message"].(string))
	}

	return  allMessages
}

func firebaseInit() *firestore.Client {
	opt := option.WithCredentialsFile("secret.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		fmt.Println("error initializing app: ", err)
		return nil
	}
	
	client, err := app.Firestore(context.Background())
	if err != nil {
		fmt.Println("error initializing Firestore client: ", err)
		return nil
	}

	return client
}