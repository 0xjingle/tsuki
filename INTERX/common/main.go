package common

import (
	"io/ioutil"
	"os"
	"sync"

	interx "github.com/TsukiCore/tsuki/INTERX/config"
	"github.com/TsukiCore/tsuki/INTERX/types"
	"google.golang.org/grpc/grpclog"
)

// Mutex will be used for Sync
var Mutex = sync.Mutex{}

// RPCMethods is a variable for rpc methods
var RPCMethods = make(map[string]map[string]types.RPCMethod)

// AddRPCMethod is a function to add a RPC method
func AddRPCMethod(method string, url string, description string, canCache bool) {
	newMethod := types.RPCMethod{}
	newMethod.Description = description
	newMethod.Enabled = true
	newMethod.CachingEnabled = true
	newMethod.CachingDuration = interx.Config.CachingDuration

	if conf, ok := interx.Config.RPC.API[method][url]; ok {
		newMethod.Enabled = !conf.Disable
		newMethod.CachingEnabled = !conf.CachingDisable
		newMethod.RateLimit = conf.RateLimit
		newMethod.AuthRateLimit = conf.AuthRateLimit
		if conf.CachingDuration != 0 {
			newMethod.CachingDuration = conf.CachingDuration
		}
	}

	if !canCache {
		newMethod.CachingEnabled = false
	}

	if _, ok := RPCMethods[method]; !ok {
		RPCMethods[method] = map[string]types.RPCMethod{}
	}
	RPCMethods[method][url] = newMethod
}

var logger = grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard)

// GetLogger is a function to get logger
func GetLogger() grpclog.LoggerV2 {
	return logger
}

// NodeStatus is a struct to be used for node status
var NodeStatus struct {
	Chainid   string `json:"chain_id"`
	Block     int64  `json:"block"`
	Blocktime string `json:"block_time"`
}
