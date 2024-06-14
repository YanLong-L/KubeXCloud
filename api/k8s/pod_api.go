package k8s

import (
	"KubeXCloud/global"
	pod_req "KubeXCloud/model/pod/request"
	"KubeXCloud/response"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
)

type PodApi struct {
}

// GetPodList 测试用
func (*PodApi) GetPodList(c *gin.Context) {
	ctx := context.TODO()
	list, err := global.KubeConfigSet.CoreV1().Pods("").List(ctx, metav1.ListOptions{})
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, item := range list.Items {
		fmt.Println(item.Namespace, item.Name)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// CreateOrUpdatePod 创建或更新pod
func (*PodApi) CreateOrUpdatePod(c *gin.Context) {
	var podReq pod_req.Pod
	if err := c.ShouldBind(&podReq); err != nil {
		response.FailWithMessage(c, "参数解析失败，detail："+err.Error())
		return
	}
	// 校验必填项
	if err := podValidate.Validate(&podReq); err != nil {
		response.FailWithMessage(c, "参数验证失败，detail："+err.Error())
	}
	if msg, err := podService.CreateOrUpdatePod(podReq); err != nil {
		response.FailWithMessage(c, msg)
	} else {
		response.SuccessWithMessage(c, msg)
	}
}
