FROM golang:latest

ENV PROJECT_DIR=/app \
    GO111MODULE=on \
    CGO_ENABLED=1

WORKDIR /app
RUN mkdir "/build"
RUN 

COPY . .
# RUN go mod init
RUN go mod tidy

COPY --from=itinance/swag /root/swag /usr/local/bin

RUN go get github.com/githubnemo/CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon
RUN go install github.com/swaggo/swag/cmd/swag@latest
ENTRYPOINT CompileDaemon -directory="./app" -build="go build -o /build/app" -command="/build/app" -polling=true