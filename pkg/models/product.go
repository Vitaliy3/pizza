package models

import (
	"agile/pkg/dbManager"
	"fmt"
	"strconv"
	"strings"
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
	Amount     int       `json:"amount"`
	Inbound    bool      `json:"inbound"`
	ItemId     int       `json:"itemId"`
	FK_Product int       `json:"fk_product"`
	Location   []float64 `json:"location"`
	Phone      string    `json:"phone"`
	Text       string    `json:"text"`
	Called     bool      `json:"called"`
}

func (p *Product) Save() {
	_, err := dbManager.Get().Exec(`insert into public.product(title,description,price,image,fk_category) values ($1,$2,$3,$4,$5)`, p.Title, p.Description, p.Price, p.Image, p.CategoryId)
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
	fmt.Println("location:", location)

	_, err := dbManager.Get().Exec(`insert into public.buy(amount,inbound,fk_product,c_location,telnumber,c_text) values ($1,$2,$3,$4,$5,$6)`, b.Amount, true, b.ItemId, location, b.Phone, b.Text)
	if err != nil {
		fmt.Println("product.save err:", err)
	}
	return err
}

func (b *Buy) BuyGetAll() ([]Buy, error) {

	rows, err := dbManager.Get().Query(`select id,amount,inbound,fk_product,c_location,telnumber,c_text,called from public.buy`)
	if err != nil {
		fmt.Println("buy getall err:", err)
	}

	buys := make([]Buy, 0)
	for rows.Next() {
		location := ""
		buy := Buy{}
		rows.Scan(&buy.ItemId, &buy.Amount, &buy.Inbound, &buy.FK_Product, &location, &buy.Phone, &buy.Text, &buy.Called)
		l := strings.Split(location, ":")
		l1, err := strconv.ParseFloat(l[0], 64)
		if err != nil {
			fmt.Println("l1 err:", err)
		}
		l2, err := strconv.ParseFloat(l[1], 64)
		if err != nil {
			fmt.Println("l1 err:", err)
		}
		buy.Location = append(buy.Location, l1, l2)
		buys = append(buys, buy)

	}

	return buys, err
}

func (b *Buy) StopTracking(item Buy) error {
	_, err := dbManager.Get().Exec(`update public.buy set inbound=$3, called=$2 where id=$1`, item.ItemId, item.Called, item.Inbound)
	if err != nil {
		fmt.Println("buy.StopTracking err:", err)
	}
	return err
}
