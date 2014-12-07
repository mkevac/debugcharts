all: bindata.go

bindata.go: static/index.html static/main.js
	go-bindata -pkg='bindata' -o bindata/bindata.go static/

clean:
	rm bindata/bindata.go
