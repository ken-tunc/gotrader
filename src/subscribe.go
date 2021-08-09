package gotrader

import (
	"errors"
	"log"
	"time"

	"github.com/ken-tunc/gotrader/src/api/bitflyer"
	"gorm.io/gorm"
)

func SubscribeTicker(productCode string, tradeDuration CandleDuration, apiClient *bitflyer.Client, db *gorm.DB) {
	ch := make(chan bitflyer.Ticker)
	go apiClient.Realtime.SubscribeTicker(productCode, ch)

	for ticker := range ch {
		for _, duration := range CandleDurations() {
			created, err := feedCandlesByTicker(ticker, duration, db)
			if err != nil {
				log.Printf("cannot feed candles: %s", err)
			}
			if created && duration == tradeDuration {
				// TODO: trade
				log.Printf("trade")
			}
		}
	}
}

func feedCandlesByTicker(ticker bitflyer.Ticker, duration CandleDuration, db *gorm.DB) (bool, error) {
	created := false
	err := db.Transaction(func(tx *gorm.DB) error {
		candle := new(Candle)
		cond := &struct {
			ProductCode string
			Duration    CandleDuration
			Time        time.Time
		}{
			ProductCode: ticker.ProductCode,
			Duration:    duration,
			Time:        TruncateDateTime(ticker.DateTime(), duration),
		}
		err := tx.Take(candle, cond).Error

		if errors.Is(err, gorm.ErrRecordNotFound) {
			created = true
			return tx.Create(NewCandle(ticker, duration)).Error
		} else if err != nil {
			return err
		}

		candle.Extend(ticker)
		return tx.Save(candle).Error
	})

	return created, err
}
