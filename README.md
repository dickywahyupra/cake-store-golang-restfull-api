# API for Cake Store

## Getting Started

1. - Install Docker
   - Install docker-compose
2. Run `cp .env.example .env`
4. Run `docker-compose -f compose-db.yml up` or `docker-compose -f compose-db.yml up -d` for background running database container 
5. Run `go run main.go`
5. Run `go test ./test -v` for run unit test