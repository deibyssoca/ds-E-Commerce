package routers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/deibyssoca/ds-E-Commerce/bd"
	"github.com/deibyssoca/ds-E-Commerce/models"
)

func InsertCategory(body string, User string) (int, string) {
	var mc models.Category
	fmt.Println("Inicio InsertCategory")
	// convierto a Body en un slice de byte
	if err := json.Unmarshal([]byte(body), &mc); err != nil {
		return http.StatusBadRequest, "Error in data received " + err.Error()
	}

	if len(mc.CategName) == 0 {
		return http.StatusBadRequest, "You must specify the Name of the category"
	}

	if len(mc.CategPath) == 0 {
		return http.StatusBadRequest, "You must specify the Path of the category"
	}

	if isAdmin, msg := bd.UserIsAdmin(User); !isAdmin {
		return http.StatusBadRequest, msg
	}

	result, err := bd.InsertCategory(mc)
	if err != nil {
		return http.StatusBadRequest, "An error occurred while trying to register the category " + mc.CategName + " - " + mc.CategPath
	}
	return http.StatusOK, "{ CategID: " + strconv.Itoa(int(result)) + "}"
}

func UpdateCategory(body string, User string, id int) (int, string) {
	var mc models.Category
	fmt.Println("Inicio UpdateCategory")
	// convierto a Body en un slice de byte
	if err := json.Unmarshal([]byte(body), &mc); err != nil {
		return http.StatusBadRequest, "Error in data received " + err.Error()
	}

	if len(mc.CategName) == 0 && len(mc.CategPath) == 0 {
		return http.StatusBadRequest, "You must specify the Name and Path of the category"
	}

	if isAdmin, msg := bd.UserIsAdmin(User); !isAdmin {
		return http.StatusBadRequest, msg
	}
	mc.CategId = id
	if err := bd.UpdateCategory(mc); err != nil {
		return http.StatusBadRequest, "An error occurred while trying to UPDATE the category " + strconv.Itoa(id) + " - " + err.Error()
	}
	return http.StatusOK, "Update OK"
}

func DeleteCategory(body string, User string, id int) (int, string) {
	if id == 0 {
		return http.StatusBadRequest, "You must specify the Id of the category to delete"
	}

	if isAdmin, msg := bd.UserIsAdmin(User); !isAdmin {
		return http.StatusBadRequest, msg
	}

	if err := bd.DeleteCategory(id); err != nil {
		return http.StatusBadRequest, "An error occurred while trying to DELETE the category " + strconv.Itoa(id) + " - " + err.Error()
	}
	return http.StatusOK, "Delete OK"
}

func SelectCategories(body string, request events.APIGatewayV2HTTPRequest) (int, string) {
	var err error
	var categId int
	var path string
	var name string

	if len(request.QueryStringParameters["categId"]) > 0 {
		categId, err = strconv.Atoi(request.QueryStringParameters["categId"])
		if err != nil {
			return http.StatusInternalServerError, "An error occurred when converting the value " + request.QueryStringParameters["categId"] + " to integer"
		}
	} else {
		if len(request.QueryStringParameters["path"]) > 0 {
			path = request.QueryStringParameters["path"]
		}
		if len(request.QueryStringParameters["name"]) > 0 {
			name = request.QueryStringParameters["name"]
		}
	}

	collection, err2 := bd.SelectCategories(categId, name, path)
	if err2 != nil {
		return http.StatusBadRequest, "An error occurred in getting the category(ies) " + err2.Error()
	}

	Categ, err3 := json.Marshal(collection)
	if err3 != nil {
		return http.StatusBadRequest, "An error occurred in convert to JSON " + err3.Error()
	}
	return http.StatusOK, string(Categ)

}
