package methods

import (
	"baoxiu/db"
	"baoxiu/models"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var data_list *gorm.DB

type Record struct{}

type Tlist struct {
	Id      int
	Tagname string
	Feature string
}
type IndexList struct {
	Id      int
	Nianji  string
	Banji   string
	Taglist []Jsontag
	Bx_time string
	Bx_tel  string
}
type Newlist struct {
	Id            int
	Nianji        string
	Banji         string
	TagList       []Tlist
	Descrip       string
	Yuyue_time    string
	Uuid          string
	Bx_teacher    string
	Bx_tel        string
	Bx_time       string
	Photo         string
	Complete_time string
	Is_complete   int
	Is_broken     int
}

func init() {
	data_list = db.DB
}

func Newrecord() Record {
	return Record{}
}

func (rr *Record) Index(c *gin.Context) {
	t := time.Now().String()

	t = t[:10]
	//c.JSON(200, gin.H{"wel": t})
	//获取 我的报修记录
	uuid, err := c.Cookie("uuid")
	var mylist IndexList
	count := 0
	if err != nil {
		mylist = IndexList{}
	} else {
		var temp models.Bxrecord
		data.Where("uuid=?", uuid).Find(&temp)
		arr := temp.Taginfo
		p1 := []Jsontag{}

		err := json.Unmarshal([]byte(arr), &p1)
		if err != nil {
			fmt.Println(err)
			c.Redirect(302, "/baoxiu")
			return
		}
		mylist.Banji = temp.Banji
		mylist.Bx_tel = temp.Bx_tel
		mylist.Bx_time = temp.Bx_time
		mylist.Id = temp.Id
		mylist.Nianji = temp.Nianji
		mylist.Taglist = p1
		count = 1
	}

	//获取所有的记录
	var list []models.Bxrecord
	rs_all := data_list.Where("is_complete=? AND is_broken=?", 0, 0).Find(&list)
	if rs_all.RowsAffected == 0 {
		fmt.Println("nill")
	}
	indexlist := []IndexList{}

	for _, val := range list {
		var one IndexList
		one.Id = val.Id
		one.Nianji = val.Nianji
		one.Banji = val.Banji
		one.Bx_tel = val.Bx_tel
		one.Bx_time = val.Bx_time
		p := []Jsontag{}
		ss := val.Taginfo
		err := json.Unmarshal([]byte(ss), &p)
		if err != nil {
			fmt.Println(err)
			return
		}
		one.Taglist = p

		indexlist = append(indexlist, one)

	}
	//获取list中对应的tag的信息，并且更新内容
	// for _, val := range list {
	// 	tl := val.Taginfo1
	// 	//将字符串转换成数组

	// 	tlist := []Tlist{}
	// 	for _, v1 := range tl {

	// 	}
	// 	var onerec models.Taginfo
	// 	data_list.Where("id=?")
	// }
	c.HTML(200, "front/index.html", gin.H{"wel": t, "list": indexlist, "mylist": mylist, "count": count})
}

// func (rr *Record) Detail(c *gin.Context) {
// 	id, b := c.Params.Get("id")
// 	t := time.Now().String()

// 	t = t[:10]
// 	if !b {
// 		//c.HTML(200, "index.html", "数据不匹配")
// 		fmt.Println("数据不匹配")
// 		c.Redirect(301, "/")
// 		return
// 	}
// 	c.HTML(200, "front/detail.html", gin.H{"wel": t})
// 	fmt.Println(id)
// }
