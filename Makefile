.PHONY: run
run:
	export <(cat .env | xargs)
	go run main.go
