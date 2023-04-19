package main

import (
	"github.com/WALL-EEEEEEE/gagdets/core"
	_ "github.com/WALL-EEEEEEE/gagdets/sources"
	log "github.com/sirupsen/logrus"
)

func main() {
	go core.Exec.Run()
	core.Exec.Output(func(collector *core.Collector) {
		cnt := 1
		for item := range *collector {
			log.Debug(item)
			cnt++
		}
		log.Infof("Found %d topics!", cnt)
	})
}
