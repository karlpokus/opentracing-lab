package logs

import (
	"log"
	"os"
)

var (
	Stdout *log.Logger = log.New(os.Stdout, "", 0)
	Stderr *log.Logger = log.New(os.Stderr, "", 0)
)
