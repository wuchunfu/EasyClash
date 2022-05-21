package clash

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Dreamacro/clash/config"
	C "github.com/Dreamacro/clash/constant"
	"github.com/Dreamacro/clash/hub"
	"github.com/Dreamacro/clash/hub/executor"
	"github.com/Dreamacro/clash/log"
	"go.uber.org/automaxprocs/maxprocs"
)

func RunClash(testConfig bool, configFile string) error {
	maxprocs.Set(maxprocs.Logger(func(string, ...any) {}))

	if configFile != "" {
		if !filepath.IsAbs(configFile) {
			currentDir, _ := os.Getwd()
			configFile = filepath.Join(currentDir, configFile)
		}
		C.SetConfig(configFile)
	} else {
		configFile := filepath.Join(C.Path.HomeDir(), C.Path.Config())
		C.SetConfig(configFile)
	}

	if err := config.Init(C.Path.HomeDir()); err != nil {
		return fmt.Errorf("initial configuration directory error: %s", err.Error())
	}

	if testConfig {
		if _, err := executor.Parse(); err != nil {
			log.Errorln(err.Error())
			return fmt.Errorf("configuration file %s test failed\n", C.Path.Config())
		}
	}

	var options []hub.Option
	//if flagset["ext-ui"] {
	//	options = append(options, hub.WithExternalUI(externalUI))
	//}
	//if flagset["ext-ctl"] {
	//	options = append(options, hub.WithExternalController(externalController))
	//}
	//if flagset["secret"] {
	//	options = append(options, hub.WithSecret(secret))
	//}

	if err := hub.Parse(options...); err != nil {
		return fmt.Errorf("parse config error: %s", err.Error())
	}

	return nil
}

func CheckConf(configFile string) (bool, error) {
	if configFile != "" {
		if !filepath.IsAbs(configFile) {
			currentDir, _ := os.Getwd()
			configFile = filepath.Join(currentDir, configFile)
		}
		C.SetConfig(configFile)
	} else {
		configFile := filepath.Join(C.Path.HomeDir(), C.Path.Config())
		C.SetConfig(configFile)
	}
	if _, err := executor.Parse(); err != nil {
		log.Errorln(err.Error())
		return false, fmt.Errorf("configuration file %s test failed\n", C.Path.Config())
	}

	return true, nil
}
