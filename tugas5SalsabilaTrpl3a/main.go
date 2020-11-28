package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	cm "pnp/Framework/git/order/common"
	"pnp/Framework/git/order/middleware"
	"pnp/Framework/git/order/services"
	"pnp/Framework/git/order/transport"

	log "github.com/Sirupsen/logrus"
	httptransport "github.com/go-kit/kit/transport/http"
)

func initHandlers() {
	var svc services.PaymentServices

	svc = services.PaymentService{}
	svc = middleware.BasicMiddleware()(svc)

	root := cm.Config.RootURL

	http.Handle(fmt.Sprintf("%s/orders", root), httptransport.NewServer(
		transport.OrderEndpoint(svc), transport.DecodeRequest, transport.EncodeResponse,
	))

	//Handler Customer
	http.Handle(fmt.Sprintf("%s/customer", root), httptransport.NewServer(
		transport.CustomerEndpoint(svc), transport.DecodeCustomerRequest, transport.EncodeResponse,
	))

	//Handler Product
	http.Handle(fmt.Sprintf("%s/product", root), httptransport.NewServer(
		transport.ProductEndpoint(svc), transport.DecodeProductRequest, transport.EncodeResponse,
	))

	//fastpay handler
	http.Handle(fmt.Sprintf("%s/fastpay", root), httptransport.NewServer(
		transport.FastEndpoint(svc), transport.DecodeFastPayRequest, transport.EncodeResponse,
	))

	//Call handler
	http.Handle(fmt.Sprintf("%s/call", root), httptransport.NewServer(
		transport.CallEndpoint(svc), transport.DecodeCallRequest, transport.EncodeResponse,
	))

	//InquiryStatus handler
	http.Handle(fmt.Sprintf("%s/status", root), httptransport.NewServer(
		transport.StatusEndpoint(svc), transport.DecodeStatusRequest, transport.EncodeResponse,
	))

}

var logger *log.Entry

func initLogger() {
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05.999",
	})

	//log.SetReportCaller(true)
}

func main() {
	configFile := flag.String("conf", "conf-dev.yml", "main configuration file")
	flag.Parse()
	initLogger()
	log.WithField("file", *configFile).Info("Loading configuration file")
	cm.LoadConfigFromFile(configFile)
	initHandlers()

	var err error
	if cm.Config.RootURL != "" || cm.Config.ListenPort != "" {
		err = http.ListenAndServe(cm.Config.ListenPort, nil)
	}

	if err != nil {
		log.WithField("error", err).Error("Unable to start the server")
		os.Exit(1)
	}
}
