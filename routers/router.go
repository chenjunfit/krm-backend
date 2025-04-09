package routers

import (
	"github.com/gin-gonic/gin"
	"krm-backend/routers/auth"
	"krm-backend/routers/cluster"
	"krm-backend/routers/configmap"
	"krm-backend/routers/cronjob"
	"krm-backend/routers/daemonset"
	"krm-backend/routers/deployment"
	"krm-backend/routers/ingress"
	"krm-backend/routers/ingressclass"
	"krm-backend/routers/namespace"
	"krm-backend/routers/node"
	"krm-backend/routers/pod"
	service "krm-backend/routers/pv"
	"krm-backend/routers/pvc"
	"krm-backend/routers/replicaset"
	"krm-backend/routers/secret"
	pv "krm-backend/routers/service"
	"krm-backend/routers/statefulset"
	"krm-backend/routers/storageclass"
	"krm-backend/routers/tools"
)

func RegisterRouters(r *gin.Engine) {
	apiGroup := r.Group("/api")
	auth.RegisterSubRouter(apiGroup)
	cluster.RegisterSubRouter(apiGroup)
	namespace.RegisterSubRouter(apiGroup)
	pod.RegisterSubRouter(apiGroup)
	deployment.RegisterSubRouter(apiGroup)
	statefulset.RegisterSubRouter(apiGroup)
	daemonset.RegisterSubRouter(apiGroup)
	cronjob.RegisterSubRouter(apiGroup)
	replicaset.RegisterSubRouter(apiGroup)
	node.RegisterSubRouter(apiGroup)
	service.RegisterSubRouter(apiGroup)
	secret.RegisterSubRouter(apiGroup)
	ingress.RegisterSubRouter(apiGroup)
	ingressclass.RegisterSubRouter(apiGroup)
	configmap.RegisterSubRouter(apiGroup)
	pv.RegisterSubRouter(apiGroup)
	pvc.RegisterSubRouter(apiGroup)
	storageclass.RegisterSubRouter(apiGroup)
	tools.RegisterSubRouter(apiGroup)
}
