package models

type Bxrecord struct {
	Id            int    `json:"id"`
	Nianji        string `json:"nianji"`
	Banji         string `json:"banji"`
	Taginfo       string
	Descrip       string `json:"descrip"`
	Yuyue_time    string `json:"yuyue"`
	Uuid          string
	Bx_teacher    string `json:"teacher"`
	Bx_tel        string `json:"tel"`
	Bx_time       string `json:"bx_time"`
	Photo         string `json:"photo"`
	Complete_time string
	Is_complete   int
	Is_broken     int
}

func (b *Bxrecord) TableName() string {
	return "bxrecord"
}

type Tag struct {
	Id          int
	Tagname     string
	Who_create  int
	Feature     string
	Create_time string
	Update_time string
}

func (t *Tag) TableName() string {
	return "tag"
}
