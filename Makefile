all: bindata.go

bindata.go: static/*.html static/*.js
	go-bindata -pkg='bindata' -o bindata/bindata.go static/

clean:
	rm bindata/bindata.go
