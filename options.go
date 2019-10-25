package main

import "fmt"

// These would be defined by the thing to be DI'd
type NeedsInt interface {
	getInt() int
}
type NeedsString interface {
	getString() string
}
type NeedsFloat interface {
	getFloat() float32
}

// Funciton or Library A defines the options it requires
type OptionsA interface {
	NeedsInt
	NeedsString
	NeedsFloat
}

func NeedsA(op OptionsA) {
	fmt.Println(op.getInt(), op.getString(), op.getFloat())
	// I can call code with a subset of my options
	NeedsB(op)
}

// Funciton or Library B defines the options it requires
type OptionsB interface {
	NeedsInt
	NeedsString
}

func NeedsB(op OptionsB) {
	fmt.Println(op.getInt(), op.getString())
	// Won't compile
	// fmt.Printf(op.getInt(), op.getString(), op.getFloat()
	// Won't compile
	// NeedsA(op)
}

// My code uses the options
type MyOptions struct {
	i int
	s string
	f float32
}

func (op MyOptions) getInt() int {
	return op.i
}

func (op MyOptions) getString() string {
	return op.s
}

func (op MyOptions) getFloat() float32 {
	return op.f
}

// Some other code uses different options, with a shared subset
type OtherOptions struct {
	other_i int
	other_s string
}

func (op OtherOptions) getInt() int {
	return op.other_i
}
func (op OtherOptions) getString() string {
	return op.other_s
}

// If the options I have are sufficient I just pass them it.
func main() {
	a := MyOptions{1, "hi", 3.14}
	b := OtherOptions{2, "bye"}

	NeedsA(a)
	NeedsB(a)
	// Won't compile
	// NeedsA(b)
	NeedsB(b)
}
