package config

import (
	"io"
	"log"
	"os"
	"strconv"
	"time"

	gotrader "github.com/ken-tunc/gotrader/src"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type AppConfig struct {
	ProductCode   string
	TradeDuration gotrader.CandleDuration

	DataSourceName string
	GormConfig     *gorm.Config

	BitflyerKey    string
	BitflyerSecret string

	HttpTimeout time.Duration
	WsTimeout   time.Duration
}

var Config AppConfig

func init() {
	logFile, ok := os.LookupEnv("GOTRADER_LOGFILE")
	if ok {
		initLogging(logFile)
	}

	location, ok := os.LookupEnv("TZ")
	if ok {
		initLocation(location)
	}

	Config.ProductCode = mustLoadEnvStr("GOTRADER_PRODUCT_CODE")
	Config.TradeDuration = gotrader.CandleDuration(mustLoadEnvStr("GOTRADER_TRADE_DURATION"))

	Config.DataSourceName = loadEnvStr("GOTRADER_DSN", ":memory:")
	Config.GormConfig = &gorm.Config{
		Logger: logger.New(log.Default(), logger.Config{IgnoreRecordNotFoundError: true}),
	}

	Config.BitflyerKey = mustLoadEnvStr("GOTRADER_BITFLYER_KEY")
	Config.BitflyerSecret = mustLoadEnvStr("GOTRADER_BITFLYER_SECRET")

	Config.HttpTimeout = time.Second * time.Duration(loadEnvInt("GOTRADER_HTTP_TIMEOUT", 5))
	Config.WsTimeout = time.Second * time.Duration(loadEnvInt("GOTRADER_WS_TIMEOUT", 10))
}

func mustLoadEnvStr(key string) string {
	value, ok := os.LookupEnv(key)
	if ok {
		return value
	}
	log.Fatalf("cannot load config key=%s", key)
	return ""
}

func loadEnvStr(key string, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if ok {
		return value
	}
	return defaultValue
}

func loadEnvInt(key string, defaultValue int) int {
	valueStr, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		log.Printf("cannot convert config to int. key=%s, error=%s", key, err)
		return defaultValue
	}
	return value
}

func initLogging(logFile string) {
	f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("cannot initialize f: %s", err)
	}
	multiLogFile := io.MultiWriter(os.Stdout, f)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)
}

func initLocation(location string) {
	loc, err := time.LoadLocation(location)
	if err != nil {
		log.Printf("cannot load location: %s", location)
		return
	}
	time.Local = loc
}
