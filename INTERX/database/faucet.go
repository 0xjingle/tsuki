package database

import (
	"time"

	interx "github.com/TsukiCore/tsuki/INTERX/config"
	"github.com/sonyarouje/simdb/db"
)

// FaucetClaim is a struct for facuet claim.
type FaucetClaim struct {
	Address string    `json:"address"`
	Claim   time.Time `json:"claim"`
}

// ID is a field for facuet claim struct.
func (c FaucetClaim) ID() (jsonField string, value interface{}) {
	value = c.Address
	jsonField = "address"
	return
}

func getFaucetDbDriver() *db.Driver {
	driver, err := db.New(interx.GetDbCacheDir() + "faucet")
	if err != nil {
		panic(err)
	}

	return driver
}

func isClaimExist(address string) bool {
	DisableStdout()

	data := FaucetClaim{}
	err := faucetDb.Open(FaucetClaim{}).Where("address", "=", address).First().AsEntity(&data)

	EnableStdout()

	if err != nil {
		return false
	}

	return true
}

func getClaim(address string) time.Time {
	DisableStdout()

	data := FaucetClaim{}
	err := faucetDb.Open(FaucetClaim{}).Where("address", "=", address).First().AsEntity(&data)

	DisableStdout()

	if err != nil {
		panic(err)
	}

	return data.Claim
}

// GetClaimTimeLeft is a function to get left time for next claim
func GetClaimTimeLeft(address string) int64 {
	if !isClaimExist(address) {
		return 0
	}

	diff := time.Now().Unix() - getClaim(address).Unix()

	if diff > interx.Config.Faucet.TimeLimit {
		return 0
	}

	return interx.Config.Faucet.TimeLimit - diff
}

// AddNewClaim is a function to add current claim time
func AddNewClaim(address string, claim time.Time) {
	DisableStdout()

	data := FaucetClaim{
		Address: address,
		Claim:   claim,
	}

	if isClaimExist(address) {
		err := faucetDb.Open(FaucetClaim{}).Update(data)
		if err != nil {
			panic(err)
		}
	} else {
		err := faucetDb.Open(FaucetClaim{}).Insert(data)
		if err != nil {
			panic(err)
		}
	}

	EnableStdout()
}

var (
	faucetDb *db.Driver = getFaucetDbDriver()
)
