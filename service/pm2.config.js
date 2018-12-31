module.exports = {
	apps: [
		{
			name: "api ",
			script: "./api/bin/api",
			watch: "./api/bin/api",
			interpreter: "none",
			log_date_format: 'YYYY-MM-DD HH:mm'
		},
		{
			name: "auth",
			script: "./auth/server.js",
			watch: "./auth",
			ignore_watch: "node_modules",
			log_date_format: 'YYYY-MM-DD HH:mm'
		},
		{
			name: "pets",
			script: "./pets/bin/pets",
			watch: "./pets/bin/pets",
			interpreter: "none",
			log_date_format: 'YYYY-MM-DD HH:mm'
		}
	]
};
