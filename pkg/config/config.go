package config

import (
	"encoding/json"
	"fmt"
	"github.com/general252/gout/uapp"
	"github.com/general252/gout/ulog"
	"io/ioutil"
	"os"
)

// JsonConfig 配置
var JsonConfig = newJsonConfig()

type jsonConfig struct {
	fileName string `json:"-"` // 配置文件
	HttpPort int    `json:"http_port"`
	DB       jsonDB `json:"db"`
}

type jsonDB struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Name     string `json:"name"`
	Password string `json:"password"`
	DBName   string `json:"db_name"`
	Charset  string `json:"charset"`
}

func (c *jsonDB) ConnectionString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", c.Name, c.Password, c.Host, c.Port, c.DBName, c.Charset)
}

func newJsonConfig() *jsonConfig {
	var ins = &jsonConfig{
		fileName: fmt.Sprintf("%v.json", uapp.GetExeName()),
		HttpPort: 9966,
		DB: jsonDB{
			Host:     "127.0.0.1",
			Port:     23306,
			Name:     "root",
			Password: "123456",
			DBName:   "test3",
			Charset:  "utf8",
		},
	}

	ins.Read()
	ins.Write()
	return ins
}

func (c *jsonConfig) Read() {
	data, err := ioutil.ReadFile(c.fileName)
	if err != nil {
		return
	}

	if err = json.Unmarshal(data, c); err != nil {
		ulog.Error("%v", err)
	}
}

func (c *jsonConfig) Write() {
	if data, err := json.MarshalIndent(c, "", "  "); err != nil {
		return
	} else {
		_ = ioutil.WriteFile(c.fileName, data, os.ModePerm)
	}
}
