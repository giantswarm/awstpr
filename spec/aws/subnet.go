package aws

type Subnet struct {
	//The name of the subnet
	Name string `json:"name" yaml:"name"`
	//The name of the route table associated to the subnet
	RouteTableName string `json:"routeTableName" yaml:"routeTableName"`
}
