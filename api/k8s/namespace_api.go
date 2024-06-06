package k8s

import (
	"KubeXCloud/response"
	"github.com/gin-gonic/gin"
)

type NameSpaceApi struct {
}

func (*NameSpaceApi) GetNameSpaceList(ctx *gin.Context) {
	response.Success(ctx)
}
