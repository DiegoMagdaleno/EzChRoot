package lib

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"strings"
)

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

func WriteNewPath(chrootName string, chrootPath string) {
	usr, err := user.Current()
	if err != nil {
		log.Panic("Error")
	}
	if _, err := os.Stat(usr.HomeDir + "/.config/ezchroot/paths"); err == nil {

		configPathWasOpened, err := os.Stat(usr.HomeDir + "/.config/ezchroot/paths")

		if err != nil {
			log.Panic(err)
		}

		sizeOfConfig := configPathWasOpened.Size()

		if sizeOfConfig == 0 {
			err := ioutil.WriteFile(usr.HomeDir+"/.config/ezchroot/paths", []byte(chrootName+":"+chrootPath+"\n"), 0644)
			if err != nil {
				log.Panic(err)
			}
		} else {
			file, err := os.OpenFile(usr.HomeDir+"/.config/ezchroot/paths", os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				log.Panic(err)
			}

			defer file.Close()

			if _, err := file.WriteString(chrootName + ":" + chrootPath + "\n"); err != nil {
				log.Panic(err)
			}
		}

	} else {
		directoryExists := checkIfConfigPathExists(usr.HomeDir + "/.config/ezchroot")
		if !directoryExists {
			log.Panic("Directory of configuration was not created do you have proper perms?")
		}

		configFile, err := os.Create(usr.HomeDir + "/.config/ezchroot/paths")

		if err != nil {
			log.Panic("Could not create configuration file")
		}

		configFile.Close()

		WriteNewPath(chrootName, chrootPath)

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
