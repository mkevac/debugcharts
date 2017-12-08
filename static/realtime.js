var chart1;
var chart2;
var chart3;

$(function () {

	var data1 = [
		{
			label: "GC",
			values: []
		}
	];
	var data2 = [
		{
			label: "BytesAllocated",
			values: []
		}
	];
	var data3 = [
		{
			label: "User",
			values: []
		},
		{
			label: "Sys",
			values: []
		}
	];

	chart1 = $('#container1').epoch({
		type: 'time.bar',
		data: data1,
		axes: ["right", "bottom", "left"],
		historySize: 600,
		windowSize: 300
	});
	chart2 = $('#container2').epoch({
		type: 'time.area',
		data: data2,
		axes: ["right", "bottom", "left"],
		historySize: 600,
		windowSize: 300
	});
	chart3 = $('#container3').epoch({
		type: 'time.bar',
		data: data3,
		axes: ["right", "bottom", "left"],
		historySize: 600,
		windowSize: 300
	});

	function wsurl() {
		var l = window.location;
		return ((l.protocol === "https:") ? "wss://" : "ws://") + l.hostname + (((l.port != 80) && (l.port != 443)) ? ":" + l.port : "") + "/debug/charts/data-feed";
	}

	ws = new WebSocket(wsurl());
	ws.onopen = function () {
		ws.onmessage = function (evt) {
			var data = JSON.parse(evt.data);
			now = Date.now()
			if (data.GcPause != 0) {
				chart1.push([{time: now, y: data.GcPause}]);
			}
			chart2.push([{time: now, y: data.BytesAllocated}]);
			chart3.push([{time: now, y: data.CpuUser}, {time: now, y: data.CpuSys}]);
		}
	};
})
