abigen:
	abigen --abi=./contract/contract.abi --pkg contract --out ./contract/contract.go

run:
	docker-compose up  --remove-orphans --build

run_client:
	go run cmd/client/main.go

run_app:
	./run_app.sh

lint:
	golangci-lint run --fix
