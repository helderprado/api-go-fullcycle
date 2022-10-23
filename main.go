package main

import (
	http "go-api/http"
	"go-api/model"

	uuid "github.com/satori/go.uuid"
)

func main() {
	produto1 := model.Product{
		ID:   uuid.NewV4().String(),
		Name: "Carrinho",
	}

	produto2 := model.Product{
		ID:   uuid.NewV4().String(),
		Name: "Boneco",
	}

	produto3 := model.Product{
		ID:   uuid.NewV4().String(),
		Name: "Boneca",
	}

	products := model.Products{}
	products.Add(produto1)
	products.Add(produto2)
	products.Add(produto3)

	server := http.NewWebServer()
	server.Products = &products
	server.Serve()

}

// func main() {
// 	fmt.Println("...listening port 9095")
// 	http.HandleFunc("/product", ProductHandler)
// 	http.ListenAndServe(":9095", nil)
// }

// func ProductHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("<h1>Olá mundo</h1>"))
// }
