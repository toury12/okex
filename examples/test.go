package main

import (
	"context"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/toury/okex"
	"github.com/toury/okex/api"
	"github.com/toury/okex/requests/rest/market"
	"log"
	"time"
)

func MustJson(v interface{}) string {
	value, _ := jsoniter.MarshalToString(v)
	return value
}

func main() {
	apiKey := "****"
	secretKey := "****"
	passphrase := "****"
	dest := okex.NormalServer // The main API server
	ctx := context.Background()
	client, err := api.NewClient(ctx, apiKey, secretKey, passphrase, dest)
	if err != nil {
		log.Fatalln(err)
	}

	response, err := client.Rest.Account.GetConfig()
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Account Config %s", MustJson(response))

	//response2, err := client.Rest.PublicData.GetInstruments(public.GetInstruments{
	//	Uly:      "",
	//	InstID:   "",
	//	InstType: "SPOT",
	//})
	//log.Printf("public instruments %s", MustJson(response2))

	//response3, err := client.Rest.Market.GetTicker(market.GetTicker{
	//	InstId: "BTC-USDT",
	//})
	//log.Printf("Marker tickers %s", MustJson(response3))

	start := time.Now()
	fmt.Println(time.Now().Add(-10 * time.Hour).UnixMilli())
	response4, err := client.Rest.Market.GetCandlesticksHistory(market.GetCandlesticks{
		InstID: "BTC-USDT",
		Before: time.Now().Add(-10 * time.Hour).UnixMilli(),
		After:  time.Now().UnixMilli(),
		Limit:  50,
		Bar:    okex.Bar1m,
	})
	log.Printf("Marker tickers %s", MustJson(response4))
	fmt.Println(time.Now().Sub(start))

	//client.Rest.Market.GetCandlesticks(market.GetCandlesticks{
	//	InstID: "",
	//	After:  time.Now().Add(-10 * time.Hour).Unix(),
	//	Before: time.Now().Unix(),
	//	Limit:  500,
	//	Bar:    "",
	//})

	//errChan := make(chan *events.Error)
	//subChan := make(chan *events.Subscribe)
	//uSubChan := make(chan *events.Unsubscribe)
	//sucCh := make(chan *events.Success)
	//lCh := make(chan *events.Login)
	//oCh := make(chan *private.Order)
	//iCh := make(chan *public.Instruments)
	//
	//// to receive unique events individually in separated channels
	//client.Ws.SetChannels(errChan, subChan, uSubChan, lCh, sucCh)
	//
	//// subscribe into orders private channel
	//// it will do the login process and wait until authorization confirmed
	//err = client.Ws.Private.Order(ws_private_requests.Order{
	//	InstType: okex.SwapInstrument,
	//}, oCh)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//// subscribe into instruments public channel
	//// it doesn't need any authorization
	//err = client.Ws.Public.Instruments(ws_public_requests.Instruments{
	//	InstType: okex.SwapInstrument,
	//}, iCh)
	//if err != nil {
	//	log.Fatalln("Instruments", err)
	//}
	//
	//// starting on listening
	//for {
	//	select {
	//	case <-lCh:
	//		log.Print("[Authorized]")
	//	case sub := <-subChan:
	//		channel, _ := sub.Arg.Get("channel")
	//		log.Printf("[Subscribed]\t%s", channel)
	//	case uSub := <-uSubChan:
	//		channel, _ := uSub.Arg.Get("channel")
	//		log.Printf("[Unsubscribed]\t%s", channel)
	//	case err := <-client.Ws.ErrChan:
	//		log.Printf("[Error]\t%+v", err)
	//	case o := <-oCh:
	//		log.Print("[Event]\tOrder")
	//		for _, p := range o.Orders {
	//			log.Printf("\t%+v", p)
	//		}
	//	case i := <-iCh:
	//		log.Print("[Event]\tInstrument")
	//		for _, p := range i.Instruments {
	//			log.Printf("\t%+v", p)
	//		}
	//	case e := <-client.Ws.StructuredEventChan:
	//		log.Printf("[Event] STRUCTED:\t%+v", e)
	//		v := reflect.TypeOf(e)
	//		switch v {
	//		case reflect.TypeOf(events.Error{}):
	//			log.Printf("[Error] STRUCTED:\t%+v", e)
	//		case reflect.TypeOf(events.Subscribe{}):
	//			log.Printf("[Subscribed] STRUCTED:\t%+v", e)
	//		case reflect.TypeOf(events.Unsubscribe{}):
	//			log.Printf("[Unsubscribed] STRUCTED:\t%+v", e)
	//		}
	//	case e := <-client.Ws.RawEventChan:
	//		log.Printf("[Event] RAW:\t%+v", e)
	//	case b := <-client.Ws.DoneChan:
	//		log.Printf("[End]:\t%v", b)
	//		return
	//	}
	//}
}
