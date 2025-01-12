build:
	GOOS=linux GOARCH=amd64 go build -trimpath -o target/powermenu-linux-amd64/powermenu .

clean:
	go clean
	rm -rf powermenu target

install: build
	mv target/powermenu-linux-amd64/powermenu /usr/bin

install-local: build
	mv target/powermenu-linux-amd64/powermenu /usr/local/bin

test:
	go test ./...

uninstall:
	rm -f /usr/bin/powermenu /usr/local/bin/powermenu
