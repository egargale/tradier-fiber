package rest

import (
	"log"

	"github.com/egargale/tradier-fiber/internals/util"

	"github.com/davecgh/go-spew/spew"
	"github.com/valyala/fasthttp"
	// "github.com/spf13/viper"
)

func MyTradier_Stream() {
	params := DefaultParams(util.MyConfig.TradierKey)
	client := NewClient(params)
	client.SelectAccount(util.MyConfig.TradierAccount)
	stream, err := client.GetSessionID()
	if err != nil {
		log.Printf("Opening Session: Error is %s", err)
	}
	spew.Dump(stream)
}

func MyTradier_Market() {
	params := DefaultParams(util.MyConfig.TradierKey)
	client := NewClient(params)
	client.SelectAccount(util.MyConfig.TradierAccount)
	market, err := client.GetMarketState()
	if err != nil {
		log.Printf("Getting Market Status: Error is %s", err)
	}
	spew.Dump(market)
}
func MyTradierSandboxPosition() {
	params := ClientParams{
		Endpoint:   SandboxEndpoint,
		AuthToken:  util.MyConfig.TradierSandboxKey,
		Client:     &fasthttp.Client{},
		RetryLimit: defaultRetries,
		Account:    util.MyConfig.TradierSandboxAccount,
	}
	client := NewClient(params)
	//spew.Dump(params)
	market, err := client.GetAccountPositions()
	if err != nil {
		log.Printf("Getting Market Status: Error is %s", err)
	}
	spew.Dump(market)
}
func MyTradierProductionPosition() {
	params := DefaultParams(util.MyConfig.TradierKey)
	client := NewClient(params)
	client.SelectAccount(util.MyConfig.TradierAccount)
	//spew.Dump(params)
	market, err := client.GetAccountPositions()
	if err != nil {
		log.Printf("Getting Market Status: Error is %s", err)
	}
	spew.Dump(market)
}
