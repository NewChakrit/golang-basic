package customer

import "fmt"

var Name = "New"
var name = "Chakrit"

func Customer() {
	fmt.Println("Hello Customer")

}

type Person2 struct {
	Name    string
	Surname string
	Age     int
	gda     float32
	Gda     float32
}

type Person3 struct {
	name string
	age  int
}

func (p Person3) GetName() string {
	return p.name
}

func (p *Person3) SetName(name string) {
	p.name = name
}

func (p Person3) GetAge() int {
	return p.age
}
