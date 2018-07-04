/*
Package redis ...
  use:

    status := RedisClient.FlushAll()
    log.Println(status)

    RedisClient.Set("name", "max", 0)
    value := RedisClient.Get("name")
    log.Println(value)
*/
package redis

import (
  "app/config"
  "github.com/go-redis/redis"
  "log"
  "strconv"
  "time"
)

/*
RedisClient ...
*/
var RedisClient *redis.Client

/*
RedisNewClient ...
*/
func init() {
  redisConfig := config.Redis()
  addr := redisConfig["addr"].(string)
  password := redisConfig["password"].(string)
  db, _ := strconv.Atoi(redisConfig["db"].(string))
  RedisClient = redis.NewClient(&redis.Options{
    Addr:     addr,
    Password: password, // no password set
    DB:       db,       // use default DB
  })

  pong, err := RedisClient.Ping().Result()
  if err != nil {
    log.Fatal("Redis connect fail. ", err)
  } else {
    log.Println("Redis connection success. ", pong)
  }
  return
}

/*
Set ...
  value : use json to encode. import "encoding/json"
*/
func Set(key string, value interface{}, expiration time.Duration) {
  StringCmd := RedisClient.Set(key, value, expiration)
  log.Println("set : ", StringCmd)
}

/*
Get ...
  val : use json to decode. import "encoding/json"
*/
func Get(key string) (StringCmd *redis.StringCmd) {
  StringCmd = RedisClient.Get(key)
  log.Println("get : ", StringCmd)
  return
}

/*
flushAll ...
*/
func flushAll() {
  RedisClient.FlushAll()
}
