package testunitario

import (
	"fmt"
	"testing"
)

func TestSuma(t *testing.T) {
	total := Suma(5, 5)

	if total != 10 {
		fmt.Println(t.Error("Error"))
	}
}
