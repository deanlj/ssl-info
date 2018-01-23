all: pre-build windows mac linux

pre-build:
	dep ensure

windows:
	GOOS=windows GOARCH=amd64 go build -o bin/windows/ssl-info.exe
	shasum -a 256 bin/windows/ssl-info.exe > sha256sums.txt

linux:
	GOOS=linux GOARCH=amd64 go build -o bin/linux/ssl-info
	shasum -a 256 bin/linux/ssl-info >> sha256sums.txt
mac:
	GOOS=darwin GOARCH=amd64 go build -o bin/mac/ssl-info
	shasum -a 256 bin/mac/ssl-info >> sha256sums.txt
