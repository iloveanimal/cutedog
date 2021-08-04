package utils

import (
	"fmt"
	"os"
	"strings"
)

const (
	BETFAIR_TOPIC            = "topic_betfair_data"
	BETFAIR_GROUP_ID         = "group_python_betfair_odds"
	BETFAIR_ORDER_BOOK_TOPIC = "topic_betfair_order_book_ws"
	BETFAIR_PRODUCT_TOPIC    = "topic_betfair_product"
)

type Config struct {
	KafkaHost          string
	DbDriver           string
	OdiallDbSource     string
	CkexDbSource       string
	TestOdiallDbSource string
	TestCkexDbSource   string
}

func LoadConfig() (config Config) {
	return Config{
		KafkaHost:          os.Getenv("KAFKA_HOST"),
		DbDriver:           os.Getenv("DB_DRIVER"),
		OdiallDbSource:     MakeDbSource("ODIALL", false),
		CkexDbSource:       MakeDbSource("CKEX", false),
		TestOdiallDbSource: MakeDbSource("ODIALL", true),
		TestCkexDbSource:   MakeDbSource("CKEX", true),
	}
}

func MakeDbSource(dbname string, forTest bool) string {
	dbname_upper := strings.ToUpper(dbname)
	host := os.Getenv("POSTGRES_HOST")
	if forTest {
		host = os.Getenv("TEST_POSTGRES_HOST")
	}
	db := os.Getenv(dbname_upper + "_DB")
	user := os.Getenv(dbname_upper + "_POSTGRES_USER")
	password := os.Getenv(dbname_upper + "_POSTGRES_PASSWORD")
	return fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable", user, password, host, db)
}
