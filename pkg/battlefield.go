package pkg

import "fmt"

type Battlefield struct {
	Armies []Army
}

func (a *Army) IsDefeated() bool {
	for _, soldier := range a.Soldiers {
		if soldier.Health > 0 {
			return false
		}
	}
	return true
}

func (bf *Battlefield) StartBattle() {
	for {
		// Army 1's turn to attack
		for _, soldier := range bf.Armies[0].Soldiers {
			if soldier.Health > 0 {
				bf.Armies[1].Soldiers[0].TakeDamage(soldier.Attack) // Simplified; only attacks first soldier of enemy army
			}
		}

		// Check if Army 2 is defeated
		if bf.Armies[1].IsDefeated() {
			fmt.Println("Army of the South is defeated!")
			break
		}

		// Army 2's turn to attack
		for _, soldier := range bf.Armies[1].Soldiers {
			if soldier.Health > 0 {
				bf.Armies[0].Soldiers[0].TakeDamage(soldier.Attack) // Simplified; only attacks first soldier of enemy army
			}
		}

		// Check if Army 1 is defeated
		if bf.Armies[0].IsDefeated() {
			fmt.Println("Army of the North is defeated!")
			break
		}
	}
}
