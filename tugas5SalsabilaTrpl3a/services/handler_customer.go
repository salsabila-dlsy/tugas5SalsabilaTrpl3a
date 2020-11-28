package services

import (
	"context"
	"database/sql"
	"fmt"

	cm "pnp/Framework/git/order/common"

	_ "github.com/go-sql-driver/mysql"
)

func (PaymentService) CustomerHandler(ctx context.Context, req cm.Customer) (res cm.Customer) {

	var db *sql.DB
	var err error

	defer panicRecovery()

	host := cm.Config.Connection.Host
	port := cm.Config.Connection.Port
	user := cm.Config.Connection.User
	pass := cm.Config.Connection.Password
	data := cm.Config.Connection.Database

	var mySQL = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", user, pass, host, port, data)

	db, err = sql.Open("mysql", mySQL)

	if err != nil {
		panic(err.Error())
	}

	res.CustomerID = req.CustomerID

	var customer cm.Customer

	sql := `SELECT
				CustomerID,
				IFNULL(CompanyName,'') CompanyName,
				IFNULL(ContactName,'') ContactName,
				IFNULL(ContactTitle,'') ContactTitle,
				IFNULL(Address,'') Address,
				IFNULL(City,'') City,
				IFNULL(Country,'') Country,	
				IFNULL(Phone,'') Phone,
				IFNULL(PostalCode,'') PostalCode			
			FROM customers WHERE CustomerID = ?`

	result, err := db.Query(sql, req.CustomerID)

	defer result.Close()

	if err != nil {
		panic(err.Error())
	}

	for result.Next() {

		err := result.Scan(&customer.CustomerID, &customer.CompanyName, &customer.ContactName, &customer.ContactTitle,
			&customer.Address, &customer.City, &customer.Country, &customer.Phone, &customer.PostalCode)

		if err != nil {
			panic(err.Error())
		}

	}

	res = customer

	return
}
