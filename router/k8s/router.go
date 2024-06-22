package k8s

import (
	"KubeXCloud/api"
	"github.com/gin-gonic/gin"
)

type K8sRouter struct {
}

func (*K8sRouter) InitK8SRouter(r *gin.Engine) {
	group := r.Group("/k8s")
	apiGroup := api.ApiGroupApp.K8SApiGroup
	group.GET("/listPod", apiGroup.GetPodList)
	group.POST("/pod", apiGroup.CreateOrUpdatePod)
	group.GET("/pod/:namespace", apiGroup.GetPodListOrDetail)
	group.DELETE("/pod/:namespace/:name", apiGroup.DeletePod)
	group.GET("/namespace", apiGroup.GetNamespaceList)
}
