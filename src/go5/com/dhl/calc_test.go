package main
import "testing"
func TestAddSum(t *testing.T){
	sum := AddSum(10)
	if  sum == 55 {
		t.Logf("与期望值相同：")
	}else {
		t.Fatalf("与期望值不同：")
	}
}
