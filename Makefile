.PHONY: build
build:
	go build -v -o ./bin/tg-bot ./cmd/main.go

.PHONY: run
run:
	go run ./cmd/main.go

.PHONY: docker-run
docker-run:
	docker build --build-arg token=$(TELEGRAM_TOKEN) -t tg-bot-image  .
	docker run -it --rm --name tg-bot-container tg-bot-image
