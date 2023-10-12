package actions

import (
	"fin/common"
	"fin/repositories"
	"fmt"
)

type ListCostCenterAction struct {
	Prompt     common.Prompt
	Repository repositories.CostCenterRepository
}

func (lcca ListCostCenterAction) GetActionName() string {
	return "list-cost-centers"
}

func (lcca ListCostCenterAction) Execute() {
	cc := lcca.Repository.List()
	for _, ccr := range cc {
		fmt.Println(ccr.Id, "->", ccr.Description)
	}
}
