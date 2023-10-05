package pkg

type Army struct {
	Name     string
	Soldiers []Soldier
}

func (a *Army) AddSoldier(s Soldier) {
	a.Soldiers = append(a.Soldiers, s)
}

func (a *Army) Attack(enemy *Army) {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	for _, attackingSoldier := range a.Soldiers {
		if attackingSoldier.Health <= 0 {
			continue // Skip if this soldier is defeated
		}

		// Pick a random enemy soldier
		enemyIndex := rand.Intn(len(enemy.Soldiers))
		targetSoldier := &enemy.Soldiers[enemyIndex]

		// Skip if the target enemy soldier is already defeated
		if targetSoldier.Health <= 0 {
			continue
		}

		// Calculate damage (simplified)
		damage := attackingSoldier.Attack - targetSoldier.Defence
		if damage < 0 {
			damage = 0
		}

		// Inflict damage
		targetSoldier.Health -= damage
		if targetSoldier.Health < 0 {
			targetSoldier.Health = 0
		}
	}
}
