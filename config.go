package mscl

import (
	"errors"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Config contains relevent parameters for Servers and ServerControllers
type Config struct {
	Username  string `yaml:"username"`
	RAMMin    int    `yaml:"ram-min"`
	RAMMax    int    `yaml:"ram-max"`
	JarFile   string `yaml:"jar-file"`
	Path      string `yaml:"root-path"`
	JavaFlags string `yaml:"java-flags"`
}

// New initializes a nil Config to default values
func (c *Config) New() error {
	if c != nil {
		return errors.New("non-nil struct object")
	}
	c = &Config{
		Username:  "minecraft",
		RAMMin:    2048,
		RAMMax:    2048,
		JarFile:   "server.jar",
		Path:      "/opt/minecraft/",
		JavaFlags: "-XX:+UseG1GC -XX:+ParallelRefProcEnabled -XX:MaxGCPauseMillis=200 -XX:+UnlockExperimentalVMOptions -XX:+DisableExplicitGC -XX:+AlwaysPreTouch -XX:G1NewSizePercent=30 -XX:G1MaxNewSizePercent=40 -XX:G1HeapRegionSize=8M -XX:G1ReservePercent=20 -XX:G1HeapWastePercent=5 -XX:G1MixedGCCountTarget=4 -XX:InitiatingHeapOccupancyPercent=15 -XX:G1MixedGCLiveThresholdPercent=90 -XX:G1RSetUpdatingPauseTimePercent=5 -XX:SurvivorRatio=32 -XX:+PerfDisableSharedMem -XX:MaxTenuringThreshold=1",
	}
	return nil
}

// Load loads config data from file at location path
func (c *Config) Load(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(file, &c)
	if err != nil {
		return err
	}

	return nil
}

// Validate validates all fields of a given config and returns either nil or an error
//func (c *Config) Validate() error {
//TODO
//}

// Save writes config data to file at location path
func (c *Config) Save(path string) error {
	file, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	ioutil.WriteFile(path, file, 0644)
	if err != nil {
		return err
	}

	return nil
}
