package handlers

import (
	"agile/pkg/models"
	"encoding/json"
	"net/http"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "POST" {
		models.SaveImage(w, r)
		return
	}

	p := make([]models.Product, 0)
	p = append(p, models.Product{
		Id:          1,
		Title:       "title1",
		Description: "desc1 desc1 desc1 desc1 desc1 desc1",
		Price:       12.5,
		Image:       "pain.jpg",
	})
	p = append(p, models.Product{
		Id:          2,
		Title:       "title1",
		Description: "desc1 desc1 desc1 desc1 desc1 desc1",
		Price:       12.55,
		Image:       "pain.jpg",
	})
	p = append(p, models.Product{
		Id:          3,
		Title:       "title1",
		Description: "desc1 desc1 desc1 desc1 desc1 desc1",
		Price:       12.55,
		Image:       "pain.jpg",
	})
	p = append(p, models.Product{
		Id:          4,
		Title:       "title1",
		Description: "desc1 desc1 desc1 desc1 desc1 desc1",
		Price:       12.5,
		Image:       "pain.jpg",
	})
	p = append(p, models.Product{
		Id:          5,
		Title:       "title1",
		Description: "desc1 desc1 desc1 desc1 desc1 desc1",
		Price:       12.5,
		Image:       "pain.jpg",
	})

	//data, err := p.GetProducts()
	//if err != nil {
	//	fmt.Println(err)
	//	w.Write([]byte(err.Error()))
	//}

	b, _ := json.Marshal(p)
	w.Write(b)

}
