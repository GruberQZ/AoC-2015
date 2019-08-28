var fs = require("fs");

// Read from input file
fs.readFile("input.txt", function(err, data) {
	// Convert data to string and split by line
	var directions = data.toString('utf-8');
	// Tracker for visited houses
	var visited = new Set();
	var visited2 = new Set();
	// Start at the origin
	var part1Position = {
		x: 0,
		y: 0
	};
	var santaPosition = {
		x: 0,
		y: 0
	};
	var roboSantaPosition = {
		x: 0,
		y: 0
	};
	// Starting location gets a present before directions begin
	visited.add(part1Position.x.toString() + "_" + part1Position.y.toString());
	visited2.add(santaPosition.x.toString() + "_" + santaPosition.y.toString());
	// Iterate through all of the directions
	for (var i = 0; i < directions.length; i++) {
		var dir = directions[i];
		var santaTurn = i % 2;
		if (dir === "^") {
			part1Position.y += 1;
			if (santaTurn) {
				santaPosition.y += 1;
			} else {
				roboSantaPosition.y += 1;
			}
		} else if (dir === "v") {
			part1Position.y -= 1;
			if (santaTurn) {
				santaPosition.y -= 1;
			} else {
				roboSantaPosition.y -= 1;
			}
		} else if (dir === "<") {
			part1Position.x -= 1;
			if (santaTurn) {
				santaPosition.x -= 1;
			} else {
				roboSantaPosition.x -= 1;
			}
		} else if (dir === ">") {
			part1Position.x += 1;
			if (santaTurn) {
				santaPosition.x += 1;
			} else {
				roboSantaPosition.x += 1;
			}
		}
		visited.add(part1Position.x.toString() + "_" + part1Position.y.toString());
		if (santaTurn) {
			visited2.add(santaPosition.x.toString() + "_" + santaPosition.y.toString());
		} else {
			visited2.add(roboSantaPosition.x.toString() + "_" + roboSantaPosition.y.toString());
		}
	}
	console.log("Part 1: This year,", visited.size, "houses receive presents");
	console.log("Part 2: The next year,", visited2.size, "houses will receive presents");
});