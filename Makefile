.PHONY: build test run clean

BINARY_NAME=slack-notifier

build:
	go build -o $(BINARY_NAME) main.go

test:
	go test ./...

run:
	@echo "メッセージを引数として指定してください。例: make run MSG='Hello Slack'"
	@[ "${MSG}" ] || (echo "MSGが指定されていません。"; exit 1)
	./$(BINARY_NAME) "${MSG}"

clean:
	rm -f $(BINARY_NAME)
