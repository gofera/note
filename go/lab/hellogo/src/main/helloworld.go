package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func (person *Person) CompareAge(another *Person) string {
	olderAge := person.Age - another.Age
	if olderAge < 0 {
		return fmt.Sprintf("%s is %d years younger than %s", person.Name, -olderAge, another.Name)
	} else if olderAge > 0 {
		return fmt.Sprintf("%s is %d years older than %s", person.Name, olderAge, another.Name)
	} else {
		return fmt.Sprintf("%s and %s has the same age.", person.Name, another.Name)
	}
}

type Cat struct {
}

type Eater interface {
	Eat(food string)
}

func (person *Person) Eat(food string) {
	fmt.Println(person.Name, "is eating", food)
}

func (cat *Cat) Eat(food string) {
	fmt.Println("Cat is eating", food)
}

func main() {
	fmt.Printf("hi %s\n", "wenzhe")
	wenzhe := Person{"Wenzhe", 18}
	qiqi := Person{"QiQi", 1}
	doudou := Person{"DouDou", 1}
	fmt.Println(wenzhe.CompareAge(&qiqi))
	fmt.Println(qiqi.CompareAge(&wenzhe))
	fmt.Println(qiqi.CompareAge(&doudou))

	helloKity := Cat{}

	eaters := [...]Eater{&qiqi, &helloKity}
	for _, eater := range eaters {
		eater.Eat("milk")
	}

}
