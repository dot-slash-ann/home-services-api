# home-services-api

A RESTful API for managing home services build using Go and the Gin framework.

## features

- keep track of monetary transactions
- categorize transactions
- tag transactions

## Prerequisites

- Go 1.23+
- PostgresSQL

## Installation

1. clone the repo:

```bash
git clone https://github.com/dot-slash-ann/home-services-api

cd home-services-api
```

2. install dependencies

```bash
go mod download
```

3. install CompileDaemon

```bash
go install github.com/githubnemo/CompileDaemon
```

4. run tests
```bash
go test ./...
```


5. Configure environment variables for your database setup

6. run the app

```bash
CompileDaemon -command="./home-services-api"
```

## License

MIT
