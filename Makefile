test:
	go test -coverprofile cp.out ./...

coverage: test
	go tool cover -html=cp.out