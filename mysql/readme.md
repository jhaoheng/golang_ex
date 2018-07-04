# example : Create(insert), read(fetch), update and delete

## how to implement
```
import (
    "app/module/mysql"
    "log"
)

db := mysql.Open()
var query string
query = "select id, token from cvr_token"
data, err := mysql.Fetchall(db, query)
if err != nil {
    log.Fatal(err)
}

fmt.Println(data)
```

# CRUD
## fetch (read)

> 取得到的 row data 要使用 Scan 來取得內容資料

``` fetchOne
db, err := mysql.Open()
var args = []interface{}{
  "{value_1}", "value_2",
}
query := "select * from table where item1=? and item_2=?"
r, err := mysql.FetchOne(db, query, args...)
fmt.Println(r["id"])
```

``` fetchAll
db, err := mysql.Open()
var args = []interface{}{
  "{value_1}", "value_2",
}
query := "select * from table where item1=? and item_2=?"
r, err := mysql.FetchAll(db, query, args...)
fmt.Println(r[0]["id"])
```

## Exec (insert / update / delete)
```
db, _ := mysql.Open()
var args = []interface{}{
  "sid1", "token1",
}
query := "INSERT INTO cvr_auth (`sid`, `token`) VALUES (?, ?)"
mysql.Exec(db, query, args...)
```

## Exec : insert id
```
result, err := mysql.Exec(db, query)
if err != nil {
    log.Fatal(err)
}
LastInsertId, err := mysql.LastInsertId(result)
if err != nil {
    log.Fatal(err)
}
fmt.Println(LastInsertId)
```

## Exec : Affect
```
result, err := mysql.Exec(db, query)
if err != true {
    log.Fatal(err)
}
affect, err := mysql.RowsAffected(result)
if err != nil {
    log.Fatal(err)
}
fmt.Println(affect)
```