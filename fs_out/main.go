package functions

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"log"
	"os"
	"time"
)

type PubSubMessage struct {
	Data []byte `json:"data"`
}

type Info struct {
	Name     string    `json:"name" firestore::"NAME"`
	Place    string    `json:"place" firestore:"PLACE"`
	Datetime time.Time `firestore:"DATETIME"`
}

func TriggerPubSubToFirestore(ctx context.Context, m PubSubMessage) error {
	var i Info

	err := json.Unmarshal(m.Data, &i)
	if err != nil {
		log.Printf("メッセージ変換エラー Error:%T message: %v", err, err)
		return nil
	}
	PutFirestore(ctx, &i)
	return nil
}

func PutFirestore(ctx context.Context, i *Info) {
	projectID := os.Getenv("GCP_PROJECT")
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Printf("Firestore接続エラー Error: %T message %v", err, err)
		return
	}
	defer client.Close()
	i.Datetime = time.Now()
	_, _, err = client.Collention("NAMES").Add(ctx, i)
	if err != nil {
		log.Printf("データ書き込みエラー Error: %T message: %v", err, err)
		return
	}
	log.Printf("データ書き込み成功")
}
