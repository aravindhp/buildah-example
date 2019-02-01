package main

import (
	"github.com/containers/buildah"
	log "github.com/sirupsen/logrus"
)

func main() {
	isolation := buildah.IsolationDefault
	log.Info("%s", isolation)
}
