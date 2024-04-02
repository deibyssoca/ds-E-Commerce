package routers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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
