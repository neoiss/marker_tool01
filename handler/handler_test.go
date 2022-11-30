package handler

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"testing"
)

var endpoint = "http://127.0.0.1:7445"

func Test_getMgrMaintainerAddress(t *testing.T) {
	getMgrMaintainerAddress(endpoint)
}

func Test_setMgrMaintainerAddress(t *testing.T) {
	from := common.HexToAddress("")
	target := common.HexToAddress("")
	privateKey, err := crypto.ToECDSA(common.FromHex(""))
	if err != nil {
		t.Fatal(err)
	}

	setMgrMaintainerAddress(endpoint, from, target, privateKey)
}

func Test_getTargetEpochPayment(t *testing.T) {
	getTargetEpochPayment(endpoint)
}

func Test_setTargetEpochPayment(t *testing.T) {
	from := common.HexToAddress("")
	target := new(big.Int).Mul(big.NewInt(50000), big.NewInt(1e18))
	privateKey, err := crypto.ToECDSA(common.FromHex(""))
	if err != nil {
		t.Fatal(err)
	}

	setTargetEpochPayment(endpoint, from, target, privateKey)
}

func Test_getElectableValidators(t *testing.T) {
	getElectableValidators(endpoint)
}

// INFO [08-26|16:55:35.641] getElectableValidators                   minElectableValidators=1 maxElectableValidators=100
// INFO [08-26|17:00:42.247] getElectableValidators                   minElectableValidators=1 maxElectableValidators=50

func Test_setElectableValidators(t *testing.T) {
	from := common.HexToAddress("")
	privateKey, err := crypto.ToECDSA(common.FromHex(""))
	if err != nil {
		t.Fatal(err)
	}
	min := big.NewInt(1)
	max := big.NewInt(100)
	setElectableValidators(endpoint, from, privateKey, min, max)
}

func Test_getCommissionUpdateDelay(t *testing.T) {
	getCommissionUpdateDelay(endpoint)
}

func Test_setCommissionUpdateDelay(t *testing.T) {
	from := common.HexToAddress("")
	privateKey, err := crypto.ToECDSA(common.FromHex(""))
	if err != nil {
		t.Fatal(err)
	}
	delayBlock := big.NewInt(10)
	setCommissionUpdateDelay(endpoint, from, privateKey, delayBlock)
}

func Test_getUnlockingPeriod(t *testing.T) {
	getUnlockingPeriod(endpoint)
}

func Test_setUnlockingPeriod(t *testing.T) {
	from := common.HexToAddress("")
	privateKey, err := crypto.ToECDSA(common.FromHex(""))
	if err != nil {
		t.Fatal(err)
	}
	period := big.NewInt(900)
	setUnlockingPeriod(endpoint, from, privateKey, period)
}

func Test_getImplAddress(t *testing.T) {
	proxyAddress := common.HexToAddress("")
	getImplAddress(endpoint, proxyAddress)
}

func Test_setImplAddress(t *testing.T) {
	proxyAddress := common.HexToAddress("")
	implAddress := common.HexToAddress("")
	from := common.HexToAddress("")
	privateKey, err := crypto.ToECDSA(common.FromHex(""))
	if err != nil {
		t.Fatal(err)
	}
	setImplAddress(endpoint, from, privateKey, proxyAddress, implAddress)
}
