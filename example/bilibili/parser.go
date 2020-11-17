package main

import (
	"fmt"
	"github.com/hearecho/MagicSpider"
	"github.com/hearecho/MagicSpider/utils"
	"os"
)

type Item struct {
	Name string
	Level int
}

func (i *Item) Process()  {
	_ = utils.IsNotExistMkDir(MagicSpider.S.RuntimePath)
	f, _ := utils.Open(MagicSpider.S.RuntimePath+"result.csv", os.O_CREATE|os.O_APPEND, 0777)
	item := fmt.Sprintf("%v,%v\n", i.Name, i.Level)
	f.WriteString(item)
	f.Close()
}

func ResParse(r *MagicSpider.Response) MagicSpider.ParseResult {
	res := &MagicSpider.ParseResult{}
	data := r.Doc.(map[string]interface{})["data"].(map[string]interface{})
	r.Meta.(*Item).Name = data["name"].(string)
	r.Meta.(*Item).Level = int(data["level"].(float64))
	res.Items = append(res.Items,r.Meta.(*Item))
	return *res
}


