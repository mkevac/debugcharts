var chart1;
var chart2;
var ws;

var lastGC = 0;
var lastTotalAllocated = 0;

var areaConfDefault = {
	chart: {
		type: 'spline',
		renderTo: 'container2',
		zoomType: 'x',
	},
	loading: {
		labelStyle: {
			top: '45%'
		}
	},
	title: {
		text: 'Memory in use'
	},
	subtitle: {
		text: document.ontouchstart === undefined ?
			'Click and drag in the plot area to zoom in' :
			'Pinch the chart to zoom in'
	},
	yAxis: {
		title: {
			text: 'bytes'
		}
	},
	legend: {
		enabled: false
	},
	plotOptions: {
		area: {
			fillColor: {
				linearGradient: { x1: 0, y1: 0, x2: 0, y2: 1},
				stops: [
					[0, Highcharts.getOptions().colors[0]],
					[1, Highcharts.Color(Highcharts.getOptions().colors[0]).setOpacity(0).get('rgba')]
				]
			},
			marker: {
				radius: 2
			},
			lineWidth: 1,
			states: {
				hover: {
					lineWidth: 1
				}
			},
			threshold: null
		}
	},

	series: [{
		type: 'area',
		name: 'Memory in use',
		data: [],
	}]
}

function updateGC(data) {
	var res = [];
	for (var i=0; i<data.GcPauses.length; i++) {
		res.push([data.GcPauses[i].T, data.GcPauses[i].C])
	}
	chart1.series[0].setData(res, true);
}

function updateMemAllocated(data) {
	var res = [];
	for (var i=0; i<data.BytesAllocated.length; i++) {
		res.push([data.BytesAllocated[i].T, data.BytesAllocated[i].C])
	}
	chart2.series[0].setData(res, true);
}

function requestData() {
	$.ajax({
		url: "/debug/charts/data",
		success: function(data) {
			updateGC(data);
			updateMemAllocated(data);
		},
		cache: false
	});
}

function loadCallback() {
	setTimeout(requestData, 100);
}

$(window).unload(function() {
	ws.close();
});

$(document).ready(function() {
	var conf = $.extend(true, {}, areaConfDefault);

	conf.chart.renderTo = 'container1';
	conf.chart.events = {load: loadCallback};
	conf.title.text = 'GC pauses';
	conf.yAxis.title.text = 'nanoseconds';
	conf.xAxis = {type: 'datetime'};
	conf.series[0].name = 'GC pauses';

	chart1 = new Highcharts.Chart(conf);

	var conf = $.extend(true, {}, areaConfDefault);

	conf.chart.renderTo = 'container2';
	conf.title.text = 'Memory in use';
	conf.yAxis.title.text = 'bytes';
	conf.xAxis = {type: 'datetime'};
	conf.series[0].name = 'Memory in use';

	chart2 = new Highcharts.Chart(conf);

	var conf = $.extend(true, {}, areaConfDefault);

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
	}
})
