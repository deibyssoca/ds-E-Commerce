package bd

import (
	"database/sql"
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

func DeleteCategory(id int) error {
	fmt.Println("DeleteCategory start")
	if err := DbConnect(); err != nil {
		return err
	}
	defer Db.Close()
	query := "DELETE from category where Categ_Id = " + strconv.Itoa(id)
	fmt.Printf("Query : %s", query)

	_, err := Db.Exec(query)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Delete Category OK")
	return nil
}

func SelectCategories(pCategId int, pName string, pPath string) ([]models.Category, error) {
	fmt.Println("SelectCategories start")
	var categ []models.Category

	if err := DbConnect(); err != nil {
		return categ, err
	}
	defer Db.Close()

	query := "SELECT Categ_Id, Categ_Name, Categ_Path from category "
	if pCategId > 0 {
		query += "WHERE Categ_Id = " + strconv.Itoa(pCategId)
	} else {
		if len(pPath) > 0 {
			query += "WHERE Categ_Path LIKE '%" + pPath + "%' "
			if len(pName) > 0 {
				query += "AND Categ_Name LIKE '%" + pName + "%' "
			}
		} else {
			if len(pName) > 0 {
				query += "WHERE Categ_Name LIKE '%" + pName + "%' "
			}
		}
	}

	fmt.Println(query)
	var rows *sql.Rows
	rows, err = Db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var c models.Category
		var categId sql.NullInt32
		var categName sql.NullString
		var categPath sql.NullString

		if err := rows.Scan(&categId, &categName, &categPath); err != nil {
			return categ, err
		}
		// Se hace así para poder trabajar con los campos nulos en la DB
		c.CategId = int(categId.Int32)
		c.CategName = categName.String // si es nulo devuelve ""
		c.CategPath = categPath.String

		categ = append(categ, c)
	}
	fmt.Println("SelectCategory > OK")
	return categ, nil
}
