run:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
	docker build -t go-docker-containerid .
	docker run -i go-docker-containerid
