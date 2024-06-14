package convert

import "KubeXCloud/convert/pod"

type ConvertGroup struct {
	PodConvert pod.PodConvertGroup
}

var ConvertGroupApp = new(ConvertGroup)
