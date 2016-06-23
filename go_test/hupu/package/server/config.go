package main

import (
	"flag"
	"fmt"
)

import (
	"github.com/Unknwon/goconfig"
)

var (
	section = "param"
	keys    = []string{
		"version",
		"installdriver",
		"installwithgui",
		"whithoutjapan",
		"filedes",
		"prodname",
		"copyright",
		"aboutcomp",
		"officnet",
	}
)

var (
	configName string

	version        string
	installdriver  string
	installwithgui string
	whithoutjapan  string
	prodname       string
	company        string
	officnet       string
)

func init() {
	goconfig.PrettyFormat = false

	flag.StringVar(&configName, "config", "buildconfig.ini", "set buildconfig.ini path.")
	flag.StringVar(&version, "version", "3.32.41.1", "set package version.")
	flag.StringVar(&installdriver, "installdriver", "0", "set installdriver.")
	flag.StringVar(&installwithgui, "installwithgui", "0", "set installwithgui.")
	flag.StringVar(&whithoutjapan, "whithoutjapan", "0", "set whithoutjapan.")
	flag.StringVar(&prodname, "prodname", "iMan", "set prodname.")
	flag.StringVar(&company, "company", "Hupu.Info.Tec.Ltd", "set company.")
	flag.StringVar(&officnet, "officnet", "www.hupu.net", "set officnet.")
}

type BuildConfig struct {
	configFile *goconfig.ConfigFile
}

func NewBuildConfig() *BuildConfig {
	var err error
	buildConfig := new(BuildConfig)
	configName = fmt.Sprintf("%s\\%s", helperLocalPath, configName)
	buildConfig.configFile, err = goconfig.LoadConfigFile(configName)
	checkError(err)

	return buildConfig
}

func (b *BuildConfig) SetValue() error {
	fmt.Printf("%sIs to modify packaging configuration file-->[ %s ].\n", INFO, configName)
	var resultErr bool
	for _, key := range keys {
		switch key {
		case "version":
			resultErr = b.setValue(key, version)
		case "installdriver":
			resultErr = b.setValue(key, installdriver)
		case "installwithgui":
			resultErr = b.setValue(key, installwithgui)
		case "whithoutjapan":
			resultErr = b.setValue(key, whithoutjapan)
		case "filedes", "copyright", "aboutcomp":
			resultErr = b.setValue(key, company)
		case "prodname":
			resultErr = b.setValue(key, prodname)
		case "officnet":
			resultErr = b.setValue(key, officnet)
		default:
			return fmt.Errorf("SetValue: set buildConfig not key [ %s ].\n", key)
		}

		if resultErr {
			return fmt.Errorf("SetValue: %s\n", "set buildConfig error.")
		}
	}

	return b.saveConfigFile()
}

func (b *BuildConfig) setValue(key, value string) bool {
	return b.configFile.SetValue(section, key, value)
}

func (b *BuildConfig) saveConfigFile() error {
	return goconfig.SaveConfigFile(b.configFile, configName)
}

func (b BuildConfig) String() string {
	return fmt.Sprintf("buildconfig = %s\nversion = %s\ninstalldriver = %s\ninstallwithgui = %s\nwhithoutjapan = %s\nprodname = %s\ncompany = %s\nofficent = %s\n",
		configName,
		version,
		installdriver,
		installwithgui,
		whithoutjapan,
		prodname,
		company,
		officnet)
}
