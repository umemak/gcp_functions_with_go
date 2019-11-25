package functions

import (
    "context"
    "log"
    "encoding/json"
)

type PubSubMessage struct {
    Data []byte `json:"data"`
}

type Info struct {
    Name string `json:"name"`
    Place string `json:"place"`
}

func TriggerPubSub(ctx context.Context, m PubSubMessage) error {
    var i Info
    err := json.Unmarshal(m.Data, &i)
    if err != nil {
        log.Printf("Error: %T message: %v", err, err)
        return nil
    }
    log.Printf("こんにちは %s さん %s へ Cloud Pub/Subから", i.Name, i.Place)
    return nil
}
