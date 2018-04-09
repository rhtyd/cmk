build:
	go build -ldflags='-s -w' -o cmk cmk.go

run:
	go run cmk.go

test:
	go test

install: build
	@echo Copied to ~/bin
	@cp cmk ~/bin

debug:
	go build -gcflags='-N -l' -o cmk cmk.go &&  dlv --listen=:2345 --headless=true --api-version=2 exec ./cmk

dist:
	rm -fr dist
	mkdir -p dist
	GOOS=linux GOARCH=amd64 go build -o dist/cmk-linux-amd64 cmk.go
	GOOS=linux GOARCH=386 go build -o dist/cmk-linux-i386 cmk.go
	GOOS=linux GOARCH=arm64 go build -o dist/cmk-linux-arm64 cmk.go
	GOOS=linux GOARCH=arm go build -o dist/cmk-linux-arm cmk.go
	GOOS=windows GOARCH=amd64 go build -o dist/cmk-x64.exe cmk.go
	GOOS=windows GOARCH=386 go build -o dist/cmk-x32.exe cmk.go
	GOOS=darwin GOARCH=amd64 go build -o dist/cmk-mac64.bin cmk.go
	GOOS=darwin GOARCH=386 go build -o dist/cmk-mac32.bin cmk.go

clean:
	@rm -f cmk

