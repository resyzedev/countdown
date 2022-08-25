package main

import (
	"fmt"
	"time"
)

const (
	countdownInit = 50
)

type Delayer interface {
	Delay()
}

type ConfigDelay struct {
	duration time.Duration
	delay    func(time.Duration)
}

func (config *ConfigDelay) Delay() {
	config.delay(config.duration)
}

func Countdown(delayer Delayer) {
	for i := countdownInit; i > 0; i-- {
		delayer.Delay()
		fmt.Println(i)
	}
	fmt.Println("OK")
}

func main() {
	Countdown(&ConfigDelay{1 * time.Second, time.Sleep})
}
