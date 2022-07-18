FROM golang:1.18

COPY . /app

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download
# Run test before starting 
RUN go test ./test

RUN go build -o /main cmd/main.go

EXPOSE 8080

CMD [ "/main" ]