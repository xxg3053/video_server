package dbops

import "database/sql"

var (
	dbConn *sql.DB
	err error
)

func init()  {
	dbConn, err = sql.Open("mysql", "root:root@tcp(172.16.92.218:3306)/video_server?charset=utf8")
	if err != nil{
		panic(err.Error())
	}

}