var fs = require("fs");

// Regular expression for extracting dimensions
var re = /(\d+)x(\d+)x(\d+)/;

// Read from input file
fs.readFile("input.txt", function(err, data) {
	// Convert data to string and split by line
	var dimSets = data.toString('utf-8').split("\n");
	// Keep track of total needed
	var wrappingPaper = 0;
	var ribbon = 0;
	// Iterate through sets of dimensions
	for (var i = 0; i < dimSets.length; i++) {
		// Match the dimensions in this set to the regular expression
		var match = re.exec(dimSets[i]);
		var l = parseInt(match[1]);
		var w = parseInt(match[2]);
		var h = parseInt(match[3]);
		// Side calculations
		var lw = l*w;
		var lh = l*h;
		var wh = w*h;
		var lwh = l*w*h;
		var smallestTwoSides = findSmallestTwoValues([l, w, h]);
		// WP and Ribbon calculations
		wrappingPaper += 2*lw + 2*wh + 2*lh + smallestTwoSides.smallest*smallestTwoSides.secondSmallest;
		ribbon += lwh + 2*smallestTwoSides.smallest + 2*smallestTwoSides.secondSmallest;
	}
	// Print result
	console.log("Part 1: The elves should order", wrappingPaper, "square feet of wrapping paper");
	console.log("Part 2: The elves should order", ribbon, "feet of ribbon");
});

// Function to find the smallest two dimensions
var findSmallestTwoValues = function(sides) {
	var first = Number.MAX_SAFE_INTEGER, second = Number.MAX_SAFE_INTEGER;
	for (var i = 0; i < sides.length; i++) {
		if (sides[i] < first) {
			second = first;
			first = sides[i];
		} else if (sides[i] < second) {
			second = sides[i];
		}
	}
	return {
		smallest: first,
		secondSmallest: second
	};
}
