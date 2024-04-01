package bd

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql" // Importa el controlador MySQL

	"github.com/deibyssoca/ds-E-Commerce/models"
	"github.com/deibyssoca/ds-E-Commerce/secretm"
)

var SecretModel models.SecretRDSJson
var err error
var Db *sql.DB

func ReadSecret() error {
	SecretModel, err = secretm.GetSecret(os.Getenv("SecretName"))
	return err
}

func DbConnect() error {
	Db, err = sql.Open("mysql", ConnStr(SecretModel))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = Db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("Connection DB OK")
	return nil
}

func ConnStr(keys models.SecretRDSJson) string {
	var dbUser, authToken, dbEndpoint, dbName string
	dbUser = keys.Username
	authToken = keys.Password
	dbEndpoint = keys.Host
	dbName = "ds_aws_db_01" //TODO: pasar a variable de entorno
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswords=true", dbUser, authToken, dbEndpoint, dbName)
	fmt.Println(dsn)
	return dsn
}

func UserIsAdmin(userUUID string) (bool, string) {
	if err := DbConnect(); err != nil {
		return false, err.Error()
	}
	query := "SELECT 1 FROM users WHERE User_UUID='" + userUUID + "' AND User_Statud = 0"
	fmt.Printf("Query : %s", query)

	rows, err := Db.Query(query)
	if err != nil {
		return false, err.Error()
	}

	var value string
	rows.Next()
	rows.Scan(&value)
	fmt.Printf("UserIsAdmin => %s", value)

	if value == "1" {
		return true, ""
	}

	return false, "User is not Admin"
}

func InsertCategory(mc models.Category) (int64, error) {
	fmt.Println("InsertCategory registration begins")
	if err := DbConnect(); err != nil {
		return 0, err
	}
	defer Db.Close()
	query := "INSERT INTO category (Categ_Name, Categ_Path) VALUES ('" + mc.CategName + "','" + mc.CategPath + " ')"
	fmt.Printf("Query : %s", query)

	result, err := Db.Exec(query)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}

	LastInsertId, err2 := result.LastInsertId()
	if err2 != nil {
		fmt.Println(err2.Error())
		return 0, err2
	}
	fmt.Println("InsertCategory registration OK")
	return LastInsertId, nil
}
