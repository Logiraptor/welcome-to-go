package main

// INTERFACE OMIT
type Speaker interface {
	Speak() string
}

type Dog struct {
	/*...*/
}

func (d Dog) Speak() string {
	return "bark!"
}

type Cat struct {
	/*...*/
}

func (c Cat) Speak() string {
	return "meow!"
}

// END INTERFACE OMIT

func main() {

}
