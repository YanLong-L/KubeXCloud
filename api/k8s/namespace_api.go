package k8s

import (
	"KubeXCloud/global"
	namespace_res "KubeXCloud/model/namespace/response"
	"KubeXCloud/response"
	"context"
	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type NamespaceApi struct {
}

// GetNamespaceList 获取k8s集群的namespace 列表
func (*NamespaceApi) GetNamespaceList(c *gin.Context) {
	ctx := context.Background()
	list, err := global.KubeConfigSet.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		response.FailWithMessage(c, err.Error())
	}
	namespaceList := make([]namespace_res.Namespace, 0)
	for _, item := range list.Items {
		namespaceList = append(namespaceList, namespace_res.Namespace{
			Name:              item.Name,
			CreationTimestamp: item.CreationTimestamp.Unix(),
			Status:            string(item.Status.String()),
		})
	}
	response.SuccessWithDetailed(c, "获取成功", namespaceList)
}
