all: bindata.go

bindata.go: static/index.html static/main.js
	go-bindata -pkg='debugcharts' static/

clean:
	rm bindata.go
