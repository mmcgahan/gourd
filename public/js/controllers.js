var datasets = [{
	label: "A",
	fillColor: "rgba(220,220,220,0.2)",
	strokeColor: "rgba(220,220,220,1)",
	pointColor: "rgba(220,220,220,1)",
	pointStrokeColor: "#fff",
	pointHighlightFill: "#fff",
	pointHighlightStroke: "rgba(220,220,220,1)",
	data: [6, 5, 8, 8, 5, 5, 4, 6, 5, 8, 8, 5, 5, 4]
}, {
	label: "B",
	fillColor: "rgba(151,187,205,0.2)",
	strokeColor: "rgba(151,187,205,1)",
	pointColor: "rgba(151,187,205,1)",
	pointStrokeColor: "#fff",
	pointHighlightFill: "#fff",
	pointHighlightStroke: "rgba(151,187,205,1)",
	data: [2, 4, 4, 1, 8, 2, 9, 6, 5, 8, 8, 5, 5, 4]
}, {
	label: "C",
	fillColor: "rgba(205,187,151,0.2)",
	strokeColor: "rgba(151,187,205,1)",
	pointColor: "rgba(151,187,205,1)",
	pointStrokeColor: "#fff",
	pointHighlightFill: "#fff",
	pointHighlightStroke: "rgba(151,187,205,1)",
	data: [2, 6, 2, 4, 1, 2, 7, 6, 5, 8, 8, 5, 5, 4]
}];

var dataMap = {
	"A": 0,
	"B": 1,
	"C": 2
};

var data = {
	labels: [-12,-11,-10,-9,-8,-7,-6,-5,-4,-3,-2,-1,0],
	datasets: datasets
};
var $chart = $("<canvas>").attr({width:"800", height:"400"});
var ctx = $chart.get(0).getContext("2d");
var myLineChart = new Chart(ctx).Line(data, { animation: false });

var app = angular.module('app', []);

app.controller('displays', function($scope) {
	$scope.displays = { "A": { "x": 1, "y": 2 } };

	var point,
		points,
		oldpoint;
	// Create a socket
	var socket = new WebSocket('ws://'+window.location.host+'/socket');

	socket.onmessage = function(event) {
		// read point object
		point = JSON.parse(event.data);

		Chart.helpers.each(myLineChart.datasets, function (dataset) {
			oldpoint = dataset.points.shift();
			if (point.Label === dataset.label) {
				oldpoint.value = point.XVal;
			} else {
				oldpoint.value = dataset.points[dataset.points.length-1].value;
			}
			dataset.points.push(oldpoint);
		});

		myLineChart.update();

		// call $apply with function arg so that angular captures errors
		// and updates FE bindings
		$scope.$apply(function() {
			$scope.displays[point.Label] = {
				"x": point.XVal,
				"y": point.YVal
			};
		});

	};

	$('.send').click(function(e) {
		socket.send(this.value);
	});
});

$(function() {
	$("body").prepend($chart);
});
