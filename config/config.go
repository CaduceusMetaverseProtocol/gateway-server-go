package config

import (
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "os"
    "path/filepath"
)

const (
    //////// Dev Aliyun Config ////////
      MysqlHost       = "182.92.150.56"
      MysqlPort       = "3306"
      MysqlUser       = "root"
      MysqlPassword   = "123456"
      MysqlDatabase   = "sign_service_db"
      StrAesSecretKey = "&lf4iNnzj6!FDHbvNdI1d#ymSnxbzBTi"

      //FullnodeGrpcAddr = "182.92.150.56:26669"
)

var (
    AesSecretKey []byte
    KeyStorePath string
    Config *ConfigInfo
)

type ConfigInfo struct {
    MysqlHost   string  `yaml:"mysql_host"`
    MysqlPort   string  `yaml:"mysql_port"`
    MysqlUser   string  `yaml:"mysql_user"`
    MysqlPassword   string  `yaml:"mysql_password"`
    MysqlDatabase  string `yaml:"mysql_database"`
    IsEnable      bool `yaml:"is_enable"`
}

func Init(path string) {
    if path == ""{
        path = os.Getenv("HOME")
    }
    configPath := filepath.Join(path,"config.yaml")
    yamlFile, err := ioutil.ReadFile(configPath)
    if err != nil {
        panic("yamlFile.Get err: "+ err.Error())
    }

    config:= &ConfigInfo{}
    err = yaml.Unmarshal(yamlFile,config)
    if err != nil {
        panic(err)
    }
    Config = config
    //uDir, err := homedir.Dir()
    //if err != nil {
    //    panic(err)
    //}
    //KeyStorePath = path.Join(uDir, keystore.KeyStoreScheme)
    AesSecretKey = []byte(StrAesSecretKey)
}

func EnableHmac()bool{
    return Config.IsEnable
}
