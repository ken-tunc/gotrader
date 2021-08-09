package main

import (
	"log"

	"github.com/ken-tunc/gotrader/src/config"

	gotrader "github.com/ken-tunc/gotrader/src"

	"github.com/ken-tunc/gotrader/src/api/bitflyer"
	"gorm.io/driver/sqlite"
)

func main() {
	c := config.Config

	client := bitflyer.NewClient(c.BitflyerKey, c.BitflyerSecret, c.HttpTimeout, c.WsTimeout)

	db, err := gotrader.OpenDb(sqlite.Open(c.DataSourceName), config.Config.GormConfig)
	if err != nil {
		log.Fatalf("cannot open database: %s", err)
	}

	gotrader.SubscribeTicker(c.ProductCode, c.TradeDuration, client, db)
}
