TARGET=wfmirror

build: app.go **/*.go
	go build -o $(TARGET) app.go

run:
	go run ./app.go -D

server:
	go run ./app.go -D -S

clean: $(TARGET)
	rm -f $(TARGET)
