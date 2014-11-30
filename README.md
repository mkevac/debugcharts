debugcharts
===========

Go memory debug charts.

This package uses chart library [Highcharts](http://www.highcharts.com/). It is free for personal and non-commercial use. Please buy a license otherwise.

Installation
------------
`go get -v -u github.com/mkevac/debugcharts`

Usage
-----
Just install package and start http server. There is an example program [here](https://github.com/mkevac/debugcharts/blob/master/example/example.go).

Then go to `http://localhost:8080/debug/charts`. You should see something like this:
<img src="example/screenshot.png" />

Data is updated every second. We keep data for last day.
