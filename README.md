# API for Cake Store

## Getting Started

1. - Install Docker
   - Install docker-compose
2. Run `cp .env.example .env`
3. Run `docker build . --tag=cake-store-api`
4. Run `docker-compose -f compose.yml up` or `docker-compose -f compose.yml up -d` for background running
5. Run `docker exec -it cake-store-restfull-api go test ./test -v` for run unit test

