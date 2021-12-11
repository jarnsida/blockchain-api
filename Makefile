build_app:
	go build -race cmd/app/main.go

build_client:
	go build -race cmd/client/main.go

test:
	go test -v ./...

abigen:
	abigen --abi=./contract/contract.abi --pkg contract --out ./contract/contract.go

run:
	docker-compose up  --remove-orphans --build

run_client:
	go run -race cmd/client/main.go

run_app:
	./run_app.sh

lint:
	golangci-lint run --fix
