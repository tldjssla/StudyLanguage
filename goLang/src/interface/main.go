// interface.go project main.go
package main

import (
	"fmt"
)

type Duck struct {
}

func (d Duck) quack() {
	fmt.Println("꽤액")
}

func (d Duck) feathers() {
	fmt.Println("오리는 흰색털을 가지고 있습니다.")
}

type Person struct {
}

func (p Person) quack() {
	fmt.Println("꽥 사람이 오리를 흉내냅니다.")
}

func (p Person) feathers() {
	fmt.Println("사람은 땅에서 깃털을 주워서 보여줍니다.")
}

type Quacker interface {
	quack()
	feathers()
}

func inTheForest(q Quacker) {
	q.quack()
	q.feathers()
}
func main() {
	var donald Duck
	var john Person

	inTheForest(donald)
	inTheForest(john)

	if v, ok := interface{}(donald).(Quacker); ok {
		fmt.Println(v, ok)
	}
}
