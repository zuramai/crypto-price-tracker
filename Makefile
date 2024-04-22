watch: 
	./bin/air -c .air.toml
setup:
	sh ./scripts/install_air.sh
prod:
	go build -o temp/main bin/main cmd/main.go && ./bin/main
