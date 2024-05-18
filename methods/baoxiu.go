package methods

import (
	"baoxiu/db"
	"baoxiu/models"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Baoxiu struct{}
type Jsontag struct {
	Id      int    `json:"id"`
	Tagname string `json:"tagname"`
	Feature string `json:"feature"`
}
type Myform struct {
	Nianji      string   `form:"nianji"`
	Banji       string   `form:"banji"`
	Tags_choice []string `form:"choice[]"`
	Teacher     string   `form:"teacher"`
	Tel         string   `form:"tel"`
	Descrip     string   `form:"descrip"`
}

var data *gorm.DB

func init() {
	data = db.DB
}

func NewBx() Baoxiu {
	return Baoxiu{}
}

func (b *Baoxiu) Index(c *gin.Context) {
	//get tags
	var tags []models.Tag
	data.Find(&tags)
	//fmt.Println(tags)

	c.HTML(200, "front/baoxiu.html", gin.H{"tags": tags})
}

func (b *Baoxiu) CreateBx(c *gin.Context) {
	var ff Myform
	c.ShouldBind(&ff)

	nianji := ff.Nianji

	banji := ff.Banji
	tags_choice := ff.Tags_choice
	teacher := ff.Teacher
	tel := ff.Tel
	descrip := ff.Descrip
	choices := tags_choice
	choices_cd := len(choices)
	//检查输入的字段是否是空的
	if banji == "" || nianji == "" || choices_cd == 0 {
		fmt.Println("空的字段存在！")
		c.Redirect(302, "/baoxiu")
		return
	}

	//检查 biao  Bxrecord 中  is_complete  is_broken 都为0的记录是否已经包含相同的班级再内
	var checkbj models.Bxrecord
	rs := data.Where("is_complete=? AND is_broken=? AND banji=? AND nianji=?", 0, 0, banji, nianji).Find(&checkbj)
	if rs.RowsAffected != 0 {
		fmt.Println("err in 已经存在报修记录")
		c.Redirect(302, "/baoxiu")

		return
	}

	jtags := []Jsontag{}
	for _, val := range choices {
		a, _ := strconv.Atoi(val)
		var tt models.Tag
		data.Where("id=?", a).Find(&tt)
		var one Jsontag
		one.Id = tt.Id
		one.Tagname = tt.Tagname
		one.Feature = tt.Feature
		jtags = append(jtags, one)
	}

	bb, err := json.Marshal(jtags)
	if err != nil {
		c.Redirect(302, "/baoxiu")
		return
	}

	var rec models.Bxrecord
	rec.Banji = banji
	rec.Nianji = nianji
	rec.Descrip = descrip
	rec.Bx_teacher = teacher
	rec.Taginfo = string(bb)
	rec.Bx_tel = tel

	// 加入时间
	rec.Bx_time = time.Now().String()

	//检查是否存在cookie  uuid  如果不存在 就创建uuid
	//组合形式为 uuid+"shijian"+ 随机5位数+
	uuid, err := c.Cookie("uuid")
	if err != nil {
		shijian := time.Now().UnixNano()
		shi_str := strconv.FormatInt(shijian, 10)
		rr := rand.New(rand.NewSource(time.Now().UnixNano()))
		by := make([]byte, 6)
		rr.Read(by)
		randstr := hex.EncodeToString(by)
		str_uuid := "uuid" + shi_str + randstr
		c.SetCookie("uuid", str_uuid, 3600*24*365*10, "/", "", false, true)
		uuid = str_uuid
	}
	rec.Uuid = uuid

	data.Create(&rec)

	c.Redirect(302, "/baoxiu/success")

}

func (b *Baoxiu) Success(c *gin.Context) {
	c.HTML(200, "front/success.html", nil)
}
func (b *Baoxiu) UpdateBx() {}

func (b *Baoxiu) DelBx(c *gin.Context) {
	id, bb := c.Params.Get("id")
	if !bb {
		fmt.Println("空数据")
		c.HTML(200, "index.html", "kk")
		return
	}
	fmt.Println(id)
	c.HTML(200, "index.html", "tt")
}

func (b *Baoxiu) SearchBx() {}
