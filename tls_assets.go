package awstpr

type ClusterComponent string
type TLSAssetType string

const (
	CA  TLSAssetType = "ca"
	Crt TLSAssetType = "crt"
	Key TLSAssetType = "key"

	APIComponent    ClusterComponent = "api"
	WorkerComponent ClusterComponent = "worker"
	EtcdComponent   ClusterComponent = "etcd"
	CalicoComponent ClusterComponent = "calico"

	ComponentLabel string = "clusterComponent"
	ClusterIDLabel string = "clusterID"
)

type AssetsBundle map[ClusterComponent]map[TLSAssetType][]byte

var ClusterComponents []ClusterComponent = []ClusterComponent{APIComponent, WorkerComponent, EtcdComponent, CalicoComponent}
var TLSAssetTypes []TLSAssetType = []TLSAssetType{CA, Crt, Key}

func ValidComponent(el ClusterComponent) bool {
	for _, v := range ClusterComponents {
		if el == v {
			return true
		}
	}
	return false
}

func NewAssetsBundle() AssetsBundle {
	bundle := make(map[ClusterComponent]map[TLSAssetType][]byte, len(ClusterComponents))

	for k, _ := range bundle {
		bundle[k] = make(map[TLSAssetType][]byte, len(TLSAssetTypes))
	}

	return bundle
}
