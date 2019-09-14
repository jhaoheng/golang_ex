package main

import (
	"fmt"
	"time"
)

type TimerObj struct {
	timeStart   time.Time
	resetTimeIs time.Duration
	t           time.Timer
}

func main() {

	timerObj := TimerObj{
		timeStart:   time.Now(),
		resetTimeIs: 10,
		t:           *time.NewTimer(2 * time.Second),
	}

	c := make(chan string)
	go func() {
		for {
			var stop string
			fmt.Scanln(&stop)
			c <- "reset"
		}
	}()

	for {
		time.Sleep(1 * time.Second)
		select {
		case <-c:
			showCurrentTime(&timerObj)
			fmt.Println(". Interupt Timer!!!!!")
			resetTimer(&timerObj)
		case <-timerObj.t.C: // 時間到期
			showCurrentTime(&timerObj)
			fmt.Println(". Time up, auto reset Timer")
			resetTimer(&timerObj)
		default:
			showCurrentTime(&timerObj)
			fmt.Println()
		}
	}
}

func showCurrentTime(timerObj *TimerObj) {
	fmt.Printf("=====> %v", time.Now().Sub(timerObj.timeStart))
}

func resetTimer(timerObj *TimerObj) {
	if timerObj.t.Stop() {
		fmt.Printf("Timer Stop Success!!  ")
	}
	fmt.Printf("Timer Reset %v\n", timerObj.resetTimeIs*time.Second)
	timerObj.t.Reset(timerObj.resetTimeIs * time.Second)
	timerObj.timeStart = time.Now()
}
