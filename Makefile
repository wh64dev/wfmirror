TARGET=wfmirror

build: app.go **/*.go
	go build -o $(TARGET) app.go
	pnpm -C ./frontend build

run:
	go run ./app.go -D

clean: $(TARGET)
	rm -f $(TARGET)
