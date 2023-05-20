package main

import (
	"fmt"
)

type coffeeMachine struct {
	water uint
	milk  uint
	beans uint
	cups  uint
	money uint
}

func (machine coffeeMachine) String() string {
	return fmt.Sprintf(`
The coffee machine has:
%v of water
%v of milk
%v of coffee beans
%v of disposable cups
%v of money
`,
		machine.water, machine.milk, machine.beans, machine.cups, machine.money,
	)
}

func (machine *coffeeMachine) maxAmount(id recepieId) uint {

	recepie := menu[id]
	waterMax := machine.water / recepie.waterRequirement
	milkMax := machine.milk / recepie.milkRequirement
	beansMax := machine.beans / recepie.beansRequirement
	cupsMax := machine.cups / recepie.cupsRequirement

	return reduce(
		func(acc, cur uint) uint {
			if cur < acc {
				return cur
			} else {
				return acc
			}
		},
		waterMax, milkMax, beansMax, cupsMax,
	)
}

func (machine *coffeeMachine) buy() {
	
	fmt.Println("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino:")
	var id recepieId;
	fmt.Scanln(&id)
	recepie := menu[id]
	machine.water -= recepie.waterRequirement
	machine.milk -= recepie.milkRequirement
	machine.beans -= recepie.beansRequirement
	machine.cups -= recepie.cupsRequirement
	machine.money += recepie.valueRequested

}

func (machine *coffeeMachine) take() {
	fmt.Printf("I gave you $%v\n", machine.money)
	machine.money = 0
}

func (machine *coffeeMachine) fill() {

	fmt.Println("Write how many ml of water you want to add:")
	var waterToAdd uint
	fmt.Scanln(&waterToAdd)
	machine.water += waterToAdd

	fmt.Println("Write how many ml of milk you want to add:")
	var milkToAdd uint
	fmt.Scanln(&milkToAdd)
	machine.milk += milkToAdd

	fmt.Println("Write how many grams of coffee beans you want to add:")
	var beansToAdd uint
	fmt.Scanln(&beansToAdd)
	machine.beans += beansToAdd

	fmt.Println("Write how many disposable coffee cups you want to add:")
	var cupsToAdd uint
	fmt.Scanln(&cupsToAdd)
	machine.cups += cupsToAdd
}

type recepie struct {
	waterRequirement uint
	milkRequirement  uint
	beansRequirement uint
	valueRequested   uint
	cupsRequirement  uint
}

type recepieId uint

const (
	espresso recepieId = iota + 1
	latte
	cappuccino
)

var menu = map[recepieId]recepie{
	espresso: recepie{
		waterRequirement: 250,
		milkRequirement:  0,
		beansRequirement: 16,
		valueRequested:   4,
		cupsRequirement:  1,
	},
	latte: recepie{
		waterRequirement: 350,
		milkRequirement:  75,
		beansRequirement: 20,
		valueRequested:   7,
		cupsRequirement:  1,
	},
	cappuccino: recepie{
		waterRequirement: 200,
		milkRequirement:  100,
		beansRequirement: 12,
		valueRequested:   6,
		cupsRequirement:  1,
	},
}

var actions = map[string]func(){}

func main() {

	machine := coffeeMachine{
		water: 400,
		milk:  540,
		beans: 120,
		cups:  9,
		money: 550,
	}

	initActionsMap(&machine)

	fmt.Println(machine)
	action := getAction()

	actions[action]()

	fmt.Println(machine)

}

func initActionsMap(machine *coffeeMachine) {
	actions["fill"] = machine.fill
	actions["take"] = machine.take
	actions["buy"] = machine.buy
}

func getAction() string {
	var action string
	fmt.Println("Write action (buy, fill, take):")
	fmt.Scanln(&action)
	return action
}

func reduce(function func(uint, uint) uint, initial uint, sequence ...uint) uint {
	result := initial

	for _, v := range sequence {
		result = function(result, v)
	}

	return result
}
