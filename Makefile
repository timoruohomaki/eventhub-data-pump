createContainer:
	docker run -d -p 27017:27017 --name=EventHubPump

startContainer:
	docker start EventHubPump

run:
	@go run main.go

build:
	@go build main.go
