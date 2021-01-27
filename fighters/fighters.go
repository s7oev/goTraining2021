package fighters

type Fighter struct {
	name string
	health int
	damagePerAttack int
}

func fight(f1, f2 Fighter) string {
	hasDeadFighter := false

	for hasDeadFighter != true {
		f1.health = f1.health - f2.damagePerAttack
		f2.health = f2.health - f1.damagePerAttack

		if f1.health <= 0 && f2.health > 0 {
			return f2.name
		}

		if f2.health <= 0 && f1.health > 0 {
			return f1.name
		}

		if f1.health <= 0 || f2.health <= 0 {
			hasDeadFighter = true
		}
	}

	return "draw"
}

func(f *Fighter) fight(otherFighter *Fighter) string {
	hasDeadFighter := false

	for hasDeadFighter != true {
		f.health = f.health - otherFighter.damagePerAttack
		otherFighter.health = otherFighter.health - f.damagePerAttack

		if f.health <= 0 && otherFighter.health > 0 {
			return otherFighter.name
		}

		if otherFighter.health <= 0 && f.health > 0 {
			return f.name
		}

		if f.health <= 0 || otherFighter.health <= 0 {
			hasDeadFighter = true
		}
	}

	return "draw"
}
