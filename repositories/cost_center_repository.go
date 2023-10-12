package repositories

import (
	"fin/common"
	"fin/models"
)

type CostCenterRepository struct {
	Database common.Database
}

func (ccr CostCenterRepository) Save(description string) {
	ccr.Database.Insert("cost_centers", "description", description)
}

func (ccr CostCenterRepository) List() []models.CostCenter {
	rows := ccr.Database.Query("cost_centers")
	cc := models.CostCenter{}.ParseCostCenters(rows)
	return cc
}
