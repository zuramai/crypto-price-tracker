watch: 
	air -c air.toml
prod:
	go build -o bin/main cmd/main.go && ./bin/main
