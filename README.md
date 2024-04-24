# Crypto Price Tracker

REST API for tracking cryptocurrency prices. Cryptocurrencies data taken from https://docs.coincap.io

## Live Demo

Access http://crypto.saugi.me for live demo

## Installation

### Install via Docker

1. Build the image

```sh
docker build -t crypto .
```

2. Run image

```
docker run -p 3000:3000 crypto
```

### Install locally

1. Clone repository
```
git clone https://github.com/zuramai/crypto-price-tracker
cd crypto-price-tracker
```

2. Run the app
```
make build && ./bin/main
```

### Development mode

Install `golang-migrate` and `air` package. This can be done by running shell command.

```
./scripts/setup_dev.sh
```

To migrate the database up:
```
make migrate-up
```

To start development mode:
```
make watch
```

## API Documentation

You can use the [OpenAPI Spec](./api/openapi.yml) or directly use [postman collection provided](./api/crypto_price_tracker.postman_collection.json).

## License

MIT License
