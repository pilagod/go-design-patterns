package composite

import "testing"

func TestCompositeSwimmer(t *testing.T) {
	swimmer := AlthleteCompositeSwimmer{
		Swim: Swim,
	}
	swimmer.Althlete.Train()
	swimmer.Swim()
}

func TestCompositeShark(t *testing.T) {
	shark := Shark{
		Swim: Swim,
	}
	shark.Eat()
	shark.Swim()
}

func TestCompositeInterfaceSwimmer(t *testing.T) {
	swimmer := InterfaceCompositeSwimmer{
		&Athlete{},
		&SwimmerImpl{},
	}
	swimmer.Train()
	swimmer.Swim()
}
