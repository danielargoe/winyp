build:
	echo "Building binary exe..."
	go build -o bin/winyp cmd/winyp/main.go
	echo "Binary winyp located: bin/winyp"

clean:
	rm bin/*