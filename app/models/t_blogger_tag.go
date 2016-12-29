package models

import (
	"blog/app/support"
	"fmt"
)

//BloggerTag model
type BloggerTag struct {
	Id     int    `xorm:"not null pk autoincr INT(11)"`
	Type   int    `xorm:"not null INT(11)"`
	Name   string `xorm:"not null VARCHAR(20)"`
	Parent int    `xorm:"INT(11)`
}

func (t *BloggerTag) TableName() string {
	return "t_tag"
}

// Query all tag
func (b *BloggerTag) FindList() ([]BloggerTag, error) {
	bt := make([]BloggerTag, 0)
	err := support.Xorm.Find(&bt)
	return bt, err
}

func (b *BloggerTag) GetByIdent(ident string) int64 {
	tag := &BloggerTag{}
	has, _ := support.Xorm.Where("ident = ?", ident).Get(tag)
	if has {
		return int64(tag.Id)
	}
	return 0
}

// Add new tag
func (b *BloggerTag) New() (bool, error) {

	bt := new(BloggerTag)
	bt.Type = b.Type
	bt.Name = b.Name
	bt.Type = b.Parent
	has, err := support.Xorm.InsertOne(bt)

	return has > 0, err
}

// QueryTags to Search for tag
// 根据用户输入的单词匹配 tag
func (t *BloggerTag) QueryTags(str string) ([]map[string][]byte, error) {
	sql := "SELECT name,id FROM t_tag WHERE name LIKE \"%" + str + "%\" ORDER BY LENGTH(name)-LENGTH(\"" + str + "\") ASC LIMIT 10"
	//sql := "SELECT name FROM t_tag"
	ress, err := support.Xorm.Query(sql)
	fmt.Println("res: ", ress)
	if err != nil {
		fmt.Println("err: ", err)
		return ress, err
	}
	return ress, nil
}
