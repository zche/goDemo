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

//func TestOdd(t *testing.T) {
//	if !Odd(11) {
//		t.Error("11 must be odd")
//	}
//	if Odd(70) {
//		t.Error("70 is not odd")
//	}
//}

func BenchmarkAdd(b *testing.B) {
	input := struct {
		a,b,c int
	}{
		3001,4002,7003,
	}
	for i:=0;i<b.N;i++ {
		if result := Add(input.a,input.b); result != input.c {
			b.Errorf("Add(%d,%d) expected result = %d,actual result=%d\n",input.a,input.b,input.c,result)
		}
	}
}
