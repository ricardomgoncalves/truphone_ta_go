test:
	go test -coverprofile cp.out ./...

coverage: test
	go tool cover -html=cp.out

test_integration:
	go test -coverprofile cpi.out -tags=test_all  ./...

coverage_integration: test_integration
	go tool cover -html=cpi.out

build_migrate:
	go build -o ./build/migrate cmd/migrate/main.go

migrate:
	./build/migrate -purl //localhost:5432 -pssl disable -pdb postgres -puser postgres -ppw postgres

docker_build:
	docker build --tag truphone_go:lastest .

vendor:
	go mod download

run:
	go run cmd/service/main.go -purl //localhost:5432 -pssl disable -pdb postgres -puser postgres -ppw postgres
