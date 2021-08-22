pretty:
	go fmt `go list ./...`

build:
	docker build . -t pursuit-gateway-dock

run:
	docker run --rm --net pursuit_network --name gateway -p 5003:5003 pursuit-gateway-dock

test:
	go test `go list ./... | grep -v cmd | grep -v vendor`
