package main

import (
	"log"

	gotrader "github.com/ken-tunc/gotrader/src"
	"gorm.io/driver/sqlite"

	"github.com/ken-tunc/gotrader/src/config"

	"github.com/ken-tunc/gotrader/src/api/bitflyer"
)

func main() {
	c := config.LoadConfig()

	client := bitflyer.NewClient(c.BitflyerKey, c.BitflyerSecret, c.HttpTimeout, c.WsTimeout)

	db, err := gotrader.OpenDb(sqlite.Open(c.DataSourceName), c.GormConfig)
	if err != nil {
		log.Fatalf("cannot open database: %s", err)
	}

	gotrader.SubscribeTicker(c.ProductCode, c.TradeDuration, client, db)
}
