// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// VFinanceRegistryContractRecord is an auto generated low-level Go binding around an user-defined struct.
type VFinanceRegistryContractRecord struct {
	RegistryId     [32]byte
	ContractNumber [32]byte
	ContractDate   uint32
	MetadataHash   [32]byte
	Timestamp      uint32
	RegisteredBy   common.Address
	Active         bool
}

// VFinanceRegistryVehicleCore is an auto generated low-level Go binding around an user-defined struct.
type VFinanceRegistryVehicleCore struct {
	Chassis      [32]byte
	LicensePlate [32]byte
	TotalValue   *big.Int
	BrandId      uint64
	ModelId      uint64
}

// BindingsMetaData contains all meta data concerning the Bindings contract.
var BindingsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"MAX_SUPPLY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"apiServerAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"brands\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"chassisHashToTokenId\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"contracts\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"registryId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"contractNumber\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"contractDate\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"metadataHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"timestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"registeredBy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"doesContractExist\",\"inputs\":[{\"name\":\"registryId\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"doesHashExist\",\"inputs\":[{\"name\":\"metadataHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"emergencyWithdrawToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getActiveContracts\",\"inputs\":[{\"name\":\"offset\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"tokenIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getApproved\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBrandName\",\"inputs\":[{\"name\":\"brandId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getContract\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structVFinanceRegistry.ContractRecord\",\"components\":[{\"name\":\"registryId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"contractNumber\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"contractDate\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"metadataHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"timestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"registeredBy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structVFinanceRegistry.VehicleCore\",\"components\":[{\"name\":\"chassis\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"licensePlate\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"totalValue\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"brandId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"modelId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getContractByChassis\",\"inputs\":[{\"name\":\"chassis\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structVFinanceRegistry.ContractRecord\",\"components\":[{\"name\":\"registryId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"contractNumber\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"contractDate\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"metadataHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"timestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"registeredBy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structVFinanceRegistry.VehicleCore\",\"components\":[{\"name\":\"chassis\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"licensePlate\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"totalValue\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"brandId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"modelId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getContractByHash\",\"inputs\":[{\"name\":\"metadataHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structVFinanceRegistry.ContractRecord\",\"components\":[{\"name\":\"registryId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"contractNumber\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"contractDate\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"metadataHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"timestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"registeredBy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structVFinanceRegistry.VehicleCore\",\"components\":[{\"name\":\"chassis\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"licensePlate\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"totalValue\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"brandId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"modelId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getContractByRegistryId\",\"inputs\":[{\"name\":\"registryId\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structVFinanceRegistry.ContractRecord\",\"components\":[{\"name\":\"registryId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"contractNumber\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"contractDate\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"metadataHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"timestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"registeredBy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structVFinanceRegistry.VehicleCore\",\"components\":[{\"name\":\"chassis\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"licensePlate\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"totalValue\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"brandId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"modelId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMetadataUrl\",\"inputs\":[{\"name\":\"metadataHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMetadataUrlByRegistryId\",\"inputs\":[{\"name\":\"registryId\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getModelName\",\"inputs\":[{\"name\":\"modelId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_metadataBaseUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_apiServerAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isApprovedForAll\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"metadataBaseUrl\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"metadataHashToTokenId\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"models\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ownerOf\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerBrand\",\"inputs\":[{\"name\":\"brandName\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"brandId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerContract\",\"inputs\":[{\"name\":\"registryId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"contractNumber\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"contractDate\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"chassis\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"licensePlate\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"totalValue\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"brandName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"modelName\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"metadataHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerModel\",\"inputs\":[{\"name\":\"modelName\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"modelId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registryIdHashToTokenId\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"safeTransferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"safeTransferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setApprovalForAll\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenURI\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateMetadataHash\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"newMetadataHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateServerConfig\",\"inputs\":[{\"name\":\"newMetadataBaseUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"newApiServerAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateStatus\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"vehicleCores\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"chassis\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"licensePlate\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"totalValue\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"brandId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"modelId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ApprovalForAll\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BrandRegistered\",\"inputs\":[{\"name\":\"brandId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"brandName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ContractRegistered\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"registryIdHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"chassisHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"metadataHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"timestamp\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MetadataUpdated\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"oldHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"timestamp\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ModelRegistered\",\"inputs\":[{\"name\":\"modelId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"modelName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StatusChanged\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"active\",\"type\":\"bool\",\"indexed\":true,\"internalType\":\"bool\"},{\"name\":\"timestamp\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SystemConfigured\",\"inputs\":[{\"name\":\"apiServer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"metadataUrl\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"VehicleTracked\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"chassis\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"licensePlate\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"totalValue\",\"type\":\"uint128\",\"indexed\":false,\"internalType\":\"uint128\"},{\"name\":\"brandId\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"modelId\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ContractAlreadyExists\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ContractNotFound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DuplicateHash\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ERC721IncorrectOwner\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721InsufficientApproval\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC721InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721InvalidOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721InvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721NonexistentToken\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInput\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MaxSupplyExceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenNotFound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"UnauthorizedAccess\",\"inputs\":[]}]",
}

// BindingsABI is the input ABI used to generate the binding from.
// Deprecated: Use BindingsMetaData.ABI instead.
var BindingsABI = BindingsMetaData.ABI

// Bindings is an auto generated Go binding around an Ethereum contract.
type Bindings struct {
	BindingsCaller     // Read-only binding to the contract
	BindingsTransactor // Write-only binding to the contract
	BindingsFilterer   // Log filterer for contract events
}

