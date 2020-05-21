package common

import (
	"time"
)

const (
	KingfisherPath = "/kingfisher/api/v1.10/"
	K8SPath        = "/kubernetes/api/v1.10/"
	Istio          = "/istio/api/v1.10/"
	InspectPath    = "/inspect/api/v1.10/"
	PresetPath     = "/preset/api/v1.10/"
	REGISTRYURL    = "registry.kingfisher.com/kingfisher/"

	KubeConfigPath   = "/etc/"
	Signing          = "kingfisher"
	ExpiresAt        = 30 * 24 * time.Hour
	KingfisherURL    = "http://192.168.10.11:30090/kubernetes/api/v1.10/token"
	KingfisherOut    = "http://192.168.10.11:30090/login"
	Kingfisher       = "http://192.168.10.11:30090/home"
	KingfisherDomain = ""
	HeaderSigning    = "X-Signing"

	// k8s各种服务资源
	Namespace          = "namespace"
	Node               = "node"
	Service            = "service"
	Pod                = "pod"
	Image              = "image"
	Deployment         = "deployment"
	StatefulSet        = "stateful_set"
	DaemonSet          = "daemon_set"
	Secret             = "secret"
	ConfigMap          = "config_map"
	Ingress            = "ingress"
	Role               = "Role"
	ClusterRole        = "cluster_role"
	RoleBinding        = "role_binding"
	ClusterRoleBinding = "cluster_role_binding"
	HPA                = "hpa"
	Cluster            = "cluster"
	PVC                = "pvc"
	PV                 = "pv"
	SC                 = "storage_classes"
	ServiceAccount     = "service_account"
	StorageClasses     = "storage_classes"
	ReplicaSet         = "replica_set"
	Product            = "product"
	LimitRange         = "limit_range"
	ResourceQuota      = "resource-quota"
	VirtualServices    = "virtual_services"
	Gateways           = "gateways"
	DestinationRules   = "destination_rules"
	EnvoyFilters       = "envoy_filters"
	ServiceEntries     = "service_entries"
	Sidecars           = "sidecars"
	Rules              = "rules"
	Instances          = "instances"
	Plugin             = "plugin"
	Inspect            = "inspect"
	PlatformRole       = "platform_role"
	ClusterPlugin      = "cluster_plugin"
	Template           = "template"
	Config             = "config"

	// 数据库表
	DataField          = "data"
	ProductTreeTable   = "product_tree"
	ProductTable       = "product"
	AuditLogTable      = "audit_log"
	ClusterTable       = "cluster"
	UserTable          = "user"
	NamespaceTable     = "namespace"
	PluginTable        = "plugin"
	InspectTable       = "inspect"
	InspectInfoTable   = "inspect_info"
	ClusterPluginTable = "cluster_plugin"
	PlatformRoleTable  = "platform_role"
	TemplateTable      = "template"
	ConfigTable        = "config"

	// 操作类型
	Create                   = ActionType("create")
	Delete                   = ActionType("delete")
	Start                    = ActionType("start")
	Patch                    = ActionType("patch")
	Scale                    = ActionType("scale")
	PatchImage               = ActionType("patch_image")
	PatchSyncImage           = ActionType("patch_sync_image")
	WatchPodIP               = ActionType("watch_pod_ip")
	PatchStepResume          = ActionType("patch_step_resume")
	PatchAllResume           = ActionType("patch_all_resume")
	PatchPause               = ActionType("patch_pause")
	Update                   = ActionType("update")
	SaveTemplate             = ActionType("save_template")
	Stop                     = ActionType("stop")
	Reboot                   = ActionType("reboot")
	Evict                    = ActionType("evict")
	Get                      = ActionType("get")
	List                     = ActionType("list")
	ListAll                  = ActionType("list_all")
	Log                      = ActionType("log")
	ListPodByNode            = ActionType("list_pod_by_node")
	ListPodByService         = ActionType("list_pod_by_service")
	ListPodByController      = ActionType("list_pod_by_controller")
	NodeMetric               = ActionType("node_metric")
	GetChart                 = ActionType("get_chart")
	GetNamespaceIsExistLabel = ActionType("get_namespace_is_exist_label")
	Debug                    = ActionType("debug")
	Rescue                   = ActionType("rescue")
	Kubectl                  = ActionType("Kubectl")
	UnKubectl                = ActionType("unkubectl")
	Status                   = ActionType("status")
	Form                     = ActionType("form")
	Yaml                     = ActionType("yaml")
	SaveAsTemplate           = ActionType("save as template")
	DebugPodIPByPod          = ActionType("debug podIP by pod")

	// 错误类型
	ContentTypeError         = "unknown content type, please specify the type json or yaml"
	K8SClientSetError        = "k8s clientSet Initialization failure: "
	IstioClientSetError      = "isito clientSet Initialization failure: "
	PrometheusClientSetError = "prometheus clientSet Initialization failure: "
	ClusterNotExistError     = "sql: no rows in result set"

	// kubectl相关常量
	ServiceAccountName     = "kingfisher"
	ClusterRoleBindingName = "kingfisher-cluster-role-binding"
	KubectlPodName         = "king-kubectl-pod"
	KubectlDeploymentName  = "king-kubectl"

	// grpc server port
	GRPCPort         = ":50000"
	KingK8S          = "king-k8s:50000"
	KingIstio        = "king-istio:50000"
	KingKf           = "king-kf:50000"
	KubectlNamespace = "default"

	// MQ Topic
	UpdateKubeConfig = "Update-Kube-Config"

	ConfigID = "c_000000000001"
)

