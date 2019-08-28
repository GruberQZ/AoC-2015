var fs = require("fs");

// Read from input file
fs.readFile("input.txt", function(err, data) {
	// Convert data to string and split by line
	var strings = data.toString('utf-8').split("\n");
	// Nice string counter
	var niceStringsPart1 = 0;
	var niceStringsPart2 = 0;
	// Test each string for niceness
	for (var i = 0; i < strings.length; i++) {
		s = strings[i];
		if (containsThreeVowels(s) && containsDouble(s) && !containsBadStrings(s)) {
			niceStringsPart1++;
		}
		if (containsRepeatingDouble(s) && containsDoubleWithSeparation(s)) {
			niceStringsPart2++;
		}
	}
	console.log("Part 1: There are", niceStringsPart1, "nice strings");
	console.log("Part 2: There are", niceStringsPart2, "nice strings");
});

var containsThreeVowels = function(s) {
	var vowelCount = 0;
	for (var i = 0; i < s.length; i++) {
		var char = s[i];
		if (char === "a" || char === "e" || char === "i" || char === "o" || char === "u") {
			vowelCount++;
			if (vowelCount === 3) {
				return true;
			}
		}
	}
	return false;
}

var containsRepeatingDouble = function(s) {
	for (var i = 0; i < s.length - 1; i++) {
		var firstPair = s.substring(i, i + 2);
		for (var j = i + 2; j < s.length - 1; j++) {
			if (firstPair === s.substring(j, j + 2)) {
				return true;
			}
		}
	}
	return false;
}

var containsDouble = function(s) {
	for (var i = 0; i < s.length - 1; i++) {
		if (s[i] === s[i + 1]) {
			return true;
		}
	}
	return false;
}

var containsDoubleWithSeparation = function(s) {
	for (var i = 0; i < s.length - 2; i++) {
		if (s[i] === s[i + 2]) {
			return true;
		}
	}
	return false;
}

var containsBadStrings = function(s) {
	for (var i = 0; i < s.length - 1; i++) {
		var pair = s.substring(i, i + 2);
		if (pair === "ab" || pair === "cd" || pair === "pq" || pair === "xy") {
			return true;
		}
	}
	return false;
}