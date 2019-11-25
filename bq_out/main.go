package functions

import (
	"cloud.google.com/go/bigquery"
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
	Name     string    `json:"name" bigquery:"NAME"`
	Place    string    `json:"place" bigquery:"PLACE"`
	Datetime time.Time `bigquery:"DATETIME"`
}

func TriggerPubSubToBigQuery(ctx context.Context, m PubSubMessage) error {
	var i Info

	err := json.Unmarshal(m.Data, &i)
	if err != nil {
		log.Printf("メッセージ変換エラー Error:%T message: %v", err, err)
		return nil
	}
	InsertBigQuery(ctx, &i)
	return nil
}

func InsertBigQuery(ctx context.Context, i *Info) {
	projectID := os.Getenv("GCP_PROJECT")
	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		log.Printf("BigQuery接続エラー Error:%T message: %v", err, err)
		return
	}
	defer client.Close()
	u := client.Dataset("GREETINGS").Table("NAMES").Uploader()

	i.Datetime = time.Now()
	items := []*Info{i}
	err = u.Put(ctx, items)
	if err != nil {
		log.Printf("データ書き込みエラー Error:%T message: %v", err, err)
		return
	}
	log.Printf("データ書き込み成功")
}
