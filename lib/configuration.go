package lib

import (
	"fmt"
	"os"

	"github.com/adrg/xdg"
	"gopkg.in/yaml.v2"
)

/*ChrootConfig type allows us to
 * unify storing into a data structure
 * the information of a chroot, like the bins
 * libraries, name and path
 * its meant to make post configuration of a chroot
 * easier.
 */
type ChrootConfig struct {
	Bins []string `yaml:"bins"`
}

func GetConfig(cfg *ChrootConfig) {
	configPath := fmt.Sprintf("%s/ezchroot/config.yml", xdg.ConfigHome)

	file, err := os.Open(configPath)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(cfg)
	if err != nil {
		panic(err)
	}
}
