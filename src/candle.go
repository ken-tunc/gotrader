package gotrader

import (
	"log"
	"math"
	"time"

	"github.com/ken-tunc/gotrader/src/api/bitflyer"
)

type CandleDuration string

const (
	MINUTE CandleDuration = "MINUTE"
	HOUR   CandleDuration = "HOUR"
	DAY    CandleDuration = "DAY"
)

func CandleDurations() []CandleDuration {
	return []CandleDuration{MINUTE, HOUR, DAY}
}

func TruncateDateTime(dateTime time.Time, duration CandleDuration) time.Time {
	switch duration {
	case MINUTE:
		return dateTime.Truncate(time.Minute)
	case HOUR:
		return dateTime.Truncate(time.Hour)
	case DAY:
		return dateTime.Truncate(time.Hour * 24)
	}

	log.Fatalf("Invalid CandleDuration: %s", duration)
	return time.Time{}
}

type Candle struct {
	ProductCode string         `gorm:"primaryKey"`
	Duration    CandleDuration `gorm:"primaryKey"`
	Time        time.Time      `gorm:"primaryKey"`
	Open        float64
	Close       float64
	High        float64
	Low         float64
	Volume      float64
}

func NewCandle(ticker bitflyer.Ticker, duration CandleDuration) *Candle {
	price := ticker.MidPrice()
	return &Candle{
		ProductCode: ticker.ProductCode,
		Duration:    duration,
		Time:        TruncateDateTime(ticker.DateTime(), duration),
		Open:        price,
		Close:       price,
		High:        price,
		Low:         price,
		Volume:      ticker.Volume,
	}
}

func (c *Candle) Extend(ticker bitflyer.Ticker) {
	newPrice := ticker.MidPrice()
	c.High = math.Max(c.High, newPrice)
	c.Low = math.Min(c.Low, newPrice)
	c.Close = newPrice
	c.Volume += ticker.Volume
}
