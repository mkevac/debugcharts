var chart1;
var chart2;

$(function() {
	var x = new Date();

	Highcharts.setOptions({
		global: {
			timezoneOffset: x.getTimezoneOffset()
		}
	})

	$.getJSON('/debug/charts/data?callback=?', function(data) {
		chart1 = new Highcharts.StockChart({
			chart: {
				renderTo: 'container1',
				zoomType: 'x'
			},
			title: {
				text: 'GC pauses'
			},
			yAxis: {
				title: {
					text: 'Nanoseconds'
				}
			},
			scrollbar: {
				enabled: false
			},
			rangeSelector: {
				buttons: [{
					type: 'second',
					count: 5,
					text: '5s'
				}, {
					type: 'second',
					count: 30,
					text: '30s'
				}, {
					type: 'minute',
					count: 1,
					text: '1m'
				}, {
					type: 'all',
					text: 'All'
				}],
				selected: 3
			},
			series: [{
				name: "GC pauses",
				data: data.GcPauses,
				type: 'area',
				tooltip: {
					valueSuffix: 'ns'
				}
			}]
		});
		chart2 = new Highcharts.StockChart({
			chart: {
				renderTo: 'container2',
				zoomType: 'x'
			},
			title: {
				text: 'Memory allocated'
			},
			yAxis: {
				title: {
					text: 'Bytes'
				}
			},
			scrollbar: {
				enabled: false
			},
			rangeSelector: {
				buttons: [{
					type: 'second',
					count: 5,
					text: '5s'
				}, {
					type: 'second',
					count: 30,
					text: '30s'
				}, {
					type: 'minute',
					count: 1,
					text: '1m'
				}, {
					type: 'all',
					text: 'All'
				}],
				selected: 3
			},
			series: [{
				name: "Allocated",
				data: data.BytesAllocated,
				type: 'area',
				tooltip: {
					valueSuffix: 'b'
				}
			}]
		})
	});

	function wsurl() {
		var l = window.location;
		return ((l.protocol === "https:") ? "wss://" : "ws://") + l.hostname + (((l.port != 80) && (l.port != 443)) ? ":" + l.port : "") + "/debug/charts/data-feed";
	}

	ws = new WebSocket(wsurl());
	ws.onopen = function () {
		ws.onmessage = function (evt) {
			var data = JSON.parse(evt.data);
			if (data.GcPause != 0) {
				chart1.series[0].addPoint([data.Ts, data.GcPause], true);
			}
			chart2.series[0].addPoint([data.Ts, data.BytesAllocated], true);
		}
	};
})
