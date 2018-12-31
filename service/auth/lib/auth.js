// grabs user and pwd from authorization header
// returns an empty array on failure
function basicAuth(authHeader) {
	if (/^Basic/.test(authHeader)) {
		const b64 = authHeader.split(" ")[1];
		return Buffer.from(b64, 'base64').toString().split(":");
	}
	return [];
}

module.exports = basicAuth;
