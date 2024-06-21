package pod

import (
	"KubeXCloud/global"
	pod_req "KubeXCloud/model/pod/request"
	pod_res "KubeXCloud/model/pod/response"
	"context"
	"errors"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	k8serror "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"strings"
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
		// 1. 先检查pod的参数是否合理，可以先dry run一下
		k8sPodCopy := *k8sGetPod
		k8sPodCopy.Name = k8sPod.Name + "-validate"
		_, err := podApi.Create(ctx, &k8sPodCopy, metav1.CreateOptions{
			DryRun: []string{metav1.DryRunAll},
		})
		if err != nil {
			errMsg := fmt.Sprintf("Pod[namespace=%s,name=%s]更新失败，detail：%s", k8sPod.Namespace, k8sPod.Name, err.Error())
			return errMsg, err
		}
		//比如pod处于terminating状态 监听pod删除完毕之后 才开始创建pod
		var labelSelector []string
		for k, v := range k8sGetPod.Labels {
			labelSelector = append(labelSelector, fmt.Sprintf("%s=%s", k, v))
		}
		//label 格式 app=test,app2=test2
		watcher, err := podApi.Watch(ctx, metav1.ListOptions{
			LabelSelector: strings.Join(labelSelector, ","),
		})
		if err != nil {
			errMsg := fmt.Sprintf("Pod[namespace=%s,name=%s]更新失败，detail：%s", k8sPod.Namespace, k8sPod.Name, err.Error())
			return errMsg, err
		}
		// 删除旧pod
		background := metav1.DeletePropagationBackground
		var gracePeriodSeconds int64 = 0
		err = podApi.Delete(ctx, k8sPod.Name, metav1.DeleteOptions{
			GracePeriodSeconds: &gracePeriodSeconds,
			PropagationPolicy:  &background,
		})
		if err != nil {
			errMsg := fmt.Sprintf("Pod[namespace=%s,name=%s]更新失败，detail：%s", k8sPod.Namespace, k8sPod.Name, err.Error())
			return errMsg, err
		}
		for event := range watcher.ResultChan() {
			k8sPodChan := event.Object.(*corev1.Pod)
			// 在判断删除事件之前，先查询pod是否已经被删除
			_, err := podApi.Get(ctx, k8sPod.Name, metav1.GetOptions{})
			if k8serror.IsNotFound(err) {
				// pod 已经删除，重新创建
				createdPod, err := podApi.Create(ctx, k8sPod, metav1.CreateOptions{})
				if err != nil {
					errMsg := fmt.Sprintf("Pod[namespace=%s,name=%s]更新失败，detail：%s", k8sPod.Namespace, k8sPod.Name, err.Error())
					return errMsg, err
				} else {
					successMsg := fmt.Sprintf("Pod[namespace=%s,name=%s]更新成功", createdPod.Namespace, createdPod.Name)
					return successMsg, err
				}
			}
			switch event.Type {
			case watch.Deleted:
				// 如果侦听到的pod不是 待删除的pod continue掉
				if k8sPodChan.Name != k8sPod.Name {
					continue
				}
				// pod 已经删除，重新创建
				//重新创建
				if createdPod, err := podApi.Create(ctx, k8sPod, metav1.CreateOptions{}); err != nil {
					errMsg := fmt.Sprintf("Pod[namespace=%s,name=%s]更新失败，detail：%s", k8sPod.Namespace, k8sPod.Name, err.Error())
					return errMsg, err
				} else {
					successMsg := fmt.Sprintf("Pod[namespace=%s,name=%s]更新成功", createdPod.Namespace, createdPod.Name)
					return successMsg, err
				}
			}
		}

		return "", err
	} else {
		// 创建pod
		createdPod, err := podApi.Create(ctx, k8sPod, metav1.CreateOptions{})
		if err != nil {
			// pod创建失败
			errMsg := fmt.Sprintf("Pod[namespace=%s,name=%s]创建失败，detail：%s", k8sPod.Namespace, k8sPod.Name, err.Error())
			return errMsg, err
		} else {
			// pod创建成功
			successMsg := fmt.Sprintf("Pod[namespace=%s,name=%s]创建成功", createdPod.Namespace, createdPod.Name)
			return successMsg, err
		}

	}

}

// GetPodList 获取pod列表
func (*PodService) GetPodList(namespace string, keyword string) (error, []pod_res.PodListItem) {
	ctx := context.Background()
	list, err := global.KubeConfigSet.CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return err, nil
	}
	podList := make([]pod_res.PodListItem, 0)
	for _, item := range list.Items {
		if strings.Contains(item.Name, keyword) {
			podItem := podConvert.PodK8s2ItemRes(item)
			podList = append(podList, podItem)
		}
	}
	return err, podList
}

// GetPodDetail 获取pod详情
func (*PodService) GetPodDetail(namespace string, name string) (podReq pod_req.Pod, err error) {
	ctx := context.TODO()
	podApi := global.KubeConfigSet.CoreV1().Pods(namespace)
	k8sGetPod, err := podApi.Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		errMsg := fmt.Sprintf("Pod[namespace=%s,name=%s]查询失败，detail：%s", namespace, name, err.Error())
		err = errors.New(errMsg)
		return
	}
	//将k8s pod 转为 pod request
	podReq = podConvert.K8s2ReqConvert.PodK8s2Req(*k8sGetPod)
	return
}

// DeletePod 删除指定pod
func (*PodService) DeletePod(namespace string, name string) error {
	background := metav1.DeletePropagationBackground
	var gracePeriodSeconds int64 = 0
	return global.KubeConfigSet.CoreV1().Pods(namespace).Delete(context.TODO(), name,
		metav1.DeleteOptions{
			GracePeriodSeconds: &gracePeriodSeconds,
			PropagationPolicy:  &background,
		})
}
