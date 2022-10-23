package models

import (
	"agile/pkg/dbManager"
	"fmt"
)

type Product struct {
	Id          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Image       string  `json:"fileName"`
}

func (p *Product) Save() {
	_, err := dbManager.Get().Exec(`insert into public.product(title,description,price,image) values ($1,$2,$3,$4)`, p.Title, p.Description, p.Price, p.Image)
	if err != nil {
		fmt.Println("product.save err:", err)
	}
}

func (p *Product) GetOne(id int) (err error, product Product) {
	err = dbManager.Get().QueryRow(`select title,description,price,image from public.product where id=$1`, id).Scan(&product.Title, &product.Description, &product.Price, &product.Image)
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}

func (p *Product) GetAll() (products []Product, err error) {
	rows, err := dbManager.Get().Query(`select title,description,price,image from public.product`)
	if err != nil {
		fmt.Println("product.GetAll err:", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var product Product
		rows.Scan(&product.Title, &product.Description, &product.Price, &product.Image)
		products = append(products, product)
	}

	return
}

func (p *Product) Remove(id int) {
	_, err := dbManager.Get().Query(`delete from table public.product where id=$1`, id)
	if err != nil {
		fmt.Println(err)
		return
	}
}
