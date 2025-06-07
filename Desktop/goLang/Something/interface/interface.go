package Interface

type Allrounder interface {
	Batting(string, int) (string, error)
	Bowling(string, int) (string, error)
}
