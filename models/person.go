package models

type Person struct{
	ID string
	Name string
	Type string
}
type PersonUpdate struct{
	Name string
	Type string
}