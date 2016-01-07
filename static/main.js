var chart1;
var chart2;
var chart2Div = document.getElementById('container2');

$(function () {

	$.getJSON('/debug/charts/data?callback=?', function (data) {
		var pDataChart1 = [{
			x: [],
			y: [],
			type: "scatter"
		}];
		for (i = 0; i < data.GcPauses.length; i++) {
			var d = moment(data.GcPauses[i][0]).format('HH:mm:ss');
			pDataChart1[0].x.push(d);
			pDataChart1[0].y.push(data.GcPauses[i][1]);
		}
		chart1 = Plotly.newPlot('container1', pDataChart1, {
            title: "GC Pauses",
            yaxis: {
                title: "Nanoseconds"
            }
        });

		var pDataChart2 = [{
			x: [],
			y: [],
			type: "scatter"
		}];
		for (i = 0; i < data.BytesAllocated.length; i++) {
			var d = moment(data.BytesAllocated[i][0]).format('HH:mm:ss');
			pDataChart2[0].x.push(d);
			pDataChart2[0].y.push(data.BytesAllocated[i][1]);
		}
		chart2 = Plotly.newPlot('container2', pDataChart2, {
            title: "Memory Allocated",
            yaxis: {
                title: "Bytes"
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
            var d = moment(data.Ts).format('HH:mm:ss');
			if (data.GcPause != 0) {
                Plotly.extendTraces('container1', {x:[[d]],y:[[data.GcPause]]}, [0], 86400);
			}
            Plotly.extendTraces('container2', {x:[[d]],y:[[data.BytesAllocated]]}, [0], 86400);
		}
	};
})
