package routers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/deibyssoca/ds-E-Commerce/bd"
	"github.com/deibyssoca/ds-E-Commerce/models"
)

func InsertProduct(body string, User string) (int, string) {
	var mp models.Product
	fmt.Println("Start InsertProducto")
	// convierto a Body en un slice de byte
	if err := json.Unmarshal([]byte(body), &mp); err != nil {
		return http.StatusBadRequest, "Error in data received " + err.Error()
	}

	if len(mp.ProdTitle) == 0 {
		return http.StatusBadRequest, "You must specify the Title of the product"
	}

	if isAdmin, msg := bd.UserIsAdmin(User); !isAdmin {
		return http.StatusBadRequest, msg
	}

	result, err := bd.InsertProduct(mp)
	if err != nil {
		return http.StatusBadRequest, "An error occurred while trying to register the product " + mp.ProdTitle
	}
	return http.StatusOK, "{ ProductID: " + strconv.Itoa(int(result)) + "}"
}
