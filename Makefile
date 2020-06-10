test:
	go test -coverprofile cp.out ./...

coverage: test
	go tool cover -html=cp.out

build_migrate:
	go build -o ./build/migrate cmd/migrate/main.go

migrate:
	./build/migrate -purl //localhost:5432 -pssl disable -pdb postgres -puser postgres -ppw postgres