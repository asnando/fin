package actions

import (
	"fin/common"
	"fin/repositories"
)

type CreateCostCenterAction struct {
	Prompt     common.Prompt
	Repository repositories.CostCenterRepository
}

func (ccca CreateCostCenterAction) GetActionName() string {
	return "create-cost-center"
}

func (cca CreateCostCenterAction) Execute() {
	cc := cca.Prompt.ReadInputWithMessage("Enter cost center description:")
	cca.Repository.Save(cc)
}
