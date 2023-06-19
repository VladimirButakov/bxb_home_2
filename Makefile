build:
	go build -o downloader main.go

run:
	go run main.go

clean:
	rm -f downloader