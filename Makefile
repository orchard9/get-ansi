.PHONY: test demo lint clean

test:
	go test -v .

demo:
	go run examples/demo/main.go

lint:
	go fmt ./...
	go vet ./...

clean:
	go clean
