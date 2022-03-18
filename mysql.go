package main

import (
	"botgo/dto"
	"botgo/log"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //导入包但不使用，init()
)

var db *sql.DB

//Go连接Mysql示例
func initDB() (err error) {
	//用户名:密码啊@tcp(ip:端口)/数据库的名字
	dsn := "root:pigkang@tcp(127.0.0.1:3306)/botgo"

	//连接数据集
	db, err = sql.Open("mysql", dsn) //open不会检验用户名和密码
	if err != nil {
		fmt.Printf("dsn:%s invalid,err:%v\n", dsn, err)
		return
	}

	err = db.Ping() //尝试连接数据库
	if err != nil {
		fmt.Printf("open %s faild,err:%v\n", dsn, err)
		return
	}
	fmt.Println("连接数据库成功~")

	//设置数据库连接池的最大连接数
	db.SetMaxIdleConns(10)

	return
}

// 查询随机一个梗数据
func query() dto.Meme {
	rowObj := db.QueryRow("SELECT id as Id,name as Name,description as Description FROM meme ORDER BY RAND() LIMIT 1") //从连接池里取一个连接出来去数据库查询单挑记录
	var meme dto.Meme
	rowObj.Scan(&meme.Id, &meme.Name, &meme.Description)
	return meme
}

// 根据梗tag查询一条数据
func queryTagMeme(tag int) dto.Meme {
	querySQL := "SELECT id as Id,name as Name,description as Description FROM meme WHERE 1=1 AND tag = ? ORDER BY RAND() LIMIT 1"
	rowObj := db.QueryRow(querySQL, tag) //从连接池里取一个连接出来去数据库查询单挑记录
	var meme dto.Meme
	rowObj.Scan(&meme.Id, &meme.Name, &meme.Description)
	return meme
}

// 搜索梗名查出数据
func queryMemeByName(memeName string) []dto.Meme {
	var meme []dto.Meme

	querySQL := "SELECT id as Id,code as Code,name as Name,description as Description FROM meme WHERE 1=1 AND name like ? limit 3"
	rows, err := db.Query(querySQL, "%"+memeName+"%")
	if err != nil {
		fmt.Printf("%s query failed,err:%v\n", querySQL, err)
	}
	defer rows.Close()
	var memeTemp dto.Meme
	for rows.Next() {
		rows.Scan(&memeTemp.Id, &memeTemp.Code, &memeTemp.Name, &memeTemp.Description)
		meme = append(meme, memeTemp)
	}
	log.Info(meme)
	return meme
}
