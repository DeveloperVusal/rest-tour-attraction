package combines

import (
	"attrtour/app/services"
)

type ServiceCombine struct {
	Items []interface{}
}

func (sc *ServiceCombine) Build() {
	sc.Items = append(sc.Items, &services.UserService{})
}
