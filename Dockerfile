FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./

COPY go.sum ./

RUN go mod download

COPY *.go ./

COPY . .

RUN go build -o /cake-store-api

EXPOSE 8001

CMD [ "/cake-store-api" ]