// BindingsCaller is an auto generated read-only Go binding around an Ethereum contract.
type BindingsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BindingsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BindingsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BindingsSession struct {
	Contract     *Bindings         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BindingsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BindingsCallerSession struct {
	Contract *BindingsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BindingsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BindingsTransactorSession struct {
	Contract     *BindingsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BindingsRaw is an auto generated low-level Go binding around an Ethereum contract.
type BindingsRaw struct {
	Contract *Bindings // Generic contract binding to access the raw methods on
}

// BindingsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BindingsCallerRaw struct {
	Contract *BindingsCaller // Generic read-only contract binding to access the raw methods on
}

// BindingsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BindingsTransactorRaw struct {
	Contract *BindingsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBindings creates a new instance of Bindings, bound to a specific deployed contract.
func NewBindings(address common.Address, backend bind.ContractBackend) (*Bindings, error) {
	contract, err := bindBindings(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bindings{BindingsCaller: BindingsCaller{contract: contract}, BindingsTransactor: BindingsTransactor{contract: contract}, BindingsFilterer: BindingsFilterer{contract: contract}}, nil
}

// NewBindingsCaller creates a new read-only instance of Bindings, bound to a specific deployed contract.
func NewBindingsCaller(address common.Address, caller bind.ContractCaller) (*BindingsCaller, error) {
	contract, err := bindBindings(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BindingsCaller{contract: contract}, nil
}

// NewBindingsTransactor creates a new write-only instance of Bindings, bound to a specific deployed contract.
func NewBindingsTransactor(address common.Address, transactor bind.ContractTransactor) (*BindingsTransactor, error) {
	contract, err := bindBindings(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BindingsTransactor{contract: contract}, nil
}

// NewBindingsFilterer creates a new log filterer instance of Bindings, bound to a specific deployed contract.
func NewBindingsFilterer(address common.Address, filterer bind.ContractFilterer) (*BindingsFilterer, error) {
	contract, err := bindBindings(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BindingsFilterer{contract: contract}, nil
}

// bindBindings binds a generic wrapper to an already deployed contract.
func bindBindings(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BindingsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bindings *BindingsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bindings.Contract.BindingsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bindings *BindingsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.Contract.BindingsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bindings *BindingsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bindings.Contract.BindingsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bindings *BindingsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bindings.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bindings *BindingsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bindings *BindingsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bindings.Contract.contract.Transact(opts, method, params...)
}

// MAXSUPPLY is a free data retrieval call binding the contract method 0x32cb6b0c.
//
// Solidity: function MAX_SUPPLY() view returns(uint256)
func (_Bindings *BindingsCaller) MAXSUPPLY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "MAX_SUPPLY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXSUPPLY is a free data retrieval call binding the contract method 0x32cb6b0c.
//
// Solidity: function MAX_SUPPLY() view returns(uint256)
func (_Bindings *BindingsSession) MAXSUPPLY() (*big.Int, error) {
	return _Bindings.Contract.MAXSUPPLY(&_Bindings.CallOpts)
}

// MAXSUPPLY is a free data retrieval call binding the contract method 0x32cb6b0c.
//
// Solidity: function MAX_SUPPLY() view returns(uint256)
func (_Bindings *BindingsCallerSession) MAXSUPPLY() (*big.Int, error) {
	return _Bindings.Contract.MAXSUPPLY(&_Bindings.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Bindings *BindingsCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Bindings *BindingsSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Bindings.Contract.UPGRADEINTERFACEVERSION(&_Bindings.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Bindings *BindingsCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Bindings.Contract.UPGRADEINTERFACEVERSION(&_Bindings.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Bindings *BindingsCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Bindings *BindingsSession) VERSION() (string, error) {
	return _Bindings.Contract.VERSION(&_Bindings.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Bindings *BindingsCallerSession) VERSION() (string, error) {
	return _Bindings.Contract.VERSION(&_Bindings.CallOpts)
}

// ApiServerAddress is a free data retrieval call binding the contract method 0xbc33768e.
//
// Solidity: function apiServerAddress() view returns(address)
func (_Bindings *BindingsCaller) ApiServerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "apiServerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ApiServerAddress is a free data retrieval call binding the contract method 0xbc33768e.
//
// Solidity: function apiServerAddress() view returns(address)
func (_Bindings *BindingsSession) ApiServerAddress() (common.Address, error) {
	return _Bindings.Contract.ApiServerAddress(&_Bindings.CallOpts)
}

// ApiServerAddress is a free data retrieval call binding the contract method 0xbc33768e.
//
// Solidity: function apiServerAddress() view returns(address)
func (_Bindings *BindingsCallerSession) ApiServerAddress() (common.Address, error) {
	return _Bindings.Contract.ApiServerAddress(&_Bindings.CallOpts)
}

// Approve is a free data retrieval call binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address , uint256 ) pure returns()
func (_Bindings *BindingsCaller) Approve(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) error {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "approve", arg0, arg1)

	if err != nil {
		return err
	}

	return err

}

// Approve is a free data retrieval call binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address , uint256 ) pure returns()
func (_Bindings *BindingsSession) Approve(arg0 common.Address, arg1 *big.Int) error {
	return _Bindings.Contract.Approve(&_Bindings.CallOpts, arg0, arg1)
}

// Approve is a free data retrieval call binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address , uint256 ) pure returns()
func (_Bindings *BindingsCallerSession) Approve(arg0 common.Address, arg1 *big.Int) error {
	return _Bindings.Contract.Approve(&_Bindings.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Bindings *BindingsCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Bindings *BindingsSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Bindings.Contract.BalanceOf(&_Bindings.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Bindings *BindingsCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Bindings.Contract.BalanceOf(&_Bindings.CallOpts, owner)
}

// Brands is a free data retrieval call binding the contract method 0xa636733c.
//
// Solidity: function brands(uint64 ) view returns(string)
func (_Bindings *BindingsCaller) Brands(opts *bind.CallOpts, arg0 uint64) (string, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "brands", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Brands is a free data retrieval call binding the contract method 0xa636733c.
//
// Solidity: function brands(uint64 ) view returns(string)
func (_Bindings *BindingsSession) Brands(arg0 uint64) (string, error) {
	return _Bindings.Contract.Brands(&_Bindings.CallOpts, arg0)
}

// Brands is a free data retrieval call binding the contract method 0xa636733c.
//
// Solidity: function brands(uint64 ) view returns(string)
func (_Bindings *BindingsCallerSession) Brands(arg0 uint64) (string, error) {
	return _Bindings.Contract.Brands(&_Bindings.CallOpts, arg0)
}

// ChassisHashToTokenId is a free data retrieval call binding the contract method 0xabb47e10.
//
// Solidity: function chassisHashToTokenId(bytes32 ) view returns(uint256)
func (_Bindings *BindingsCaller) ChassisHashToTokenId(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "chassisHashToTokenId", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChassisHashToTokenId is a free data retrieval call binding the contract method 0xabb47e10.
//
// Solidity: function chassisHashToTokenId(bytes32 ) view returns(uint256)
func (_Bindings *BindingsSession) ChassisHashToTokenId(arg0 [32]byte) (*big.Int, error) {
	return _Bindings.Contract.ChassisHashToTokenId(&_Bindings.CallOpts, arg0)
}

// ChassisHashToTokenId is a free data retrieval call binding the contract method 0xabb47e10.
//
// Solidity: function chassisHashToTokenId(bytes32 ) view returns(uint256)
func (_Bindings *BindingsCallerSession) ChassisHashToTokenId(arg0 [32]byte) (*big.Int, error) {
	return _Bindings.Contract.ChassisHashToTokenId(&_Bindings.CallOpts, arg0)
}

// Contracts is a free data retrieval call binding the contract method 0x474da79a.
//
// Solidity: function contracts(uint256 ) view returns(bytes32 registryId, bytes32 contractNumber, uint32 contractDate, bytes32 metadataHash, uint32 timestamp, address registeredBy, bool active)
func (_Bindings *BindingsCaller) Contracts(opts *bind.CallOpts, arg0 *big.Int) (struct {
	RegistryId     [32]byte
	ContractNumber [32]byte
	ContractDate   uint32
	MetadataHash   [32]byte
	Timestamp      uint32
	RegisteredBy   common.Address
	Active         bool
}, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "contracts", arg0)

	outstruct := new(struct {
		RegistryId     [32]byte
		ContractNumber [32]byte
		ContractDate   uint32
		MetadataHash   [32]byte
		Timestamp      uint32
		RegisteredBy   common.Address
		Active         bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RegistryId = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.ContractNumber = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.ContractDate = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.MetadataHash = *abi.ConvertType(out[3], new([32]byte)).(*[32]byte)
	outstruct.Timestamp = *abi.ConvertType(out[4], new(uint32)).(*uint32)
	outstruct.RegisteredBy = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.Active = *abi.ConvertType(out[6], new(bool)).(*bool)

	return *outstruct, err

}

// Contracts is a free data retrieval call binding the contract method 0x474da79a.
//
// Solidity: function contracts(uint256 ) view returns(bytes32 registryId, bytes32 contractNumber, uint32 contractDate, bytes32 metadataHash, uint32 timestamp, address registeredBy, bool active)
func (_Bindings *BindingsSession) Contracts(arg0 *big.Int) (struct {
	RegistryId     [32]byte
	ContractNumber [32]byte
	ContractDate   uint32
	MetadataHash   [32]byte
	Timestamp      uint32
	RegisteredBy   common.Address
	Active         bool
}, error) {
	return _Bindings.Contract.Contracts(&_Bindings.CallOpts, arg0)
}

// Contracts is a free data retrieval call binding the contract method 0x474da79a.
//
// Solidity: function contracts(uint256 ) view returns(bytes32 registryId, bytes32 contractNumber, uint32 contractDate, bytes32 metadataHash, uint32 timestamp, address registeredBy, bool active)
func (_Bindings *BindingsCallerSession) Contracts(arg0 *big.Int) (struct {
	RegistryId     [32]byte
	ContractNumber [32]byte
	ContractDate   uint32
	MetadataHash   [32]byte
	Timestamp      uint32
	RegisteredBy   common.Address
	Active         bool
}, error) {
	return _Bindings.Contract.Contracts(&_Bindings.CallOpts, arg0)
}

// DoesContractExist is a free data retrieval call binding the contract method 0x0df205f0.
//
// Solidity: function doesContractExist(string registryId) view returns(bool)
func (_Bindings *BindingsCaller) DoesContractExist(opts *bind.CallOpts, registryId string) (bool, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "doesContractExist", registryId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DoesContractExist is a free data retrieval call binding the contract method 0x0df205f0.
//
// Solidity: function doesContractExist(string registryId) view returns(bool)
func (_Bindings *BindingsSession) DoesContractExist(registryId string) (bool, error) {
	return _Bindings.Contract.DoesContractExist(&_Bindings.CallOpts, registryId)
}

// DoesContractExist is a free data retrieval call binding the contract method 0x0df205f0.
//
// Solidity: function doesContractExist(string registryId) view returns(bool)
func (_Bindings *BindingsCallerSession) DoesContractExist(registryId string) (bool, error) {
	return _Bindings.Contract.DoesContractExist(&_Bindings.CallOpts, registryId)
}

// DoesHashExist is a free data retrieval call binding the contract method 0x1ee66770.
//
// Solidity: function doesHashExist(bytes32 metadataHash) view returns(bool)
func (_Bindings *BindingsCaller) DoesHashExist(opts *bind.CallOpts, metadataHash [32]byte) (bool, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "doesHashExist", metadataHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DoesHashExist is a free data retrieval call binding the contract method 0x1ee66770.
//
// Solidity: function doesHashExist(bytes32 metadataHash) view returns(bool)
func (_Bindings *BindingsSession) DoesHashExist(metadataHash [32]byte) (bool, error) {
	return _Bindings.Contract.DoesHashExist(&_Bindings.CallOpts, metadataHash)
}

// DoesHashExist is a free data retrieval call binding the contract method 0x1ee66770.
//
// Solidity: function doesHashExist(bytes32 metadataHash) view returns(bool)
func (_Bindings *BindingsCallerSession) DoesHashExist(metadataHash [32]byte) (bool, error) {
	return _Bindings.Contract.DoesHashExist(&_Bindings.CallOpts, metadataHash)
}

// GetActiveContracts is a free data retrieval call binding the contract method 0xf672cbc5.
//
// Solidity: function getActiveContracts(uint256 offset, uint256 limit) view returns(uint256[] tokenIds)
func (_Bindings *BindingsCaller) GetActiveContracts(opts *bind.CallOpts, offset *big.Int, limit *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getActiveContracts", offset, limit)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetActiveContracts is a free data retrieval call binding the contract method 0xf672cbc5.
//
// Solidity: function getActiveContracts(uint256 offset, uint256 limit) view returns(uint256[] tokenIds)
func (_Bindings *BindingsSession) GetActiveContracts(offset *big.Int, limit *big.Int) ([]*big.Int, error) {
	return _Bindings.Contract.GetActiveContracts(&_Bindings.CallOpts, offset, limit)
}

// GetActiveContracts is a free data retrieval call binding the contract method 0xf672cbc5.
//
// Solidity: function getActiveContracts(uint256 offset, uint256 limit) view returns(uint256[] tokenIds)
func (_Bindings *BindingsCallerSession) GetActiveContracts(offset *big.Int, limit *big.Int) ([]*big.Int, error) {
	return _Bindings.Contract.GetActiveContracts(&_Bindings.CallOpts, offset, limit)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Bindings *BindingsCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Bindings *BindingsSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Bindings.Contract.GetApproved(&_Bindings.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Bindings *BindingsCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Bindings.Contract.GetApproved(&_Bindings.CallOpts, tokenId)
}

// GetBrandName is a free data retrieval call binding the contract method 0x89d3e203.
//
// Solidity: function getBrandName(uint64 brandId) view returns(string)
func (_Bindings *BindingsCaller) GetBrandName(opts *bind.CallOpts, brandId uint64) (string, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getBrandName", brandId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetBrandName is a free data retrieval call binding the contract method 0x89d3e203.
//
// Solidity: function getBrandName(uint64 brandId) view returns(string)
func (_Bindings *BindingsSession) GetBrandName(brandId uint64) (string, error) {
	return _Bindings.Contract.GetBrandName(&_Bindings.CallOpts, brandId)
}

// GetBrandName is a free data retrieval call binding the contract method 0x89d3e203.
//
// Solidity: function getBrandName(uint64 brandId) view returns(string)
func (_Bindings *BindingsCallerSession) GetBrandName(brandId uint64) (string, error) {
	return _Bindings.Contract.GetBrandName(&_Bindings.CallOpts, brandId)
}

// GetContract is a free data retrieval call binding the contract method 0x6ebc8c86.
//
// Solidity: function getContract(uint256 tokenId) view returns((bytes32,bytes32,uint32,bytes32,uint32,address,bool), (bytes32,bytes32,uint128,uint64,uint64))
func (_Bindings *BindingsCaller) GetContract(opts *bind.CallOpts, tokenId *big.Int) (VFinanceRegistryContractRecord, VFinanceRegistryVehicleCore, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getContract", tokenId)

	if err != nil {
		return *new(VFinanceRegistryContractRecord), *new(VFinanceRegistryVehicleCore), err
	}

	out0 := *abi.ConvertType(out[0], new(VFinanceRegistryContractRecord)).(*VFinanceRegistryContractRecord)
	out1 := *abi.ConvertType(out[1], new(VFinanceRegistryVehicleCore)).(*VFinanceRegistryVehicleCore)

	return out0, out1, err

}

// GetContract is a free data retrieval call binding the contract method 0x6ebc8c86.
//
// Solidity: function getContract(uint256 tokenId) view returns((bytes32,bytes32,uint32,bytes32,uint32,address,bool), (bytes32,bytes32,uint128,uint64,uint64))
func (_Bindings *BindingsSession) GetContract(tokenId *big.Int) (VFinanceRegistryContractRecord, VFinanceRegistryVehicleCore, error) {
	return _Bindings.Contract.GetContract(&_Bindings.CallOpts, tokenId)
}

// GetContract is a free data retrieval call binding the contract method 0x6ebc8c86.
//
// Solidity: function getContract(uint256 tokenId) view returns((bytes32,bytes32,uint32,bytes32,uint32,address,bool), (bytes32,bytes32,uint128,uint64,uint64))
func (_Bindings *BindingsCallerSession) GetContract(tokenId *big.Int) (VFinanceRegistryContractRecord, VFinanceRegistryVehicleCore, error) {
	return _Bindings.Contract.GetContract(&_Bindings.CallOpts, tokenId)
}

// GetContractByChassis is a free data retrieval call binding the contract method 0xd54f9a74.
//
// Solidity: function getContractByChassis(string chassis) view returns((bytes32,bytes32,uint32,bytes32,uint32,address,bool), (bytes32,bytes32,uint128,uint64,uint64))
func (_Bindings *BindingsCaller) GetContractByChassis(opts *bind.CallOpts, chassis string) (VFinanceRegistryContractRecord, VFinanceRegistryVehicleCore, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getContractByChassis", chassis)

	if err != nil {
		return *new(VFinanceRegistryContractRecord), *new(VFinanceRegistryVehicleCore), err
	}

	out0 := *abi.ConvertType(out[0], new(VFinanceRegistryContractRecord)).(*VFinanceRegistryContractRecord)
	out1 := *abi.ConvertType(out[1], new(VFinanceRegistryVehicleCore)).(*VFinanceRegistryVehicleCore)

	return out0, out1, err

}

// GetContractByChassis is a free data retrieval call binding the contract method 0xd54f9a74.
//
// Solidity: function getContractByChassis(string chassis) view returns((bytes32,bytes32,uint32,bytes32,uint32,address,bool), (bytes32,bytes32,uint128,uint64,uint64))
func (_Bindings *BindingsSession) GetContractByChassis(chassis string) (VFinanceRegistryContractRecord, VFinanceRegistryVehicleCore, error) {
	return _Bindings.Contract.GetContractByChassis(&_Bindings.CallOpts, chassis)
}

// GetContractByChassis is a free data retrieval call binding the contract method 0xd54f9a74.
//
// Solidity: function getContractByChassis(string chassis) view returns((bytes32,bytes32,uint32,bytes32,uint32,address,bool), (bytes32,bytes32,uint128,uint64,uint64))
func (_Bindings *BindingsCallerSession) GetContractByChassis(chassis string) (VFinanceRegistryContractRecord, VFinanceRegistryVehicleCore, error) {
	return _Bindings.Contract.GetContractByChassis(&_Bindings.CallOpts, chassis)
}

// GetContractByHash is a free data retrieval call binding the contract method 0x6933a44c.
//
// Solidity: function getContractByHash(bytes32 metadataHash) view returns((bytes32,bytes32,uint32,bytes32,uint32,address,bool), (bytes32,bytes32,uint128,uint64,uint64))
func (_Bindings *BindingsCaller) GetContractByHash(opts *bind.CallOpts, metadataHash [32]byte) (VFinanceRegistryContractRecord, VFinanceRegistryVehicleCore, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getContractByHash", metadataHash)

	if err != nil {
		return *new(VFinanceRegistryContractRecord), *new(VFinanceRegistryVehicleCore), err
	}

	out0 := *abi.ConvertType(out[0], new(VFinanceRegistryContractRecord)).(*VFinanceRegistryContractRecord)
	out1 := *abi.ConvertType(out[1], new(VFinanceRegistryVehicleCore)).(*VFinanceRegistryVehicleCore)

	return out0, out1, err

}

// GetContractByHash is a free data retrieval call binding the contract method 0x6933a44c.
//
// Solidity: function getContractByHash(bytes32 metadataHash) view returns((bytes32,bytes32,uint32,bytes32,uint32,address,bool), (bytes32,bytes32,uint128,uint64,uint64))
func (_Bindings *BindingsSession) GetContractByHash(metadataHash [32]byte) (VFinanceRegistryContractRecord, VFinanceRegistryVehicleCore, error) {
	return _Bindings.Contract.GetContractByHash(&_Bindings.CallOpts, metadataHash)
}

// GetContractByHash is a free data retrieval call binding the contract method 0x6933a44c.
//
// Solidity: function getContractByHash(bytes32 metadataHash) view returns((bytes32,bytes32,uint32,bytes32,uint32,address,bool), (bytes32,bytes32,uint128,uint64,uint64))
func (_Bindings *BindingsCallerSession) GetContractByHash(metadataHash [32]byte) (VFinanceRegistryContractRecord, VFinanceRegistryVehicleCore, error) {
	return _Bindings.Contract.GetContractByHash(&_Bindings.CallOpts, metadataHash)
}

// GetContractByRegistryId is a free data retrieval call binding the contract method 0x456830ff.
//
// Solidity: function getContractByRegistryId(string registryId) view returns((bytes32,bytes32,uint32,bytes32,uint32,address,bool), (bytes32,bytes32,uint128,uint64,uint64))
func (_Bindings *BindingsCaller) GetContractByRegistryId(opts *bind.CallOpts, registryId string) (VFinanceRegistryContractRecord, VFinanceRegistryVehicleCore, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getContractByRegistryId", registryId)

	if err != nil {
		return *new(VFinanceRegistryContractRecord), *new(VFinanceRegistryVehicleCore), err
	}

	out0 := *abi.ConvertType(out[0], new(VFinanceRegistryContractRecord)).(*VFinanceRegistryContractRecord)
	out1 := *abi.ConvertType(out[1], new(VFinanceRegistryVehicleCore)).(*VFinanceRegistryVehicleCore)

	return out0, out1, err

}

// GetContractByRegistryId is a free data retrieval call binding the contract method 0x456830ff.
//
// Solidity: function getContractByRegistryId(string registryId) view returns((bytes32,bytes32,uint32,bytes32,uint32,address,bool), (bytes32,bytes32,uint128,uint64,uint64))
func (_Bindings *BindingsSession) GetContractByRegistryId(registryId string) (VFinanceRegistryContractRecord, VFinanceRegistryVehicleCore, error) {
	return _Bindings.Contract.GetContractByRegistryId(&_Bindings.CallOpts, registryId)
}

// GetContractByRegistryId is a free data retrieval call binding the contract method 0x456830ff.
//
// Solidity: function getContractByRegistryId(string registryId) view returns((bytes32,bytes32,uint32,bytes32,uint32,address,bool), (bytes32,bytes32,uint128,uint64,uint64))
func (_Bindings *BindingsCallerSession) GetContractByRegistryId(registryId string) (VFinanceRegistryContractRecord, VFinanceRegistryVehicleCore, error) {
	return _Bindings.Contract.GetContractByRegistryId(&_Bindings.CallOpts, registryId)
}

// GetMetadataUrl is a free data retrieval call binding the contract method 0xec4358ea.
//
// Solidity: function getMetadataUrl(bytes32 metadataHash) view returns(string)
func (_Bindings *BindingsCaller) GetMetadataUrl(opts *bind.CallOpts, metadataHash [32]byte) (string, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getMetadataUrl", metadataHash)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetMetadataUrl is a free data retrieval call binding the contract method 0xec4358ea.
//
// Solidity: function getMetadataUrl(bytes32 metadataHash) view returns(string)
func (_Bindings *BindingsSession) GetMetadataUrl(metadataHash [32]byte) (string, error) {
	return _Bindings.Contract.GetMetadataUrl(&_Bindings.CallOpts, metadataHash)
}

// GetMetadataUrl is a free data retrieval call binding the contract method 0xec4358ea.
//
// Solidity: function getMetadataUrl(bytes32 metadataHash) view returns(string)
func (_Bindings *BindingsCallerSession) GetMetadataUrl(metadataHash [32]byte) (string, error) {
	return _Bindings.Contract.GetMetadataUrl(&_Bindings.CallOpts, metadataHash)
}

// GetMetadataUrlByRegistryId is a free data retrieval call binding the contract method 0xeed47f2e.
//
// Solidity: function getMetadataUrlByRegistryId(string registryId) view returns(string)
func (_Bindings *BindingsCaller) GetMetadataUrlByRegistryId(opts *bind.CallOpts, registryId string) (string, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getMetadataUrlByRegistryId", registryId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetMetadataUrlByRegistryId is a free data retrieval call binding the contract method 0xeed47f2e.
//
// Solidity: function getMetadataUrlByRegistryId(string registryId) view returns(string)
func (_Bindings *BindingsSession) GetMetadataUrlByRegistryId(registryId string) (string, error) {
	return _Bindings.Contract.GetMetadataUrlByRegistryId(&_Bindings.CallOpts, registryId)
}

// GetMetadataUrlByRegistryId is a free data retrieval call binding the contract method 0xeed47f2e.
//
// Solidity: function getMetadataUrlByRegistryId(string registryId) view returns(string)
func (_Bindings *BindingsCallerSession) GetMetadataUrlByRegistryId(registryId string) (string, error) {
	return _Bindings.Contract.GetMetadataUrlByRegistryId(&_Bindings.CallOpts, registryId)
}

// GetModelName is a free data retrieval call binding the contract method 0xddb3069c.
//
// Solidity: function getModelName(uint64 modelId) view returns(string)
func (_Bindings *BindingsCaller) GetModelName(opts *bind.CallOpts, modelId uint64) (string, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getModelName", modelId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetModelName is a free data retrieval call binding the contract method 0xddb3069c.
//
// Solidity: function getModelName(uint64 modelId) view returns(string)
func (_Bindings *BindingsSession) GetModelName(modelId uint64) (string, error) {
	return _Bindings.Contract.GetModelName(&_Bindings.CallOpts, modelId)
}

// GetModelName is a free data retrieval call binding the contract method 0xddb3069c.
//
// Solidity: function getModelName(uint64 modelId) view returns(string)
func (_Bindings *BindingsCallerSession) GetModelName(modelId uint64) (string, error) {
	return _Bindings.Contract.GetModelName(&_Bindings.CallOpts, modelId)
}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() pure returns(string)
func (_Bindings *BindingsCaller) GetVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() pure returns(string)
func (_Bindings *BindingsSession) GetVersion() (string, error) {
	return _Bindings.Contract.GetVersion(&_Bindings.CallOpts)
}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() pure returns(string)
func (_Bindings *BindingsCallerSession) GetVersion() (string, error) {
	return _Bindings.Contract.GetVersion(&_Bindings.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Bindings *BindingsCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Bindings *BindingsSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Bindings.Contract.IsApprovedForAll(&_Bindings.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Bindings *BindingsCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Bindings.Contract.IsApprovedForAll(&_Bindings.CallOpts, owner, operator)
}

// MetadataBaseUrl is a free data retrieval call binding the contract method 0x49f3f1a6.
//
// Solidity: function metadataBaseUrl() view returns(string)
func (_Bindings *BindingsCaller) MetadataBaseUrl(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "metadataBaseUrl")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// MetadataBaseUrl is a free data retrieval call binding the contract method 0x49f3f1a6.
//
// Solidity: function metadataBaseUrl() view returns(string)
func (_Bindings *BindingsSession) MetadataBaseUrl() (string, error) {
	return _Bindings.Contract.MetadataBaseUrl(&_Bindings.CallOpts)
}

// MetadataBaseUrl is a free data retrieval call binding the contract method 0x49f3f1a6.
//
// Solidity: function metadataBaseUrl() view returns(string)
func (_Bindings *BindingsCallerSession) MetadataBaseUrl() (string, error) {
	return _Bindings.Contract.MetadataBaseUrl(&_Bindings.CallOpts)
}

// MetadataHashToTokenId is a free data retrieval call binding the contract method 0xa69e14b1.
//
// Solidity: function metadataHashToTokenId(bytes32 ) view returns(uint256)
func (_Bindings *BindingsCaller) MetadataHashToTokenId(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "metadataHashToTokenId", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MetadataHashToTokenId is a free data retrieval call binding the contract method 0xa69e14b1.
//
// Solidity: function metadataHashToTokenId(bytes32 ) view returns(uint256)
func (_Bindings *BindingsSession) MetadataHashToTokenId(arg0 [32]byte) (*big.Int, error) {
	return _Bindings.Contract.MetadataHashToTokenId(&_Bindings.CallOpts, arg0)
}

// MetadataHashToTokenId is a free data retrieval call binding the contract method 0xa69e14b1.
//
// Solidity: function metadataHashToTokenId(bytes32 ) view returns(uint256)
func (_Bindings *BindingsCallerSession) MetadataHashToTokenId(arg0 [32]byte) (*big.Int, error) {
	return _Bindings.Contract.MetadataHashToTokenId(&_Bindings.CallOpts, arg0)
}

// Models is a free data retrieval call binding the contract method 0x861d297f.
//
// Solidity: function models(uint64 ) view returns(string)
func (_Bindings *BindingsCaller) Models(opts *bind.CallOpts, arg0 uint64) (string, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "models", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Models is a free data retrieval call binding the contract method 0x861d297f.
//
// Solidity: function models(uint64 ) view returns(string)
func (_Bindings *BindingsSession) Models(arg0 uint64) (string, error) {
	return _Bindings.Contract.Models(&_Bindings.CallOpts, arg0)
}

// Models is a free data retrieval call binding the contract method 0x861d297f.
//
// Solidity: function models(uint64 ) view returns(string)
func (_Bindings *BindingsCallerSession) Models(arg0 uint64) (string, error) {
	return _Bindings.Contract.Models(&_Bindings.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Bindings *BindingsCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Bindings *BindingsSession) Name() (string, error) {
	return _Bindings.Contract.Name(&_Bindings.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Bindings *BindingsCallerSession) Name() (string, error) {
	return _Bindings.Contract.Name(&_Bindings.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bindings *BindingsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bindings *BindingsSession) Owner() (common.Address, error) {
	return _Bindings.Contract.Owner(&_Bindings.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bindings *BindingsCallerSession) Owner() (common.Address, error) {
	return _Bindings.Contract.Owner(&_Bindings.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Bindings *BindingsCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Bindings *BindingsSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Bindings.Contract.OwnerOf(&_Bindings.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Bindings *BindingsCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Bindings.Contract.OwnerOf(&_Bindings.CallOpts, tokenId)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Bindings *BindingsCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Bindings *BindingsSession) ProxiableUUID() ([32]byte, error) {
	return _Bindings.Contract.ProxiableUUID(&_Bindings.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Bindings *BindingsCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Bindings.Contract.ProxiableUUID(&_Bindings.CallOpts)
}

// RegistryIdHashToTokenId is a free data retrieval call binding the contract method 0xfffc02bc.
//
// Solidity: function registryIdHashToTokenId(bytes32 ) view returns(uint256)
func (_Bindings *BindingsCaller) RegistryIdHashToTokenId(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "registryIdHashToTokenId", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RegistryIdHashToTokenId is a free data retrieval call binding the contract method 0xfffc02bc.
//
// Solidity: function registryIdHashToTokenId(bytes32 ) view returns(uint256)
func (_Bindings *BindingsSession) RegistryIdHashToTokenId(arg0 [32]byte) (*big.Int, error) {
	return _Bindings.Contract.RegistryIdHashToTokenId(&_Bindings.CallOpts, arg0)
}

// RegistryIdHashToTokenId is a free data retrieval call binding the contract method 0xfffc02bc.
//
// Solidity: function registryIdHashToTokenId(bytes32 ) view returns(uint256)
func (_Bindings *BindingsCallerSession) RegistryIdHashToTokenId(arg0 [32]byte) (*big.Int, error) {
	return _Bindings.Contract.RegistryIdHashToTokenId(&_Bindings.CallOpts, arg0)
}

// SetApprovalForAll is a free data retrieval call binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address , bool ) pure returns()
func (_Bindings *BindingsCaller) SetApprovalForAll(opts *bind.CallOpts, arg0 common.Address, arg1 bool) error {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "setApprovalForAll", arg0, arg1)

	if err != nil {
		return err
	}

	return err

}

// SetApprovalForAll is a free data retrieval call binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address , bool ) pure returns()
func (_Bindings *BindingsSession) SetApprovalForAll(arg0 common.Address, arg1 bool) error {
	return _Bindings.Contract.SetApprovalForAll(&_Bindings.CallOpts, arg0, arg1)
}

// SetApprovalForAll is a free data retrieval call binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address , bool ) pure returns()
func (_Bindings *BindingsCallerSession) SetApprovalForAll(arg0 common.Address, arg1 bool) error {
	return _Bindings.Contract.SetApprovalForAll(&_Bindings.CallOpts, arg0, arg1)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Bindings *BindingsCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Bindings *BindingsSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Bindings.Contract.SupportsInterface(&_Bindings.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Bindings *BindingsCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Bindings.Contract.SupportsInterface(&_Bindings.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Bindings *BindingsCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Bindings *BindingsSession) Symbol() (string, error) {
	return _Bindings.Contract.Symbol(&_Bindings.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Bindings *BindingsCallerSession) Symbol() (string, error) {
	return _Bindings.Contract.Symbol(&_Bindings.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Bindings *BindingsCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Bindings *BindingsSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Bindings.Contract.TokenURI(&_Bindings.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Bindings *BindingsCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Bindings.Contract.TokenURI(&_Bindings.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Bindings *BindingsCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Bindings *BindingsSession) TotalSupply() (*big.Int, error) {
	return _Bindings.Contract.TotalSupply(&_Bindings.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Bindings *BindingsCallerSession) TotalSupply() (*big.Int, error) {
	return _Bindings.Contract.TotalSupply(&_Bindings.CallOpts)
}

// VehicleCores is a free data retrieval call binding the contract method 0x65950af4.
//
// Solidity: function vehicleCores(uint256 ) view returns(bytes32 chassis, bytes32 licensePlate, uint128 totalValue, uint64 brandId, uint64 modelId)
func (_Bindings *BindingsCaller) VehicleCores(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Chassis      [32]byte
	LicensePlate [32]byte
	TotalValue   *big.Int
	BrandId      uint64
	ModelId      uint64
}, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "vehicleCores", arg0)

	outstruct := new(struct {
		Chassis      [32]byte
		LicensePlate [32]byte
		TotalValue   *big.Int
		BrandId      uint64
		ModelId      uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Chassis = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.LicensePlate = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.TotalValue = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.BrandId = *abi.ConvertType(out[3], new(uint64)).(*uint64)
	outstruct.ModelId = *abi.ConvertType(out[4], new(uint64)).(*uint64)

	return *outstruct, err

}

// VehicleCores is a free data retrieval call binding the contract method 0x65950af4.
//
// Solidity: function vehicleCores(uint256 ) view returns(bytes32 chassis, bytes32 licensePlate, uint128 totalValue, uint64 brandId, uint64 modelId)
func (_Bindings *BindingsSession) VehicleCores(arg0 *big.Int) (struct {
	Chassis      [32]byte
	LicensePlate [32]byte
	TotalValue   *big.Int
	BrandId      uint64
	ModelId      uint64
}, error) {
	return _Bindings.Contract.VehicleCores(&_Bindings.CallOpts, arg0)
}

// VehicleCores is a free data retrieval call binding the contract method 0x65950af4.
//
// Solidity: function vehicleCores(uint256 ) view returns(bytes32 chassis, bytes32 licensePlate, uint128 totalValue, uint64 brandId, uint64 modelId)
func (_Bindings *BindingsCallerSession) VehicleCores(arg0 *big.Int) (struct {
	Chassis      [32]byte
	LicensePlate [32]byte
	TotalValue   *big.Int
	BrandId      uint64
	ModelId      uint64
}, error) {
	return _Bindings.Contract.VehicleCores(&_Bindings.CallOpts, arg0)
}

// EmergencyWithdrawToken is a paid mutator transaction binding the contract method 0xa4c3b091.
//
// Solidity: function emergencyWithdrawToken(address token, uint256 amount) returns()
func (_Bindings *BindingsTransactor) EmergencyWithdrawToken(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "emergencyWithdrawToken", token, amount)
}

// EmergencyWithdrawToken is a paid mutator transaction binding the contract method 0xa4c3b091.
//
// Solidity: function emergencyWithdrawToken(address token, uint256 amount) returns()
func (_Bindings *BindingsSession) EmergencyWithdrawToken(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.EmergencyWithdrawToken(&_Bindings.TransactOpts, token, amount)
}

// EmergencyWithdrawToken is a paid mutator transaction binding the contract method 0xa4c3b091.
//
// Solidity: function emergencyWithdrawToken(address token, uint256 amount) returns()
func (_Bindings *BindingsTransactorSession) EmergencyWithdrawToken(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.EmergencyWithdrawToken(&_Bindings.TransactOpts, token, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0x5c6d8da1.
//
// Solidity: function initialize(string name, string symbol, string _metadataBaseUrl, address _apiServerAddress) returns()
func (_Bindings *BindingsTransactor) Initialize(opts *bind.TransactOpts, name string, symbol string, _metadataBaseUrl string, _apiServerAddress common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "initialize", name, symbol, _metadataBaseUrl, _apiServerAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x5c6d8da1.
//
// Solidity: function initialize(string name, string symbol, string _metadataBaseUrl, address _apiServerAddress) returns()
func (_Bindings *BindingsSession) Initialize(name string, symbol string, _metadataBaseUrl string, _apiServerAddress common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.Initialize(&_Bindings.TransactOpts, name, symbol, _metadataBaseUrl, _apiServerAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x5c6d8da1.
//
// Solidity: function initialize(string name, string symbol, string _metadataBaseUrl, address _apiServerAddress) returns()
func (_Bindings *BindingsTransactorSession) Initialize(name string, symbol string, _metadataBaseUrl string, _apiServerAddress common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.Initialize(&_Bindings.TransactOpts, name, symbol, _metadataBaseUrl, _apiServerAddress)
}

// RegisterBrand is a paid mutator transaction binding the contract method 0x9f9aab6f.
//
// Solidity: function registerBrand(string brandName) returns(uint64 brandId)
func (_Bindings *BindingsTransactor) RegisterBrand(opts *bind.TransactOpts, brandName string) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "registerBrand", brandName)
}

// RegisterBrand is a paid mutator transaction binding the contract method 0x9f9aab6f.
//
// Solidity: function registerBrand(string brandName) returns(uint64 brandId)
func (_Bindings *BindingsSession) RegisterBrand(brandName string) (*types.Transaction, error) {
	return _Bindings.Contract.RegisterBrand(&_Bindings.TransactOpts, brandName)
}

// RegisterBrand is a paid mutator transaction binding the contract method 0x9f9aab6f.
//
// Solidity: function registerBrand(string brandName) returns(uint64 brandId)
func (_Bindings *BindingsTransactorSession) RegisterBrand(brandName string) (*types.Transaction, error) {
	return _Bindings.Contract.RegisterBrand(&_Bindings.TransactOpts, brandName)
}

// RegisterContract is a paid mutator transaction binding the contract method 0xa35614d5.
//
// Solidity: function registerContract(string registryId, string contractNumber, uint32 contractDate, string chassis, string licensePlate, uint128 totalValue, string brandName, string modelName) returns(uint256 tokenId, bytes32 metadataHash)
func (_Bindings *BindingsTransactor) RegisterContract(opts *bind.TransactOpts, registryId string, contractNumber string, contractDate uint32, chassis string, licensePlate string, totalValue *big.Int, brandName string, modelName string) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "registerContract", registryId, contractNumber, contractDate, chassis, licensePlate, totalValue, brandName, modelName)
}

// RegisterContract is a paid mutator transaction binding the contract method 0xa35614d5.
//
// Solidity: function registerContract(string registryId, string contractNumber, uint32 contractDate, string chassis, string licensePlate, uint128 totalValue, string brandName, string modelName) returns(uint256 tokenId, bytes32 metadataHash)
func (_Bindings *BindingsSession) RegisterContract(registryId string, contractNumber string, contractDate uint32, chassis string, licensePlate string, totalValue *big.Int, brandName string, modelName string) (*types.Transaction, error) {
	return _Bindings.Contract.RegisterContract(&_Bindings.TransactOpts, registryId, contractNumber, contractDate, chassis, licensePlate, totalValue, brandName, modelName)
}

// RegisterContract is a paid mutator transaction binding the contract method 0xa35614d5.
//
// Solidity: function registerContract(string registryId, string contractNumber, uint32 contractDate, string chassis, string licensePlate, uint128 totalValue, string brandName, string modelName) returns(uint256 tokenId, bytes32 metadataHash)
func (_Bindings *BindingsTransactorSession) RegisterContract(registryId string, contractNumber string, contractDate uint32, chassis string, licensePlate string, totalValue *big.Int, brandName string, modelName string) (*types.Transaction, error) {
	return _Bindings.Contract.RegisterContract(&_Bindings.TransactOpts, registryId, contractNumber, contractDate, chassis, licensePlate, totalValue, brandName, modelName)
}

// RegisterModel is a paid mutator transaction binding the contract method 0x9dfdc7de.
//
// Solidity: function registerModel(string modelName) returns(uint64 modelId)
func (_Bindings *BindingsTransactor) RegisterModel(opts *bind.TransactOpts, modelName string) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "registerModel", modelName)
}

// RegisterModel is a paid mutator transaction binding the contract method 0x9dfdc7de.
//
// Solidity: function registerModel(string modelName) returns(uint64 modelId)
func (_Bindings *BindingsSession) RegisterModel(modelName string) (*types.Transaction, error) {
	return _Bindings.Contract.RegisterModel(&_Bindings.TransactOpts, modelName)
}

// RegisterModel is a paid mutator transaction binding the contract method 0x9dfdc7de.
//
// Solidity: function registerModel(string modelName) returns(uint64 modelId)
func (_Bindings *BindingsTransactorSession) RegisterModel(modelName string) (*types.Transaction, error) {
	return _Bindings.Contract.RegisterModel(&_Bindings.TransactOpts, modelName)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bindings *BindingsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bindings *BindingsSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bindings.Contract.RenounceOwnership(&_Bindings.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bindings *BindingsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bindings.Contract.RenounceOwnership(&_Bindings.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Bindings *BindingsTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Bindings *BindingsSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SafeTransferFrom(&_Bindings.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Bindings *BindingsTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SafeTransferFrom(&_Bindings.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Bindings *BindingsTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Bindings *BindingsSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Bindings.Contract.SafeTransferFrom0(&_Bindings.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Bindings *BindingsTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Bindings.Contract.SafeTransferFrom0(&_Bindings.TransactOpts, from, to, tokenId, data)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Bindings *BindingsTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Bindings *BindingsSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.TransferFrom(&_Bindings.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Bindings *BindingsTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.TransferFrom(&_Bindings.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bindings *BindingsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bindings *BindingsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.TransferOwnership(&_Bindings.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bindings *BindingsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.TransferOwnership(&_Bindings.TransactOpts, newOwner)
}

// UpdateMetadataHash is a paid mutator transaction binding the contract method 0x07f63700.
//
// Solidity: function updateMetadataHash(uint256 tokenId, bytes32 newMetadataHash) returns()
func (_Bindings *BindingsTransactor) UpdateMetadataHash(opts *bind.TransactOpts, tokenId *big.Int, newMetadataHash [32]byte) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "updateMetadataHash", tokenId, newMetadataHash)
}

// UpdateMetadataHash is a paid mutator transaction binding the contract method 0x07f63700.
//
// Solidity: function updateMetadataHash(uint256 tokenId, bytes32 newMetadataHash) returns()
func (_Bindings *BindingsSession) UpdateMetadataHash(tokenId *big.Int, newMetadataHash [32]byte) (*types.Transaction, error) {
	return _Bindings.Contract.UpdateMetadataHash(&_Bindings.TransactOpts, tokenId, newMetadataHash)
}

// UpdateMetadataHash is a paid mutator transaction binding the contract method 0x07f63700.
//
// Solidity: function updateMetadataHash(uint256 tokenId, bytes32 newMetadataHash) returns()
func (_Bindings *BindingsTransactorSession) UpdateMetadataHash(tokenId *big.Int, newMetadataHash [32]byte) (*types.Transaction, error) {
	return _Bindings.Contract.UpdateMetadataHash(&_Bindings.TransactOpts, tokenId, newMetadataHash)
}

// UpdateServerConfig is a paid mutator transaction binding the contract method 0x31d1cbda.
//
// Solidity: function updateServerConfig(string newMetadataBaseUrl, address newApiServerAddress) returns()
func (_Bindings *BindingsTransactor) UpdateServerConfig(opts *bind.TransactOpts, newMetadataBaseUrl string, newApiServerAddress common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "updateServerConfig", newMetadataBaseUrl, newApiServerAddress)
}

// UpdateServerConfig is a paid mutator transaction binding the contract method 0x31d1cbda.
//
// Solidity: function updateServerConfig(string newMetadataBaseUrl, address newApiServerAddress) returns()
func (_Bindings *BindingsSession) UpdateServerConfig(newMetadataBaseUrl string, newApiServerAddress common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.UpdateServerConfig(&_Bindings.TransactOpts, newMetadataBaseUrl, newApiServerAddress)
}

// UpdateServerConfig is a paid mutator transaction binding the contract method 0x31d1cbda.
//
// Solidity: function updateServerConfig(string newMetadataBaseUrl, address newApiServerAddress) returns()
func (_Bindings *BindingsTransactorSession) UpdateServerConfig(newMetadataBaseUrl string, newApiServerAddress common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.UpdateServerConfig(&_Bindings.TransactOpts, newMetadataBaseUrl, newApiServerAddress)
}

// UpdateStatus is a paid mutator transaction binding the contract method 0x387008c2.
//
// Solidity: function updateStatus(uint256 tokenId, bool active) returns()
func (_Bindings *BindingsTransactor) UpdateStatus(opts *bind.TransactOpts, tokenId *big.Int, active bool) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "updateStatus", tokenId, active)
}

// UpdateStatus is a paid mutator transaction binding the contract method 0x387008c2.
//
// Solidity: function updateStatus(uint256 tokenId, bool active) returns()
func (_Bindings *BindingsSession) UpdateStatus(tokenId *big.Int, active bool) (*types.Transaction, error) {
	return _Bindings.Contract.UpdateStatus(&_Bindings.TransactOpts, tokenId, active)
}

// UpdateStatus is a paid mutator transaction binding the contract method 0x387008c2.
//
// Solidity: function updateStatus(uint256 tokenId, bool active) returns()
func (_Bindings *BindingsTransactorSession) UpdateStatus(tokenId *big.Int, active bool) (*types.Transaction, error) {
	return _Bindings.Contract.UpdateStatus(&_Bindings.TransactOpts, tokenId, active)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Bindings *BindingsTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Bindings *BindingsSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Bindings.Contract.UpgradeToAndCall(&_Bindings.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Bindings *BindingsTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Bindings.Contract.UpgradeToAndCall(&_Bindings.TransactOpts, newImplementation, data)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Bindings *BindingsTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Bindings *BindingsSession) Receive() (*types.Transaction, error) {
	return _Bindings.Contract.Receive(&_Bindings.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Bindings *BindingsTransactorSession) Receive() (*types.Transaction, error) {
	return _Bindings.Contract.Receive(&_Bindings.TransactOpts)
}

// BindingsApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Bindings contract.
type BindingsApprovalIterator struct {
	Event *BindingsApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingsApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingsApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingsApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsApproval represents a Approval event raised by the Bindings contract.
type BindingsApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Bindings *BindingsFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*BindingsApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &BindingsApprovalIterator{contract: _Bindings.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Bindings *BindingsFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BindingsApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsApproval)
				if err := _Bindings.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Bindings *BindingsFilterer) ParseApproval(log types.Log) (*BindingsApproval, error) {
	event := new(BindingsApproval)
	if err := _Bindings.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Bindings contract.
type BindingsApprovalForAllIterator struct {
	Event *BindingsApprovalForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingsApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsApprovalForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingsApprovalForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingsApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsApprovalForAll represents a ApprovalForAll event raised by the Bindings contract.
type BindingsApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Bindings *BindingsFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*BindingsApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &BindingsApprovalForAllIterator{contract: _Bindings.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Bindings *BindingsFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *BindingsApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsApprovalForAll)
				if err := _Bindings.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Bindings *BindingsFilterer) ParseApprovalForAll(log types.Log) (*BindingsApprovalForAll, error) {
	event := new(BindingsApprovalForAll)
	if err := _Bindings.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsBrandRegisteredIterator is returned from FilterBrandRegistered and is used to iterate over the raw logs and unpacked data for BrandRegistered events raised by the Bindings contract.
type BindingsBrandRegisteredIterator struct {
	Event *BindingsBrandRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingsBrandRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsBrandRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingsBrandRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingsBrandRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsBrandRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsBrandRegistered represents a BrandRegistered event raised by the Bindings contract.
type BindingsBrandRegistered struct {
	BrandId   uint64
	BrandName string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBrandRegistered is a free log retrieval operation binding the contract event 0xd6a561f68236f1785e9c45009eba5ac6bfb0ae92da53a5dbc4995423886d320d.
//
// Solidity: event BrandRegistered(uint64 indexed brandId, string brandName)
func (_Bindings *BindingsFilterer) FilterBrandRegistered(opts *bind.FilterOpts, brandId []uint64) (*BindingsBrandRegisteredIterator, error) {

	var brandIdRule []interface{}
	for _, brandIdItem := range brandId {
		brandIdRule = append(brandIdRule, brandIdItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "BrandRegistered", brandIdRule)
	if err != nil {
		return nil, err
	}
	return &BindingsBrandRegisteredIterator{contract: _Bindings.contract, event: "BrandRegistered", logs: logs, sub: sub}, nil
}

// WatchBrandRegistered is a free log subscription operation binding the contract event 0xd6a561f68236f1785e9c45009eba5ac6bfb0ae92da53a5dbc4995423886d320d.
//
// Solidity: event BrandRegistered(uint64 indexed brandId, string brandName)
func (_Bindings *BindingsFilterer) WatchBrandRegistered(opts *bind.WatchOpts, sink chan<- *BindingsBrandRegistered, brandId []uint64) (event.Subscription, error) {

	var brandIdRule []interface{}
	for _, brandIdItem := range brandId {
		brandIdRule = append(brandIdRule, brandIdItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "BrandRegistered", brandIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsBrandRegistered)
				if err := _Bindings.contract.UnpackLog(event, "BrandRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBrandRegistered is a log parse operation binding the contract event 0xd6a561f68236f1785e9c45009eba5ac6bfb0ae92da53a5dbc4995423886d320d.
//
// Solidity: event BrandRegistered(uint64 indexed brandId, string brandName)
func (_Bindings *BindingsFilterer) ParseBrandRegistered(log types.Log) (*BindingsBrandRegistered, error) {
	event := new(BindingsBrandRegistered)
	if err := _Bindings.contract.UnpackLog(event, "BrandRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsContractRegisteredIterator is returned from FilterContractRegistered and is used to iterate over the raw logs and unpacked data for ContractRegistered events raised by the Bindings contract.
type BindingsContractRegisteredIterator struct {
	Event *BindingsContractRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingsContractRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsContractRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingsContractRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingsContractRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsContractRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsContractRegistered represents a ContractRegistered event raised by the Bindings contract.
type BindingsContractRegistered struct {
	TokenId        *big.Int
	RegistryIdHash [32]byte
	ChassisHash    [32]byte
	MetadataHash   [32]byte
	Timestamp      uint32
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterContractRegistered is a free log retrieval operation binding the contract event 0x08763d7ed3c1fd5bc57e914d68af9baaed2ab99ca4d797aca811079add788b92.
//
// Solidity: event ContractRegistered(uint256 indexed tokenId, bytes32 indexed registryIdHash, bytes32 indexed chassisHash, bytes32 metadataHash, uint32 timestamp)
func (_Bindings *BindingsFilterer) FilterContractRegistered(opts *bind.FilterOpts, tokenId []*big.Int, registryIdHash [][32]byte, chassisHash [][32]byte) (*BindingsContractRegisteredIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var registryIdHashRule []interface{}
	for _, registryIdHashItem := range registryIdHash {
		registryIdHashRule = append(registryIdHashRule, registryIdHashItem)
	}
	var chassisHashRule []interface{}
	for _, chassisHashItem := range chassisHash {
		chassisHashRule = append(chassisHashRule, chassisHashItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "ContractRegistered", tokenIdRule, registryIdHashRule, chassisHashRule)
	if err != nil {
		return nil, err
	}
	return &BindingsContractRegisteredIterator{contract: _Bindings.contract, event: "ContractRegistered", logs: logs, sub: sub}, nil
}

// WatchContractRegistered is a free log subscription operation binding the contract event 0x08763d7ed3c1fd5bc57e914d68af9baaed2ab99ca4d797aca811079add788b92.
//
// Solidity: event ContractRegistered(uint256 indexed tokenId, bytes32 indexed registryIdHash, bytes32 indexed chassisHash, bytes32 metadataHash, uint32 timestamp)
func (_Bindings *BindingsFilterer) WatchContractRegistered(opts *bind.WatchOpts, sink chan<- *BindingsContractRegistered, tokenId []*big.Int, registryIdHash [][32]byte, chassisHash [][32]byte) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var registryIdHashRule []interface{}
	for _, registryIdHashItem := range registryIdHash {
		registryIdHashRule = append(registryIdHashRule, registryIdHashItem)
	}
	var chassisHashRule []interface{}
	for _, chassisHashItem := range chassisHash {
		chassisHashRule = append(chassisHashRule, chassisHashItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "ContractRegistered", tokenIdRule, registryIdHashRule, chassisHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsContractRegistered)
				if err := _Bindings.contract.UnpackLog(event, "ContractRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseContractRegistered is a log parse operation binding the contract event 0x08763d7ed3c1fd5bc57e914d68af9baaed2ab99ca4d797aca811079add788b92.
//
// Solidity: event ContractRegistered(uint256 indexed tokenId, bytes32 indexed registryIdHash, bytes32 indexed chassisHash, bytes32 metadataHash, uint32 timestamp)
func (_Bindings *BindingsFilterer) ParseContractRegistered(log types.Log) (*BindingsContractRegistered, error) {
	event := new(BindingsContractRegistered)
	if err := _Bindings.contract.UnpackLog(event, "ContractRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Bindings contract.
type BindingsInitializedIterator struct {
	Event *BindingsInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingsInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingsInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingsInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsInitialized represents a Initialized event raised by the Bindings contract.
type BindingsInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Bindings *BindingsFilterer) FilterInitialized(opts *bind.FilterOpts) (*BindingsInitializedIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BindingsInitializedIterator{contract: _Bindings.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Bindings *BindingsFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BindingsInitialized) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsInitialized)
				if err := _Bindings.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Bindings *BindingsFilterer) ParseInitialized(log types.Log) (*BindingsInitialized, error) {
	event := new(BindingsInitialized)
	if err := _Bindings.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsMetadataUpdatedIterator is returned from FilterMetadataUpdated and is used to iterate over the raw logs and unpacked data for MetadataUpdated events raised by the Bindings contract.
type BindingsMetadataUpdatedIterator struct {
	Event *BindingsMetadataUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingsMetadataUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsMetadataUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingsMetadataUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingsMetadataUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsMetadataUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsMetadataUpdated represents a MetadataUpdated event raised by the Bindings contract.
type BindingsMetadataUpdated struct {
	TokenId   *big.Int
	OldHash   [32]byte
	NewHash   [32]byte
	Timestamp uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMetadataUpdated is a free log retrieval operation binding the contract event 0x886980e9630c81aec383adb83bdf9ec93f47b6a360a6e4f9ad631813ea336c5c.
//
// Solidity: event MetadataUpdated(uint256 indexed tokenId, bytes32 indexed oldHash, bytes32 indexed newHash, uint32 timestamp)
func (_Bindings *BindingsFilterer) FilterMetadataUpdated(opts *bind.FilterOpts, tokenId []*big.Int, oldHash [][32]byte, newHash [][32]byte) (*BindingsMetadataUpdatedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var oldHashRule []interface{}
	for _, oldHashItem := range oldHash {
		oldHashRule = append(oldHashRule, oldHashItem)
	}
	var newHashRule []interface{}
	for _, newHashItem := range newHash {
		newHashRule = append(newHashRule, newHashItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "MetadataUpdated", tokenIdRule, oldHashRule, newHashRule)
	if err != nil {
		return nil, err
	}
	return &BindingsMetadataUpdatedIterator{contract: _Bindings.contract, event: "MetadataUpdated", logs: logs, sub: sub}, nil
}

// WatchMetadataUpdated is a free log subscription operation binding the contract event 0x886980e9630c81aec383adb83bdf9ec93f47b6a360a6e4f9ad631813ea336c5c.
//
// Solidity: event MetadataUpdated(uint256 indexed tokenId, bytes32 indexed oldHash, bytes32 indexed newHash, uint32 timestamp)
func (_Bindings *BindingsFilterer) WatchMetadataUpdated(opts *bind.WatchOpts, sink chan<- *BindingsMetadataUpdated, tokenId []*big.Int, oldHash [][32]byte, newHash [][32]byte) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var oldHashRule []interface{}
	for _, oldHashItem := range oldHash {
		oldHashRule = append(oldHashRule, oldHashItem)
	}
	var newHashRule []interface{}
	for _, newHashItem := range newHash {
		newHashRule = append(newHashRule, newHashItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "MetadataUpdated", tokenIdRule, oldHashRule, newHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsMetadataUpdated)
				if err := _Bindings.contract.UnpackLog(event, "MetadataUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMetadataUpdated is a log parse operation binding the contract event 0x886980e9630c81aec383adb83bdf9ec93f47b6a360a6e4f9ad631813ea336c5c.
//
// Solidity: event MetadataUpdated(uint256 indexed tokenId, bytes32 indexed oldHash, bytes32 indexed newHash, uint32 timestamp)
func (_Bindings *BindingsFilterer) ParseMetadataUpdated(log types.Log) (*BindingsMetadataUpdated, error) {
	event := new(BindingsMetadataUpdated)
	if err := _Bindings.contract.UnpackLog(event, "MetadataUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsModelRegisteredIterator is returned from FilterModelRegistered and is used to iterate over the raw logs and unpacked data for ModelRegistered events raised by the Bindings contract.
type BindingsModelRegisteredIterator struct {
	Event *BindingsModelRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingsModelRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsModelRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingsModelRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingsModelRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsModelRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsModelRegistered represents a ModelRegistered event raised by the Bindings contract.
type BindingsModelRegistered struct {
	ModelId   uint64
	ModelName string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterModelRegistered is a free log retrieval operation binding the contract event 0x67a7beac1d9610bc204daae1a6b821d21024a7ab7a5a5984f2083572c713a961.
//
// Solidity: event ModelRegistered(uint64 indexed modelId, string modelName)
func (_Bindings *BindingsFilterer) FilterModelRegistered(opts *bind.FilterOpts, modelId []uint64) (*BindingsModelRegisteredIterator, error) {

	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "ModelRegistered", modelIdRule)
	if err != nil {
		return nil, err
	}
	return &BindingsModelRegisteredIterator{contract: _Bindings.contract, event: "ModelRegistered", logs: logs, sub: sub}, nil
}

// WatchModelRegistered is a free log subscription operation binding the contract event 0x67a7beac1d9610bc204daae1a6b821d21024a7ab7a5a5984f2083572c713a961.
//
// Solidity: event ModelRegistered(uint64 indexed modelId, string modelName)
func (_Bindings *BindingsFilterer) WatchModelRegistered(opts *bind.WatchOpts, sink chan<- *BindingsModelRegistered, modelId []uint64) (event.Subscription, error) {

	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "ModelRegistered", modelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsModelRegistered)
				if err := _Bindings.contract.UnpackLog(event, "ModelRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseModelRegistered is a log parse operation binding the contract event 0x67a7beac1d9610bc204daae1a6b821d21024a7ab7a5a5984f2083572c713a961.
//
// Solidity: event ModelRegistered(uint64 indexed modelId, string modelName)
func (_Bindings *BindingsFilterer) ParseModelRegistered(log types.Log) (*BindingsModelRegistered, error) {
	event := new(BindingsModelRegistered)
	if err := _Bindings.contract.UnpackLog(event, "ModelRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Bindings contract.
type BindingsOwnershipTransferredIterator struct {
	Event *BindingsOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingsOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsOwnershipTransferred represents a OwnershipTransferred event raised by the Bindings contract.
type BindingsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bindings *BindingsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BindingsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BindingsOwnershipTransferredIterator{contract: _Bindings.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bindings *BindingsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BindingsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsOwnershipTransferred)
				if err := _Bindings.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bindings *BindingsFilterer) ParseOwnershipTransferred(log types.Log) (*BindingsOwnershipTransferred, error) {
	event := new(BindingsOwnershipTransferred)
	if err := _Bindings.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsStatusChangedIterator is returned from FilterStatusChanged and is used to iterate over the raw logs and unpacked data for StatusChanged events raised by the Bindings contract.
type BindingsStatusChangedIterator struct {
	Event *BindingsStatusChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingsStatusChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsStatusChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingsStatusChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingsStatusChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsStatusChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsStatusChanged represents a StatusChanged event raised by the Bindings contract.
type BindingsStatusChanged struct {
	TokenId   *big.Int
	Active    bool
	Timestamp uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStatusChanged is a free log retrieval operation binding the contract event 0xf3af971fc7f1836a845935ae73f7a6d6c219a218fad3a750456741df99e05781.
//
// Solidity: event StatusChanged(uint256 indexed tokenId, bool indexed active, uint32 timestamp)
func (_Bindings *BindingsFilterer) FilterStatusChanged(opts *bind.FilterOpts, tokenId []*big.Int, active []bool) (*BindingsStatusChangedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var activeRule []interface{}
	for _, activeItem := range active {
		activeRule = append(activeRule, activeItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "StatusChanged", tokenIdRule, activeRule)
	if err != nil {
		return nil, err
	}
	return &BindingsStatusChangedIterator{contract: _Bindings.contract, event: "StatusChanged", logs: logs, sub: sub}, nil
}

// WatchStatusChanged is a free log subscription operation binding the contract event 0xf3af971fc7f1836a845935ae73f7a6d6c219a218fad3a750456741df99e05781.
//
// Solidity: event StatusChanged(uint256 indexed tokenId, bool indexed active, uint32 timestamp)
func (_Bindings *BindingsFilterer) WatchStatusChanged(opts *bind.WatchOpts, sink chan<- *BindingsStatusChanged, tokenId []*big.Int, active []bool) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var activeRule []interface{}
	for _, activeItem := range active {
		activeRule = append(activeRule, activeItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "StatusChanged", tokenIdRule, activeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsStatusChanged)
				if err := _Bindings.contract.UnpackLog(event, "StatusChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStatusChanged is a log parse operation binding the contract event 0xf3af971fc7f1836a845935ae73f7a6d6c219a218fad3a750456741df99e05781.
//
// Solidity: event StatusChanged(uint256 indexed tokenId, bool indexed active, uint32 timestamp)
func (_Bindings *BindingsFilterer) ParseStatusChanged(log types.Log) (*BindingsStatusChanged, error) {
	event := new(BindingsStatusChanged)
	if err := _Bindings.contract.UnpackLog(event, "StatusChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsSystemConfiguredIterator is returned from FilterSystemConfigured and is used to iterate over the raw logs and unpacked data for SystemConfigured events raised by the Bindings contract.
type BindingsSystemConfiguredIterator struct {
	Event *BindingsSystemConfigured // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingsSystemConfiguredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsSystemConfigured)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingsSystemConfigured)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingsSystemConfiguredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsSystemConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsSystemConfigured represents a SystemConfigured event raised by the Bindings contract.
type BindingsSystemConfigured struct {
	ApiServer   common.Address
	MetadataUrl string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSystemConfigured is a free log retrieval operation binding the contract event 0xfc545c0151b00b9a0ac4d74abb5522892064539eb246b848b96edccd3302e4be.
//
// Solidity: event SystemConfigured(address indexed apiServer, string metadataUrl)
func (_Bindings *BindingsFilterer) FilterSystemConfigured(opts *bind.FilterOpts, apiServer []common.Address) (*BindingsSystemConfiguredIterator, error) {

	var apiServerRule []interface{}
	for _, apiServerItem := range apiServer {
		apiServerRule = append(apiServerRule, apiServerItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "SystemConfigured", apiServerRule)
	if err != nil {
		return nil, err
	}
	return &BindingsSystemConfiguredIterator{contract: _Bindings.contract, event: "SystemConfigured", logs: logs, sub: sub}, nil
}

// WatchSystemConfigured is a free log subscription operation binding the contract event 0xfc545c0151b00b9a0ac4d74abb5522892064539eb246b848b96edccd3302e4be.
//
// Solidity: event SystemConfigured(address indexed apiServer, string metadataUrl)
func (_Bindings *BindingsFilterer) WatchSystemConfigured(opts *bind.WatchOpts, sink chan<- *BindingsSystemConfigured, apiServer []common.Address) (event.Subscription, error) {

	var apiServerRule []interface{}
	for _, apiServerItem := range apiServer {
		apiServerRule = append(apiServerRule, apiServerItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "SystemConfigured", apiServerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsSystemConfigured)
				if err := _Bindings.contract.UnpackLog(event, "SystemConfigured", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSystemConfigured is a log parse operation binding the contract event 0xfc545c0151b00b9a0ac4d74abb5522892064539eb246b848b96edccd3302e4be.
//
// Solidity: event SystemConfigured(address indexed apiServer, string metadataUrl)
func (_Bindings *BindingsFilterer) ParseSystemConfigured(log types.Log) (*BindingsSystemConfigured, error) {
	event := new(BindingsSystemConfigured)
	if err := _Bindings.contract.UnpackLog(event, "SystemConfigured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Bindings contract.
type BindingsTransferIterator struct {
	Event *BindingsTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingsTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingsTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingsTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsTransfer represents a Transfer event raised by the Bindings contract.
type BindingsTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Bindings *BindingsFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*BindingsTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &BindingsTransferIterator{contract: _Bindings.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Bindings *BindingsFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BindingsTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsTransfer)
				if err := _Bindings.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Bindings *BindingsFilterer) ParseTransfer(log types.Log) (*BindingsTransfer, error) {
	event := new(BindingsTransfer)
	if err := _Bindings.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Bindings contract.
type BindingsUpgradedIterator struct {
	Event *BindingsUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingsUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingsUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingsUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsUpgraded represents a Upgraded event raised by the Bindings contract.
type BindingsUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Bindings *BindingsFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*BindingsUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &BindingsUpgradedIterator{contract: _Bindings.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Bindings *BindingsFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *BindingsUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsUpgraded)
				if err := _Bindings.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Bindings *BindingsFilterer) ParseUpgraded(log types.Log) (*BindingsUpgraded, error) {
	event := new(BindingsUpgraded)
	if err := _Bindings.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsVehicleTrackedIterator is returned from FilterVehicleTracked and is used to iterate over the raw logs and unpacked data for VehicleTracked events raised by the Bindings contract.
type BindingsVehicleTrackedIterator struct {
	Event *BindingsVehicleTracked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingsVehicleTrackedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsVehicleTracked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingsVehicleTracked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingsVehicleTrackedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsVehicleTrackedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsVehicleTracked represents a VehicleTracked event raised by the Bindings contract.
type BindingsVehicleTracked struct {
	TokenId      *big.Int
	Chassis      [32]byte
	LicensePlate [32]byte
	TotalValue   *big.Int
	BrandId      uint64
	ModelId      uint64
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterVehicleTracked is a free log retrieval operation binding the contract event 0xe3de43c7d5e7c3deb7350fadd7772fbadda69da2d4ab42b3ecbc95446f2aed2a.
//
// Solidity: event VehicleTracked(uint256 indexed tokenId, bytes32 indexed chassis, bytes32 indexed licensePlate, uint128 totalValue, uint64 brandId, uint64 modelId)
func (_Bindings *BindingsFilterer) FilterVehicleTracked(opts *bind.FilterOpts, tokenId []*big.Int, chassis [][32]byte, licensePlate [][32]byte) (*BindingsVehicleTrackedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var chassisRule []interface{}
	for _, chassisItem := range chassis {
		chassisRule = append(chassisRule, chassisItem)
	}
	var licensePlateRule []interface{}
	for _, licensePlateItem := range licensePlate {
		licensePlateRule = append(licensePlateRule, licensePlateItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "VehicleTracked", tokenIdRule, chassisRule, licensePlateRule)
	if err != nil {
		return nil, err
	}
	return &BindingsVehicleTrackedIterator{contract: _Bindings.contract, event: "VehicleTracked", logs: logs, sub: sub}, nil
}

// WatchVehicleTracked is a free log subscription operation binding the contract event 0xe3de43c7d5e7c3deb7350fadd7772fbadda69da2d4ab42b3ecbc95446f2aed2a.
//
// Solidity: event VehicleTracked(uint256 indexed tokenId, bytes32 indexed chassis, bytes32 indexed licensePlate, uint128 totalValue, uint64 brandId, uint64 modelId)
func (_Bindings *BindingsFilterer) WatchVehicleTracked(opts *bind.WatchOpts, sink chan<- *BindingsVehicleTracked, tokenId []*big.Int, chassis [][32]byte, licensePlate [][32]byte) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var chassisRule []interface{}
	for _, chassisItem := range chassis {
		chassisRule = append(chassisRule, chassisItem)
	}
	var licensePlateRule []interface{}
	for _, licensePlateItem := range licensePlate {
		licensePlateRule = append(licensePlateRule, licensePlateItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "VehicleTracked", tokenIdRule, chassisRule, licensePlateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsVehicleTracked)
				if err := _Bindings.contract.UnpackLog(event, "VehicleTracked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVehicleTracked is a log parse operation binding the contract event 0xe3de43c7d5e7c3deb7350fadd7772fbadda69da2d4ab42b3ecbc95446f2aed2a.
//
// Solidity: event VehicleTracked(uint256 indexed tokenId, bytes32 indexed chassis, bytes32 indexed licensePlate, uint128 totalValue, uint64 brandId, uint64 modelId)
func (_Bindings *BindingsFilterer) ParseVehicleTracked(log types.Log) (*BindingsVehicleTracked, error) {
	event := new(BindingsVehicleTracked)
	if err := _Bindings.contract.UnpackLog(event, "VehicleTracked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
