package services

import (
	"context"
	"database/sql"
	"fmt"

	cm "pnp/Framework/git/order/common"

	_ "github.com/go-sql-driver/mysql"
)

func (PaymentService) InquiryStatusHandler(ctx context.Context, req cm.InquiryStatusReq) (res cm.InquiryStatusRes) {

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

	var statResponse cm.InquiryStatusRes

	sql := `SELECT
				IFNULL(merchant,''),
				IFNULL(payment_reff,''),
				IFNULL(payment_date,''),
				IFNULL(payment_status_code,''),
				IFNULL(payment_status_desc,''),
				IFNULL(response_code,''),
				IFNULL(response_desc,'')
			FROM transaksi WHERE trx_id = ?`

	result, err := db.Query(sql, req.TransactionID)

	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	for result.Next() {

		err := result.Scan(&statResponse.Merchant, &statResponse.PaymentRefference, &statResponse.PaymentDate,
			&statResponse.PaymentStatusCode, &statResponse.PaymentStatusDesc, &statResponse.ResponseCode,
			&statResponse.ResponseDesc)

		if err != nil {
			panic(err.Error())
		}

	}

	statResponse.Response = req.Request
	statResponse.TransactionID = req.TransactionID
	statResponse.MerchantID = req.MerchantID
	statResponse.BillNO = req.BillNO

	res = statResponse

	return
}
