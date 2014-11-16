var chart1;
var lastGC = 0;

function requestData() {
	$.ajax({
		url: "/debug/vars",
		success: function(data) {
			var pauses = [];

			if (lastGC == 0) {
				lastGC = data.memstats.NumGC-256;
				if (lastGC < 0) {
					lastGC = 1;
				}
			}

			for (i=lastGC; i<=data.memstats.NumGC; i++) {
				pauses[i] = data.memstats.PauseNs[(i+255)%256];
				console.log("added pause", i)
			}

			for (i in pauses) {
				chart1.series[0].addPoint(pauses[i], false, false);
			}

			lastGC = data.memstats.NumGC+1;

			chart1.redraw();

			// memory in use

			var now = (new Date()).getTime();
			var memAllocated = data.memstats.Alloc;
			chart2.series[0].addPoint([now, memAllocated], false, false);
			chart2.redraw();

			setTimeout(requestData, 10000);
		},
		cache: false
	});
}

$(document).ready(function() {
	chart1 = new Highcharts.Chart(
			{
				chart: {
					renderTo: 'container1',
					zoomType: 'x',
					events: {
						load: requestData
					}
				},
				title: {
					text: 'GC pauses'
				},
				subtitle: {
					text: document.ontouchstart === undefined ?
						'Click and drag in the plot area to zoom in' :
						'Pinch the chart to zoom in'
				},
				yAxis: {
					title: {
						text: 'nanoseconds'
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
					name: 'GC pauses',
					data: [],
				}]
			}
	);
	chart2 = new Highcharts.Chart(
			{
				chart: {
					type: 'spline',
					renderTo: 'container2',
					zoomType: 'x',
					events: {
						load: requestData
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
				xAxis: {
					type: 'datetime'
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
	);
})
