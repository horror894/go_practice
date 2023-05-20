package old

import (
	"fmt"
)

const ( // const for promts
	requestAvalibleWater = "Write how many ml of water the coffee machine has:"
	requestAvalibleMilk  = "Write how many ml of milk the coffee machine has:"
	requestAvalibleBeans = "Write how many grams of coffee beans the coffee machine has:"
)

const ( // const prompts for err
	inputErr = "Wrong input: %d \n"
)

const (
	WaterPerCup = 200
	MilkPerCup  = 50
	BeansPerCup = 15
)

type coffeMachineLimit struct {
	whaterLimit int
	milkLimit   int
	beansLimit  int
	cupsLimit   int
	moneyLimit  int
}

func reqestAvalibleResources() (coffeMachineLimit, error) {
	newLimits := coffeMachineLimit{}
	var err error
	fmt.Println(requestAvalibleWater)
	if _, err := fmt.Scanln(&newLimits.whaterLimit); err != nil {
		fmt.Printf(inputErr, err)
		return newLimits, err
	}
	fmt.Println(requestAvalibleMilk)
	if _, err := fmt.Scanln(&newLimits.milkLimit); err != nil {
		fmt.Printf(inputErr, err)
		return newLimits, err
	}
	fmt.Println(requestAvalibleBeans)
	if _, err := fmt.Scanln(&newLimits.beansLimit); err != nil {
		fmt.Printf(inputErr, err)
		return newLimits, err
	}

	return newLimits, err
}

func nomain() {
	// write your code here
	var cupCount int
	if newCoffeMachine, err := reqestAvalibleResources(); err == nil {
		fmt.Println("Write how many cups of coffee you will need:")
		if _, err := fmt.Scanln(&cupCount); err == nil {
			cupResourses := make([]int, 3)
			cupResourses[0] = newCoffeMachine.whaterLimit / 200
			cupResourses[1] = newCoffeMachine.milkLimit / 50
			cupResourses[2] = newCoffeMachine.beansLimit / 15
			min := cupResourses[0]
			for _, element := range cupResourses {
				if element < min {
					min = element
				}
			}

			if cupCount > min {
				fmt.Printf("No, I can make only %d cups of coffee", min)
			} else if cupCount < min {
				additionalCups := min - cupCount
				fmt.Printf("Yes, I can make that amount of coffee (and even %d more than that)", additionalCups)
			} else if cupCount == min {
				fmt.Printf("Yes, I can make that amount of coffee")
			}
		}

	} else {
		fmt.Printf("Wrong limit parameters: %d\n", err)
	}

}
