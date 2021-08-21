pretty:
	go fmt `go list ./...`

build:
	docker build . -t pursuit-gateway-dock

run:
	docker run --net pursuit_network -p 5003:5003 pursuit-gateway-dock

test:
	go test `go list ./... | grep -v cmd | grep -v vendor`
