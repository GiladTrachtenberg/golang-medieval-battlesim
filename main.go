package main

import (
	"fmt"

	"github.com/GiladTrachtenberg/golang-medieval-battlesim/pkg"
)

func main() {
	fmt.Println("Starting Medieval Battle Simulator")

	// Initialize two armies
	army1 := pkg.Army{Name: "Army of the North"}
	army2 := pkg.Army{Name: "Army of the South"}

	// Add soldiers to both armies
	soldier1 := pkg.Soldier{Name: "John", Health: 100, Attack: 10, Defence: 5}
	soldier2 := pkg.Soldier{Name: "Rob", Health: 100, Attack: 12, Defence: 4}

	army1.AddSoldier(soldier1)
	army2.AddSoldier(soldier2)

	// Place them on the battlefield
	battlefield := pkg.Battlefield{}
	battlefield.Armies = append(battlefield.Armies, army1, army2)

	// Start the battle
	battlefield.StartBattle()
}
