package tips

import "testing"

func TestExpression(t *testing.T) {
	A := Three(true).If(1).Else(0).Value()
	if Int(A) == 0 {
		t.Error("Three Expresion was error should 1 ")
	}
	B := Three(false).If(1).Else(0).Value()
	if Int(B) != 0 {
		t.Error("Three Expresion was error should 10")
	}

	C := Three(true).If(100).Else(10).Int()
	if C != 100 {
		t.Error("Three Expresion value should is 100")
	}

	D := Three(true).If("yes").Else("no").String()
	if D != "yes" {
		t.Error("Three Expresion value should is yes")
	}
}

func TestIfThree(t *testing.T) {
	A := IfThree(true)(1, 0)
	if Int(A) != 1 {
		t.Error("IfThree value should be 1")
	}

	B := IfThree(false)("yes", "no")
	if String(B) != "no" {
		t.Error("ifThree value should be `no` when condition is false")
	}

}
