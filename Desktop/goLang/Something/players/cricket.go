package players

import (
	"fmt"
)

type CricketPlayer struct {
	Name    string
	Age     int
	Address string
}

func (c CricketPlayer) Batting(name string, age int) (string, error) {
	fmt.Println("I am a cricket player")

	return "I am cricket player who is batting ", nil
}

func (c CricketPlayer) Bowling(naem string, age int) (string, error) {

	return "I am cricket player who is bowling ", nil
}

func NewCricketPlayer(name string, age int) CricketPlayer {
	return CricketPlayer{
		Name: name,
		Age:  age,
	}
}
