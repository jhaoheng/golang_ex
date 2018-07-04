1. 執行 `docker-compose up -d`
2. 進入容器 `docker exec -it myapp /bin/bash`
3. 在 `app/main.go` 中，填妥 key/secret
4. 執行 `go run main.go`, 預設是列表既有的 bucket

ps : 要測試 s3 其他功能, 參考 main.go 中的 func