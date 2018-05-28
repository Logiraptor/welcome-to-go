package main

// START OMIT
type Age int

type Person struct {
	FirstName string
	LastName string
	age Age
}
func (p Person) DoPersonThings() {}

type PersonPointer *Person
type PersonMap map[string]Person
type PersonSlice []Person
type PersonArray [5]Person
type PersonChan chan Person
type PersonFunc func(Person) bool
type PersonInterface interface {
	DoPersonThings()
}
// END OMIT
