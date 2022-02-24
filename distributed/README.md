`go run persist/server/itemsaver.go -port 9999`
`go run worker/server/worker.go -port 9000`
`go run main.go -itemsaver_host ":9999" -worker_hosts ":9000,:9001`
