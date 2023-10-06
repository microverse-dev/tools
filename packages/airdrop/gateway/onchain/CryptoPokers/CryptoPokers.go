// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package CryptoPokers

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

// IERC721ATokenOwnership is an auto generated low-level Go binding around an user-defined struct.
type IERC721ATokenOwnership struct {
	Addr           common.Address
	StartTimestamp uint64
	Burned         bool
	ExtraData      *big.Int
}

// CryptoPokersMetaData contains all meta data concerning the CryptoPokers contract.
var CryptoPokersMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_tokenName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_tokenSymbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_cost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxMintAmountPerTx\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_hiddenMetadataUri\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_royaltyBasis\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ApprovalCallerNotOwnerNorApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ApprovalQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BalanceQueryForZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitialRegistryAddressCannotBeZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidQueryRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintERC2309QuantityExceedsLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintZeroQuantity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnershipNotInitializedForExtraData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RegistryHasBeenRevoked\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferCallerNotOwnerNorApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFromIncorrectOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToNonERC721ReceiverImplementer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"URIQueryForNonexistentToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toTokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"ConsecutiveTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newRegistry\",\"type\":\"address\"}],\"name\":\"OperatorFilterRegistryAddressUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"OperatorFilterRegistryRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"PermanentURI\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"alClaimed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_mintAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_maxMintAmount\",\"type\":\"uint8\"},{\"internalType\":\"bytes32[]\",\"name\":\"_merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"alMint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"alMintEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_fromId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_toId\",\"type\":\"uint256\"}],\"name\":\"emitPermanentURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"explicitOwnershipOf\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"startTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"burned\",\"type\":\"bool\"},{\"internalType\":\"uint24\",\"name\":\"extraData\",\"type\":\"uint24\"}],\"internalType\":\"structIERC721A.TokenOwnership\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"explicitOwnershipsOf\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"startTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"burned\",\"type\":\"bool\"},{\"internalType\":\"uint24\",\"name\":\"extraData\",\"type\":\"uint24\"}],\"internalType\":\"structIERC721A.TokenOwnership[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hiddenMetadataUri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOperatorFilterRegistryRevoked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxMintAmountPerTx\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"merkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_mintAmount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_mintAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"mintForAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operatorFilterRegistry\",\"outputs\":[{\"internalType\":\"contractIOperatorFilterRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"permanentURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revealed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revokeOperatorFilterRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"royaltyBasis\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salePrice\",\"type\":\"uint256\"}],\"name\":\"royaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"royaltyAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"royaltyWallet\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_state\",\"type\":\"bool\"}],\"name\":\"setALMintEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_cost\",\"type\":\"uint256\"}],\"name\":\"setCost\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_hiddenMetadataUri\",\"type\":\"string\"}],\"name\":\"setHiddenMetadataUri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxMintAmountPerTx\",\"type\":\"uint256\"}],\"name\":\"setMaxMintAmountPerTx\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_merkleRoot\",\"type\":\"bytes32\"}],\"name\":\"setMerkleRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_state\",\"type\":\"bool\"}],\"name\":\"setPaused\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_state\",\"type\":\"bool\"}],\"name\":\"setRevealed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_royaltyBasis\",\"type\":\"uint256\"}],\"name\":\"setRoyaltyBasis\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_royaltyWallet\",\"type\":\"address\"}],\"name\":\"setRoyaltyWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_uriPrefix\",\"type\":\"string\"}],\"name\":\"setUriPrefix\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_uriSuffix\",\"type\":\"string\"}],\"name\":\"setUriSuffix\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_withdrawalWallet\",\"type\":\"address\"}],\"name\":\"setWithdrawalWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"tokensOfOwner\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stop\",\"type\":\"uint256\"}],\"name\":\"tokensOfOwnerIn\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newRegistry\",\"type\":\"address\"}],\"name\":\"updateOperatorFilterRegistryAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uriPermanent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uriPrefix\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uriSuffix\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawalWallet\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// CryptoPokersABI is the input ABI used to generate the binding from.
// Deprecated: Use CryptoPokersMetaData.ABI instead.
var CryptoPokersABI = CryptoPokersMetaData.ABI

// CryptoPokers is an auto generated Go binding around an Ethereum contract.
type CryptoPokers struct {
	CryptoPokersCaller     // Read-only binding to the contract
	CryptoPokersTransactor // Write-only binding to the contract
	CryptoPokersFilterer   // Log filterer for contract events
}

