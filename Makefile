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

migrate: build_migrate
	./build/migrate -purl //localhost:5432 -pssl disable -pdb postgres -puser postgres -ppw postgres

build_app:
	docker build --tag truphone_go:latest .

run_postgres:
	docker-compose up -d postgres

run: run_postgres
	docker-compose up -d truphone

vendor:
	go mod download

run_local:
	go run cmd/service/main.go -purl //localhost:5432 -pssl disable -pdb postgres -puser postgres -ppw postgres
