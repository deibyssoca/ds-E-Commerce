package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/deibyssoca/ds-E-Commerce/auth"
	"github.com/deibyssoca/ds-E-Commerce/routers"
)

func Handlers(path string, method string, body string, headers map[string]string, request events.APIGatewayV2HTTPRequest) (int, string) {
	fmt.Println("Voy a procesar " + path + " > " + method)
	id := request.PathParameters["id"]
	idn, _ := strconv.Atoi(id)

	isOk, statusCode, user := validateAuthorization(path, method, headers)
	if !isOk {
		return statusCode, user
	}
	switch path[0:4] {
	case "user":
		return userProcess(body, path, method, user, id, request)
	case "prod":
		return productProcess(body, path, method, user, idn, request)
	case "stoc":
		return stockProcess(body, path, method, user, idn, request)
	case "addr":
		return addressProcess(body, path, method, user, idn, request)
	case "cate":
		return categoryProcess(body, method, user, idn, request)
	case "orde":
		return orderProcess(body, path, method, user, idn, request)
	}

	return http.StatusBadRequest, "Method invalid"
}

func userProcess(body string, path string, method string, user string, id string, request events.APIGatewayV2HTTPRequest) (int, string) {
	return http.StatusBadRequest, "Method invalid"
}
func productProcess(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	return http.StatusBadRequest, "Method invalid"
}
func categoryProcess(body string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	fmt.Println("In categoryProcess")
	switch method {
	case "POST":
		fmt.Println("Case POST")
		return routers.InsertCategory(body, user)
	case "PUT":
		fmt.Println("Case PUT")
		return routers.UpdateCategory(body, user, id)
	}
	return http.StatusBadRequest, "Method invalid"
}
func stockProcess(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	return http.StatusBadRequest, "Method invalid"
}
func addressProcess(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	return http.StatusBadRequest, "Method invalid"
}
func orderProcess(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	return http.StatusBadRequest, "Method invalid"
}

func validateAuthorization(path string, method string, headers map[string]string) (bool, int, string) {
	if (path == "product" && method == "GET") || (path == "category" && method == "GET") {
		return true, http.StatusOK, ""
	}
	token := headers["authorization"]
	if len(token) == 0 {
		return false, http.StatusUnauthorized, "Token is required"
	}

	tokenOK, msg, err := auth.ValidToken(token)
	if !tokenOK {
		if err != nil {
			fmt.Println("Error in token" + err.Error())
			return false, http.StatusUnauthorized, err.Error()
		} else {
			fmt.Println("Error in token" + msg)
			return false, http.StatusUnauthorized, msg
		}
	}
	fmt.Println("Token OK")
	return true, http.StatusOK, msg
}
