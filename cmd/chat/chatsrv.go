package chat

import (
	"github.com/SoleMer/golang/internal/config/config.go"
)

func main() {

	cfg, err := config.LoadConfig("config.yaml")

	if err != nil {
		fmt.Println(err.Error)
		os.Exit(1)
	}

	fmt.Println(cfg.DB.Driver)
	fmt.Println(cfg.Version)

}
