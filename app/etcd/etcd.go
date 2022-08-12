// Describe:
package etcd

type Etcd struct {
	Endpoints string `json:"endpoints" validate:"required,url" label:"endpoints"`
	CaPath    string `json:"caPath" validate:"required,file" label:"cacert"`
	CertPath  string `json:"certPath" validate:"required,file" label:"cert"`
	KeyPath   string `json:"keyPath" validate:"required,file" label:"key"`
}
