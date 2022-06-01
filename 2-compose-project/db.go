package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

// CheckSalesRate is never invoked - is here just to showcase the logic we will use in the tests
func CheckSalesRate() {
	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", dbUser, dbPass, dbHost, dbSchema)

	db, err := sql.Open("postgres", connStr)
	checkErr(err)
	defer db.Close()

	shopDB := ShopDB{db}
	sr, err := calculateSalesRate(&shopDB)
	checkErr(err)

	fmt.Printf(sr)
}

func calculateSalesRate(sm ShopModel) (string, error) {
	since := time.Now().Add(-24 * time.Hour)

	sales, err := sm.CountSales(since)
	if err != nil {
		return "", err
	}

	customers, err := sm.CountCustomers(since)
	if err != nil {
		return "", err
	}

	rate := float64(sales) / float64(customers)

	return fmt.Sprintf("%.2f", rate), nil
}

type ShopModel interface {
	CountCustomers(time.Time) (int, error)
	CountSales(time.Time) (int, error)
}

type ShopDB struct {
	*sql.DB
}

func (s *ShopDB) CountCustomers(since time.Time) (int, error) {
	return s.count("customers", since)
}

func (s *ShopDB) CountSales(since time.Time) (int, error) {
	return s.count("sales", since)
}

func (s *ShopDB) count(table string, since time.Time) (int, error) {
	var count int
	err := s.QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE timestamp > $1", table), since).Scan(&count)
	return count, err
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
