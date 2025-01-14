build:
	GOOS=linux GOARCH=amd64 go build -trimpath -o target/powermenu-linux-amd64/powermenu .

package: build
	mkdir ./dist
	tar -czf ./dist/powermenu-linux-amd64.tar.gz --transform 's/^./powermenu-linux-amd64/' -C ./target/powermenu-linux-amd64 .
	sha256sum ./dist/powermenu-linux-amd64.tar.gz | awk '{print $$1}' > ./dist/powermenu-linux-amd64.tar.gz.checksum

clean:
	go clean
	rm -rf powermenu target dist

install: build
	mv target/powermenu-linux-amd64/powermenu /usr/bin

install-local: build
	mv target/powermenu-linux-amd64/powermenu /usr/local/bin

test:
	go test ./...

uninstall:
	rm -f /usr/bin/powermenu /usr/local/bin/powermenu
