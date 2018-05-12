package flyweight

import "testing"

func TestTeamFlyweightFactory(t *testing.T) {
	t.Run("GetTeam", func(t *testing.T) {
		factory := TeamFlyweightFactory()
		teamA1 := factory.GetTeam(TeamA)

		if teamA1 == nil {
			t.Error("The pointer to the TeamA was nil")
		}
		teamA2 := factory.GetTeam(TeamA)

		if teamA2 == nil {
			t.Error("The pointer to the TeamA was nil")
		}
		if teamA1 != teamA2 {
			t.Error("TeamA pointers weren't the same")
		}
		if factory.GetNumberOfObjects() != 1 {
			t.Errorf("The number of objects created was not 1, got: %d\n", factory.GetNumberOfObjects())
		}
	})

	t.Run("GetTeam with high volume", func(t *testing.T) {
		factory := TeamFlyweightFactory()
		teams := make([]*Team, 500000*2)

		for i := 0; i < 500000; i++ {
			teams[i] = factory.GetTeam(TeamA)
		}
		for i := 500000; i < 500000*2; i++ {
			teams[i] = factory.GetTeam(TeamB)
		}
		if factory.GetNumberOfObjects() != 2 {
			t.Errorf("The number of objects created was not 2, got: %d", factory.GetNumberOfObjects())
		}
	})
}
