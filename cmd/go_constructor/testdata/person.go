package example

import "time"

//genconstructor
type Person struct {
	id        string `required:""`
	name      string `required:""`
	tags      []string
	createdAt time.Time `required:"time.Now()"`
}

//genconstructor -p
type PersonService struct {
	id string `required:""`
}

//genconstructor
type ArmedWarrior struct {
	Person   `required:""`
	armament string `required:""`
}
