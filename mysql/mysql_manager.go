/*
Package mysql ... refer readme.md
*/
package mysql

import (
  "app/config"
  "database/sql"
  // "encoding/json"
  // "fmt"
  m "github.com/go-sql-driver/mysql"
  "log"
  // "reflect"
  "github.com/happierall/l"
)

/*
DataSourceName ... save connection info
*/
var DataSourceName string

/*
CVRDB ...
*/
var CVRDB *sql.DB

/*
AUTHDB ...
*/
var AUTHDB *sql.DB

/*
init ...
*/
func init() {
  var err error
  CVRDB, err = Open("")
  if err != nil {
    log.Fatal(err, ". CVR database connect fail")
  }

  AUTHDB, err = Open("auth")
  if err != nil {
    log.Fatal(err, ". Auth database connect fail")
  }
}

/*
Open ...
*/
func Open(dbset string) (db *sql.DB, err error) {
  host, port, username, password, dbname, charset := config.Database(dbset)
  // [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
  DataSourceName = username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=" + charset

  // db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:6612)/cvr?parseTime=true")
  l.Log("Connecting to mysql db")
  db, err = sql.Open("mysql", DataSourceName)
  // defer db.Close()
  if err != nil {
    log.Fatalln(err)
  }
  db.SetMaxIdleConns(20) // 數據庫空間連接閒置數量
  db.SetMaxOpenConns(20) // 最大連接數
  db.SetMaxIdleConns(0)

  err = db.Ping()
  // if err = db.Ping(); err != nil {
  //     // log.Fatalln(err)
  //     return
  // }
  return
}

/*
ShowDBConfig ...
    use this func to show current database config
*/
func ShowDBConfig() {
  cfg, _ := m.ParseDSN(DataSourceName)
  // log.Println(cfg)
  log.Println(
    "\n --- Addr    : ", cfg.Addr,
    "\n --- User    : ", cfg.User,
    "\n --- Pw      : ", cfg.Passwd,
    "\n --- DBName  : ", cfg.DBName)
  return
}

/*
Close ...
*/
func Close(db *sql.DB) {
  defer db.Close()
}

// row.Scan(&struct.item1, $struct.item2) 來取得內容
/* example
err := row.Scan(&id, &email, &name)
if err != nil {
    if err == sql.ErrNoRows {
                   // no such user id
    } else {
        panic(err)
    }
} else {
    user = User{Id: id, Name: name, Email: email}
}
*/

/*
FetchOne ...
    fmt.Println(tableData)
    fmt.Println(tableData["{item}"])
*/
func FetchOne(db *sql.DB, query string, args ...interface{}) (tableData map[string]interface{}, err error) {
  tableDatas, err := FetchAll(db, query, args...)
  if err != nil {
    return
  }
  if len(tableDatas) == 0 {
    return
  }
  tableData = tableDatas[0]
  return
}

/*
FetchAll ...
    fmt.Println(tableDatas)
    fmt.Println(tableDatas[0])
    fmt.Println(tableDatas[0]["{item}"])
    fmt.Println(tableDatas[1]["{item}"])
*/
func FetchAll(db *sql.DB, query string, args ...interface{}) (tableDatas []map[string]interface{}, err error) {
  rows, err := db.Query(query, args...)
  if err != nil {
    log.Println(err)
    return
  }
  defer rows.Close()
  columns, err := rows.Columns()
  if err != nil {
    return
  }
  count := len(columns)
  tableDatas = make([]map[string]interface{}, 0)
  values := make([]interface{}, count)
  valuePtrs := make([]interface{}, count)
  for rows.Next() {
    for i := 0; i < count; i++ {
      valuePtrs[i] = &values[i]
    }
    rows.Scan(valuePtrs...)
    entry := make(map[string]interface{})
    for i, col := range columns {
      var v interface{}
      val := values[i]
      b, ok := val.([]byte)
      if ok {
        v = string(b)
      } else {
        v = val
      }
      entry[col] = v
    }
    tableDatas = append(tableDatas, entry)
  }
  return
}

/*
Exec ...
    query ex : query := "insert into cvr_token (token) values ('maxhu')"
    insert / update / delete
*/
func Exec(db *sql.DB, query string, args ...interface{}) (sql.Result, error) {
  result, err := db.Exec(query, args...)
  return result, err
}

/*
LastInsertId ... exec:insert 後使用
*/
func LastInsertId(result sql.Result) (int64, error) {
  id, err := result.LastInsertId()
  return id, err
}

/*
RowsAffected exec:update/delete 後使用
*/
func RowsAffected(result sql.Result) (int64, error) {
  affect, err := result.RowsAffected()
  return affect, err
}
