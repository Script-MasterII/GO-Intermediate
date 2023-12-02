package main

import (
	"testing"
)

/*
	 func TestSum(t *testing.T) {
		total := Sum(5, 5)
		if total != 11 {
			t.Errorf("Sum was incorrect, got %d expectd %d", total, 10)
		}
	}
*/
func TestSum(t *testing.T) {
	tables := []struct {
		a int
		b int
		c int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{25, 26, 51},
	}
	for _, item := range tables {
		total := Sum(item.a, item.b)
		if total != item.c {
			t.Errorf("Sum was incorrect, got %d expected %d", total, item.c)
		}
	}
}
func TestMax(t *testing.T) {
	tables := []struct {
		a int
		b int
		c int
	}{
		{4, 2, 4},
		{3, 2, 3},
		{3, 5, 5},
	}

	for _, item := range tables {
		max := GetMax(item.a, item.b)
		if max != item.c {
			t.Errorf("GEtMax  was incorrect, got %d, expectd %d", max, item.c)
		}
	}
}
