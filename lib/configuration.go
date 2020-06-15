package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"strings"
)

/*ChrootConfig type allows us to
 * unify storing into a data structure
 * the information of a chroot, like the bins
 * libraries, name and path
 * its meant to make post configuration of a chroot
 * easier.
 */
type ChrootConfig struct {
	Name string
	Path string
	Bins []string
	Libs []string
}

func checkIfConfigPathExists(configPath string) bool {
	src, err := os.Stat(configPath)

	if os.IsNotExist(err) {
		errDir := os.Mkdir(configPath, 0755)
		if errDir != nil {
			log.Panic(err)
		}
		return true
	}

	if src.Mode().IsRegular() {
		fmt.Println(configPath, "Exists but it is a file!")
		return false
	}

	if src.Mode().IsDir() {
		fmt.Println("Config directory is present!")
		return true
	}
	return false
}

func checkConfigFile(configFile string) error {
	_, err := os.Stat(configFile)
	if os.IsNotExist(err) {
		_, err := os.Create(configFile)
		if err != nil {
			return err
		}
	}
	return nil
}

/*WriteNewPath allows us to write the configuration
 * of the Chroot from memory to a file, so we can
 * modify it later.
 * It is stored on a JSON file, usignt the Chroot configstruct
 * The way this works, is by writing our bytes into a file
 * Currently it is unable to have multiple chroots due to the nature
 * of JSON and not using append mode, however I do plan on
 * fixing this
 */
func WriteNewPath(chrootName string, chrootPath string, chrootBins []string, chrootLibs []string) {

	var usr *user.User
	newChroot := &ChrootConfig{
		Name: chrootName,
		Path: chrootPath,
		Bins: chrootBins,
		Libs: chrootLibs,
	}
	configRelativePath := "/.config/ezchroot/roots"
	existingConfigData := []ChrootConfig{}

	existingConfigData = append(existingConfigData, *newChroot)

	switch realUser := os.Getuid(); realUser {
	case 0:
		sudoer := os.Getenv("SUDO_USER")
		usr, _ = user.Lookup(sudoer)
	default:
		usr, _ = user.Current()
	}

	existingConfigDataToBytes, err := json.Marshal(existingConfigData)
	if err != nil {
		log.Panic(err)
	}

	err = checkConfigFile(usr.HomeDir + configRelativePath)
	if err != nil {
		log.Panic(err)
	}

	err = ioutil.WriteFile(usr.HomeDir+configRelativePath, existingConfigDataToBytes, 0644)
	if err != nil {
		log.Panic(err)
	}

}

func ManageConfig(name string) (bool, string) {
	var eachConfigValue []string
	var splitConfig []string
	usr, err := user.Current()
	if err != nil {
		log.Panic("Error")
	}
	content, err := ioutil.ReadFile(usr.HomeDir + "/.config/ezchroot/paths")

	if err != nil {
		log.Panic(err)
	}

	lines := strings.Split(string(content), "\n")

	for i := 0; i <= len(lines)-1; i++ {
		splitConfig = strings.Split(lines[i], ":")
		for j := 0; j <= len(splitConfig)-1; j++ {
			eachConfigValue = append(eachConfigValue, splitConfig[j])
		}
	}

	elementMap := make(map[string]string)
	for i := 0; i < len(eachConfigValue)-1; i += 2 {
		elementMap[eachConfigValue[i]] = eachConfigValue[i+1]
	}

	value, ok := elementMap[name]
	if ok {
		return true, value
	} else {
		return false, name
	}

}
