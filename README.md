#  Temperature zipcode

A tool for retrieving the weather of a location using its ZIP code

## Configuration

```bash
go mod tidy
```

### How to run

You must configure your [Weather API](https://weatherapi.com/) key on `.env` file like the example `.env.example`

```bash
go run cmd/main.go
```
#### Or
```bash
docker compose up -d
```

The server will be running on port `:8080`

## Usage

### Local

```bash
curl http://localhost:8080/28928728/temperature
```

### GCP
```bash
curl https://prd1---temperature-zipcode-srv-5an7noxrva-uc.a.run.app/28928728/temperature
```

The path parameter zipcode (28928728 on example), is required.