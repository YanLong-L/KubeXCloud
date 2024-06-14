package pod

import (
	"KubeXCloud/global"
	pod_req "KubeXCloud/model/pod/request"
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PodService struct {
}

// CreateOrUpdatePod 创建或更新pod
func (*PodService) CreateOrUpdatePod(podReq pod_req.Pod) (msg string, err error) {
	k8sPod := podConvert.Req2K8sConvert.PodReq2K8s(podReq)
	// 定义context
	ctx := context.Background()
	// 获取操作pod的接口
	podApi := global.KubeConfigSet.CoreV1().Pods(podReq.Base.Namespace)
	// 先查一下当前pod是否已经存在
	k8sGetPod, err := podApi.Get(ctx, podReq.Base.Name, metav1.GetOptions{})
	if err == nil {
		// 更新pod

		return "", err
	} else {
		// 创建pod
		createdPod, err := podApi.Create(ctx, k8sPod, metav1.CreateOptions{})
		if err != nil {
			// pod创建失败
			errMsg := fmt.Sprintf("Pod[namespace=%s],")
			return errMsg, err
		} else {
			// pod创建成功
			successMsg := fmt.Sprintf("Pod[namespace=%s,name=%s]创建成功", createdPod.Namespace, createdPod.Name)
			return successMsg, err
		}

	}

}
