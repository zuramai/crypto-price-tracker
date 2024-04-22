watch: 
	./bin/air -c .air.toml
setup:
	sh ./scripts/install_air.sh
prod:
	go build -o temp/main bin/main cmd/main.go && ./bin/main
migrate: 
	migrate -path "./database/migrations" -database "sqlite://database/main.db" up 
drop: 
	migrate -path "./database/migrations" -database "sqlite://database/main.db" drop
setup:
	sh ./scripts/setup_dev.sh