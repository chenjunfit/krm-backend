package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"k8s.io/client-go/kubernetes"
)

const (
	TimeFormat                    string = "2006-01-02 15:04:02"
	ClusterConfigSecretLabelKey   string = "kubeasy.com/cluster.matadata"
	ClusterConfigSecretLabelValue string = "true"
)

var (
	Port          string
	JwtSignKey    string
	JwtExpireTime int64
	AdminUserName string
	AdminPassword string

	MetaNamespace     string
	ClientSet         *kubernetes.Clientset
	ClusterKubeConfig map[string]string
	ProtectNameSpace  []string

	GinMode   string
	InCluster string
)

type ReturnData struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

func NewReturnData() ReturnData {
	return ReturnData{
		Status: 200,
		Data:   make(map[string]interface{}),
	}
}

func init() {
	viper.AutomaticEnv()
	viper.SetDefault("LOG_LEVEL", "debug")
	viper.SetDefault("PORT", ":8080")
	viper.SetDefault("JWT_SIGN_KEY", "tiantianmoyu")
	viper.SetDefault("JWT_EXPIRE_TIME", 1200000)
	viper.SetDefault("GIN_MODE", "debug")
	//admin,admin
	viper.SetDefault("ADMIN_USER_NAME", "admin")
	viper.SetDefault("ADMIN_PASSWORD", "admin")
	viper.SetDefault("META_NAMESPACE", "meta-namespace")
	viper.SetDefault("ProtectNameSpace", []string{"kube-system", "kube-flannel"})
	viper.SetDefault("IN_CLUSTER", "false")
	logLevel := viper.GetString("LOG_LEVEL")
	Port = viper.GetString("PORT")
	JwtSignKey = viper.GetString("JWT_SIGN_KEY")
	JwtExpireTime = viper.GetInt64("JWT_EXPIRE_TIME")
	AdminUserName = viper.GetString("ADMIN_USER_NAME")
	AdminPassword = viper.GetString("ADMIN_PASSWORD")
	MetaNamespace = viper.GetString("META_NAMESPACE")
	ProtectNameSpace = viper.GetStringSlice("ProtectNameSpace")
	GinMode = viper.GetString("GIN_MODE")
	InCluster = viper.GetString("IN_CLUSTER")
	initLogConfig(logLevel)
}
func initLogConfig(level string) {
	if level == "debug" {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.DebugLevel)
	}
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.JSONFormatter{TimestampFormat: TimeFormat})
}
