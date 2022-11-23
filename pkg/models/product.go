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
	Count       int     `json:"count"`
	Image       string  `json:"fileName"`
	Discount    float64 `json:"discount"`
	CategoryId  int64   `json:"categoryId"`
	UserId      int64   `json:"userId"`
}

type Buy struct {
	Amount   int       `json:"amount"`
	Inbound  bool      `json:"inbound"`
	ItemId   int       `json:"itemId"`
	Location []float64 `json:"location"`
	Phone    string    `json:"phone"`
	Text     string    `json:"text"`
}

func (p *Product) Save() {
	_, err := dbManager.Get().Exec(`insert into public.product(title,description,price,image,fk_category,fk_user) values ($1,$2,$3,$4,$5,$6)`, p.Title, p.Description, p.Price, p.Image, p.CategoryId, p.UserId)
	if err != nil {
		fmt.Println("product.save err:", err)
	}
}

func (p *Product) GetAll() (products []Product, err error) {
	rows, err := dbManager.Get().Query(`select id,title,description,price,image,fk_category,fk_user from public.product`)
	if err != nil {
		fmt.Println("product.GetAll err:", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var product Product
		rows.Scan(&product.Id, &product.Title, &product.Description, &product.Price, &product.Image, &product.CategoryId, &product.UserId)
		products = append(products, product)
	}

	return
}

func (b *Buy) Buy(userId int) error {
	location := fmt.Sprintf("%f:%f", b.Location[0], b.Location[1])
	_, err := dbManager.Get().Exec(`insert into public.buy(amount,inbound,fk_product,fk_user,c_location,telnumber,c_text) values ($1,$2,$3,$4,$5,$6,$7)`, b.Amount, b.Inbound, b.ItemId, userId, location, b.Text, b.Phone)
	if err != nil {
		fmt.Println("product.save err:", err)
	}
	return err
}
