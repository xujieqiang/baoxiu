package methods

import (
	"baoxiu/models"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Detail struct{}

func NewDetail() Detail {
	return Detail{}
}

func (d *Detail) Index(c *gin.Context) {
	//
	id, b := c.Params.Get("id")
	if !b {
		c.Redirect(200, "/wrong") //gin.H{"a": "网页已经丢失"}
	}
	fmt.Println(id)
	var rec models.Bxrecord
	data.Where("id=?", id).Find(&rec)
	tlist := []Jsontag{}
	ss := rec.Taginfo
	err := json.Unmarshal([]byte(ss), &tlist)
	if err != nil {
		c.Redirect(302, "/")
		return
	}

	c.HTML(200, "front/detail.html", gin.H{"rec": rec, "taglist": tlist})
	//
}
