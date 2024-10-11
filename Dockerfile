# 使用官方的 Go 基礎映像
FROM golang:1.20-alpine

# 設置工作目錄
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
# 下載並安裝應用所需的依賴
RUN go mod tidy

# 將當前目錄的所有文件拷貝到容器的 /app 目錄下
COPY . .

# 編譯 Go 應用
RUN go build -o main ./cmd/server

# 容器將會監聽 8080 埠
EXPOSE 8080

# 設置應用執行時的指令
CMD ["./main"]

