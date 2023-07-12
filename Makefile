
test:
	go test ./tests -v -p 1

run:
	go run main.go

build:
	go build -o main ./main.go