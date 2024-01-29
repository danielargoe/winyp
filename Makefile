build:
	@echo "Building binary exe..."
	go build -o bin/winyp cmd/winyp/main.go
	@sudo cp ./bin/winyp /usr/local/bin/
	@echo "Binary winyp location: ~/usr/local/bin/"

clean:
	@echo "Removing executable from /usr/local/bin/"
	@rm ./bin/winyp
	@sudo rm /usr/local/bin/winyp