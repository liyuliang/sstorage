package service

import (
	"queue-services"
)

func Start() {

	initQueue()

	services.AddSingleProcessTask("Pull Job", func(workerNum int) (err error) {
		return
	})

	services.Service().Start(true)
}

func initQueue() {

}
