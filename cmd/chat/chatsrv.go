package chat

import (
	"github.com/SoleMer/golang/internal/config/config.go"
)

func main() {

	cfg, err := config.LoadConfig("config.yaml")

	if err != nil {
		return nil, err
	}

	var c := &Config{}
	err = yaml.Umarshal(file, c)
	if err!= nil {
		return nil, err
	}
	return c, nil
}
