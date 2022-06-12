package test

import (
	"cloud-disk/core/models"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
	"xorm.io/xorm"
)

func TestXormTest(t *testing.T) {

	engine, err := xorm.NewEngine("mysql", "root:123456@tcp(localhost:3306)/cloud-disk?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		t.Fatal(err)
	}
	//data := make([]*models.UserBasic, 0)
	//err = engine.Find(&data)
	//if err != nil {
	//	t.Fatal(err)
	//}

	//user := new(models.UserBasic)
	//info, err := engine.TableInfo(user)
	//fmt.Println("================")
	//fmt.Println(info)
	//if err != nil {
	//	t.Fatal(err)
	//}

	//user.Name = "myname1"
	//user.Identity="22222221"
	//user.Password="2222221"
	//user.Email="222222223@111.com"
	//insert, err := engine.Insert(user)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//fmt.Println("==insert:",insert)
	user := make([]*models.UserBasic, 0)
	err = engine.Find(&user)
	if err != nil {
		t.Fatal(err)
	}
	b, err := json.Marshal(user)
	if err != nil {
		t.Fatal(err)
	}
	//dst := new(bytes.Buffer)
	//err = json.Indent(dst, b, "", "  ")
	//if err != nil {
	//	t.Fatal(err)
	//}
	fmt.Println(string(b))
}
