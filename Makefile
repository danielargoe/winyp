build:
	@echo "Building binary exe..."
	go build -o bin/winyp cmd/winyp/main.go
	@echo "Binary winyp location: bin/winyp"

clean:
	@echo "Removing executables from bin/"
	rm bin/*