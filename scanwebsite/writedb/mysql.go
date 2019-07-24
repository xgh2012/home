package writedb

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strconv"
	"time"
)

const (
	DB_Driver_DEV = "app:rwy123@tcp(10.0.3.3:3306)/CloudDB_Wifi?charset=utf8"
	DB_Driver_PRD = "app:rwy#1218@tcp(10.0.0.249:3306)/CloudDB_Wifi?charset=utf8"
	//DB_Driver_DEV = "xgh:Ads_55821284@tcp(127.0.0.1:3306)/gop?charset=utf8"
)
var (
	Db *sql.DB
	opend bool
)

type ActInfo struct {
	Title string
	Contents string
	Href string
	Laiyuan string
	Author string
	Created string
	From string
	Images string
}

type ActInfoArticle struct {
	title string
	domain string
	article_type int//1图文 2视频
	logo_type int//1单图小 2单图大 3多图
	pv_false int
	pv int
	content string
	create_time string
	status int
	author_type int //0无  1自动 2其他
	author_name string //作者名称
	from string //来源
	from_url string //来源地址
	is_scanwrite int //来源地址
	logos string //缩略图
}

func init()  {
	Db,opend= OpenDB()
}

func InsertToDB (info ActInfo) {
	//return
	nowTimeStr := ""
	stmt, err := Db.Prepare("insert fapp_find_article_pre set href=?,title=?,contents=?,created=?,laiyuan=?,author=?,images=?")
	CheckErr(err)
	if info.Created == ""{
		nowTimeStr = GetTime()
	}else{
		nowTimeStr =info.Created
	}
	_,err1 := stmt.Exec(info.Href, info.Title, info.Contents, nowTimeStr,info.Laiyuan,info.Author,info.Images)
	CheckErr(err1)
	if err1 != nil {
		log.Fatal(err1)
		return
	}
	AddToFindArticle(info)
}

//写入到发现文章里面
//TODO 换一种写法 写数据库
func AddToFindArticle(info ActInfo)  {
	data :=ActInfoArticle{
		title:info.Title,
		domain:"",
		article_type:1,
		logo_type:3,
		author_type:2,
		author_name:info.Author,
		pv_false:1000,
		pv:0,
		content:info.Contents,
		create_time:info.Created,
		status:0,
		from:info.From,
		from_url:info.Laiyuan,
		is_scanwrite:1,
		logos:info.Images,
	}
	Db.Exec("INSERT INTO fapp_find_article(title,article_type,logo_type,pv_false,pv,content,create_time,status,`from`,from_url,is_scanwrite,author_type,author_name,logos) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?)",
		data.title,data.article_type,data.logo_type,data.pv_false,data.pv,data.content,data.create_time,data.status,data.from,data.from_url,data.is_scanwrite,data.author_type,data.author_name,data.logos)
}

func QueryIsExist(str string) bool  {
	rows, err := Db.Query("SELECT COUNT(href) as num FROM fapp_find_article_pre where href=?",str)
	CheckErr(err)
	if err != nil {
		log.Fatalln("error:", err)
		return false
	}
	var numbers = 0
	for rows.Next(){
		rows.Scan(&numbers)
	}
	if numbers==0{
		return false
	}
	return true
}

func QueryFromDB(str string) {
	rows, err := Db.Query("SELECT COUNT(href) as num FROM tab_contents where href=?",str)
	CheckErr(err)
	if err != nil {
		fmt.Println("error:", err)
	}
	var numbers = 0
	for rows.Next(){
		rows.Scan(&numbers)
	}
}

func OpenDB() (db *sql.DB,success bool) {
	var isOpen bool

	env_ss := DB_Driver_PRD

	db, err := sql.Open("mysql", env_ss)
	if err != nil {
		isOpen = false
	} else {
		isOpen = true
	}
	CheckErr(err)
	return db,isOpen
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
		fmt.Println("err:", err)
	}
}

func GetTime() string {
	const shortForm = "2006-01-02 15:04:05"
	t := time.Now()
	temp := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.Local)
	str := temp.Format(shortForm)
	return str
}

func GetMD5Hash(text string) string {
	haser := md5.New()
	haser.Write([]byte(text))
	return hex.EncodeToString(haser.Sum(nil))
}

func GetNowtimeMD5() string {
	t := time.Now()
	timestamp := strconv.FormatInt(t.UTC().UnixNano(), 10)
	return GetMD5Hash(timestamp)
}
