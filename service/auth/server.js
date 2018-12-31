#!/usr/bin/env node

const mongo = require('mongodb').MongoClient;
const mongoConnString = 'mongodb://localhost:27017';
const mongoConnOpts = { useNewUrlParser: true };
const http = require('http');
const server = http.createServer();
const PORT = process.env.PORT || 9112;
const userCredentials = require('./lib/auth');
const success = res => {
	log('success');
	res.writeHead(200);
	res.end('ok');
};
const fail = (err, res) => {
	log(err.message, 'error');
	res.writeHead(401);
	res.end('Unauthorized');
};
const exit = mongoClient => {
	if (mongoClient) {
		mongoClient.close();
		log('mongoClient closed');
	}
	process.exit(0);
};

function log(msg, type) {
	//const prefix = `auth ${ new Date().toISOString() }`;
	console[type || 'log'](msg);
}

function authenticate(users, req, res) {
	//log(`${ req.method } ${ req.url }`);

	const [name, pwd] = userCredentials(req.headers.authorization);
	if (name && pwd) {
		return users.findOne({ name, pwd })
			.then(doc => {
				if (!doc) {
					return Promise.reject(new Error(`ERR no match for ${ name }:${ pwd }`));
				}
				success(res);
			})
			.catch(err => fail(err, res));
	}
	fail(new Error('ERR user credentials missing from auth header'), res);
}

function startServer(users) {
	server
		.on('request', authenticate.bind(null, users))
		.listen(PORT, () => log(`listening on port ${ PORT }`))
}

mongo.connect(mongoConnString, mongoConnOpts)
	.then(mongoClient => {
		log(`connected to ${ mongoConnString }`);
		process.on('SIGINT', exit.bind(null, mongoClient));
		return Promise.resolve(mongoClient.db('pets').collection('users'));
	})
	.then(startServer)
	.catch(err => {
		throw err;
		process.exit(1);
	});
