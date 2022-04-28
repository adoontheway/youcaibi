package conf

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	LBAddr  string `json:"lb_addr"` //负载均衡地址
	OssAddr string `json:"oss_addr"`
}

var configuration *Configuration

func init() {
	fmt.Println(os.Getwd())
	file, _ := os.Open("../conf/conf.json")
	defer file.Close()

	decoder := json.NewDecoder(file)
	configuration = &Configuration{}

	err := decoder.Decode(configuration)
	if err != nil {
		panic(err)
	}
}

func GetLbAddr() string {
	return configuration.LBAddr
}

func GetOssAddr() string {
	return configuration.OssAddr
}
