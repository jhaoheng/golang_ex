> https://golang.org/pkg/time/#Timer


# type Timer

```
type Timer struct {
	C <-chan Time
}
```

- Timer 類型是一種事件，當 Timer 到期，目前的時間會送到 `C` 中，除非 Timer 被 `AfterFunc` 重新建立。
- Timer 必須透過 NewTimer or AfterFunc 建立 

## Table list
- `func AfterFunc` : 建立一個 Timer 並在時間到後，執行 Func
- `func NewTimer` : 建立一個 Timer
- `func (*Timer) Reset` : 重設置 Timer
- `func (*Timer) Stop` : 命令 Timer 停止
- 範例
	- Timer 第一次到期為兩秒後, 到期後自動 t.Reset(10), 再度自動倒數
	- 每次執行 t.Reset(10) 後, 到期都會自動倒數
	- 按下 Enter 後, Timer 會直接中斷倒數, 並且重新執行 t.Reset(10)
	- ctrl+c 停止運行

## `func AfterFunc`
> `func AfterFunc(d Duration, f func()) *Timer`

- `func AfterFunc` : 建立一個 Timer, 當時間到的時候，會呼叫 f() 在 Timer 自己的 goroutine 中。可以用 Stop 來停止。 

```
func main() {
	var t *time.Timer
	f := func() {
		fmt.Printf("Expiration time : %v.\n", time.Now())
		fmt.Printf("C`s len: %d\n", len(t.C))
	}
	t = time.AfterFunc(2*time.Second, f)

	var stop string
	fmt.Scanln(&stop)
}
```

## `func NewTimer`
> `func NewTimer(d Duration) *Timer`

- `func NewTimer` : 建立一個 new Timer, 會將目前的時間送到 C 中 (C 是 Timer struct 的 channel)
- 所以可以透過 `<-t.C` 來等待時間到期

```
func main() {
	// init
	t := time.NewTimer(2 * time.Second)
	// Current time
	now := time.Now()
	fmt.Printf("time : %v.\n", now)

	expire := <-t.C
	fmt.Printf("time : %v.\n", expire)
}
```

## `func (*Timer) Stop`
> `func (t *Timer) Stop() bool`

- `func (*Timer) Stop` : 用於'主動'停止計時器, 當 Timer 尚未停止的時候調用
- 若 timer 已經停止，則 t.Stop() 會返回 false
- stop 並不會 close the channel, 為了防止不正確地讀取 channel, 所以為了要確保在呼叫 stop() 後 channel 是空的, 必須確認 stop() 的 return, 若是 false, 則必須讓 channel 為空。


```
if !t.Stop() {
	<-t.C // 釋放掉 channel 的物件
}
``` 

- Stop() 並不會停止其他 Timer's channel
- 對於透過 AfterFunc(d,f) 產生的 Timer, 若 t.Stop() 返回 false (代表 timer 已經過期, 且 f 已經在他自己的 goroutine 啟動), Stop 並不會等待 f 完成. 若需要知道 f 是否已經完成, 必須在 f 中做一些協調處理.

## `func (*Timer) Reset`
> `func (t *Timer) Reset(d Duration) bool`

- Reset 的調用, 有兩個條件
	1. timer 必須停止或者過期
	2. channel 必須為空
- Reset 的使用前, 若 Timer 的狀態是過期或已經停止, 則必須確定 t.C 的 channel 為空.
- 程式可以透過 `<-t.C` 來判斷是否時間到期，此時 channel 為空, 在此狀況下, t.Reset() 可以直接被使用.

```
func main() {
	now := time.Now()
	fmt.Printf("           time: %v.\n", now)

	t := time.NewTimer(2 * time.Second)
	expire := <-t.C
	fmt.Printf("Expiration time: %v.\n", expire)
	t.Reset(5 * time.Second)
	expire = <-t.C
	fmt.Printf("Expiration time: %v.\n", expire)
}
```

- 若程式沒有從 `<-t.C` 中接收參數, 但 timer 卻已經使用 stop(), 此時必須讓 channel 為空, 才可以使用.

```
if !t.Stop() {
	<-t.C
}
t.Reset(d)
```

# Example

> 無法透過 go playground, 因為無法讀取 Enter 鍵

```
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
			fmt.Println(". Interrupt Timer!!!!!")
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
```

[執行結果]
=====> 1.002051028s
=====> 2.006385502s. Time up, auto reset Timer
Timer Reset 10s
=====> 1.002283105s
=====> 2.005053556s
=====> 3.007041044s
=====> 4.009081465s

=====> 5.010377978s. Interrupt Timer!!!!!
Timer Stop Success!!  Timer Reset 10s
=====> 1.003484026s
^Csignal: interrupt