type ResponseData struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

type PatchData struct {
	Op    string      `json:"op"`
	Path  string      `json:"path"`
	Value interface{} `json:"value"`
}

type PatchJson struct {
	Patches []PatchData `json:"patches"`
}

type PostYaml struct {
	Context string `json:"context"`
}

type Projects struct {
	ProjectId         int         `json:"project_id"`
	OwnerId           int         `json:"owner_id"`
	Name              string      `json:"name"`
	CreationTime      interface{} `json:"creation_time"`
	UpdateTime        interface{} `json:"update_time"`
	Deleted           int         `json:"deleted"`
	OwnerName         string      `json:"owner_name"`
	Togglable         bool        `json:"togglable"`
	CurrentUserRoleId int         `json:"current_user_role_id"`
	RepoCount         int         `json:"repo_count"`
	Metadata          Metadata
}

type Metadata struct {
	AutoScan           string `json:"auto_scan"`
	EnableContentTrust string `json:"enable_content_trust"`
	PreventVul         string `json:"prevent_vul"`
	Public             string `json:"public"`
	Severity           string `json:"severity"`
}

type ImageList struct {
	Id           int         `json:"id"`
	Name         string      `json:"name"`
	ProjectId    int         `json:"project_id"`
	Description  string      `json:"description"`
	PullCount    int         `json:"pull_count"`
	StarCount    int         `json:"star_count"`
	TagsCount    int         `json:"tags_count"`
	CreationTime interface{} `json:"creation_time"`
	UpdateTime   interface{} `json:"update_time"`
}

type ImageTags struct {
	Digest        string      `json:"digest"`
	Name          string      `json:"name"`
	Size          int         `json:"size"`
	Architecture  string      `json:"architecture"`
	Os            string      `json:"os"`
	DockerVersion string      `json:"docker_version"`
	Author        string      `json:"author"`
	Created       interface{} `json:"created"`
	Config        interface{} `json:"config"`
}

type ProductTree struct {
	Id            string `json:"id"`
	Uuid          string `json:"uuid"`
	Name          string `json:"name"`
	Lev           string `json:"lev"`
	TechLeader    string `json:"tech_leader"`
	ProductLeader string `json:"product_leader"`
	VicePresident string `json:"vice_president"`
	Status        string `json:"status"`
	IsPlatform    string `json:"is_platform"`
}

type ActionType string

type AuditLog struct {
	Type       string      `json:"type"`
	Id         string      `json:"id"`
	Name       string      `json:"name"`
	Result     bool        `json:"result"`
	ProductId  string      `json:"product_id"`
	Cluster    string      `json:"cluster"`
	ActionType ActionType  `json:"action_type"`
	PostType   ActionType  `json:"post_type"`
	ActionTime int64       `json:"action_time"`
	User       string      `json:"user"`
	Msg        string      `json:"msg"`
	Json       interface{} `json:"json"`
	Namespace  string      `json:"namespace"`
}

