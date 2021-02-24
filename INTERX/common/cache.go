package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/TsukiCore/tsuki/INTERX/config"
	"github.com/TsukiCore/tsuki/INTERX/global"
	"github.com/TsukiCore/tsuki/INTERX/types"
)

// PutCache is a function to save value to cache
func PutCache(chainIDHash string, endpointHash string, requestHash string, value types.InterxResponse) error {
	GetLogger().Info("[cache] Saving interx response")

	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	folderPath := fmt.Sprintf("%s/%s/%s", config.GetResponseCacheDir(), chainIDHash, endpointHash)
	filePath := fmt.Sprintf("%s/%s", folderPath, requestHash)

	global.Mutex.Lock()
	err = os.MkdirAll(folderPath, os.ModePerm)
	if err != nil {
		global.Mutex.Unlock()

		GetLogger().Error("[cache] Unable to create a folder: ", folderPath)
		return err
	}

	err = ioutil.WriteFile(filePath, data, 0644)
	global.Mutex.Unlock()

	if err != nil {
		GetLogger().Error("[cache] Unable to save response: ", filePath)
	}

	return err
}

// GetCache is a function to get value from cache
func GetCache(chainIDHash string, endpointHash string, requestHash string) (types.InterxResponse, error) {
	filePath := fmt.Sprintf("%s/%s/%s/%s", config.GetResponseCacheDir(), chainIDHash, endpointHash, requestHash)

	response := types.InterxResponse{}

	GetLogger().Info("start here")
	data, err := ioutil.ReadFile(filePath)
	GetLogger().Info("end here")

	if err != nil {
		return response, nil
	}

	err = json.Unmarshal([]byte(data), &response)

	return response, err
}
