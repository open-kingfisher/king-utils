package access

import (
	"github.com/open-kingfisher/king-utils/common"
	"github.com/open-kingfisher/king-utils/common/log"
	"github.com/open-kingfisher/king-utils/db"
	"github.com/open-kingfisher/king-utils/kit"
	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"k8s.io/client-go/discovery"
	cacheddiscovery "k8s.io/client-go/discovery/cached"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/restmapper"
	"k8s.io/client-go/tools/clientcmd"
	metrics "k8s.io/metrics/pkg/client/clientset/versioned"
	customMetrics "k8s.io/metrics/pkg/client/custom_metrics"
	"time"
)

func Access(clusterId string) (*kubernetes.Clientset, error) {
	config, err := GetConfig(clusterId)
	if err != nil {
		return nil, err
	}
	// 创建 clientSet
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Errorf("New clientSet error: %s", err)
		return nil, err
	}
	return clientSet, err
}

func MetricsClient(clusterId string) (*metrics.Clientset, error) {
	config, err := GetConfig(clusterId)
	if err != nil {
		return nil, err
	}
	// 创建 clientSet
	clientSet, err := metrics.NewForConfig(config)
	if err != nil {
		log.Errorf("New clientSet error: %s", err)
		return nil, err
	}
	return clientSet, err
}

func DynamicClient(clusterId string) (dynamic.Interface, error) {
	config, err := GetConfig(clusterId)
	if err != nil {
		return nil, err
	}
	// 创建 clientSet
	clientSet, err := dynamic.NewForConfig(config)
	if err != nil {
		log.Errorf("New clientSet error: %s", err)
		return nil, err
	}
	return clientSet, err
}

func CustomMetricsClient(clusterId string) (customMetrics.CustomMetricsClient, error) {
	config, err := GetConfig(clusterId)
	if err != nil {
		return nil, err
	}
	// 创建 Custom Metrics
	// https://github.com/kubernetes/kubernetes/blob/master/test/e2e/instrumentation/monitoring/custom_metrics_stackdriver.go
	discoveryClient := discovery.NewDiscoveryClientForConfigOrDie(config)
	cachedDiscoClient := cacheddiscovery.NewMemCacheClient(discoveryClient)
	restMapper := restmapper.NewDeferredDiscoveryRESTMapper(cachedDiscoClient)
	restMapper.Reset()
	availableAPIsGetter := customMetrics.NewAvailableAPIsGetter(discoveryClient)
	clientSet := customMetrics.NewForConfig(config, restMapper, availableAPIsGetter)
	return clientSet, err
}

func InFormer(clusterId string) (informers.SharedInformerFactory, error) {
	config, err := GetConfig(clusterId)
	if err != nil {
		return nil, err
	}
	// 创建 clientSet
	clientSet, err := kubernetes.NewForConfig(config)
	informerSet := informers.NewSharedInformerFactory(clientSet, 10*time.Second)
	if err != nil {
		log.Errorf("New clientSet error: %s", err)
		return nil, err
	}
	return informerSet, nil
}

func PrometheusClient() (v1.API, error) {
	// https://godoc.org/github.com/prometheus/client_golang/api/prometheus/v1#example-API--Query
	config := api.Config{
		Address: "http://prometheus.kingfisher.com:9090", // 线上prometheus地址
	}
	client, err := api.NewClient(config)
	if err != nil {
		log.Errorf("New prometheus client error: %s", err)
		return nil, err
	}
	return v1.NewAPI(client), nil
}

func GetConfig(clusterId string) (*restclient.Config, error) {
	cluster := common.ClusterDB{}
	if err := db.GetById(common.Cluster, clusterId, &cluster); err != nil {
		return nil, err
	}
	kubeconfig := common.KubeConfigPath + cluster.Id
	// 判断配置文件是否存在，存在不在创建
	if !kit.IsExist(kubeconfig) {
		// 从数据库中读取kubeconfig内容创建kubeconfig配置文件
		if err := kit.CreateConfig(cluster, kubeconfig); err != nil {
			log.Errorf("Create kubeconfig error: %s", err)
			return nil, err
		}
	}
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Errorf("Build kubeconfig error: %s", err)
		return nil, err
	}
	return config, nil
}
