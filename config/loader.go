package config

import (
	"github.com/spf13/viper"
	"path/filepath"
)

func Load(path string)(config *Config, err error)  {
	var (
		dir      = filepath.Dir(path)
		filename = filepath.Base(path)
		ext      = filepath.Ext(path)[1:]
	)

	viper.AddConfigPath(dir)
	viper.SetConfigName(filename)
	viper.SetConfigType(ext)

	if err = viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err = viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	return
}
