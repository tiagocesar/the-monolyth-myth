package main

import (
	"testing"
	"time"
)

type ShopDBMock struct{}

func TestCalculateSalesRate(t *testing.T) {
	shopDBMock := ShopDBMock{}
	sr, err := calculateSalesRate(&shopDBMock)
	if err != nil {
		t.Fatal(err)
	}

	exp := "0.33"
	if sr != exp {
		t.Fatalf("got %v; expected %v", sr, exp)
	}
}

func (s *ShopDBMock) CountCustomers(_ time.Time) (int, error) {
	return 1000, nil
}

func (s *ShopDBMock) CountSales(_ time.Time) (int, error) {
	return 333, nil
}
