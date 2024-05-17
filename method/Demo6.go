package main

type T5 struct {
	a int
}

func (t T5) M1() {
	t.a = 10
}

func (t *T5) M2() {
	t.a = 11
}

func main() {
	var t T5
	println(t.a) // 0

	t.M1()
	println(t.a) // 0

	p := &t
	p.M2()
	println(t.a) // 11
}
