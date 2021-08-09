package main

import (
	"log"

	"github.com/ken-tunc/gotrader/src/config"

	"github.com/ken-tunc/gotrader/src/api/bitflyer"
)

func main() {
	c := config.LoadConfig()

	client := bitflyer.NewClient(c.BitflyerKey, c.BitflyerSecret, c.HttpTimeout, c.WsTimeout)
	com, err := client.Commission.GetCommissionRate("BTC_JPY")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(com)

	//db, err := gotrader.OpenDb(sqlite.Open(c.DataSourceName), c.GormConfig)
	//if err != nil {
	//	log.Fatalf("cannot open database: %s", err)
	//}
	//
	//gotrader.SubscribeTicker(c.ProductCode, c.TradeDuration, client, db)
}
