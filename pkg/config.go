package pkg

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type ScriptConfig struct {
	ScriptName map[string]int `mapstructure:"scriptName"`
}

// 读取配置
func ReadFile(path string) ([]string, []int) {
	vp := viper.New()
	vp.SetConfigFile(path)
	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("Recovered from panic: %v\n", r)
		}
	}()
	err := vp.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed to read config file: %s", err))
	}

	var config ScriptConfig
	// 映射到结构体以确保有序
	err = vp.Unmarshal(&config)
	if err != nil {
		panic(fmt.Errorf("failed to unmarshal configuration: %s", err))
	}

	var resKey []string
	var resValue []int

	for k, v := range config.ScriptName {
		resKey = append(resKey, k)
		resValue = append(resValue, v)
	}
	return resKey, resValue
}

// func ReadFile(path string) ([]string, []int) {
// 	vp := viper.New()
// 	vp.SetConfigFile(path)
// 	defer func() {
// 		if r := recover(); r != nil {
// 			log.Fatalf("Recovered from panic: %v\n", r)
// 		}
// 	}()
// 	err := vp.ReadInConfig()
// 	if err != nil {
// 		panic(fmt.Errorf("failed to read config file: %s", err))
// 	}

// 	scriptNameConfig, ok := vp.Get("scriptName").(map[string]interface{})
// 	if !ok {
// 		panic(fmt.Errorf("failed to retrieve scriptName configuration"))
// 	}

// 	var resKey []string
// 	var resValue []int

// 	for k := range scriptNameConfig {
// 		resKey = append(resKey, k)
// 	}
// 	//确保有序
// 	sort.Strings(resKey)

// 	for _, v := range resKey {
// 		if val, ok := scriptNameConfig[v].(int); ok {
// 			resValue = append(resValue, val)
// 		} else {
// 			panic(fmt.Errorf("invalid type for key %s in scriptName configuration", v))
// 		}
// 	}
// 	return resKey, resValue
// }
