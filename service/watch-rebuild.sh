#!/bin/bash

# https://github.com/karlpokus/pibox/blob/master/watch.sh
# https://github.com/kimmobrunfeldt/chokidar-cli
# -c "if [ '{event}' = 'change' ]; then npm run build-less -- {path}; fi;"

chokidar "**/*.go" -c "go build -o api/bin/api api/*.go; go build -o pets/bin/pets pets/*.go "

# prints to stdout even build errors
