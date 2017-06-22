var crypto = require("crypto");

var input = "iwrupvqb";

var count = -1;

var md5hash = "";

while (!md5hash.startsWith("00000")) {
	count++;
	md5hash = crypto.createHash("md5").update(input + count).digest("hex");
}

console.log("Part 1:", count);

while (!md5hash.startsWith("000000")) {
	count++;
	md5hash = crypto.createHash("md5").update(input + count).digest("hex");
}

console.log("Part 2:", count);