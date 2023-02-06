package main

func f() {
	panic("stop")
}
func g() {
	f()
}
func h() {
	g()
}
func main() {
	h()
}
