package db

import (
	"MagicSpider/spider/log"
	"MagicSpider/spider/setting"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

/**
数据库操作,从数据库中读取配置，以及从存储结果
 */
var DB *sql.DB

func InitDB()  {
	DB,err := sql.Open("mysql",setting.DBurl)
	if err != nil {
		log.Error.Printf("init database error! === %v\n",err.Error())
		os.Exit(1)
	}
	DB.SetConnMaxLifetime(100)
	DB.SetMaxIdleConns(10)
	if err := DB.Ping(); err != nil {
		log.Error.Printf("init database error! === %v\n",err.Error())
		os.Exit(1)
	}
	log.Info.Println("inti database success")
}

func QuerySetting(name string) string {
	stmt,err := DB.Prepare("select val from setting where setName = ?")
	if err != nil {
		log.Error.Printf("query database error! === %v\n",err.Error())
		os.Exit(2)
	}
	query,err := stmt.Query(name)
	if err != nil {
		log.Error.Printf("query database error! === %v\n",err.Error())
		os.Exit(2)
	}
	var res string
	err = query.Scan(res)
	if err != nil {
		log.Error.Printf("query database error! === %v\n",err.Error())
		os.Exit(2)
	}
	return res
}



