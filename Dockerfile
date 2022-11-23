FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./

COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN mkdir -p mysql && chmod -R 777 mysql/

COPY . .

RUN go build -o /cake-store-api

# EXPOSE 8081

CMD [ "/cake-store-api" ]