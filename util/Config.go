package util

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Config Allows streamlined file-based config access
type Config struct {
	properties map[string]*interface{}
}

func (c *Config) Read(filename string) {

	data, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
		return
	}

	var properties map[string]*interface{}
	err = json.Unmarshal(data, &properties)

	if err != nil {
		panic(err)
		return
	}

	c.properties = properties
}

func (c *Config) GetFloat(key string, defaultVal float64) float64 {
	val := c.properties[key]
	if val == nil {
		return defaultVal
	}
	i, success := (*val).(float64)
	if !success {
		errConvert(val, "float64")
	}
	return i
}

func (c *Config) GetInt(key string, defaultVal int64) int64 {
	i := c.GetFloat(key, float64(defaultVal))
	return int64(i)
}

func (c *Config) GetString(key string, defaultVal string) string {
	val := c.properties[key]
	if val == nil {
		return defaultVal
	}
	i, success := (*val).(string)
	if !success {
		errConvert(val, "float64")
	}
	return i
}

func (c *Config) GetStringArr(key string, defaultVal *[]string) []string {
	val := c.properties[key]
	if val == nil {
		return *defaultVal
	}
	arr, success := (*val).([]string)
	if !success {
		errConvert(val, "float64")
	}
	return arr
}

func errConvert(val *interface{}, typeName string) {
	log.Fatalf("%s is not %s", *val, typeName)
}
