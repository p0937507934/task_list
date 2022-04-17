# build stage
FROM golang:alpine3.14
WORKDIR /server
COPY . .
RUN go mod tidy
RUN GOOS=linux GOARCH=arm go build -o main main.go
ENTRYPOINT ./main
