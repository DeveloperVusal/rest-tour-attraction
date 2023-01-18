package combines

import (
	"attrtour/app/models"
)

type ModelCombine struct {
	Items []interface{}
}

func (mc *ModelCombine) Build() {
	mc.Items = append(mc.Items, &models.Group{})
	mc.Items = append(mc.Items, &models.User{})
	mc.Items = append(mc.Items, &models.Auth{})
	mc.Items = append(mc.Items, &models.UserBanned{})
	mc.Items = append(mc.Items, &models.Language{})
	mc.Items = append(mc.Items, &models.Continent{})
	mc.Items = append(mc.Items, &models.Country{})
	mc.Items = append(mc.Items, &models.Location{})
	mc.Items = append(mc.Items, &models.MainImage{})
}
