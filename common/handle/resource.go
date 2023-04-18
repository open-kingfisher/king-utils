package handle

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/duiniwukenaihe/king-utils/common"
	"github.com/duiniwukenaihe/king-utils/db"
	"github.com/duiniwukenaihe/king-utils/kit"
	"github.com/duiniwukenaihe/king-utils/middleware/jwt"
	"k8s.io/client-go/kubernetes"
	"net/http"
	"time"
)

type Resources struct {
	Namespace   string
	Cluster     string
	Product     string
	Name        string
	User        *jwt.CustomClaims
	ClientSet   *kubernetes.Clientset
	PatchData   *common.PatchJson
	DataType    string
	Uid         string
	Kind        string
	KindName    string
	MetricsName string
	LabelKey    string
	MetricsKind string
	Controller  string
	Scale       string
	Step        string
	PostType    string
	Signing     string
	Time        string
}

// 获取HTTP请求相关通用参数
func GenerateCommonParams(c *gin.Context, clientSet *kubernetes.Clientset) (commonParams *Resources) {
	// 获取用户信息
	user := c.MustGet("user").(*jwt.CustomClaims)
	commonParams = &Resources{
		Namespace:   c.Query("namespace"),
		Cluster:     c.Query("cluster"),
		Product:     c.Query("productId"),
		Name:        c.Param("name"),
		DataType:    c.Query("type"),
		User:        user,
		ClientSet:   clientSet,
		Uid:         c.Query("uid"),
		Kind:        c.Query("kind"),
		KindName:    c.Query("kindName"),
		MetricsName: c.Param("metricsName"),
		LabelKey:    c.Query("labelKey"),
		MetricsKind: c.Param("metricsKind"),
		Controller:  c.Param("controller"),
		Scale:       c.Query("scale"),
		Step:        c.Query("step"),
		PostType:    c.Query("postType"),
		Signing:     c.GetHeader(common.HeaderSigning),
		Time:        c.Query("time"),
	}
	return
}

type AuditLog struct {
	Kind        string
	ActionType  common.ActionType
	PostType    common.ActionType
	Resources   *Resources
	PostData    interface{}
	Name        string
	ClusterData *common.ClusterDB
	ProductData *common.ProductDB
}

func (a *AuditLog) InsertAuditLog() (err error) {
	var jsonData interface{}
	if a.ActionType == common.Delete || a.ActionType == common.Patch {
		a.Name = a.Resources.Name
		jsonData = a.Resources.PatchData
	} else {
		jsonData = a.PostData
	}
	auditLog := common.AuditLog{}
	if a.Kind == common.Cluster && a.ActionType != common.Delete {
		auditLog = common.AuditLog{
			Type:       a.Kind,
			Id:         a.ClusterData.Id,
			Name:       a.ClusterData.Name,
			User:       a.Resources.User.Name,
			ProductId:  a.Resources.Product,
			Cluster:    "",
			Json:       a.ClusterData,
			ActionTime: time.Now().Unix(),
			ActionType: a.ActionType,
			PostType:   a.PostType,
			Namespace:  a.Resources.Namespace,
			Result:     true,
			Msg:        "",
		}
	} else if a.Kind == common.Product && a.ActionType != common.Delete {
		auditLog = common.AuditLog{
			Type:       a.Kind,
			Id:         a.ProductData.Id,
			Name:       "",
			User:       a.Resources.User.Name,
			ProductId:  a.Resources.Product,
			Cluster:    "",
			Json:       a.ProductData,
			ActionTime: time.Now().Unix(),
			ActionType: a.ActionType,
			PostType:   a.PostType,
			Namespace:  a.Resources.Namespace,
			Result:     true,
			Msg:        "",
		}
	} else {
		auditLog = common.AuditLog{
			Type:       a.Kind,
			Id:         "",
			Name:       a.Name,
			User:       a.Resources.User.Name,
			ProductId:  a.Resources.Product,
			Cluster:    a.Resources.Cluster,
			Json:       jsonData,
			ActionTime: time.Now().Unix(),
			ActionType: a.ActionType,
			PostType:   a.PostType,
			Namespace:  a.Resources.Namespace,
			Result:     true,
			Msg:        "",
		}
	}
	return db.Insert(common.AuditLogTable, auditLog)
}

func CreateTemplate(template *common.TemplateDB) error {
	templateList := make([]*common.TemplateDB, 0)
	if err := db.List(common.DataField, common.TemplateTable, &templateList, "WHERE data-> '$.name'=? and data-> '$.kind'=?", template.Name, template.Kind); err == nil {
		if len(templateList) > 0 {
			return errors.New("the template name already exists")
		}
	} else {
		return err
	}
	template.Id = kit.UUID("t")
	template.CreateTime = time.Now().Unix()
	template.ModifyTime = template.CreateTime
	if err := db.Insert(common.TemplateTable, template); err != nil {
		return err
	}
	return nil
}

func HandlerResponse(result interface{}, err error) *common.ResponseData {
	responseData := &common.ResponseData{}
	if err == nil {
		responseData.Msg = ""
		responseData.Data = result
		responseData.Code = http.StatusOK
	} else {
		responseData.Msg = err.Error()
		responseData.Data = ""
		responseData.Code = http.StatusInternalServerError
	}
	return responseData
}
