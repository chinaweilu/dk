package conf

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/larspensjo/config"
)

var ini map[string]map[string]string

func init() {
	ini = make(map[string]map[string]string)
	path, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	cfg, err := config.ReadDefault(path + "\\dk.ini")
	if err != nil {
		fmt.Println(err)
		return
	}
	head := []string{"redisServer", "redisConnect"}
	for _, iniKey := range head {
		if !cfg.HasSection(iniKey) {
			continue
		}
		list, err := cfg.SectionOptions(iniKey)
		if err != nil {
			fmt.Println(err)
			return
		}
		var iniMap = make(map[string]string)
		for _, v := range list {
			value, err := cfg.String(iniKey, v)
			if err != nil {
				fmt.Println(err)
				continue
			}
			iniMap[v] = value
		}
		ini[iniKey] = iniMap
	}
}

func GetRedisServer() []string {
	v, ok := ini["redisServer"]
	if !ok {
		return nil
	}

	result := make([]string, len(v))

	i := 0
	for _, address := range v {
		result[i] = address
		i++
	}
	return result
}

func GetIni(h, k string) string {
	conf, ok := ini[h]
	if !ok {
		return ""
	}

	value, ok := conf[k]
	if !ok {
		return ""
	}
	return value
}
