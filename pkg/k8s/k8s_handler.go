package k8s

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	v1 "k8s.io/api/apps/v1"
	v2 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"

	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
)

func InitK8s() {
	K8sCli = &K8sClient{}
	err := K8sCli.NewKubernetesClient(&Config{
		Host:  "Host",
		Port:  6643,
		Token: "Token",
	})
	if err != nil {
		panic(err)
	}
}

// CreateDeployments 创建Deployments
func CreateDeployments(cameraId uint32) {
	var (
		deployYaml []byte
		deployJson []byte
		err        error
		deployment = &v1.Deployment{}
	)

	// 读取YAML
	if deployYaml, err = ioutil.ReadFile("E:\\rm\\smartfacotry_server\\detection.yaml"); err != nil {
		return
	}

	// YAML转JSON
	if deployJson, err = yaml.ToJSON(deployYaml); err != nil {
		return
	}

	// JSON转struct
	if err = json.Unmarshal(deployJson, &deployment); err != nil {
		return
	}

	deployment.ObjectMeta.Name = fmt.Sprintf("detect%d", cameraId)
	deployment.Name = fmt.Sprintf("detect-k8s%d", cameraId)

	// 修改cameraId参数
	deployment.Spec.Template.Spec.Containers[0].Env[1].Value = fmt.Sprintf("%d", cameraId)

	// 查询k8s是否有该deployment
	if _, err = K8sCli.ClientSet.AppsV1().Deployments("default").Get(context.TODO(), deployment.Name, metaV1.GetOptions{}); err != nil {
		if !errors.IsNotFound(err) {
			return
		}
		// 不存在则创建
		if _, err = K8sCli.ClientSet.AppsV1().Deployments("default").Create(context.TODO(), deployment, metaV1.CreateOptions{}); err != nil {
			fmt.Println(err)
			return
		}
	} else {
		// 已存在则更新
		if _, err = K8sCli.ClientSet.AppsV1().Deployments("default").Update(context.TODO(), deployment, metaV1.UpdateOptions{}); err != nil {
			return
		}
	}

	return
}

// DeleteDeployments 创建deployments
func DeleteDeployments(name string) error {
	return K8sCli.ClientSet.AppsV1().Deployments("default").Delete(context.TODO(), name, metaV1.DeleteOptions{})
}

// CreateService 创建service
func CreateService() {
	var (
		deployYaml []byte
		deployJson []byte
		err        error
		service    = &v2.Service{}
	)

	// 读取YAML
	if deployYaml, err = ioutil.ReadFile("./nginx-srv.yaml"); err != nil {
		panic(err)
	}

	// YAML转JSON
	if deployJson, err = yaml.ToJSON(deployYaml); err != nil {
		panic(err)
	}

	// JSON转struct
	if err = json.Unmarshal(deployJson, &service); err != nil {
		panic(err)
	}

	_, err = K8sCli.ClientSet.CoreV1().Services("test99").Create(context.TODO(), service, metaV1.CreateOptions{})
	if err != nil {
		panic(err.Error())
	}
}

// DeleteService 删除service
func DeleteService(name string) error {
	return K8sCli.ClientSet.CoreV1().Services("test99").Delete(context.TODO(), name, metaV1.DeleteOptions{})
}
