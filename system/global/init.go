package global

import (
	"fmt"
	"github.com/spf13/viper"
	"math/rand"
	"sync"
	"time"
)

//把一些全局的东西都加载进来

var once sync.Once

func Init()  {
	once.Do(func() {
		//加入随机种子
		rand.Seed(time.Now().UnixNano())
		//设置配置
		viper.SetConfigName("config")
		viper.AddConfigPath("config/")
		viper.SetConfigType("yaml")
		//读取配置文件
		err := viper.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
	})
}
