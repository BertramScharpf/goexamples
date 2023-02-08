package main

import "fmt"

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func (t *Tree) insertAll(vals ...int) *Tree {
	for _, val := range vals {
		t = t.insert(val)
	}
	return t
}

func (t *Tree) insert(v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if v <= t.Value {
		t.Left = t.Left.insert(v)
	} else {
		t.Right = t.Right.insert(v)
	}
	return t
}

func (t *Tree) String() string {
	s := ""
	if t != nil {
		if t.Left != nil {
			s += t.Left.String() + " "
		}
		s += fmt.Sprint(t.Value)
		if t.Right != nil {
			s += " " + t.Right.String()
		}
	}
	return "(" + s + ")"
}

func (t *Tree) Walk(ch chan int) {
	var walkImpl func(*Tree, chan int)
	walkImpl = func(t *Tree, ch chan int) {
		if t == nil {
			return
		}
		walkImpl(t.Left, ch)
		ch <- t.Value
		walkImpl(t.Right, ch)
	}
	walkImpl(t, ch)
	close(ch)
}

// NOTE: The implementation leaks goroutines when trees are different.
// See binarytrees_quit.go for a better solution.
func Same(t1, t2 *Tree) bool {
	w1, w2 := make(chan int), make(chan int)

	go t1.Walk(w1)
	go t2.Walk(w2)

	for {
		v1, ok1 := <-w1
		v2, ok2 := <-w2
		if !ok1 || !ok2 {
			return !ok1 && !ok2
		}
		if v1 != v2 {
			return false
		}
	}
}

func main() {
	var s, t *Tree
	s = s.insertAll(3, 1, 8, 1, 2, 5, 13)
	t = t.insertAll(8, 3, 13, 1, 5, 1, 2)
	fmt.Println(s)
	fmt.Println(t)
	if Same(s, t) {
		fmt.Println("Inhalt ist gleich.")
	} else {
		fmt.Println("Inhalt ist unterschiedlich.")
	}
}
