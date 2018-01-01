# SnakeCoin Go
Golang implementation of the _tiniest blockchain_ as described at 
https://medium.com/crypto-currently/lets-build-the-tiniest-blockchain-e70965a248bg

## Usage
Start the server:
```bash
$ go run node/main.go
```

## Endpoints
Make a transaction:
```bash
$ curl -X POST http://localhost:8000/txn \
    -H 'Content-Type: application/json' \
    -d '{
    "from": "71238uqirbfh894-random-public-key-a-alkjdflakjfewn204ij",
    "to": "93j4ivnqiopvh43-random-public-key-b-qjrgvnoeirbnferinfo",
    "amount": 3
  }'
```

Mine a SnakeCoin:
```bash
$ curl -X GET http://localhost:8000/mine
```

Get the current blockchain:
```bash
$ curl -X GET http://localhost:8000/blocks
```