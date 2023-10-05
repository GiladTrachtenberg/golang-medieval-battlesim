package pkg

type Army struct {
	Name     string
	Soldiers []Soldier
}

func (a *Army) AddSoldier(s Soldier) {
	a.Soldiers = append(a.Soldiers, s)
}

func (a *Army) Attack(enemy *Army) {
	// Implement attack logic
}
