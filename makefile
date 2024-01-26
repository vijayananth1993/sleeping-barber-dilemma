vet:
	go vet ./...

fmt:
	go fmt ./...

tidy:
	go mod tidy

pre_commit: vet tidy fmt