package main

import (
	"os"

	"github.com/sirupsen/logrus"

	"github.com/boltdb/bolt"
)

// NewBoldClient return a instance of the database
func NewBoldClient(dbname string) *bolt.DB {
	db, err := bolt.Open(dbname+".db", 0644, nil)
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}

	return db
}
