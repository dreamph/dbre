package main

import (
	"fmt"
	"log"

	"github.com/dreamph/dbre/adapters/bun/utils"
)

type Test struct {
	ID         int    `bun:"id,pk" gorm:"primary_key;column:id;" json:"id"`
	Name       string `db:"name_1" gorm:"column:name;" json:"name_j"`
	NonGormTag string `db:"name_2" json:"-"`
}

func main() {
	t := Test{}
	pkFields, dataFields, err := utils.GetDbFields(t)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(pkFields)
	fmt.Println(dataFields)
}
