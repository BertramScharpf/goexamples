package main

import "fmt"

var m = map[string]int{
	"foo": 33,
	"bar": 44,
	"baz": 55,
	"a": 101,
	"b": 102,
	"c": 103,
	"d": 104,
	"e": 105,
	"f": 106,
	"g": 107,
	"h": 108,
	"i": 109,
	"j": 110,
	"k": 111,
	"l": 112,
	"m": 113,
	"n": 114,
	"o": 115,
	"p": 116,
	"q": 117,
	"r": 118,
	"s": 119,
	"t": 120,
	"u": 121,
	"v": 122,
	"w": 123,
	"x": 124,
	"y": 125,
	"z": 126,
	"A": 201,
	"B": 202,
	"C": 203,
	"D": 204,
	"E": 205,
	"F": 206,
	"G": 207,
	"H": 208,
	"I": 209,
	"J": 210,
	"K": 211,
	"L": 212,
	"M": 213,
	"N": 214,
	"O": 215,
	"P": 216,
	"Q": 217,
	"R": 218,
	"S": 219,
	"T": 220,
	"U": 221,
	"V": 222,
	"W": 223,
	"X": 224,
	"Y": 225,
	"Z": 226,
}

func main() {
	for k, v := range m {
		fmt.Printf("%v: %v\n", k, v)
	}
}
