package models

type SecretRDSJson struct {
	Username            string `json:"username"`
	Password            string `json:"password"`
	Engine              string `json:"engine"`
	Host                string `json:"host"`
	Port                int    `json:"port"`
	DbClusterIdentifier string `json:"dbClusterIdentifier"`
}

type SignUp struct {
	UserEmail string `json:"userEmail"`
	UserUUID  string `json:"userUUID"`
}

type Category struct {
	CategId   int    `json:"categID"`
	CategName string `json:"categName"`
	CategPath string `json:"categPath"`
}
