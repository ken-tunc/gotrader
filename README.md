# gotrader

[![Build and Test](https://github.com/ken-tunc/gotrader/actions/workflows/build-and-test.yml/badge.svg)](https://github.com/ken-tunc/gotrader/actions/workflows/build-and-test.yml)

A simple cryptocurrency trading bot written in golang.

## Environment variables

| name                     | description                                                                                            | mandatory | example      | 
| ------------------------ | ------------------------------------------------------------------------------------------------------ | --------- | ------------ | 
| TZ                       | Location name corresponding to a file in the IANA Time Zone database. Defaults to the system timezone. |           | Asia/Tokyo   | 
| GOTRADER_LOGFILE         | File path to application log file. If not set, logs will be output to the console.                     |           | gotrader.log | 
| GOTRADER_PRODUCT_CODE    | Product code for tickers to subscribe to and feed for each duration.                                   | Y         | BTC_JPY      | 
| GOTRADER_TRADE_DURATION  | Duration of trading opportunity. Must be one of `CandleDuration`.                                      | Y         | HOUR         | 
| GOTRADER_DSN             | Data source name of SQLite. If not set, in memory database will be used.                               |           | gotrader.db  | 
| GOTRADER_BITFLYER_KEY    | API key of the bitflyer lightning API.                                                                 | Y         |              | 
| GOTRADER_BITFLYER_SECRET | API secret of the bitflyer lightning API.                                                              | Y         |              | 
| GOTRADER_HTTP_TIMEOUT    | Timeout seconds of http requests. Defaults to 5 sec.                                                   |           | 10           | 
| GOTRADER_WS_TIMEOUT      | Timeout seconds of websocket connection. Defaults to 10 sec.                                           |           | 10           | 
