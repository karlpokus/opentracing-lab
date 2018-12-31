// http://mongodb.github.io/node-mongodb-native/3.1/api/MongoClient.html
const mongo = require('mongodb').MongoClient;
const opts = { useNewUrlParser: true };
//let client;

/*function exit() {
	if (client) {
		client.close();
		console.log('client closed');
	}
	process.exit(0);
}

mongo.connect(url, opts)
	.then(c => {
		client = c;
		const users = client.db(db).collection(collection);

		return users.find().toArray()
			.then(docs => {
				docs.forEach(doc => console.log(doc.name));
				exit();
			})
	})
	.catch(err => {
		console.error(err);
		exit();
	})*/

const db = {
	start: url => mongo.connect(url, opts)
};
