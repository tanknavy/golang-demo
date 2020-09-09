package user //此包要被导入，不能使用main
//https://gorm.io/docs/index.html
import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	//"github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

const dsn string = "hive:q1w2e3r4@tcp(spark3:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"

type User struct {
	//type Tmp struct { //gorm默认使用ID作为主键，struct蛇形命名(TableName)对应DB中table_name(表名、栏位名)
	gorm.Model
	ID        int64 //db中自增
	Username  string
	Email     string
	ImageFile string
	Password  string
}

type Tabler interface { //struc名称User默认映射表名users, 自定义
	TableName() string
}

// TableName overrides the table name used by User to `profiles`
func (User) TableName() string {
	return "user" //映射表名user
}

func InitialMigration() {
	//db, err = gorm.Open("mysql","hive:q1w2e3r4@tcp(spark3:3306)/test")
	//dsn := "hive:q1w2e3r4@tcp(spark3:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect to database")
	}

	//defer db.Close()
	db.AutoMigrate(&User{}) //测试DB中只有
	//db.AutoMigrate(&Tmp{})

}

func AllUsers(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w,"all users endpoint hit")
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	//defer db.Close()

	var users []User
	//var users []Tmp
	db.Find(&users)
	//fmt.Println(users)
	json.NewEncoder(w).Encode(users)
}

func NewUser(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w,"new users endpoint hit")
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	//defer db.Close()

	vars := mux.Vars(r) //接受网页request取得参数
	name := vars["name"]
	email := vars["email"]

	db.Create(&User{Username: name, Email: email})  //插入数据
	fmt.Fprintf(w, "new user created successfully") //写出web

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "delete users endpoint hit")
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	//defer db.Close()

	vars := mux.Vars(r)//接受网页request取得参数
	name := vars["name"]

	var user User
	db.Where("username = ?", name).Find(&user)
	db.Delete(&user) //crud中变量都是使用指针，

	fmt.Fprintf(w,"User Successful deleted")

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "update users endpoint hit")
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	//defer db.Close()

	vars := mux.Vars(r) //从http request中拿到全部parameter, map类型
	name := vars["name"] //读出字段
	email := vars["email"]

	var user User 
	db.Where("name = ?", name).Find(&user)//查找并赋值给user变量
	user.Email = email //设置新值

	db.Save(&user) //保存到db
	fmt.Fprintf(w,"successfully updateded")//想网页返回信息
}
