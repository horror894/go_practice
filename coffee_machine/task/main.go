package main

import (
	"fmt"
	"os"
	"strings"
)

// How I can optimize func buy(), too many repeating of code

// Add error handling

// Create separate function for different types of coffee or one but with params

// What I need to think about - It's not cool that I take object and transmitting it through many func, I think it will
// be more transparent if func return result and I obviously transmit this result

// Add \n between output

type coffeeMachineState struct {
	water int
	milk  int
	beans int
	cups  int
	money int
}

var userBalance int

const ( // Prompts
	statusMess  = "The coffee machine has:\n"
	statusWater = "%d ml of water\n"
	statusMilk  = "%d ml of milk\n"
	statusBeans = "%d g of coffee beans\n"
	statusCups  = "%d disposable cups\n"
	statusMoney = "%d$ of money\n"

	actionRequest   = "Write action (1-fill balance, 2-show balance, 3-take rest, 4-buy, 5-fill, 6-take, 7-remaining, 8-exit):\n"
	actionBuyCoffee = "What do you want to buy? 1-espresso, 2-latte, 3-cappuccino, 4-back to main menu:\n"
	actionTake      = "I give you $%d\n"
	actionCooked    = "I have enough resources, making you a coffee!\n Take your coffee!\n"
	actionTakeRest  = "Pick up the rest: %d\n"

	fillWater = "Write how many ml of water you want to add:\n"
	fillMilk  = "Write how many ml of milk you want to add:\n"
	fillBeans = "Write how many grams of coffee beans you want to add:\n"
	fillCups  = "Write how many disposable cups you want to add:\n"

	fillBalance        = "Please put you money:\n"
	fillBalanceConfirm = "You filled you balance: %d$\n"
	fillBalanceShow    = "You balance is: %d$\n"

	errAction         = "Unknown user action: [%s]\n"
	errCoffeType      = "Wrong option\n"
	errNotEnough      = "Sorry, not enough%s!\n"
	errCheckMoney     = "Sorry can't check you money, returned it's you\n"
	errNotEnoughMoney = "Not enough money, please fill you balance\n"
)

const ( // init parameters
	waterDefault = 400
	milkDefault  = 540
	beansDefault = 120
	cupsDefault  = 9
	moneyDefault = 550
)

const ( // resources required for each type of coffee
	espressoWater = 250
	espressoMilk  = 0
	espressoBeans = 16
	espressoCost  = 4

	latteWater = 350
	latteMilk  = 75
	latteBeans = 20
	latteCost  = 7

	cappuccinoWater = 200
	cappuccinoMilk  = 100
	cappuccinoBeans = 12
	cappuccinoCost  = 6
)

func coffeeMachineConstructor() coffeeMachineState {
	return coffeeMachineState{water: waterDefault, milk: milkDefault, beans: beansDefault, cups: cupsDefault, money: moneyDefault}
}

func printState(e *coffeeMachineState) {
	fmt.Printf(statusMess)
	fmt.Printf(statusWater, e.water)
	fmt.Printf(statusMilk, e.milk)
	fmt.Printf(statusBeans, e.beans)
	fmt.Printf(statusCups, e.cups)
	fmt.Printf(statusMoney, e.money)
}

func buyCoffee(machine *coffeeMachineState, userBalance *int, water int, milk int, beans int, cost int) {
	var neededIngridiens strings.Builder
	if *userBalance < cost {
		fmt.Printf(errNotEnoughMoney)
	} else {
		if machine.water < water {
			neededIngridiens.WriteString(" water")
		}
		if machine.milk < milk {
			neededIngridiens.WriteString(" milk")
		}
		if machine.beans < beans {
			neededIngridiens.WriteString(" beans")
		}
		if machine.cups < 1 {
			neededIngridiens.WriteString(" cups")
		}
		if neededIngridiens.String() != "" {
			fmt.Printf(errNotEnough, neededIngridiens.String())
		} else {
			machine.water -= water
			machine.milk -= milk
			machine.beans -= beans
			machine.money += cost
			*userBalance -= cost
			machine.cups--
			fmt.Print(actionCooked)
		}
	}
}

func coffeeMenu(machine *coffeeMachineState, userBalance *int) {
	var option string
	fmt.Print(actionBuyCoffee)
	fmt.Scanln(&option)

	switch option {
	case "1":
		buyCoffee(machine, userBalance, espressoWater, espressoMilk, espressoBeans, espressoCost)
	case "2":
		buyCoffee(machine, userBalance, latteWater, latteMilk, latteBeans, latteCost)
	case "3":
		buyCoffee(machine, userBalance, cappuccinoWater, cappuccinoMilk, cappuccinoBeans, cappuccinoCost)
	case "4":
		mainMenu(machine, userBalance)

	default:
		fmt.Print(errCoffeType)

	}
}

func fill(machine *coffeeMachineState) {
	var water, milk, beans, cups int

	fmt.Print(fillWater)
	fmt.Scanln(&water)
	fmt.Print(fillMilk)
	fmt.Scanln(&milk)
	fmt.Print(fillBeans)
	fmt.Scanln(&beans)
	fmt.Print(fillCups)
	fmt.Scanln(&cups)
	(*machine).water += water
	(*machine).milk += milk
	(*machine).beans += beans
	(*machine).cups += cups

}

func take(machine *coffeeMachineState) {
	fmt.Printf(actionTake, machine.money)
	machine.money = 0

}

func checkThatMoneyIsReal() bool {
	return true
}

func funcFillBalance(userBalance *int) {
	var userInput int
	fmt.Print(fillBalance)
	fmt.Scanln(&userInput)
	if checkThatMoneyIsReal() {
		*userBalance += userInput
		fmt.Printf(fillBalanceConfirm, *userBalance)
	} else {
		fmt.Print(errCheckMoney)
	}
}

func showBalance(userBalance *int) {
	fmt.Printf(fillBalanceShow, *userBalance)
}

func takeRest(userBalance *int) {
	fmt.Printf(actionTakeRest, *userBalance)
	*userBalance = 0
}

func chooseAction(e string, machine *coffeeMachineState, userBalance *int) {
	switch e {
	case "1": //fill balance
		funcFillBalance(userBalance)
	case "2": //
		showBalance(userBalance)
	case "3":
		takeRest(userBalance)
	case "4": // buy
		coffeeMenu(machine, userBalance)
	case "5": // fill
		fill(machine)
	case "6": // take
		take(machine)
	case "7": // remaining
		printState(machine)
	case "8": // exit
		os.Exit(0)
	default:
		fmt.Printf(errAction, e)

	}

}

func mainMenu(machine *coffeeMachineState, balance *int) {
	var userInput string
	for {
		fmt.Print(actionRequest)
		fmt.Scanln(&userInput)
		chooseAction(userInput, machine, balance)
		userInput = ""
	}
}

func main() {
	newCoffeeMachine := coffeeMachineConstructor()
	mainMenu(&newCoffeeMachine, &userBalance)
}
