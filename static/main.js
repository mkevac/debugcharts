var chart1;
var chart2;

$(function() {
	var x = new Date();
	var BytesAllocated = [];
	var GcPauses = [];

	$.getJSON('/debug/charts/data?callback=?', function(data) {
		BytesAllocated = data.BytesAllocated;
		GcPauses = data.GcPauses;

		chart1 = $.plot("#container1", [GcPauses], {
			series: {
				lines: {show: true},
				points: {show: true},
				shadowSize: 0
			},
			xaxis: {
				mode: "time"
			}
		});
		chart2 = $.plot("#container2", [BytesAllocated], {
			series: {
				lines: {show: true},
				points: {show: true},
				shadowSize: 0
			},
			xaxis: {
				mode: "time"
			}
		});
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
				GcPauses.push([data.Ts, data.GcPause]);
				chart1.setData([GcPauses]);
				chart1.setupGrid();
				chart1.draw();
			}

			BytesAllocated.push([data.Ts, data.BytesAllocated]);
			chart2.setData([BytesAllocated]);
			chart2.setupGrid();
			chart2.draw();
		}
	};
})
