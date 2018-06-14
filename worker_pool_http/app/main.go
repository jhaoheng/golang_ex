package main

import "time"
import "fmt"
import "net/http"

/*
最高的工作池數量
*/
const maxWorker = 5

/*
worker ...
在 jobs 中取得任務
*/
func worker(workerNumber int, jobs <-chan string, results chan<- int) {
  for range jobs {
    <-jobs
    fmt.Println("Worker number is : ", workerNumber, ", started  job at", time.Now().Format("04:05"))
    time.Sleep(time.Second * 1)
    fmt.Println("Worker number is : ", workerNumber, ", finished job at", time.Now().Format("04:05"))
    // results <- workerNumber
  }
}

func main() {

  // 建立兩種 channel
  jobs := make(chan string, 100)
  results := make(chan int, 100)

  // 開啟三個 worker 在 goroutine. 因 channel 在等待參數, 所以一開始狀態是 blocked
  // 將 workerNumber 設定成 1, 得到最後的消耗時間會所有不同
  fmt.Println("Worker 數量 : ", maxWorker)
  for workerNumber := 1; workerNumber <= maxWorker; workerNumber++ {
    go worker(workerNumber, jobs, results)
  }

  http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Get request at : ", time.Now().Format("04:05"))
    jobs <- "task"
    w.WriteHeader(200)
  })

  http.ListenAndServe(":3000", nil)
}
