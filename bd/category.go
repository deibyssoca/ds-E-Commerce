package bd

import (
	"fmt"
	"strconv"

	"github.com/deibyssoca/ds-E-Commerce/models"
	"github.com/deibyssoca/ds-E-Commerce/tools"
)

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

func UpdateCategory(mc models.Category) error {
	fmt.Println("UpdateCategory begins")
	if err := DbConnect(); err != nil {
		return err
	}
	defer Db.Close()
	query := "UPDATE category SET"

	if len(mc.CategName) > 0 {
		query += " Categ_Name = '" + tools.EscapeString(mc.CategName) + "'"
	}
	if len(mc.CategPath) > 0 {
		if len(mc.CategName) > 0 {
			query += ", "
		}
		query += " Categ_Path = '" + tools.EscapeString(mc.CategPath) + "'"
	}

	query += " WHERE Categ_Id = " + strconv.Itoa(mc.CategId)

	_, err := Db.Exec(query)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Printf("Query : %s", query)

	return nil
}
