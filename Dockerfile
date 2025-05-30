# 使用官方 Golang image 作為 build base
FROM golang:1.23 AS builder

# 設定工作目錄
WORKDIR /app

# 複製 go mod 並下載依賴
COPY go.mod go.sum ./
RUN go mod download

# 複製專案所有檔案
COPY . .

# 產生 swagger 文件（若已安裝 swag 並存在 main.go 的註解）
RUN go install github.com/swaggo/swag/cmd/swag@latest && swag init

# 編譯 Go 應用
RUN go build -o task-api main.go

# 建立最小 runtime image
FROM ubuntu:22.04

# 設定工作目錄
WORKDIR /root/

# 複製執行檔與 Swagger 靜態文件
COPY --from=builder /app/task-api .
COPY --from=builder /app/docs ./docs

# 暴露 port 8080
EXPOSE 8080

# 啟動服務
CMD ["./task-api"]
