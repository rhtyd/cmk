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

clean:
	@rm -f cmk

