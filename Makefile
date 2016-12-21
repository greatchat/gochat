NAME = go-chat

all: clean deps build run

deps:
	mkdir -p _logs
	glide install

build:
	go build -o $(NAME)

run: 
	./$(NAME)

kafka-up:
	docker-compose -f docker-compose-kafka-single.yml up -d

kafka-down:
	docker-compose -f docker-compose-kafka-single.yml down

clean:
	rm -f $(HOME)/.credentials/go-chat.sainbsurys.co.uk.json

test: kafka-up
	go test -race -cover -v $(shell glide nv)
