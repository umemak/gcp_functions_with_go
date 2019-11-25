package functions

import (
    "context"
    "fmt"
    "net/http"
    "os"
    "time"
    "cloud.google.com/go/storage"
)
func TriggerHTTPToBucket(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        names, err := r.URL.Query()["name"]
        if !err || len(names[0]) < 1 {
            fmt.Fprint(w, "パラメータに\"name\"がありません\r\n")
            return
        }
        WriteBucket(w, names[0])
    default:
        http.Error(w, "405 - Method Not Allowed", http.StatusMethodNotAllowed)
    }
}
func WriteBucket(rw http.ResponseWriter, name string) {
    bucketName := os.Getenv("BUCKET_NAME")
    ctx := context.Background()
    client, err := storage.NewClient(ctx)
    if err != nil {
        fmt.Fprint(rw, "Storage接続エラー エラータイプ: %T エラーメッセージ: %S", err, err)
        return
    }
    defer client.Close()
    objectName := time.Now().Format("20060102150405")
    fw := client.Bucket(bucketName).Object(objectName).NewWriter(ctx)
    if _, err := fw.Write([]byte(name + "\r\n")); err != nil {
        fmt.Fprint(rw,"オブジェクト書き込みエラー エラータイプ: %T, エラーメッセージ: %S", err, err)
        return
    }
    fmt.Fprint(rw, "Storageに書き込みました: %s:%s", bucketName, objectName)
}

