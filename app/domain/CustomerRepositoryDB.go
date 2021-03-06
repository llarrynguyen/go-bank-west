package domain

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)
type CustomerRepositoryDB struct {
	client *sql.DB
}

func (d CustomerRepositoryDB ) FindAll()  ([]Customer,error) {


	findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"

	rows, err := d.client.Query(findAllSql)

	if err != nil {
		log.Println("Error while quering customer table" + err.Error())
	}

	customers := make([] Customer,0)
	for rows.Next(){
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City,&c.Zipcode,&c.DateOfBirth,&c.Status)

		if err != nil {
			log.Println("Error while scaning customers" + err.Error())
			return nil,err
		}
		customers = append(customers,c)
	}

	return customers, nil
}

func NewCustomerRepositoryDB() CustomerRepositoryDB{
	client, err := sql.Open("mysql", "root:Kwonwoo1!@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDB{client: client}
}