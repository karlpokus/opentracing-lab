# opentracing lab
Experimenting with the open tracing api and adding some metrics.

# wat?
We have 3 http services; A main api service that authenticates request via an auth service and if successful calls the pets service. Both the auth service and pets use mongodb for persistance.

The plan is to add tracing with jaeger and metrics with prometheus and grafana.

# requirements
- jaeger
- mongodb
- prometheus
- grafana
- pm2 for process management, unified logging and watch bin and restart
- chokidar for watch src and build

# usage
Run everything
```bash
$ pm2 start pm2.config.js && pm2 logs
$ ./watch-rebuild.sh
```
Make a request
```bash
$ curl -i -u user:pwd [-d '{"key":"value"}'] host:port/path
```
HTTP api
```
GET /api/pet/:name
	returns pet by name
GET /api/pets
	returns all pets
POST /api/pets/add {name: string, type: string, born: YYYY-MM-DDTHH:MM:SSZ}
	adds a pet
```
ports
```bash
api 9111
auth 9112
pets 9113
jaeger 16686
```

# todos
- [x] api
- [x] auth
- [x] pets
- [ ] collect metrics with prometheus
- [ ] pass ctx to rpc calls
- [ ] add debug logs
- [x] add supervisord or pm2
- [x] watch, rebuild, restart for dev
- [x] shared logs
- [ ] shared config file
- [x] better service structure

# license
MIT
