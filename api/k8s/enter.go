package k8s

import (
	"KubeXCloud/service"
	"KubeXCloud/validate"
)

type ApiGroup struct {
	PodApi
	NamespaceApi
}

var podValidate = validate.ValidateGroupApp.PodValidate
var podService = service.ServiceGroupApp.PodServiceGroup.PodService
