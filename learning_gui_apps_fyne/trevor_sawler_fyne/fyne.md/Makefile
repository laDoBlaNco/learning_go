BINARY_NAME=Markdown
APP_NAME=LadosMarkdown
VERSION=1.0.0
ICON=icon.png

## build: build binary and package app
build:
	rm -rf ${BINARY_NAME}
	rm -f fyne-md
	fyne package -appVersion ${VERSION} -name ${APP_NAME} -icon ${ICON} -release
	
## run: builds and runs the application
run:
	go run .
	
## clean: runs go clean and deletes binaries
clean:
	@echo "Cleaning..."
	@go clean
	@rm -rf ${BINARY_NAME}
	@echo "Cleaned"
	
## test: runs all tests
test:
	go test -v ./...
