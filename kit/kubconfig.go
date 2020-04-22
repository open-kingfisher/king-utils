package kit

import (
	"github.com/open-kingfisher/king-utils/common"
	"os"
)

func IsExist(kubconfig string) bool {
	_, err := os.Stat(kubconfig)
	return err == nil || os.IsExist(err)
}

func DeleteConfig(kubeconfig string) error {
	return os.Remove(kubeconfig)
}

func CreateConfig(cluster common.ClusterDB, kubeconfig string) error {
	file, err := os.Create(kubeconfig)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(cluster.KubConfig)
	if err != nil {
		return err
	}
	return nil
}
