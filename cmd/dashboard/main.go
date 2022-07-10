package main

import (
	"github.com/checkmeup/mono/internal/log"
)

func main() {
	l := log.Default()
	l.Debug("Hello %s", "World")
}
