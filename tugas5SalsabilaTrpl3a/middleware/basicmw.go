package middleware

import (
	"time"

	"context"

	cm "pnp/Framework/git/order/common"
	"pnp/Framework/git/order/services"

	log "github.com/Sirupsen/logrus"
)

func BasicMiddleware() services.ServiceMiddleware {
	return func(next services.PaymentServices) services.PaymentServices {
		return BasicMiddlewareStruct{next}
	}
}

type BasicMiddlewareStruct struct {
	services.PaymentServices
}

func (mw BasicMiddlewareStruct) OrderHandler(ctx context.Context, request cm.Message) cm.Message {
	defer func(begin time.Time) {
		log.WithField("execTime", float64(time.Since(begin).Nanoseconds())/float64(1e6)).Info("OrderHandler ends")
	}(time.Now())

	log.WithField("request", request).Info("OrderHandler begins")

	return mw.PaymentServices.OrderHandler(ctx, request)

}

func (mw BasicMiddlewareStruct) CustomerHandler(ctx context.Context, request cm.Customer) cm.Customer {
	defer func(begin time.Time) {
		log.WithField("execTime", float64(time.Since(begin).Nanoseconds())/float64(1e6)).Info("CustomerHandler ends")
	}(time.Now())

	log.WithField("request", request).Info("CustomerHandler begins")

	return mw.PaymentServices.CustomerHandler(ctx, request)

}

func (mw BasicMiddlewareStruct) ProductHandler(ctx context.Context, request cm.Product) cm.Product {
	defer func(begin time.Time) {
		log.WithField("execTime", float64(time.Since(begin).Nanoseconds())/float64(1e6)).Info("ProductHandler ends")
	}(time.Now())

	log.WithField("request", request).Info("ProductHandler begins")

	return mw.PaymentServices.ProductHandler(ctx, request)

}

func (mw BasicMiddlewareStruct) FastPayHandler(ctx context.Context, request cm.FastPayRequest) cm.FastPayResponse {
	defer func(begin time.Time) {
		log.WithField("execTime", float64(time.Since(begin).Nanoseconds())/float64(1e6)).Info("FastPayHandler ends")
	}(time.Now())

	log.WithField("request", request).Info("FastPayHandler begins")

	return mw.PaymentServices.FastPayHandler(ctx, request)

}

func (mw BasicMiddlewareStruct) CallHandler(ctx context.Context, request cm.FastPayRequest) cm.FastPayResponse {
	defer func(begin time.Time) {
		log.WithField("execTime", float64(time.Since(begin).Nanoseconds())/float64(1e6)).Info("CallHandler ends")
	}(time.Now())

	log.WithField("request", request).Info("CallHandler begins")

	return mw.PaymentServices.FastPayHandler(ctx, request)

}

func (mw BasicMiddlewareStruct) InquiryStatusHandler(ctx context.Context, request cm.InquiryStatusReq) cm.InquiryStatusRes {
	defer func(begin time.Time) {
		log.WithField("execTime", float64(time.Since(begin).Nanoseconds())/float64(1e6)).Info("InquiryStatusHandler ends")
	}(time.Now())

	log.WithField("request", request).Info("InquiryStatusHandler begins")

	return mw.PaymentServices.InquiryStatusHandler(ctx, request)

}
