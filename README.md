# HackerChat

HackerChat is a revolutionary idea comparable only to the invention of hot water or the re-engineering of the wheel. HackerChat will solve any communication issue by putting a chat right where you need it: right in the terminal. On top of that, HackerChat comes with state-of-the-art filters to facilitate communication and reduce professional and cultural misunderstandings (UK only).

For example:

```
# Dev filters:
> It's very verbose # this will be translated to its correct meaning:
> It's Java

# Company-wise filers:
> Don't worry! # will be translated to:
> Be worried!

# UK filters:
> I hear what you say # will be translated to:
> I disagree and do not want to discuss it further

> That is a very brave proposal # will be translated to:
> You are insane

```

Over-engineered, under-tested and completely unstable, HackerChat has the security of a V8 mounted on a shopping cart: a great experience for your CPU and absolute hell for your users.

*Make development great again!*

## Team Members
- Dan Bond
- Tim Blackwell
- Edo Scalafiotti

# Requirements
- Docker & Docker Compose
- Go >= 1.7

# Installation

```bash

git clone https://github.com/edoardo849/hackerchat

# Install dependencies
go get ./...

# OR, install dependencies with glide
glide install

# Edit main.go to include the IP of your Kafka instance and modify accordingly docker-compose.yml

# Run the docker container with Kafka (locally or from a remote server)
make kafka-up

# Run the app
make run

# Kill Docker
make kafka-down

```

# TODOS
- Use [NSQ](https://github.com/nsqio/nsq) instead of Kafka
- Investigate P2P because... why not? If life isn't spicy it's no fun...
- As a CLI client, maybe have a look at [tcell](https://github.com/gdamore/tcell) which is the base for [micro](https://github.com/zyedidia/micro)
