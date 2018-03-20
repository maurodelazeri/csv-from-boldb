package main

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/boltdb/bolt"
)

type OrderDone struct {
	LimitPrice float64
	Last       float64
	Profit     float64
	Side       string
	Change     float64
	Product    string
	Exchange   string
	Reason     string
	Time       time.Time
}

func main() {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open("binance-BTCUSDT.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	restoredone := []OrderDone{}

	// Executed orders
	db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte("limit"))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			dataOrder := &OrderDone{}
			err := json.Unmarshal(v, dataOrder)
			if err != nil {

			}

			AppendStringToFile("mauro", dataOrder.Time.String()+","+FloatToString(dataOrder.Profit)+"\n")

			restoredone = append(restoredone, *dataOrder)
		}

		return nil
	})

}

func FloatToString(input_num float64) string {
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}

func AppendStringToFile(path, text string) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(text)
	if err != nil {
		return err
	}
	return nil
}
