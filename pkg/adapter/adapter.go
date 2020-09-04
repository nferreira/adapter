package adapter

import (
	"github.com/nferreira/app/pkg/service"
)

type Adapter interface {
	BindRules(map[BindingRule]service.BusinessService)
	service.Service
}
