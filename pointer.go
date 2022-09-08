package tips

import (
	"github.com/fsnotify/fsnotify"
	"github.com/slclub/go-tips/logf"
	"github.com/slclub/go-tips/stringbyte"
	"github.com/spf13/viper"
	"path"
	"reflect"
)

func IsNil(i interface{}) bool {
	vi := reflect.ValueOf(i)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}
	return false
}

// 用viper 读取配置文件
func ConfigWithViper(file_name_any string) *viper.Viper {
	if file_name_any == "" {
		return nil
	}
	f_path := path.Dir(file_name_any)
	f_name := path.Base(file_name_any)
	f_ext := "yaml"
	if val := path.Ext(file_name_any); len(val) > 0 {
		f_ext = val[1:]
	}
	f_name = stringbyte.SubLeft(f_name, ".") // 去掉后缀
	config := viper.New()
	config.SetConfigName(f_name) // name of config file (without extension)
	config.SetConfigType(f_ext)  // REQUIRED if the config file does not have the extension in the name
	config.AddConfigPath(f_path) // path to look for the config file in
	//viper.AddConfigPath("$HOME/.appname") // call multiple times to add many search paths
	config.AddConfigPath(".")    // optionally look for config in the working directory
	err := config.ReadInConfig() // Find and read the config file
	if err != nil {              // Handle errors reading the config file
		logf.Printf("Fatal error config file: %v \n", err)
	}

	config.OnConfigChange(func(e fsnotify.Event) {
		logf.Printf("Config file changed:", e.Name)
	})
	config.WatchConfig()
	return config
}
