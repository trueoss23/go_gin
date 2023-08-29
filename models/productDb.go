package models

import (
	"database/sql"

	"go_gin/config"

	_ "github.com/go-sql-driver/mysql"

	"fmt"
)

func GetProducts() []Product {
	fmt.Println(config.Cfg.DbUser, config.Cfg.DbPass, config.Cfg.DbName)

	db, err := sql.Open("mysql", config.Cfg.DbUser+":"+config.Cfg.DbPass+"@tcp(127.0.0.1:3306)/"+config.Cfg.DbName)

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	defer db.Close()
	results, err := db.Query("SELECT * FROM product")

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil

	}

	products := []Product{}

	for results.Next() {
		var prod Product
		err = results.Scan(&prod.Code, &prod.Name, &prod.Qty, &prod.LastUpdated)

		if err != nil {
			panic(err.Error())
		}

		products = append(products, prod)
	}

	return products

}

func GetProduct(code string) *Product {

	db, err := sql.Open("mysql", config.Cfg.DbUser+":"+config.Cfg.DbPass+"@tcp(127.0.0.1:3306)/"+config.Cfg.DbName)

	prod := &Product{}

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	defer db.Close()
	results, err := db.Query("SELECT * FROM product where code=?", code)

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	if results.Next() {

		err = results.Scan(&prod.Code, &prod.Name, &prod.Qty, &prod.LastUpdated)

		if err != nil {
			return nil
		}
	} else {

		return nil
	}

	return prod

}

func AddProduct(product Product) {
	db, err := sql.Open("mysql", config.Cfg.DbUser+":"+config.Cfg.DbPass+"@tcp(127.0.0.1:3306)/"+config.Cfg.DbName)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	insert, err := db.Query(
		"INSERT INTO product (code,name,qty,last_updated) VALUES (?,?,?, now())",
		product.Code, product.Name, product.Qty)
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

}
