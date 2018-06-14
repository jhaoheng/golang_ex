# 目的
在 http 下, 透過不同的 worker pool, 觀察是否處理的速度有差別


# 使用
1. `docker-compose up -d`
2. `docker exec -it app /bin/bash`
3. 設定 main.go 中的 maxWorker
4. `go run main.go`
5. 執行 vegeta 進行壓力測試

# vegeta
> 須先安裝 vegeta : https://github.com/tsenart/vegeta

1. `echo "GET http://127.0.0.1:3000/test" |vegeta attack -duration=10s -timeout=11s -rate=15 >results.bin`
2. 觀看 report
  - `vegeta report -inputs=results.bin`

## 測試結果 worker : 1
> 出現 error : timeout awaiting response headers

```
Requests      [total, rate]            150, 15.10
Duration      [total, attack, wait]    20.942179623s, 9.933332234s, 11.008847389s
Latencies     [mean, 50, 95, 99, max]  5.72066402s, 6.806991181s, 11.010648448s, 11.011096082s, 11.011208737s
Bytes In      [total, mean]            0, 0.00
Bytes Out     [total, mean]            0, 0.00
Success       [ratio]                  56.00%
Status Codes  [code:count]             200:84  0:66  
Error Set:
Get http://127.0.0.1:3000/test: net/http: timeout awaiting response headers
```

## 測試結果 worker : 5
```
Requests      [total, rate]            150, 15.10
Duration      [total, attack, wait]    9.937430225s, 9.933332234s, 4.097991ms
Latencies     [mean, 50, 95, 99, max]  4.850239ms, 4.862287ms, 7.570609ms, 8.48227ms, 9.782561ms
Bytes In      [total, mean]            0, 0.00
Bytes Out     [total, mean]            0, 0.00
Success       [ratio]                  100.00%
Status Codes  [code:count]             200:150  
Error Set:
```
