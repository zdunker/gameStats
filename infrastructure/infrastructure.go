package infrastructure

type infrastructures struct{}

var infras = infrastructures{}

func Infrastructure() infrastructures {
	return infras
}

// Here I want to abandon config file design, instead, to keep infrastructure config safe,
// I would like to use ENV variables to configure the infrastructure.
// So that I could use docker workflow to deploy the entire server with presetting ENV variables to docker image.
func InitInfrastructure() error {
	// TODO
	return nil
}

// func (inf infrastructures) GetDB() (*pg.DB, error) {
// 	//TODO
// }
