NAME = go-chat

build:
	go build -o $(NAME)

run: build
	./$(NAME)

kafka-up:
	docker-compose -f docker-compose-kafka-single.yml up -d

kafka-down:
	docker-compose -f docker-compose-kafka-single.yml down
clean:
	rm $(HOME)/.credentials/go-chat.sainbsurys.co.uk.json

test: kafka-up
	go test -race -cover -v $(shell glide novendor)
