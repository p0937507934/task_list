# build stage
FROM golang:alpine3.14
WORKDIR /server
COPY . .
RUN go mod tidy
RUN GOOS=linux GOARCH=amd64 go build -o main main.go
ENTRYPOINT ./main
