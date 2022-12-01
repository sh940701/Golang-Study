package go_test

import (
	"fmt"
	clt "lecture/go-test/controller"
	"testing"
)

func TestCalc(t *testing.T) {
	fmt.Println("Hello world!")
	ressum := clt.Sum(1, 2)

	t.Errorf("sum = %d", ressum)

	if ressum ==3 {
		t.Error("fail calc")
	}
}