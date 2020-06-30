package calc

import "testing"

func TestEven(t *testing.T) {
	t.Log("Start TestEven")
	//t.Error("This is a T.error func")
	if !Even(10) {
		t.Log("10 must be even")
		t.Fail()
	}
	if Even(77) {
		t.Error("77 is not even")
	}
}

func TestAdd(t *testing.T) {
	inputs := []struct{a,b,c int} {
		{a:1,b:2,c:3},
		{a:4,b:5,c:9},
		{a:10,b:20,c:30},
	}
	for _, data := range inputs {
		if result := Add(data.a, data.b); result != data.c {
			t.Errorf("Add(%d,%d) expected result = %d,actual result=%d\n",data.a,data.b,data.c,result)
		}
	}
}
