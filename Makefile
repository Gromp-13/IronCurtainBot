.PHONY:



build:
	go build -o ./.bin/bot cmd/bot/main.g

run: build
	./.bin/bot
