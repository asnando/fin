package main

import (
	"fin/actions"
	"fin/common"
	"fmt"
	"os"
	"strconv"
)

type App struct {
	Prompt  common.Prompt
	Actions []actions.Action
}

func (a App) Start() {
	a.DisplayMenu()
	a.ListenForAction()
}

func (a App) DisplayMenu() {
	fmt.Println(`
CREATE
1. Cost center
2. Payment Method
5. Transaction

LIST
3. Cost center
4. Payment Method
6. Last transaction
7. Current month transactions

CSV
8. Import from csv
9. Export to csv

Pick an option:`)
}

func (a App) ListenForAction() {
	option := a.Prompt.ReadInput()
	opt, err := strconv.Atoi(option)

	if err != nil {
		fmt.Println("Incorrect input:", option)
		os.Exit(1)
	}

	switch opt {
	case 1:
		a.triggerAction("create-cost-center")
	case 2:
		a.triggerAction("create-payment-method")
	case 3:
		a.triggerAction("list-cost-centers")
	case 4:
		a.triggerAction("list-payment-methods")
	case 5:
		a.triggerAction("create-transaction")
	case 6:
		a.triggerAction("get-last-transaction")
	case 7:
		a.triggerAction("get-month-transactions")
	case 8:
		a.triggerAction("import-csv")
	case 9:
		a.triggerAction("export-csv")
	default:
		fmt.Println("Unrecognized option:", opt)
		os.Exit(1)
	}
}

func (a App) triggerAction(actionName string) {
	action := a.filterAction(actionName)
	if action == nil {
		fmt.Println("Unregistered action", actionName)
		os.Exit(1)
	}
	action.Execute()
}

func (a App) filterAction(actionName string) actions.Action {
	for _, action := range a.Actions {
		if action.GetActionName() == actionName {
			return action
		}
	}
	return nil
}
