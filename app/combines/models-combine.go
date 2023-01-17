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
}
