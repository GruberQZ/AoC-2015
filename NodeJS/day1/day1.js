var fs = require("fs");

// Read from input file
fs.readFile("input.txt", function(err, data) {
	var insts = data.toString('utf-8');
	var count = 0;
	var basementInst = null;
	for (var i = 0; i < insts.length; i++) {
		// Read each instruction
		if (insts[i] === "(") {
			count += 1
		} else {
			count -= 1
		}
		// Find the position of the instruction that puts Santa in the basement for the first time
		if (count === -1 && basementInst === null) {
			// Add 1 to i because first instruction is position 1 rather than 0
			basementInst = i + 1;
		}
	}
	console.log("The instructions take Santa to floor", count);
	console.log("The position of the character that causes Santa to first enter the basement is", basementInst);
});