type ClusterDB struct {
	Id         string   `json:"id"`
	Name       string   `json:"name" binding:"required,requiredValidate"`
	Describe   string   `json:"describe"`
	Token      string   `json:"token" binding:"required,requiredValidate"`
	CAHash     string   `json:"ca_hash" binding:"required,requiredValidate"`
	KubConfig  string   `json:"kub_config"`
	Product    []string `json:"product"`
	Version    string   `json:"version"`
	CreateTime int64    `json:"createTime"`
	ModifyTime int64    `json:"modifyTime"`
}

type User struct {
	Id         string   `json:"id"`
	Name       string   `json:"name" binding:"required,requiredValidate"`
	RealName   string   `json:"realName"`
	Mail       string   `json:"mail" binding:"required,requiredValidate"`
	Mobile     string   `json:"mobile"`
	Product    []string `json:"product"`
	Cluster    []string `json:"cluster"`
	Namespace  []string `json:"namespace"`
	Role       string   `json:"role"`
	CreateTime int64    `json:"createTime"`
	ModifyTime int64    `json:"modifyTime"`
}

type NamespaceDB struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Cluster    string `json:"cluster"`
	Product    string `json:"product"`
	CreateTime int64  `json:"createTime"`
	ModifyTime int64  `json:"modifyTime"`
}

type Info struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type InfoType struct {
	Value  string `json:"value"`
	Label  string `json:"label"`
	Level  int    `json:"level"`
	Title  string `json:"title"`
	Expand bool   `json:"expand"`
}

type ProductDB struct {
	Id         string   `json:"id" binding:"required,requiredValidate"`
	Cluster    []string `json:"cluster"`
	Namespace  []string `json:"namespace"`
	CreateTime int64    `json:"createTime"`
	ModifyTime int64    `json:"modifyTime"`
}

type PostType struct {
	Context string `json:"context"`
}

type ClusterPluginDB struct {
	Id        string `json:"id"`
	Plugin    string `json:"plugin"`
	Cluster   string `json:"cluster"`
	Status    int    `json:"status"`
	Timestamp int64  `json:"timestamp"`
}

type PluginDB struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Describe   string `json:"describe"`
	CreateTime int64  `json:"createTime"`
	ModifyTime int64  `json:"modifyTime"`
}

type PlatformRoleDB struct {
	Id         string   `json:"id"`
	Name       string   `json:"name"`
	Describe   string   `json:"describe"`
	Access     []string `json:"access"`
	CreateTime int64    `json:"createTime"`
	ModifyTime int64    `json:"modifyTime"`
}

type InspectDB struct {
	Id         string   `json:"id"`
	Name       string   `json:"name"`
	Describe   string   `json:"describe"`
	Cluster    string   `json:"cluster"`
	Level      []string `json:"level"`
	Namespace  []string `json:"namespace"`
	Basic      []string `json:"basic"`
	Unused     []string `json:"unused"`
	State      []string `json:"state"`
	Security   []string `json:"security"`
	User       string   `json:"user"`
	Interval   string   `json:"interval"`
	Report     string   `json:"report"`
	CreateTime int64    `json:"createTime"`
	ModifyTime int64    `json:"modifyTime"`
}

type InspectInfoDB struct {
	Id         string      `json:"id"`
	Info       interface{} `json:"info"`
	CreateTime int64       `json:"createTime"`
}

type TemplateDB struct {
	Id         string      `json:"id"`
	Name       string      `json:"name"`
	Describe   string      `json:"describe"`
	Kind       string      `json:"kind"`
	Spec       interface{} `json:"spec"`
	CreateTime int64       `json:"createTime"`
	ModifyTime int64       `json:"modifyTime"`
}

type LDAPDB struct {
	Mode           string `json:"mode"`
	URL            string `json:"url"`
	SearchDN       string `json:"searchDN"`
	SearchPassword string `json:"searchPassword"`
	BaseDN         string `json:"baseDN"`
	UserFilter     string `json:"userFilter"`
	TLS            bool   `json:"tls"`
	Attributes     string `json:"attributes"`
}

type ConfigDB struct {
	Id         string `json:"id"`
	LDAPDB     LDAPDB `json:"ldap"`
	CreateTime int64  `json:"createTime"`
	ModifyTime int64  `json:"modifyTime"`
}
