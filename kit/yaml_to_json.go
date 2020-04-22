package kit

import (
	"encoding/json"
	"github.com/ghodss/yaml"
)

type Controller struct {
	Kind string `json:"kind"`
}

func YamlToJson(context string) ([]byte, string, error) {
	j, err := yaml.YAMLToJSON([]byte(context))
	if err != nil {
		return nil, "", err
	}
	// 获取context内容中kind的值
	kind := Controller{}
	if err := json.Unmarshal(j, &kind); err != nil {
		return nil, "", err
	}
	return j, kind.Kind, nil
}

// 为了获取部署的kind
func JsonToJson(context []byte) ([]byte, string, error) {
	kind := Controller{}
	if err := json.Unmarshal(context, &kind); err != nil {
		return nil, "", err
	}
	return []byte(context), kind.Kind, nil
}
