var fs = require("fs");

var re = /([^0-9]+)\s(\d+),(\d+)\sthrough\s(\d+),(\d+)/;

// Read from input file
fs.readFile("input.txt", function(err, data) {
	// Convert data to string and split by line
	var insts = data.toString('utf-8').split("\n");
	// Make grid of lights
	// Initialize all the lights to off
	var lights = new Array(1000);
	var brightLights = new Array(1000);
	for (var i = 0; i < 1000; i++) {
		lights[i] = new Array(1000);
		brightLights[i] = new Array(1000);
		for (var j = 0; j < 1000; j++) {
			lights[i][j] = false;
			brightLights[i][j] = 0;
		}
	}
	// Iterate through all the instructions
	for (var i = 0; i < insts.length; i++) {
		var match = re.exec(insts[i]);
		var action = match[1];
		var xLow = parseInt(match[2]);
		var yLow = parseInt(match[3]);
		var xHigh = parseInt(match[4]);
		var yHigh = parseInt(match[5]);
		// Transition lights
		if (action === "turn on") {
			for (var j = xLow; j <= xHigh; j++) {
				for (var k = yLow; k <= yHigh; k++) {
					lights[j][k] = true;
					brightLights[j][k] += 1;
				}
			}
		} else if (action === "turn off") {
			for (var j = xLow; j <= xHigh; j++) {
				for (var k = yLow; k <= yHigh; k++) {
					lights[j][k] = false;
					if (brightLights[j][k] > 0) {
						brightLights[j][k] -= 1;
					}
				}
			}
		} else if (action === "toggle") {
			for (var j = xLow; j <= xHigh; j++) {
				for (var k = yLow; k <= yHigh; k++) {
					lights[j][k] = !lights[j][k];
					brightLights[j][k] += 2;
				}
			}
		} else {
			console.log("Didn't perform action", action);
		}
	}
	// Count the number of lights that are on
	var count = 0;
	var brightness = 0;
	for (var j = 0; j < 1000; j++) {
		for (var k = 0; k < 1000; k++) {
			if (lights[j][k]) {
				count++;
			}
			brightness += brightLights[j][k];
		}
	}
	console.log("Part 1: There are", count, "lights on");
	console.log("Part 2: The total brightness is", brightness);
});