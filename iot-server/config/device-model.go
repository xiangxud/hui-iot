package config

import (
	"gopkg.in/yaml.v2"
	"hui-iot/iot-server/common"
	"hui-iot/iot-server/domain"
	"io/ioutil"
	"strings"
)

// save iot-server types info
var DeviceModelMap = make(map[string]domain.DeviceModel)

func InitDeviceModels(configPath string) (bool, map[string]domain.DeviceModel) {
	//
	//	//读取配置文件目录。
	yamlFiles, err := ioutil.ReadDir(configPath)
	if err != nil {
		common.Log.Error("Read config path ${root}/base/config/ err:{}", err.Error())
		return false, nil
	}
	//cycle read config yaml files transform iot-server type info
	for i := 0; i < len(yamlFiles); i++ {
		file := yamlFiles[i]
		if !file.IsDir() && strings.Contains(file.Name(), "dm-") {
			yamlFile, err := ioutil.ReadFile(configPath + "/" + file.Name())
			if err != nil {
				common.Log.Error("read yaml file err:{}", err.Error())
				return false, nil
			}
			deviceModel := getDeviceModel(file.Name(), yamlFile)
			DeviceModelMap[deviceModel.ID] = deviceModel
		}
	}
	return true, DeviceModelMap
}

//初始化设备类型
func getDeviceModel(filename string, bytes []byte) domain.DeviceModel {
	deviceModel := domain.DeviceModel{}
	err := yaml.Unmarshal(bytes, &deviceModel)
	if err != nil {
		panic("Read config file of iot-server-type err: " + filename + "" + err.Error())
	}
	return deviceModel
}