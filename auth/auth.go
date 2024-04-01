package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type TokenJSON struct {
	Sub       string
	Event_id  string
	Token_use string
	Scope     string
	Auth_time int
	Iss       string
	Exp       int
	Iat       int
	Client_id string
	Username  string
}

func ValidToken(token string) (bool, string, error) {
	var parts []string

	if parts = strings.Split(token, "."); len(parts) != 3 {
		fmt.Println("Token is not valid")
		return false, "Token is not valid", nil
	}
	userInfo, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		fmt.Println("It is not possible to decode the token", err.Error())
		return false, err.Error(), err
	}
	var tkj TokenJSON
	if err = json.Unmarshal(userInfo, &tkj); err != nil {
		fmt.Println("It is not possible to decode the struct JSON", err.Error())
		return false, err.Error(), err
	}
	now := time.Now()
	tm := time.Unix(int64(tkj.Exp), 0)
	if tm.Before(now) {
		fmt.Println("Date expired token = " + tm.String())
		fmt.Println("Token expired!")
		return false, "Token expired!", err
	}
	return true, string(tkj.Username), nil
}
