var datasets = {
	"A": {
		label: "A",
		fillColor: "rgba(220,220,220,0.2)",
		strokeColor: "rgba(220,220,220,1)",
		pointColor: "rgba(220,220,220,1)",
		pointStrokeColor: "#fff",
		pointHighlightFill: "#fff",
		pointHighlightStroke: "rgba(220,220,220,1)",
		data: [65, 59, 80, 81, 56, 55, 40]
	},
	"B": {
		label: "B",
		fillColor: "rgba(151,187,205,0.2)",
		strokeColor: "rgba(151,187,205,1)",
		pointColor: "rgba(151,187,205,1)",
		pointStrokeColor: "#fff",
		pointHighlightFill: "#fff",
		pointHighlightStroke: "rgba(151,187,205,1)",
		data: [28, 48, 40, 19, 86, 27, 90]
	},
	"C": {
		label: "C",
		fillColor: "rgba(151,187,205,0.2)",
		strokeColor: "rgba(151,187,205,1)",
		pointColor: "rgba(151,187,205,1)",
		pointStrokeColor: "#fff",
		pointHighlightFill: "#fff",
		pointHighlightStroke: "rgba(151,187,205,1)",
		data: [28, 48, 40, 19, 86, 27, 90]
	}
};

var data = {
	labels: ["January", "February", "March", "April", "May", "June", "July"],
	datasets: [datasets.A, datasets.B, datasets.C]
};
var $chart = $("<canvas>").attr({width:"800", height:"400"});
var ctx = $chart.get(0).getContext("2d");
var myLineChart = new Chart(ctx).Line(data);

var app = angular.module('app', []);

app.controller('displays', function($scope) {
	$scope.displays = {
		"A": {
			"x": 1,
			"y": 2
		}
	};

	var point;
	// Create a socket
	var socket = new WebSocket('ws://'+window.location.host+'/socket');

	socket.onmessage = function(event) {
		point = JSON.parse(event.data);

		myLineChart.addData([4, point.XVal*10, point.YVal*10], point.Label);
		myLineChart.removeData();

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
