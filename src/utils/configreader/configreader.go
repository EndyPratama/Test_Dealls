package configreader

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Interface interface {
	ReadConfig(cfg interface{})
	AllSettings() map[string]interface{}
}

type Options struct {
	Name string
	Type string
	Path string
}

type configReader struct {
	viper *viper.Viper
	opt   Options
}

func Init(opt Options) Interface {
	vp := viper.New()
	vp.SetConfigName(opt.Name)
	vp.SetConfigType(opt.Type)
	vp.AddConfigPath(opt.Path)
	if err := vp.ReadInConfig(); err != nil {
		fmt.Println("Name :", opt.Name)
		fmt.Println("Type :", opt.Type)
		fmt.Println("Path :", opt.Path)
		pwd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("Dir now :", pwd)
		panic(fmt.Errorf("fatal error found during reading file. err: %w", err))
	}

	c := &configReader{
		viper: vp,
		opt:   opt,
	}

	return c
}

func (c *configReader) ReadConfig(cfg interface{}) {
	if err := c.viper.Unmarshal(&cfg); err != nil {
		panic(fmt.Errorf("fatal error found during unmarshaling config. err: %w", err))
	}
}

func (c *configReader) AllSettings() map[string]interface{} {
	return c.viper.AllSettings()
}