// CryptoPokersCaller is an auto generated read-only Go binding around an Ethereum contract.
type CryptoPokersCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CryptoPokersTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CryptoPokersTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CryptoPokersFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CryptoPokersFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CryptoPokersSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CryptoPokersSession struct {
	Contract     *CryptoPokers     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CryptoPokersCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CryptoPokersCallerSession struct {
	Contract *CryptoPokersCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// CryptoPokersTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CryptoPokersTransactorSession struct {
	Contract     *CryptoPokersTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// CryptoPokersRaw is an auto generated low-level Go binding around an Ethereum contract.
type CryptoPokersRaw struct {
	Contract *CryptoPokers // Generic contract binding to access the raw methods on
}

// CryptoPokersCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CryptoPokersCallerRaw struct {
	Contract *CryptoPokersCaller // Generic read-only contract binding to access the raw methods on
}

// CryptoPokersTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CryptoPokersTransactorRaw struct {
	Contract *CryptoPokersTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCryptoPokers creates a new instance of CryptoPokers, bound to a specific deployed contract.
func NewCryptoPokers(address common.Address, backend bind.ContractBackend) (*CryptoPokers, error) {
	contract, err := bindCryptoPokers(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CryptoPokers{CryptoPokersCaller: CryptoPokersCaller{contract: contract}, CryptoPokersTransactor: CryptoPokersTransactor{contract: contract}, CryptoPokersFilterer: CryptoPokersFilterer{contract: contract}}, nil
}

// NewCryptoPokersCaller creates a new read-only instance of CryptoPokers, bound to a specific deployed contract.
func NewCryptoPokersCaller(address common.Address, caller bind.ContractCaller) (*CryptoPokersCaller, error) {
	contract, err := bindCryptoPokers(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CryptoPokersCaller{contract: contract}, nil
}

// NewCryptoPokersTransactor creates a new write-only instance of CryptoPokers, bound to a specific deployed contract.
func NewCryptoPokersTransactor(address common.Address, transactor bind.ContractTransactor) (*CryptoPokersTransactor, error) {
	contract, err := bindCryptoPokers(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CryptoPokersTransactor{contract: contract}, nil
}

// NewCryptoPokersFilterer creates a new log filterer instance of CryptoPokers, bound to a specific deployed contract.
func NewCryptoPokersFilterer(address common.Address, filterer bind.ContractFilterer) (*CryptoPokersFilterer, error) {
	contract, err := bindCryptoPokers(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CryptoPokersFilterer{contract: contract}, nil
}

// bindCryptoPokers binds a generic wrapper to an already deployed contract.
func bindCryptoPokers(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CryptoPokersMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CryptoPokers *CryptoPokersRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CryptoPokers.Contract.CryptoPokersCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CryptoPokers *CryptoPokersRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CryptoPokers.Contract.CryptoPokersTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CryptoPokers *CryptoPokersRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CryptoPokers.Contract.CryptoPokersTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CryptoPokers *CryptoPokersCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CryptoPokers.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CryptoPokers *CryptoPokersTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CryptoPokers.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CryptoPokers *CryptoPokersTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CryptoPokers.Contract.contract.Transact(opts, method, params...)
}

// AlClaimed is a free data retrieval call binding the contract method 0x9f0709d8.
//
// Solidity: function alClaimed(address ) view returns(uint256)
func (_CryptoPokers *CryptoPokersCaller) AlClaimed(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CryptoPokers.contract.Call(opts, &out, "alClaimed", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AlClaimed is a free data retrieval call binding the contract method 0x9f0709d8.
//
// Solidity: function alClaimed(address ) view returns(uint256)
func (_CryptoPokers *CryptoPokersSession) AlClaimed(arg0 common.Address) (*big.Int, error) {
	return _CryptoPokers.Contract.AlClaimed(&_CryptoPokers.CallOpts, arg0)
}

// AlClaimed is a free data retrieval call binding the contract method 0x9f0709d8.
//
// Solidity: function alClaimed(address ) view returns(uint256)
func (_CryptoPokers *CryptoPokersCallerSession) AlClaimed(arg0 common.Address) (*big.Int, error) {
	return _CryptoPokers.Contract.AlClaimed(&_CryptoPokers.CallOpts, arg0)
}

// AlMintEnabled is a free data retrieval call binding the contract method 0x4c0d0a20.
//
// Solidity: function alMintEnabled() view returns(bool)
func (_CryptoPokers *CryptoPokersCaller) AlMintEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CryptoPokers.contract.Call(opts, &out, "alMintEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AlMintEnabled is a free data retrieval call binding the contract method 0x4c0d0a20.
//
// Solidity: function alMintEnabled() view returns(bool)
func (_CryptoPokers *CryptoPokersSession) AlMintEnabled() (bool, error) {
	return _CryptoPokers.Contract.AlMintEnabled(&_CryptoPokers.CallOpts)
}

// AlMintEnabled is a free data retrieval call binding the contract method 0x4c0d0a20.
//
// Solidity: function alMintEnabled() view returns(bool)
func (_CryptoPokers *CryptoPokersCallerSession) AlMintEnabled() (bool, error) {
	return _CryptoPokers.Contract.AlMintEnabled(&_CryptoPokers.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_CryptoPokers *CryptoPokersCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CryptoPokers.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_CryptoPokers *CryptoPokersSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _CryptoPokers.Contract.BalanceOf(&_CryptoPokers.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_CryptoPokers *CryptoPokersCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _CryptoPokers.Contract.BalanceOf(&_CryptoPokers.CallOpts, owner)
}

// Cost is a free data retrieval call binding the contract method 0x13faede6.
//
// Solidity: function cost() view returns(uint256)
func (_CryptoPokers *CryptoPokersCaller) Cost(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CryptoPokers.contract.Call(opts, &out, "cost")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Cost is a free data retrieval call binding the contract method 0x13faede6.
//
// Solidity: function cost() view returns(uint256)
func (_CryptoPokers *CryptoPokersSession) Cost() (*big.Int, error) {
	return _CryptoPokers.Contract.Cost(&_CryptoPokers.CallOpts)
}

// Cost is a free data retrieval call binding the contract method 0x13faede6.
//
// Solidity: function cost() view returns(uint256)
func (_CryptoPokers *CryptoPokersCallerSession) Cost() (*big.Int, error) {
	return _CryptoPokers.Contract.Cost(&_CryptoPokers.CallOpts)
}

// ExplicitOwnershipOf is a free data retrieval call binding the contract method 0xc23dc68f.
//
// Solidity: function explicitOwnershipOf(uint256 tokenId) view returns((address,uint64,bool,uint24))
func (_CryptoPokers *CryptoPokersCaller) ExplicitOwnershipOf(opts *bind.CallOpts, tokenId *big.Int) (IERC721ATokenOwnership, error) {
	var out []interface{}
	err := _CryptoPokers.contract.Call(opts, &out, "explicitOwnershipOf", tokenId)

	if err != nil {
		return *new(IERC721ATokenOwnership), err
	}

	out0 := *abi.ConvertType(out[0], new(IERC721ATokenOwnership)).(*IERC721ATokenOwnership)

	return out0, err

}

// ExplicitOwnershipOf is a free data retrieval call binding the contract method 0xc23dc68f.
//
// Solidity: function explicitOwnershipOf(uint256 tokenId) view returns((address,uint64,bool,uint24))
func (_CryptoPokers *CryptoPokersSession) ExplicitOwnershipOf(tokenId *big.Int) (IERC721ATokenOwnership, error) {
	return _CryptoPokers.Contract.ExplicitOwnershipOf(&_CryptoPokers.CallOpts, tokenId)
}

// ExplicitOwnershipOf is a free data retrieval call binding the contract method 0xc23dc68f.
//
// Solidity: function explicitOwnershipOf(uint256 tokenId) view returns((address,uint64,bool,uint24))
func (_CryptoPokers *CryptoPokersCallerSession) ExplicitOwnershipOf(tokenId *big.Int) (IERC721ATokenOwnership, error) {
	return _CryptoPokers.Contract.ExplicitOwnershipOf(&_CryptoPokers.CallOpts, tokenId)
}

// ExplicitOwnershipsOf is a free data retrieval call binding the contract method 0x5bbb2177.
//
// Solidity: function explicitOwnershipsOf(uint256[] tokenIds) view returns((address,uint64,bool,uint24)[])
func (_CryptoPokers *CryptoPokersCaller) ExplicitOwnershipsOf(opts *bind.CallOpts, tokenIds []*big.Int) ([]IERC721ATokenOwnership, error) {
	var out []interface{}
	err := _CryptoPokers.contract.Call(opts, &out, "explicitOwnershipsOf", tokenIds)

	if err != nil {
		return *new([]IERC721ATokenOwnership), err
	}

	out0 := *abi.ConvertType(out[0], new([]IERC721ATokenOwnership)).(*[]IERC721ATokenOwnership)

	return out0, err

}

// ExplicitOwnershipsOf is a free data retrieval call binding the contract method 0x5bbb2177.
//
// Solidity: function explicitOwnershipsOf(uint256[] tokenIds) view returns((address,uint64,bool,uint24)[])
func (_CryptoPokers *CryptoPokersSession) ExplicitOwnershipsOf(tokenIds []*big.Int) ([]IERC721ATokenOwnership, error) {
	return _CryptoPokers.Contract.ExplicitOwnershipsOf(&_CryptoPokers.CallOpts, tokenIds)
}

// ExplicitOwnershipsOf is a free data retrieval call binding the contract method 0x5bbb2177.
//
// Solidity: function explicitOwnershipsOf(uint256[] tokenIds) view returns((address,uint64,bool,uint24)[])
func (_CryptoPokers *CryptoPokersCallerSession) ExplicitOwnershipsOf(tokenIds []*big.Int) ([]IERC721ATokenOwnership, error) {
	return _CryptoPokers.Contract.ExplicitOwnershipsOf(&_CryptoPokers.CallOpts, tokenIds)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_CryptoPokers *CryptoPokersCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CryptoPokers.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_CryptoPokers *CryptoPokersSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _CryptoPokers.Contract.GetApproved(&_CryptoPokers.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_CryptoPokers *CryptoPokersCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _CryptoPokers.Contract.GetApproved(&_CryptoPokers.CallOpts, tokenId)
}

// HiddenMetadataUri is a free data retrieval call binding the contract method 0xa45ba8e7.
//
// Solidity: function hiddenMetadataUri() view returns(string)
func (_CryptoPokers *CryptoPokersCaller) HiddenMetadataUri(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CryptoPokers.contract.Call(opts, &out, "hiddenMetadataUri")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// HiddenMetadataUri is a free data retrieval call binding the contract method 0xa45ba8e7.
//
// Solidity: function hiddenMetadataUri() view returns(string)
func (_CryptoPokers *CryptoPokersSession) HiddenMetadataUri() (string, error) {
	return _CryptoPokers.Contract.HiddenMetadataUri(&_CryptoPokers.CallOpts)
}

// HiddenMetadataUri is a free data retrieval call binding the contract method 0xa45ba8e7.
//
// Solidity: function hiddenMetadataUri() view returns(string)
func (_CryptoPokers *CryptoPokersCallerSession) HiddenMetadataUri() (string, error) {
	return _CryptoPokers.Contract.HiddenMetadataUri(&_CryptoPokers.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_CryptoPokers *CryptoPokersCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _CryptoPokers.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_CryptoPokers *CryptoPokersSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _CryptoPokers.Contract.IsApprovedForAll(&_CryptoPokers.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_CryptoPokers *CryptoPokersCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _CryptoPokers.Contract.IsApprovedForAll(&_CryptoPokers.CallOpts, owner, operator)
}

// IsOperatorFilterRegistryRevoked is a free data retrieval call binding the contract method 0xecba222a.
//
// Solidity: function isOperatorFilterRegistryRevoked() view returns(bool)
func (_CryptoPokers *CryptoPokersCaller) IsOperatorFilterRegistryRevoked(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CryptoPokers.contract.Call(opts, &out, "isOperatorFilterRegistryRevoked")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperatorFilterRegistryRevoked is a free data retrieval call binding the contract method 0xecba222a.
//
// Solidity: function isOperatorFilterRegistryRevoked() view returns(bool)
func (_CryptoPokers *CryptoPokersSession) IsOperatorFilterRegistryRevoked() (bool, error) {
	return _CryptoPokers.Contract.IsOperatorFilterRegistryRevoked(&_CryptoPokers.CallOpts)
}

// IsOperatorFilterRegistryRevoked is a free data retrieval call binding the contract method 0xecba222a.
//
// Solidity: function isOperatorFilterRegistryRevoked() view returns(bool)
func (_CryptoPokers *CryptoPokersCallerSession) IsOperatorFilterRegistryRevoked() (bool, error) {
	return _CryptoPokers.Contract.IsOperatorFilterRegistryRevoked(&_CryptoPokers.CallOpts)
}

// MaxMintAmountPerTx is a free data retrieval call binding the contract method 0x94354fd0.
//
// Solidity: function maxMintAmountPerTx() view returns(uint256)
func (_CryptoPokers *CryptoPokersCaller) MaxMintAmountPerTx(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CryptoPokers.contract.Call(opts, &out, "maxMintAmountPerTx")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxMintAmountPerTx is a free data retrieval call binding the contract method 0x94354fd0.
//
// Solidity: function maxMintAmountPerTx() view returns(uint256)
func (_CryptoPokers *CryptoPokersSession) MaxMintAmountPerTx() (*big.Int, error) {
	return _CryptoPokers.Contract.MaxMintAmountPerTx(&_CryptoPokers.CallOpts)
}

// MaxMintAmountPerTx is a free data retrieval call binding the contract method 0x94354fd0.
//
// Solidity: function maxMintAmountPerTx() view returns(uint256)
func (_CryptoPokers *CryptoPokersCallerSession) MaxMintAmountPerTx() (*big.Int, error) {
	return _CryptoPokers.Contract.MaxMintAmountPerTx(&_CryptoPokers.CallOpts)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_CryptoPokers *CryptoPokersCaller) MaxSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CryptoPokers.contract.Call(opts, &out, "maxSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_CryptoPokers *CryptoPokersSession) MaxSupply() (*big.Int, error) {
	return _CryptoPokers.Contract.MaxSupply(&_CryptoPokers.CallOpts)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_CryptoPokers *CryptoPokersCallerSession) MaxSupply() (*big.Int, error) {
	return _CryptoPokers.Contract.MaxSupply(&_CryptoPokers.CallOpts)
}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_CryptoPokers *CryptoPokersCaller) MerkleRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CryptoPokers.contract.Call(opts, &out, "merkleRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_CryptoPokers *CryptoPokersSession) MerkleRoot() ([32]byte, error) {
	return _CryptoPokers.Contract.MerkleRoot(&_CryptoPokers.CallOpts)
}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_CryptoPokers *CryptoPokersCallerSession) MerkleRoot() ([32]byte, error) {
	return _CryptoPokers.Contract.MerkleRoot(&_CryptoPokers.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CryptoPokers *CryptoPokersCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CryptoPokers.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CryptoPokers *CryptoPokersSession) Name() (string, error) {
	return _CryptoPokers.Contract.Name(&_CryptoPokers.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CryptoPokers *CryptoPokersCallerSession) Name() (string, error) {
	return _CryptoPokers.Contract.Name(&_CryptoPokers.CallOpts)
}

// OperatorFilterRegistry is a free data retrieval call binding the contract method 0xb0ccc31e.
//
// Solidity: function operatorFilterRegistry() view returns(address)
func (_CryptoPokers *CryptoPokersCaller) OperatorFilterRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CryptoPokers.contract.Call(opts, &out, "operatorFilterRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OperatorFilterRegistry is a free data retrieval call binding the contract method 0xb0ccc31e.
//
// Solidity: function operatorFilterRegistry() view returns(address)
func (_CryptoPokers *CryptoPokersSession) OperatorFilterRegistry() (common.Address, error) {
	return _CryptoPokers.Contract.OperatorFilterRegistry(&_CryptoPokers.CallOpts)
}

// OperatorFilterRegistry is a free data retrieval call binding the contract method 0xb0ccc31e.
//
// Solidity: function operatorFilterRegistry() view returns(address)
func (_CryptoPokers *CryptoPokersCallerSession) OperatorFilterRegistry() (common.Address, error) {
	return _CryptoPokers.Contract.OperatorFilterRegistry(&_CryptoPokers.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CryptoPokers *CryptoPokersCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CryptoPokers.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CryptoPokers *CryptoPokersSession) Owner() (common.Address, error) {
	return _CryptoPokers.Contract.Owner(&_CryptoPokers.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CryptoPokers *CryptoPokersCallerSession) Owner() (common.Address, error) {
	return _CryptoPokers.Contract.Owner(&_CryptoPokers.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_CryptoPokers *CryptoPokersCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CryptoPokers.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_CryptoPokers *CryptoPokersSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _CryptoPokers.Contract.OwnerOf(&_CryptoPokers.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_CryptoPokers *CryptoPokersCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _CryptoPokers.Contract.OwnerOf(&_CryptoPokers.CallOpts, tokenId)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CryptoPokers *CryptoPokersCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CryptoPokers.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CryptoPokers *CryptoPokersSession) Paused() (bool, error) {
	return _CryptoPokers.Contract.Paused(&_CryptoPokers.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CryptoPokers *CryptoPokersCallerSession) Paused() (bool, error) {
	return _CryptoPokers.Contract.Paused(&_CryptoPokers.CallOpts)
}

// Revealed is a free data retrieval call binding the contract method 0x51830227.
//
// Solidity: function revealed() view returns(bool)
func (_CryptoPokers *CryptoPokersCaller) Revealed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CryptoPokers.contract.Call(opts, &out, "revealed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Revealed is a free data retrieval call binding the contract method 0x51830227.
//
// Solidity: function revealed() view returns(bool)
func (_CryptoPokers *CryptoPokersSession) Revealed() (bool, error) {
	return _CryptoPokers.Contract.Revealed(&_CryptoPokers.CallOpts)
}

// Revealed is a free data retrieval call binding the contract method 0x51830227.
//
// Solidity: function revealed() view returns(bool)
func (_CryptoPokers *CryptoPokersCallerSession) Revealed() (bool, error) {
	return _CryptoPokers.Contract.Revealed(&_CryptoPokers.CallOpts)
}

// RoyaltyBasis is a free data retrieval call binding the contract method 0xc6ec6909.
//
// Solidity: function royaltyBasis() view returns(uint256)
func (_CryptoPokers *CryptoPokersCaller) RoyaltyBasis(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CryptoPokers.contract.Call(opts, &out, "royaltyBasis")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RoyaltyBasis is a free data retrieval call binding the contract method 0xc6ec6909.
//
// Solidity: function royaltyBasis() view returns(uint256)
func (_CryptoPokers *CryptoPokersSession) RoyaltyBasis() (*big.Int, error) {
	return _CryptoPokers.Contract.RoyaltyBasis(&_CryptoPokers.CallOpts)
}

// RoyaltyBasis is a free data retrieval call binding the contract method 0xc6ec6909.
//
// Solidity: function royaltyBasis() view returns(uint256)
func (_CryptoPokers *CryptoPokersCallerSession) RoyaltyBasis() (*big.Int, error) {
	return _CryptoPokers.Contract.RoyaltyBasis(&_CryptoPokers.CallOpts)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_CryptoPokers *CryptoPokersCaller) RoyaltyInfo(opts *bind.CallOpts, tokenId *big.Int, salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	var out []interface{}
	err := _CryptoPokers.contract.Call(opts, &out, "royaltyInfo", tokenId, salePrice)

	outstruct := new(struct {
		Receiver      common.Address
		RoyaltyAmount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Receiver = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.RoyaltyAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_CryptoPokers *CryptoPokersSession) RoyaltyInfo(tokenId *big.Int, salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _CryptoPokers.Contract.RoyaltyInfo(&_CryptoPokers.CallOpts, tokenId, salePrice)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_CryptoPokers *CryptoPokersCallerSession) RoyaltyInfo(tokenId *big.Int, salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _CryptoPokers.Contract.RoyaltyInfo(&_CryptoPokers.CallOpts, tokenId, salePrice)
}

// RoyaltyWallet is a free data retrieval call binding the contract method 0x3f0d2ec1.
//
// Solidity: function royaltyWallet() view returns(address)
func (_CryptoPokers *CryptoPokersCaller) RoyaltyWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CryptoPokers.contract.Call(opts, &out, "royaltyWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RoyaltyWallet is a free data retrieval call binding the contract method 0x3f0d2ec1.
//
// Solidity: function royaltyWallet() view returns(address)
func (_CryptoPokers *CryptoPokersSession) RoyaltyWallet() (common.Address, error) {
	return _CryptoPokers.Contract.RoyaltyWallet(&_CryptoPokers.CallOpts)
}

// RoyaltyWallet is a free data retrieval call binding the contract method 0x3f0d2ec1.
//
// Solidity: function royaltyWallet() view returns(address)
func (_CryptoPokers *CryptoPokersCallerSession) RoyaltyWallet() (common.Address, error) {
	return _CryptoPokers.Contract.RoyaltyWallet(&_CryptoPokers.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CryptoPokers *CryptoPokersCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _CryptoPokers.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CryptoPokers *CryptoPokersSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CryptoPokers.Contract.SupportsInterface(&_CryptoPokers.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CryptoPokers *CryptoPokersCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CryptoPokers.Contract.SupportsInterface(&_CryptoPokers.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CryptoPokers *CryptoPokersCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CryptoPokers.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CryptoPokers *CryptoPokersSession) Symbol() (string, error) {
	return _CryptoPokers.Contract.Symbol(&_CryptoPokers.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CryptoPokers *CryptoPokersCallerSession) Symbol() (string, error) {
	return _CryptoPokers.Contract.Symbol(&_CryptoPokers.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_CryptoPokers *CryptoPokersCaller) TokenURI(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _CryptoPokers.contract.Call(opts, &out, "tokenURI", _tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_CryptoPokers *CryptoPokersSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _CryptoPokers.Contract.TokenURI(&_CryptoPokers.CallOpts, _tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_CryptoPokers *CryptoPokersCallerSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _CryptoPokers.Contract.TokenURI(&_CryptoPokers.CallOpts, _tokenId)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address owner) view returns(uint256[])
func (_CryptoPokers *CryptoPokersCaller) TokensOfOwner(opts *bind.CallOpts, owner common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _CryptoPokers.contract.Call(opts, &out, "tokensOfOwner", owner)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address owner) view returns(uint256[])
func (_CryptoPokers *CryptoPokersSession) TokensOfOwner(owner common.Address) ([]*big.Int, error) {
	return _CryptoPokers.Contract.TokensOfOwner(&_CryptoPokers.CallOpts, owner)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address owner) view returns(uint256[])
func (_CryptoPokers *CryptoPokersCallerSession) TokensOfOwner(owner common.Address) ([]*big.Int, error) {
	return _CryptoPokers.Contract.TokensOfOwner(&_CryptoPokers.CallOpts, owner)
}

// TokensOfOwnerIn is a free data retrieval call binding the contract method 0x99a2557a.
//
// Solidity: function tokensOfOwnerIn(address owner, uint256 start, uint256 stop) view returns(uint256[])
func (_CryptoPokers *CryptoPokersCaller) TokensOfOwnerIn(opts *bind.CallOpts, owner common.Address, start *big.Int, stop *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _CryptoPokers.contract.Call(opts, &out, "tokensOfOwnerIn", owner, start, stop)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// TokensOfOwnerIn is a free data retrieval call binding the contract method 0x99a2557a.
//
// Solidity: function tokensOfOwnerIn(address owner, uint256 start, uint256 stop) view returns(uint256[])
func (_CryptoPokers *CryptoPokersSession) TokensOfOwnerIn(owner common.Address, start *big.Int, stop *big.Int) ([]*big.Int, error) {
	return _CryptoPokers.Contract.TokensOfOwnerIn(&_CryptoPokers.CallOpts, owner, start, stop)
}

// TokensOfOwnerIn is a free data retrieval call binding the contract method 0x99a2557a.
//
// Solidity: function tokensOfOwnerIn(address owner, uint256 start, uint256 stop) view returns(uint256[])
func (_CryptoPokers *CryptoPokersCallerSession) TokensOfOwnerIn(owner common.Address, start *big.Int, stop *big.Int) ([]*big.Int, error) {
	return _CryptoPokers.Contract.TokensOfOwnerIn(&_CryptoPokers.CallOpts, owner, start, stop)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CryptoPokers *CryptoPokersCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CryptoPokers.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CryptoPokers *CryptoPokersSession) TotalSupply() (*big.Int, error) {
	return _CryptoPokers.Contract.TotalSupply(&_CryptoPokers.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CryptoPokers *CryptoPokersCallerSession) TotalSupply() (*big.Int, error) {
	return _CryptoPokers.Contract.TotalSupply(&_CryptoPokers.CallOpts)
}

// UriPermanent is a free data retrieval call binding the contract method 0x4e63b8a1.
//
// Solidity: function uriPermanent() view returns(bool)
func (_CryptoPokers *CryptoPokersCaller) UriPermanent(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CryptoPokers.contract.Call(opts, &out, "uriPermanent")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UriPermanent is a free data retrieval call binding the contract method 0x4e63b8a1.
//
// Solidity: function uriPermanent() view returns(bool)
func (_CryptoPokers *CryptoPokersSession) UriPermanent() (bool, error) {
	return _CryptoPokers.Contract.UriPermanent(&_CryptoPokers.CallOpts)
}

// UriPermanent is a free data retrieval call binding the contract method 0x4e63b8a1.
//
// Solidity: function uriPermanent() view returns(bool)
func (_CryptoPokers *CryptoPokersCallerSession) UriPermanent() (bool, error) {
	return _CryptoPokers.Contract.UriPermanent(&_CryptoPokers.CallOpts)
}

// UriPrefix is a free data retrieval call binding the contract method 0x62b99ad4.
//
// Solidity: function uriPrefix() view returns(string)
func (_CryptoPokers *CryptoPokersCaller) UriPrefix(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CryptoPokers.contract.Call(opts, &out, "uriPrefix")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UriPrefix is a free data retrieval call binding the contract method 0x62b99ad4.
//
// Solidity: function uriPrefix() view returns(string)
func (_CryptoPokers *CryptoPokersSession) UriPrefix() (string, error) {
	return _CryptoPokers.Contract.UriPrefix(&_CryptoPokers.CallOpts)
}

// UriPrefix is a free data retrieval call binding the contract method 0x62b99ad4.
//
// Solidity: function uriPrefix() view returns(string)
func (_CryptoPokers *CryptoPokersCallerSession) UriPrefix() (string, error) {
	return _CryptoPokers.Contract.UriPrefix(&_CryptoPokers.CallOpts)
}

// UriSuffix is a free data retrieval call binding the contract method 0x5503a0e8.
//
// Solidity: function uriSuffix() view returns(string)
func (_CryptoPokers *CryptoPokersCaller) UriSuffix(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CryptoPokers.contract.Call(opts, &out, "uriSuffix")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UriSuffix is a free data retrieval call binding the contract method 0x5503a0e8.
//
// Solidity: function uriSuffix() view returns(string)
func (_CryptoPokers *CryptoPokersSession) UriSuffix() (string, error) {
	return _CryptoPokers.Contract.UriSuffix(&_CryptoPokers.CallOpts)
}

// UriSuffix is a free data retrieval call binding the contract method 0x5503a0e8.
//
// Solidity: function uriSuffix() view returns(string)
func (_CryptoPokers *CryptoPokersCallerSession) UriSuffix() (string, error) {
	return _CryptoPokers.Contract.UriSuffix(&_CryptoPokers.CallOpts)
}

// WithdrawalWallet is a free data retrieval call binding the contract method 0x4a7d80b3.
//
// Solidity: function withdrawalWallet() view returns(address)
func (_CryptoPokers *CryptoPokersCaller) WithdrawalWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CryptoPokers.contract.Call(opts, &out, "withdrawalWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WithdrawalWallet is a free data retrieval call binding the contract method 0x4a7d80b3.
//
// Solidity: function withdrawalWallet() view returns(address)
func (_CryptoPokers *CryptoPokersSession) WithdrawalWallet() (common.Address, error) {
	return _CryptoPokers.Contract.WithdrawalWallet(&_CryptoPokers.CallOpts)
}

// WithdrawalWallet is a free data retrieval call binding the contract method 0x4a7d80b3.
//
// Solidity: function withdrawalWallet() view returns(address)
func (_CryptoPokers *CryptoPokersCallerSession) WithdrawalWallet() (common.Address, error) {
	return _CryptoPokers.Contract.WithdrawalWallet(&_CryptoPokers.CallOpts)
}

// AlMint is a paid mutator transaction binding the contract method 0xd609c04d.
//
// Solidity: function alMint(uint256 _mintAmount, uint8 _maxMintAmount, bytes32[] _merkleProof) payable returns()
func (_CryptoPokers *CryptoPokersTransactor) AlMint(opts *bind.TransactOpts, _mintAmount *big.Int, _maxMintAmount uint8, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _CryptoPokers.contract.Transact(opts, "alMint", _mintAmount, _maxMintAmount, _merkleProof)
}

// AlMint is a paid mutator transaction binding the contract method 0xd609c04d.
//
// Solidity: function alMint(uint256 _mintAmount, uint8 _maxMintAmount, bytes32[] _merkleProof) payable returns()
func (_CryptoPokers *CryptoPokersSession) AlMint(_mintAmount *big.Int, _maxMintAmount uint8, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _CryptoPokers.Contract.AlMint(&_CryptoPokers.TransactOpts, _mintAmount, _maxMintAmount, _merkleProof)
}

// AlMint is a paid mutator transaction binding the contract method 0xd609c04d.
//
// Solidity: function alMint(uint256 _mintAmount, uint8 _maxMintAmount, bytes32[] _merkleProof) payable returns()
func (_CryptoPokers *CryptoPokersTransactorSession) AlMint(_mintAmount *big.Int, _maxMintAmount uint8, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _CryptoPokers.Contract.AlMint(&_CryptoPokers.TransactOpts, _mintAmount, _maxMintAmount, _merkleProof)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address operator, uint256 tokenId) payable returns()
func (_CryptoPokers *CryptoPokersTransactor) Approve(opts *bind.TransactOpts, operator common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CryptoPokers.contract.Transact(opts, "approve", operator, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address operator, uint256 tokenId) payable returns()
func (_CryptoPokers *CryptoPokersSession) Approve(operator common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CryptoPokers.Contract.Approve(&_CryptoPokers.TransactOpts, operator, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address operator, uint256 tokenId) payable returns()
func (_CryptoPokers *CryptoPokersTransactorSession) Approve(operator common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CryptoPokers.Contract.Approve(&_CryptoPokers.TransactOpts, operator, tokenId)
}

// EmitPermanentURI is a paid mutator transaction binding the contract method 0xf22a6903.
//
// Solidity: function emitPermanentURI(uint256 _fromId, uint256 _toId) returns()
func (_CryptoPokers *CryptoPokersTransactor) EmitPermanentURI(opts *bind.TransactOpts, _fromId *big.Int, _toId *big.Int) (*types.Transaction, error) {
	return _CryptoPokers.contract.Transact(opts, "emitPermanentURI", _fromId, _toId)
}

// EmitPermanentURI is a paid mutator transaction binding the contract method 0xf22a6903.
//
// Solidity: function emitPermanentURI(uint256 _fromId, uint256 _toId) returns()
func (_CryptoPokers *CryptoPokersSession) EmitPermanentURI(_fromId *big.Int, _toId *big.Int) (*types.Transaction, error) {
	return _CryptoPokers.Contract.EmitPermanentURI(&_CryptoPokers.TransactOpts, _fromId, _toId)
}

// EmitPermanentURI is a paid mutator transaction binding the contract method 0xf22a6903.
//
// Solidity: function emitPermanentURI(uint256 _fromId, uint256 _toId) returns()
func (_CryptoPokers *CryptoPokersTransactorSession) EmitPermanentURI(_fromId *big.Int, _toId *big.Int) (*types.Transaction, error) {
	return _CryptoPokers.Contract.EmitPermanentURI(&_CryptoPokers.TransactOpts, _fromId, _toId)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 _mintAmount) payable returns()
func (_CryptoPokers *CryptoPokersTransactor) Mint(opts *bind.TransactOpts, _mintAmount *big.Int) (*types.Transaction, error) {
	return _CryptoPokers.contract.Transact(opts, "mint", _mintAmount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 _mintAmount) payable returns()
func (_CryptoPokers *CryptoPokersSession) Mint(_mintAmount *big.Int) (*types.Transaction, error) {
	return _CryptoPokers.Contract.Mint(&_CryptoPokers.TransactOpts, _mintAmount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 _mintAmount) payable returns()
func (_CryptoPokers *CryptoPokersTransactorSession) Mint(_mintAmount *big.Int) (*types.Transaction, error) {
	return _CryptoPokers.Contract.Mint(&_CryptoPokers.TransactOpts, _mintAmount)
}

// MintForAddress is a paid mutator transaction binding the contract method 0xefbd73f4.
//
// Solidity: function mintForAddress(uint256 _mintAmount, address _receiver) returns()
func (_CryptoPokers *CryptoPokersTransactor) MintForAddress(opts *bind.TransactOpts, _mintAmount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _CryptoPokers.contract.Transact(opts, "mintForAddress", _mintAmount, _receiver)
}

// MintForAddress is a paid mutator transaction binding the contract method 0xefbd73f4.
//
// Solidity: function mintForAddress(uint256 _mintAmount, address _receiver) returns()
func (_CryptoPokers *CryptoPokersSession) MintForAddress(_mintAmount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _CryptoPokers.Contract.MintForAddress(&_CryptoPokers.TransactOpts, _mintAmount, _receiver)
}

// MintForAddress is a paid mutator transaction binding the contract method 0xefbd73f4.
//
// Solidity: function mintForAddress(uint256 _mintAmount, address _receiver) returns()
func (_CryptoPokers *CryptoPokersTransactorSession) MintForAddress(_mintAmount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _CryptoPokers.Contract.MintForAddress(&_CryptoPokers.TransactOpts, _mintAmount, _receiver)
}

// PermanentURI is a paid mutator transaction binding the contract method 0x31b54a15.
//
// Solidity: function permanentURI() returns()
func (_CryptoPokers *CryptoPokersTransactor) PermanentURI(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CryptoPokers.contract.Transact(opts, "permanentURI")
}

// PermanentURI is a paid mutator transaction binding the contract method 0x31b54a15.
//
// Solidity: function permanentURI() returns()
func (_CryptoPokers *CryptoPokersSession) PermanentURI() (*types.Transaction, error) {
	return _CryptoPokers.Contract.PermanentURI(&_CryptoPokers.TransactOpts)
}

// PermanentURI is a paid mutator transaction binding the contract method 0x31b54a15.
//
// Solidity: function permanentURI() returns()
func (_CryptoPokers *CryptoPokersTransactorSession) PermanentURI() (*types.Transaction, error) {
	return _CryptoPokers.Contract.PermanentURI(&_CryptoPokers.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CryptoPokers *CryptoPokersTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CryptoPokers.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CryptoPokers *CryptoPokersSession) RenounceOwnership() (*types.Transaction, error) {
	return _CryptoPokers.Contract.RenounceOwnership(&_CryptoPokers.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CryptoPokers *CryptoPokersTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _CryptoPokers.Contract.RenounceOwnership(&_CryptoPokers.TransactOpts)
}

// RevokeOperatorFilterRegistry is a paid mutator transaction binding the contract method 0x5ef9432a.
//
// Solidity: function revokeOperatorFilterRegistry() returns()
func (_CryptoPokers *CryptoPokersTransactor) RevokeOperatorFilterRegistry(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CryptoPokers.contract.Transact(opts, "revokeOperatorFilterRegistry")
}

// RevokeOperatorFilterRegistry is a paid mutator transaction binding the contract method 0x5ef9432a.
//
// Solidity: function revokeOperatorFilterRegistry() returns()
func (_CryptoPokers *CryptoPokersSession) RevokeOperatorFilterRegistry() (*types.Transaction, error) {
	return _CryptoPokers.Contract.RevokeOperatorFilterRegistry(&_CryptoPokers.TransactOpts)
}

// RevokeOperatorFilterRegistry is a paid mutator transaction binding the contract method 0x5ef9432a.
//
// Solidity: function revokeOperatorFilterRegistry() returns()
func (_CryptoPokers *CryptoPokersTransactorSession) RevokeOperatorFilterRegistry() (*types.Transaction, error) {
	return _CryptoPokers.Contract.RevokeOperatorFilterRegistry(&_CryptoPokers.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) payable returns()
func (_CryptoPokers *CryptoPokersTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CryptoPokers.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) payable returns()
func (_CryptoPokers *CryptoPokersSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CryptoPokers.Contract.SafeTransferFrom(&_CryptoPokers.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) payable returns()
func (_CryptoPokers *CryptoPokersTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CryptoPokers.Contract.SafeTransferFrom(&_CryptoPokers.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) payable returns()
func (_CryptoPokers *CryptoPokersTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _CryptoPokers.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) payable returns()
func (_CryptoPokers *CryptoPokersSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _CryptoPokers.Contract.SafeTransferFrom0(&_CryptoPokers.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) payable returns()
func (_CryptoPokers *CryptoPokersTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _CryptoPokers.Contract.SafeTransferFrom0(&_CryptoPokers.TransactOpts, from, to, tokenId, data)
}

// SetALMintEnabled is a paid mutator transaction binding the contract method 0x1d0f48f7.
//
// Solidity: function setALMintEnabled(bool _state) returns()
func (_CryptoPokers *CryptoPokersTransactor) SetALMintEnabled(opts *bind.TransactOpts, _state bool) (*types.Transaction, error) {
	return _CryptoPokers.contract.Transact(opts, "setALMintEnabled", _state)
}

// SetALMintEnabled is a paid mutator transaction binding the contract method 0x1d0f48f7.
//
// Solidity: function setALMintEnabled(bool _state) returns()
func (_CryptoPokers *CryptoPokersSession) SetALMintEnabled(_state bool) (*types.Transaction, error) {
	return _CryptoPokers.Contract.SetALMintEnabled(&_CryptoPokers.TransactOpts, _state)
}

// SetALMintEnabled is a paid mutator transaction binding the contract method 0x1d0f48f7.
//
// Solidity: function setALMintEnabled(bool _state) returns()
func (_CryptoPokers *CryptoPokersTransactorSession) SetALMintEnabled(_state bool) (*types.Transaction, error) {
	return _CryptoPokers.Contract.SetALMintEnabled(&_CryptoPokers.TransactOpts, _state)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_CryptoPokers *CryptoPokersTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _CryptoPokers.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_CryptoPokers *CryptoPokersSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _CryptoPokers.Contract.SetApprovalForAll(&_CryptoPokers.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_CryptoPokers *CryptoPokersTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _CryptoPokers.Contract.SetApprovalForAll(&_CryptoPokers.TransactOpts, operator, approved)
}

// SetCost is a paid mutator transaction binding the contract method 0x44a0d68a.
//
// Solidity: function setCost(uint256 _cost) returns()
func (_CryptoPokers *CryptoPokersTransactor) SetCost(opts *bind.TransactOpts, _cost *big.Int) (*types.Transaction, error) {
	return _CryptoPokers.contract.Transact(opts, "setCost", _cost)
}

// SetCost is a paid mutator transaction binding the contract method 0x44a0d68a.
//
// Solidity: function setCost(uint256 _cost) returns()
func (_CryptoPokers *CryptoPokersSession) SetCost(_cost *big.Int) (*types.Transaction, error) {
	return _CryptoPokers.Contract.SetCost(&_CryptoPokers.TransactOpts, _cost)
}

// SetCost is a paid mutator transaction binding the contract method 0x44a0d68a.
//
// Solidity: function setCost(uint256 _cost) returns()
func (_CryptoPokers *CryptoPokersTransactorSession) SetCost(_cost *big.Int) (*types.Transaction, error) {
	return _CryptoPokers.Contract.SetCost(&_CryptoPokers.TransactOpts, _cost)
}

// SetHiddenMetadataUri is a paid mutator transaction binding the contract method 0x4fdd43cb.
//
// Solidity: function setHiddenMetadataUri(string _hiddenMetadataUri) returns()
func (_CryptoPokers *CryptoPokersTransactor) SetHiddenMetadataUri(opts *bind.TransactOpts, _hiddenMetadataUri string) (*types.Transaction, error) {
	return _CryptoPokers.contract.Transact(opts, "setHiddenMetadataUri", _hiddenMetadataUri)
}

// SetHiddenMetadataUri is a paid mutator transaction binding the contract method 0x4fdd43cb.
//
// Solidity: function setHiddenMetadataUri(string _hiddenMetadataUri) returns()
func (_CryptoPokers *CryptoPokersSession) SetHiddenMetadataUri(_hiddenMetadataUri string) (*types.Transaction, error) {
	return _CryptoPokers.Contract.SetHiddenMetadataUri(&_CryptoPokers.TransactOpts, _hiddenMetadataUri)
}

// SetHiddenMetadataUri is a paid mutator transaction binding the contract method 0x4fdd43cb.
//
// Solidity: function setHiddenMetadataUri(string _hiddenMetadataUri) returns()
func (_CryptoPokers *CryptoPokersTransactorSession) SetHiddenMetadataUri(_hiddenMetadataUri string) (*types.Transaction, error) {
	return _CryptoPokers.Contract.SetHiddenMetadataUri(&_CryptoPokers.TransactOpts, _hiddenMetadataUri)
}

// SetMaxMintAmountPerTx is a paid mutator transaction binding the contract method 0xb071401b.
//
// Solidity: function setMaxMintAmountPerTx(uint256 _maxMintAmountPerTx) returns()
func (_CryptoPokers *CryptoPokersTransactor) SetMaxMintAmountPerTx(opts *bind.TransactOpts, _maxMintAmountPerTx *big.Int) (*types.Transaction, error) {
	return _CryptoPokers.contract.Transact(opts, "setMaxMintAmountPerTx", _maxMintAmountPerTx)
}

// SetMaxMintAmountPerTx is a paid mutator transaction binding the contract method 0xb071401b.
//
// Solidity: function setMaxMintAmountPerTx(uint256 _maxMintAmountPerTx) returns()
func (_CryptoPokers *CryptoPokersSession) SetMaxMintAmountPerTx(_maxMintAmountPerTx *big.Int) (*types.Transaction, error) {
	return _CryptoPokers.Contract.SetMaxMintAmountPerTx(&_CryptoPokers.TransactOpts, _maxMintAmountPerTx)
}

// SetMaxMintAmountPerTx is a paid mutator transaction binding the contract method 0xb071401b.
//
// Solidity: function setMaxMintAmountPerTx(uint256 _maxMintAmountPerTx) returns()
func (_CryptoPokers *CryptoPokersTransactorSession) SetMaxMintAmountPerTx(_maxMintAmountPerTx *big.Int) (*types.Transaction, error) {
	return _CryptoPokers.Contract.SetMaxMintAmountPerTx(&_CryptoPokers.TransactOpts, _maxMintAmountPerTx)
}

// SetMerkleRoot is a paid mutator transaction binding the contract method 0x7cb64759.
//
// Solidity: function setMerkleRoot(bytes32 _merkleRoot) returns()
func (_CryptoPokers *CryptoPokersTransactor) SetMerkleRoot(opts *bind.TransactOpts, _merkleRoot [32]byte) (*types.Transaction, error) {
	return _CryptoPokers.contract.Transact(opts, "setMerkleRoot", _merkleRoot)
}

// SetMerkleRoot is a paid mutator transaction binding the contract method 0x7cb64759.
//
// Solidity: function setMerkleRoot(bytes32 _merkleRoot) returns()
func (_CryptoPokers *CryptoPokersSession) SetMerkleRoot(_merkleRoot [32]byte) (*types.Transaction, error) {
	return _CryptoPokers.Contract.SetMerkleRoot(&_CryptoPokers.TransactOpts, _merkleRoot)
}

// SetMerkleRoot is a paid mutator transaction binding the contract method 0x7cb64759.
//
// Solidity: function setMerkleRoot(bytes32 _merkleRoot) returns()
func (_CryptoPokers *CryptoPokersTransactorSession) SetMerkleRoot(_merkleRoot [32]byte) (*types.Transaction, error) {
	return _CryptoPokers.Contract.SetMerkleRoot(&_CryptoPokers.TransactOpts, _merkleRoot)
}

// SetPaused is a paid mutator transaction binding the contract method 0x16c38b3c.
//
// Solidity: function setPaused(bool _state) returns()
func (_CryptoPokers *CryptoPokersTransactor) SetPaused(opts *bind.TransactOpts, _state bool) (*types.Transaction, error) {
	return _CryptoPokers.contract.Transact(opts, "setPaused", _state)
}

// SetPaused is a paid mutator transaction binding the contract method 0x16c38b3c.
//
// Solidity: function setPaused(bool _state) returns()
func (_CryptoPokers *CryptoPokersSession) SetPaused(_state bool) (*types.Transaction, error) {
	return _CryptoPokers.Contract.SetPaused(&_CryptoPokers.TransactOpts, _state)
}

// SetPaused is a paid mutator transaction binding the contract method 0x16c38b3c.
//
// Solidity: function setPaused(bool _state) returns()
func (_CryptoPokers *CryptoPokersTransactorSession) SetPaused(_state bool) (*types.Transaction, error) {
	return _CryptoPokers.Contract.SetPaused(&_CryptoPokers.TransactOpts, _state)
}

// SetRevealed is a paid mutator transaction binding the contract method 0xe0a80853.
//
// Solidity: function setRevealed(bool _state) returns()
func (_CryptoPokers *CryptoPokersTransactor) SetRevealed(opts *bind.TransactOpts, _state bool) (*types.Transaction, error) {
	return _CryptoPokers.contract.Transact(opts, "setRevealed", _state)
}

// SetRevealed is a paid mutator transaction binding the contract method 0xe0a80853.
//
// Solidity: function setRevealed(bool _state) returns()
func (_CryptoPokers *CryptoPokersSession) SetRevealed(_state bool) (*types.Transaction, error) {
	return _CryptoPokers.Contract.SetRevealed(&_CryptoPokers.TransactOpts, _state)
}

// SetRevealed is a paid mutator transaction binding the contract method 0xe0a80853.
//
// Solidity: function setRevealed(bool _state) returns()
func (_CryptoPokers *CryptoPokersTransactorSession) SetRevealed(_state bool) (*types.Transaction, error) {
	return _CryptoPokers.Contract.SetRevealed(&_CryptoPokers.TransactOpts, _state)
}

// SetRoyaltyBasis is a paid mutator transaction binding the contract method 0xae40c3e0.
//
// Solidity: function setRoyaltyBasis(uint256 _royaltyBasis) returns()
func (_CryptoPokers *CryptoPokersTransactor) SetRoyaltyBasis(opts *bind.TransactOpts, _royaltyBasis *big.Int) (*types.Transaction, error) {
	return _CryptoPokers.contract.Transact(opts, "setRoyaltyBasis", _royaltyBasis)
}

// SetRoyaltyBasis is a paid mutator transaction binding the contract method 0xae40c3e0.
//
// Solidity: function setRoyaltyBasis(uint256 _royaltyBasis) returns()
func (_CryptoPokers *CryptoPokersSession) SetRoyaltyBasis(_royaltyBasis *big.Int) (*types.Transaction, error) {
	return _CryptoPokers.Contract.SetRoyaltyBasis(&_CryptoPokers.TransactOpts, _royaltyBasis)
}

// SetRoyaltyBasis is a paid mutator transaction binding the contract method 0xae40c3e0.
//
// Solidity: function setRoyaltyBasis(uint256 _royaltyBasis) returns()
func (_CryptoPokers *CryptoPokersTransactorSession) SetRoyaltyBasis(_royaltyBasis *big.Int) (*types.Transaction, error) {
	return _CryptoPokers.Contract.SetRoyaltyBasis(&_CryptoPokers.TransactOpts, _royaltyBasis)
}

// SetRoyaltyWallet is a paid mutator transaction binding the contract method 0xcdeee637.
//
// Solidity: function setRoyaltyWallet(address _royaltyWallet) returns()
func (_CryptoPokers *CryptoPokersTransactor) SetRoyaltyWallet(opts *bind.TransactOpts, _royaltyWallet common.Address) (*types.Transaction, error) {
	return _CryptoPokers.contract.Transact(opts, "setRoyaltyWallet", _royaltyWallet)
}

// SetRoyaltyWallet is a paid mutator transaction binding the contract method 0xcdeee637.
//
// Solidity: function setRoyaltyWallet(address _royaltyWallet) returns()
func (_CryptoPokers *CryptoPokersSession) SetRoyaltyWallet(_royaltyWallet common.Address) (*types.Transaction, error) {
	return _CryptoPokers.Contract.SetRoyaltyWallet(&_CryptoPokers.TransactOpts, _royaltyWallet)
}

// SetRoyaltyWallet is a paid mutator transaction binding the contract method 0xcdeee637.
//
// Solidity: function setRoyaltyWallet(address _royaltyWallet) returns()
func (_CryptoPokers *CryptoPokersTransactorSession) SetRoyaltyWallet(_royaltyWallet common.Address) (*types.Transaction, error) {
	return _CryptoPokers.Contract.SetRoyaltyWallet(&_CryptoPokers.TransactOpts, _royaltyWallet)
}

// SetUriPrefix is a paid mutator transaction binding the contract method 0x7ec4a659.
//
// Solidity: function setUriPrefix(string _uriPrefix) returns()
func (_CryptoPokers *CryptoPokersTransactor) SetUriPrefix(opts *bind.TransactOpts, _uriPrefix string) (*types.Transaction, error) {
	return _CryptoPokers.contract.Transact(opts, "setUriPrefix", _uriPrefix)
}

// SetUriPrefix is a paid mutator transaction binding the contract method 0x7ec4a659.
//
// Solidity: function setUriPrefix(string _uriPrefix) returns()
func (_CryptoPokers *CryptoPokersSession) SetUriPrefix(_uriPrefix string) (*types.Transaction, error) {
	return _CryptoPokers.Contract.SetUriPrefix(&_CryptoPokers.TransactOpts, _uriPrefix)
}

// SetUriPrefix is a paid mutator transaction binding the contract method 0x7ec4a659.
//
// Solidity: function setUriPrefix(string _uriPrefix) returns()
func (_CryptoPokers *CryptoPokersTransactorSession) SetUriPrefix(_uriPrefix string) (*types.Transaction, error) {
	return _CryptoPokers.Contract.SetUriPrefix(&_CryptoPokers.TransactOpts, _uriPrefix)
}

// SetUriSuffix is a paid mutator transaction binding the contract method 0x16ba10e0.
//
// Solidity: function setUriSuffix(string _uriSuffix) returns()
func (_CryptoPokers *CryptoPokersTransactor) SetUriSuffix(opts *bind.TransactOpts, _uriSuffix string) (*types.Transaction, error) {
	return _CryptoPokers.contract.Transact(opts, "setUriSuffix", _uriSuffix)
}

// SetUriSuffix is a paid mutator transaction binding the contract method 0x16ba10e0.
//
// Solidity: function setUriSuffix(string _uriSuffix) returns()
func (_CryptoPokers *CryptoPokersSession) SetUriSuffix(_uriSuffix string) (*types.Transaction, error) {
	return _CryptoPokers.Contract.SetUriSuffix(&_CryptoPokers.TransactOpts, _uriSuffix)
}

// SetUriSuffix is a paid mutator transaction binding the contract method 0x16ba10e0.
//
// Solidity: function setUriSuffix(string _uriSuffix) returns()
func (_CryptoPokers *CryptoPokersTransactorSession) SetUriSuffix(_uriSuffix string) (*types.Transaction, error) {
	return _CryptoPokers.Contract.SetUriSuffix(&_CryptoPokers.TransactOpts, _uriSuffix)
}

// SetWithdrawalWallet is a paid mutator transaction binding the contract method 0x75796f76.
//
// Solidity: function setWithdrawalWallet(address _withdrawalWallet) returns()
func (_CryptoPokers *CryptoPokersTransactor) SetWithdrawalWallet(opts *bind.TransactOpts, _withdrawalWallet common.Address) (*types.Transaction, error) {
	return _CryptoPokers.contract.Transact(opts, "setWithdrawalWallet", _withdrawalWallet)
}

// SetWithdrawalWallet is a paid mutator transaction binding the contract method 0x75796f76.
//
// Solidity: function setWithdrawalWallet(address _withdrawalWallet) returns()
func (_CryptoPokers *CryptoPokersSession) SetWithdrawalWallet(_withdrawalWallet common.Address) (*types.Transaction, error) {
	return _CryptoPokers.Contract.SetWithdrawalWallet(&_CryptoPokers.TransactOpts, _withdrawalWallet)
}

// SetWithdrawalWallet is a paid mutator transaction binding the contract method 0x75796f76.
//
// Solidity: function setWithdrawalWallet(address _withdrawalWallet) returns()
func (_CryptoPokers *CryptoPokersTransactorSession) SetWithdrawalWallet(_withdrawalWallet common.Address) (*types.Transaction, error) {
	return _CryptoPokers.Contract.SetWithdrawalWallet(&_CryptoPokers.TransactOpts, _withdrawalWallet)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) payable returns()
func (_CryptoPokers *CryptoPokersTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CryptoPokers.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) payable returns()
func (_CryptoPokers *CryptoPokersSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CryptoPokers.Contract.TransferFrom(&_CryptoPokers.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) payable returns()
func (_CryptoPokers *CryptoPokersTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CryptoPokers.Contract.TransferFrom(&_CryptoPokers.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CryptoPokers *CryptoPokersTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CryptoPokers.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CryptoPokers *CryptoPokersSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CryptoPokers.Contract.TransferOwnership(&_CryptoPokers.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CryptoPokers *CryptoPokersTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CryptoPokers.Contract.TransferOwnership(&_CryptoPokers.TransactOpts, newOwner)
}

// UpdateOperatorFilterRegistryAddress is a paid mutator transaction binding the contract method 0xb8d1e532.
//
// Solidity: function updateOperatorFilterRegistryAddress(address newRegistry) returns()
func (_CryptoPokers *CryptoPokersTransactor) UpdateOperatorFilterRegistryAddress(opts *bind.TransactOpts, newRegistry common.Address) (*types.Transaction, error) {
	return _CryptoPokers.contract.Transact(opts, "updateOperatorFilterRegistryAddress", newRegistry)
}

// UpdateOperatorFilterRegistryAddress is a paid mutator transaction binding the contract method 0xb8d1e532.
//
// Solidity: function updateOperatorFilterRegistryAddress(address newRegistry) returns()
func (_CryptoPokers *CryptoPokersSession) UpdateOperatorFilterRegistryAddress(newRegistry common.Address) (*types.Transaction, error) {
	return _CryptoPokers.Contract.UpdateOperatorFilterRegistryAddress(&_CryptoPokers.TransactOpts, newRegistry)
}

// UpdateOperatorFilterRegistryAddress is a paid mutator transaction binding the contract method 0xb8d1e532.
//
// Solidity: function updateOperatorFilterRegistryAddress(address newRegistry) returns()
func (_CryptoPokers *CryptoPokersTransactorSession) UpdateOperatorFilterRegistryAddress(newRegistry common.Address) (*types.Transaction, error) {
	return _CryptoPokers.Contract.UpdateOperatorFilterRegistryAddress(&_CryptoPokers.TransactOpts, newRegistry)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_CryptoPokers *CryptoPokersTransactor) Withdraw(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _CryptoPokers.contract.Transact(opts, "withdraw", _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_CryptoPokers *CryptoPokersSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _CryptoPokers.Contract.Withdraw(&_CryptoPokers.TransactOpts, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_CryptoPokers *CryptoPokersTransactorSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _CryptoPokers.Contract.Withdraw(&_CryptoPokers.TransactOpts, _amount)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() returns()
func (_CryptoPokers *CryptoPokersTransactor) WithdrawAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CryptoPokers.contract.Transact(opts, "withdrawAll")
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() returns()
func (_CryptoPokers *CryptoPokersSession) WithdrawAll() (*types.Transaction, error) {
	return _CryptoPokers.Contract.WithdrawAll(&_CryptoPokers.TransactOpts)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() returns()
func (_CryptoPokers *CryptoPokersTransactorSession) WithdrawAll() (*types.Transaction, error) {
	return _CryptoPokers.Contract.WithdrawAll(&_CryptoPokers.TransactOpts)
}

// CryptoPokersApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the CryptoPokers contract.
type CryptoPokersApprovalIterator struct {
	Event *CryptoPokersApproval // Event containing the contract specifics and raw log

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
func (it *CryptoPokersApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CryptoPokersApproval)
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
		it.Event = new(CryptoPokersApproval)
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
func (it *CryptoPokersApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CryptoPokersApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CryptoPokersApproval represents a Approval event raised by the CryptoPokers contract.
type CryptoPokersApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_CryptoPokers *CryptoPokersFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*CryptoPokersApprovalIterator, error) {

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

	logs, sub, err := _CryptoPokers.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CryptoPokersApprovalIterator{contract: _CryptoPokers.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_CryptoPokers *CryptoPokersFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CryptoPokersApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _CryptoPokers.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CryptoPokersApproval)
				if err := _CryptoPokers.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_CryptoPokers *CryptoPokersFilterer) ParseApproval(log types.Log) (*CryptoPokersApproval, error) {
	event := new(CryptoPokersApproval)
	if err := _CryptoPokers.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CryptoPokersApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the CryptoPokers contract.
type CryptoPokersApprovalForAllIterator struct {
	Event *CryptoPokersApprovalForAll // Event containing the contract specifics and raw log

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
func (it *CryptoPokersApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CryptoPokersApprovalForAll)
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
		it.Event = new(CryptoPokersApprovalForAll)
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
func (it *CryptoPokersApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CryptoPokersApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CryptoPokersApprovalForAll represents a ApprovalForAll event raised by the CryptoPokers contract.
type CryptoPokersApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_CryptoPokers *CryptoPokersFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*CryptoPokersApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _CryptoPokers.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &CryptoPokersApprovalForAllIterator{contract: _CryptoPokers.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_CryptoPokers *CryptoPokersFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *CryptoPokersApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _CryptoPokers.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CryptoPokersApprovalForAll)
				if err := _CryptoPokers.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_CryptoPokers *CryptoPokersFilterer) ParseApprovalForAll(log types.Log) (*CryptoPokersApprovalForAll, error) {
	event := new(CryptoPokersApprovalForAll)
	if err := _CryptoPokers.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CryptoPokersConsecutiveTransferIterator is returned from FilterConsecutiveTransfer and is used to iterate over the raw logs and unpacked data for ConsecutiveTransfer events raised by the CryptoPokers contract.
type CryptoPokersConsecutiveTransferIterator struct {
	Event *CryptoPokersConsecutiveTransfer // Event containing the contract specifics and raw log

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
func (it *CryptoPokersConsecutiveTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CryptoPokersConsecutiveTransfer)
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
		it.Event = new(CryptoPokersConsecutiveTransfer)
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
func (it *CryptoPokersConsecutiveTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CryptoPokersConsecutiveTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CryptoPokersConsecutiveTransfer represents a ConsecutiveTransfer event raised by the CryptoPokers contract.
type CryptoPokersConsecutiveTransfer struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	From        common.Address
	To          common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterConsecutiveTransfer is a free log retrieval operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_CryptoPokers *CryptoPokersFilterer) FilterConsecutiveTransfer(opts *bind.FilterOpts, fromTokenId []*big.Int, from []common.Address, to []common.Address) (*CryptoPokersConsecutiveTransferIterator, error) {

	var fromTokenIdRule []interface{}
	for _, fromTokenIdItem := range fromTokenId {
		fromTokenIdRule = append(fromTokenIdRule, fromTokenIdItem)
	}

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CryptoPokers.contract.FilterLogs(opts, "ConsecutiveTransfer", fromTokenIdRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CryptoPokersConsecutiveTransferIterator{contract: _CryptoPokers.contract, event: "ConsecutiveTransfer", logs: logs, sub: sub}, nil
}

// WatchConsecutiveTransfer is a free log subscription operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_CryptoPokers *CryptoPokersFilterer) WatchConsecutiveTransfer(opts *bind.WatchOpts, sink chan<- *CryptoPokersConsecutiveTransfer, fromTokenId []*big.Int, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromTokenIdRule []interface{}
	for _, fromTokenIdItem := range fromTokenId {
		fromTokenIdRule = append(fromTokenIdRule, fromTokenIdItem)
	}

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CryptoPokers.contract.WatchLogs(opts, "ConsecutiveTransfer", fromTokenIdRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CryptoPokersConsecutiveTransfer)
				if err := _CryptoPokers.contract.UnpackLog(event, "ConsecutiveTransfer", log); err != nil {
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

// ParseConsecutiveTransfer is a log parse operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_CryptoPokers *CryptoPokersFilterer) ParseConsecutiveTransfer(log types.Log) (*CryptoPokersConsecutiveTransfer, error) {
	event := new(CryptoPokersConsecutiveTransfer)
	if err := _CryptoPokers.contract.UnpackLog(event, "ConsecutiveTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CryptoPokersOperatorFilterRegistryAddressUpdatedIterator is returned from FilterOperatorFilterRegistryAddressUpdated and is used to iterate over the raw logs and unpacked data for OperatorFilterRegistryAddressUpdated events raised by the CryptoPokers contract.
type CryptoPokersOperatorFilterRegistryAddressUpdatedIterator struct {
	Event *CryptoPokersOperatorFilterRegistryAddressUpdated // Event containing the contract specifics and raw log

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
func (it *CryptoPokersOperatorFilterRegistryAddressUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CryptoPokersOperatorFilterRegistryAddressUpdated)
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
		it.Event = new(CryptoPokersOperatorFilterRegistryAddressUpdated)
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
func (it *CryptoPokersOperatorFilterRegistryAddressUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CryptoPokersOperatorFilterRegistryAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CryptoPokersOperatorFilterRegistryAddressUpdated represents a OperatorFilterRegistryAddressUpdated event raised by the CryptoPokers contract.
type CryptoPokersOperatorFilterRegistryAddressUpdated struct {
	NewRegistry common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorFilterRegistryAddressUpdated is a free log retrieval operation binding the contract event 0x9f513fe86dc42fdbac355fa4d9b1d5be7b5e6cd2df67e30db8003766568de476.
//
// Solidity: event OperatorFilterRegistryAddressUpdated(address newRegistry)
func (_CryptoPokers *CryptoPokersFilterer) FilterOperatorFilterRegistryAddressUpdated(opts *bind.FilterOpts) (*CryptoPokersOperatorFilterRegistryAddressUpdatedIterator, error) {

	logs, sub, err := _CryptoPokers.contract.FilterLogs(opts, "OperatorFilterRegistryAddressUpdated")
	if err != nil {
		return nil, err
	}
	return &CryptoPokersOperatorFilterRegistryAddressUpdatedIterator{contract: _CryptoPokers.contract, event: "OperatorFilterRegistryAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchOperatorFilterRegistryAddressUpdated is a free log subscription operation binding the contract event 0x9f513fe86dc42fdbac355fa4d9b1d5be7b5e6cd2df67e30db8003766568de476.
//
// Solidity: event OperatorFilterRegistryAddressUpdated(address newRegistry)
func (_CryptoPokers *CryptoPokersFilterer) WatchOperatorFilterRegistryAddressUpdated(opts *bind.WatchOpts, sink chan<- *CryptoPokersOperatorFilterRegistryAddressUpdated) (event.Subscription, error) {

	logs, sub, err := _CryptoPokers.contract.WatchLogs(opts, "OperatorFilterRegistryAddressUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CryptoPokersOperatorFilterRegistryAddressUpdated)
				if err := _CryptoPokers.contract.UnpackLog(event, "OperatorFilterRegistryAddressUpdated", log); err != nil {
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

// ParseOperatorFilterRegistryAddressUpdated is a log parse operation binding the contract event 0x9f513fe86dc42fdbac355fa4d9b1d5be7b5e6cd2df67e30db8003766568de476.
//
// Solidity: event OperatorFilterRegistryAddressUpdated(address newRegistry)
func (_CryptoPokers *CryptoPokersFilterer) ParseOperatorFilterRegistryAddressUpdated(log types.Log) (*CryptoPokersOperatorFilterRegistryAddressUpdated, error) {
	event := new(CryptoPokersOperatorFilterRegistryAddressUpdated)
	if err := _CryptoPokers.contract.UnpackLog(event, "OperatorFilterRegistryAddressUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CryptoPokersOperatorFilterRegistryRevokedIterator is returned from FilterOperatorFilterRegistryRevoked and is used to iterate over the raw logs and unpacked data for OperatorFilterRegistryRevoked events raised by the CryptoPokers contract.
type CryptoPokersOperatorFilterRegistryRevokedIterator struct {
	Event *CryptoPokersOperatorFilterRegistryRevoked // Event containing the contract specifics and raw log

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
func (it *CryptoPokersOperatorFilterRegistryRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CryptoPokersOperatorFilterRegistryRevoked)
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
		it.Event = new(CryptoPokersOperatorFilterRegistryRevoked)
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
func (it *CryptoPokersOperatorFilterRegistryRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CryptoPokersOperatorFilterRegistryRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CryptoPokersOperatorFilterRegistryRevoked represents a OperatorFilterRegistryRevoked event raised by the CryptoPokers contract.
type CryptoPokersOperatorFilterRegistryRevoked struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterOperatorFilterRegistryRevoked is a free log retrieval operation binding the contract event 0x51e2d870cc2e10853e38dc06fcdae46ad3c3f588f326608803dac6204541ad16.
//
// Solidity: event OperatorFilterRegistryRevoked()
func (_CryptoPokers *CryptoPokersFilterer) FilterOperatorFilterRegistryRevoked(opts *bind.FilterOpts) (*CryptoPokersOperatorFilterRegistryRevokedIterator, error) {

	logs, sub, err := _CryptoPokers.contract.FilterLogs(opts, "OperatorFilterRegistryRevoked")
	if err != nil {
		return nil, err
	}
	return &CryptoPokersOperatorFilterRegistryRevokedIterator{contract: _CryptoPokers.contract, event: "OperatorFilterRegistryRevoked", logs: logs, sub: sub}, nil
}

// WatchOperatorFilterRegistryRevoked is a free log subscription operation binding the contract event 0x51e2d870cc2e10853e38dc06fcdae46ad3c3f588f326608803dac6204541ad16.
//
// Solidity: event OperatorFilterRegistryRevoked()
func (_CryptoPokers *CryptoPokersFilterer) WatchOperatorFilterRegistryRevoked(opts *bind.WatchOpts, sink chan<- *CryptoPokersOperatorFilterRegistryRevoked) (event.Subscription, error) {

	logs, sub, err := _CryptoPokers.contract.WatchLogs(opts, "OperatorFilterRegistryRevoked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CryptoPokersOperatorFilterRegistryRevoked)
				if err := _CryptoPokers.contract.UnpackLog(event, "OperatorFilterRegistryRevoked", log); err != nil {
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

// ParseOperatorFilterRegistryRevoked is a log parse operation binding the contract event 0x51e2d870cc2e10853e38dc06fcdae46ad3c3f588f326608803dac6204541ad16.
//
// Solidity: event OperatorFilterRegistryRevoked()
func (_CryptoPokers *CryptoPokersFilterer) ParseOperatorFilterRegistryRevoked(log types.Log) (*CryptoPokersOperatorFilterRegistryRevoked, error) {
	event := new(CryptoPokersOperatorFilterRegistryRevoked)
	if err := _CryptoPokers.contract.UnpackLog(event, "OperatorFilterRegistryRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CryptoPokersOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CryptoPokers contract.
type CryptoPokersOwnershipTransferredIterator struct {
	Event *CryptoPokersOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CryptoPokersOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CryptoPokersOwnershipTransferred)
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
		it.Event = new(CryptoPokersOwnershipTransferred)
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
func (it *CryptoPokersOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CryptoPokersOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CryptoPokersOwnershipTransferred represents a OwnershipTransferred event raised by the CryptoPokers contract.
type CryptoPokersOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CryptoPokers *CryptoPokersFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CryptoPokersOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CryptoPokers.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CryptoPokersOwnershipTransferredIterator{contract: _CryptoPokers.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CryptoPokers *CryptoPokersFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CryptoPokersOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CryptoPokers.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CryptoPokersOwnershipTransferred)
				if err := _CryptoPokers.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_CryptoPokers *CryptoPokersFilterer) ParseOwnershipTransferred(log types.Log) (*CryptoPokersOwnershipTransferred, error) {
	event := new(CryptoPokersOwnershipTransferred)
	if err := _CryptoPokers.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CryptoPokersPermanentURIIterator is returned from FilterPermanentURI and is used to iterate over the raw logs and unpacked data for PermanentURI events raised by the CryptoPokers contract.
type CryptoPokersPermanentURIIterator struct {
	Event *CryptoPokersPermanentURI // Event containing the contract specifics and raw log

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
func (it *CryptoPokersPermanentURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CryptoPokersPermanentURI)
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
		it.Event = new(CryptoPokersPermanentURI)
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
func (it *CryptoPokersPermanentURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CryptoPokersPermanentURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CryptoPokersPermanentURI represents a PermanentURI event raised by the CryptoPokers contract.
type CryptoPokersPermanentURI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPermanentURI is a free log retrieval operation binding the contract event 0xa109ba539900bf1b633f956d63c96fc89b814c7287f7aa50a9216d0b55657207.
//
// Solidity: event PermanentURI(string _value, uint256 indexed _id)
func (_CryptoPokers *CryptoPokersFilterer) FilterPermanentURI(opts *bind.FilterOpts, _id []*big.Int) (*CryptoPokersPermanentURIIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _CryptoPokers.contract.FilterLogs(opts, "PermanentURI", _idRule)
	if err != nil {
		return nil, err
	}
	return &CryptoPokersPermanentURIIterator{contract: _CryptoPokers.contract, event: "PermanentURI", logs: logs, sub: sub}, nil
}

// WatchPermanentURI is a free log subscription operation binding the contract event 0xa109ba539900bf1b633f956d63c96fc89b814c7287f7aa50a9216d0b55657207.
//
// Solidity: event PermanentURI(string _value, uint256 indexed _id)
func (_CryptoPokers *CryptoPokersFilterer) WatchPermanentURI(opts *bind.WatchOpts, sink chan<- *CryptoPokersPermanentURI, _id []*big.Int) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _CryptoPokers.contract.WatchLogs(opts, "PermanentURI", _idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CryptoPokersPermanentURI)
				if err := _CryptoPokers.contract.UnpackLog(event, "PermanentURI", log); err != nil {
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

// ParsePermanentURI is a log parse operation binding the contract event 0xa109ba539900bf1b633f956d63c96fc89b814c7287f7aa50a9216d0b55657207.
//
// Solidity: event PermanentURI(string _value, uint256 indexed _id)
func (_CryptoPokers *CryptoPokersFilterer) ParsePermanentURI(log types.Log) (*CryptoPokersPermanentURI, error) {
	event := new(CryptoPokersPermanentURI)
	if err := _CryptoPokers.contract.UnpackLog(event, "PermanentURI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CryptoPokersTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the CryptoPokers contract.
type CryptoPokersTransferIterator struct {
	Event *CryptoPokersTransfer // Event containing the contract specifics and raw log

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
func (it *CryptoPokersTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CryptoPokersTransfer)
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
		it.Event = new(CryptoPokersTransfer)
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
func (it *CryptoPokersTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CryptoPokersTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CryptoPokersTransfer represents a Transfer event raised by the CryptoPokers contract.
type CryptoPokersTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_CryptoPokers *CryptoPokersFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*CryptoPokersTransferIterator, error) {

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

	logs, sub, err := _CryptoPokers.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CryptoPokersTransferIterator{contract: _CryptoPokers.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_CryptoPokers *CryptoPokersFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CryptoPokersTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _CryptoPokers.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CryptoPokersTransfer)
				if err := _CryptoPokers.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_CryptoPokers *CryptoPokersFilterer) ParseTransfer(log types.Log) (*CryptoPokersTransfer, error) {
	event := new(CryptoPokersTransfer)
	if err := _CryptoPokers.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
