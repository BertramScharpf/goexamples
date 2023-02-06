package main

func main() {
	s := "Hallöchen"
	r := []rune(s)

	println(string(r))
	println(string(r[4:6]))
	println(string(r[4:5]))
	println(string(r[:4]))
	println(string(r[5:]))
}
