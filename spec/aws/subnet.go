package aws

type Subnet struct {
	Name           string `json:"name" yaml:"name"`
	RouteTableName string `json:"routeTableName" yaml:"routeTableName"`
}
