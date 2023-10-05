package pkg

type Soldier struct {
	Name    string
	Health  int
	Attack  int
	Defence int
}

func (s *Soldier) TakeDamage(damage int) {
	actualDamage := damage - s.Defence
	if actualDamage < 0 {
		actualDamage = 0
	}
	s.Health -= actualDamage
	if s.Health < 0 {
		s.Health = 0
	}
}
