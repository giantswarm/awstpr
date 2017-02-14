package awsclusterspec

type Spec struct {
	Customer   string                `json:"customer"`
	ClusterID  string                `json:"clusterId"`
	Kubernetes kubernetes.Kubernetes `json:"kubernetes"`

	K8sVersion     string `json:"k8sVersion"`
	KubectlVersion string `json:"kubectlVersion"`

	Master Master `json:"master"`
	Worker Worker `json:"worker"`
}

type Master struct {
	Kubernetes Kubernetes `json:"kubernetes"`
	Machine
}

type Worker struct {
	Kubernetes Kubernetes `json:"kubernetes"`
	Machine
}

type Kubernetes struct {
	DnsIP            string `json:"dnsIP"`
	Domain           string `json:"domain"`
	EtcdDomainName   string `json:"etcdDomainName"`
	MasterDomainName string `json:"masterDomainName"`
	InsecurePort     string `json:"insecurePort"`
	SecurePort       string `json:"securePort"`
}

type Machine struct {
	InstanceType string `json:"instanceType"`
	Replicas     int    `json:"replicas"`
}
