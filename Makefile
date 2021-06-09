build:
	go mod download
	go build -o cmd/korai main.go
run:
	go mod download
	go run main.go
shell:
	docker run -p 9000:9000 -it --mount src=`pwd`,target=/app,type=bind  korai /bin/bash
build-image:
	docker build . -t korai
build-image-prod:
	docker build -f Production.Dockerfile . -t korai
run-prod-image:
	docker run -p 9000:9000 korai
