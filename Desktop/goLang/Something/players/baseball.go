package players

import "fmt"

type FootballPlayer struct {
	Name    string
	sex     string
	matches int
}

func (ft FootballPlayer) Batting(name string, age int) (string, error) {
	fmt.Println("I am a football player batting")
	return "My name is " + name, nil
}

func (ft FootballPlayer) Bowling(name string, age int) (string, error) {
	fmt.Println("I am a football player bowlingk")
	return "hey there", nil
}

func NewFootBallPlayer(name string, matches int) FootballPlayer {
	return FootballPlayer{
		Name:    name,
		matches: matches,
		sex:     "MALE",
	}
}
