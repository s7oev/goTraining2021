package fighters

import "testing"

func assert(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got %s but wanted %s", got, want)
	}
}

func TestFighterAttack(t *testing.T) {
	t.Run("first fighter is winner", func(t *testing.T) {
		givenF1 := Fighter{
			name:            "Peter",
			health:          100,
			damagePerAttack: 10,
		}

		givenF2 := Fighter{
			name:            "George",
			health:          100,
			damagePerAttack: 7,
		}

		got := fight(givenF1, givenF2)
		want := "Peter"

		assert(t, got, want)
	})

	t.Run("second fighter is winner", func(t *testing.T) {
		givenF1 := Fighter{
			name:            "Peter",
			health:          100,
			damagePerAttack: 7,
		}

		givenF2 := Fighter{
			name:            "George",
			health:          100,
			damagePerAttack: 10,
		}

		got := fight(givenF1, givenF2)
		want := "George"

		assert(t, got, want)
	})

	t.Run("draw", func(t *testing.T) {
		givenF1 := Fighter{
			name:            "Peter",
			health:          100,
			damagePerAttack: 10,
		}

		givenF2 := Fighter{
			name:            "George",
			health:          100,
			damagePerAttack: 10,
		}

		got := fight(givenF1, givenF2)
		want := "draw"

		assert(t, got, want)
	})
}

func TestFighterAttackMethod(t *testing.T) {
	t.Run("first fighter is winner", func(t *testing.T) {
		givenF1 := Fighter{
			name:            "Ivan",
			health:          100,
			damagePerAttack: 10,
		}

		givenF2 := Fighter{
			name:            "George",
			health:          100,
			damagePerAttack: 17,
		}

		got := (&givenF1).fight(&givenF2)
		want := "George"

		assert(t, got, want)
	})

	t.Run("draw", func(t *testing.T) {
		givenF1 := Fighter{
			name:            "Ivan",
			health:          100,
			damagePerAttack: 17,
		}

		givenF2 := Fighter{
			name:            "George",
			health:          100,
			damagePerAttack: 17,
		}

		got := givenF1.fight(&givenF2)
		want := "draw"

		assert(t, got, want)
	})
}