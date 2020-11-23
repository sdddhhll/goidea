package main

import "testing"

func TestAddUpper(t *testing.T) {
	sum := AddUpper(10)
	if sum != 55 {
		t.Fatalf("错误，期望值是55")
	}
	t.Logf("成功")
}