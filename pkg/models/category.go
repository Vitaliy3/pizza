package models

import (
	"agile/pkg/dbManager"
	"fmt"
)

type Category struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (p *Category) Save() error {
	_, err := dbManager.Get().Exec(`insert into public.category(c_name) values ($1)`, p.Name)
	if err != nil {
		fmt.Println("product.save err:", err)
		return err
	}
	return err
}

func (p *Category) GetAll() (categories []Category, err error) {
	rows, err := dbManager.Get().Query(`select id,c_name from public.category`)
	if err != nil {
		fmt.Println("Category.GetAll err:", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var cate Category
		rows.Scan(&cate.Id, &cate.Name)
		categories = append(categories, cate)
	}

	return
}
