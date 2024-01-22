// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package test

import (
	"math/big"
	"strings"

	"github.com/FISCO-BCOS/go-sdk/abi"
	"github.com/FISCO-BCOS/go-sdk/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
)

// ElectronicSealABI is the input ABI used to generate the binding from.
const ElectronicSealABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_cardId\",\"type\":\"string\"},{\"name\":\"_datetime\",\"type\":\"string\"}],\"name\":\"addSealAccount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_hash\",\"type\":\"string\"},{\"name\":\"_code\",\"type\":\"string\"},{\"name\":\"_datetime\",\"type\":\"string\"}],\"name\":\"sealSignature\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"getSealAccount\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_hash\",\"type\":\"string\"}],\"name\":\"querySignature\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"queryAllHash\",\"outputs\":[{\"name\":\"\",\"type\":\"string[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// ElectronicSealBin is the compiled bytecode used for deploying new contracts.
var ElectronicSealBin = "0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550612121806100606000396000f30060806040526004361062000073576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680635c4997ea14620000785780636d0cfa4814620000a657806370f2f50514620000d4578063c3602599146200011a578063c547120c146200015f575b600080fd5b3480156200008557600080fd5b50620000a460048036036200009e919081019062000a21565b620001a3565b005b348015620000b357600080fd5b50620000d26004803603620000cc919081019062000a21565b620002f9565b005b348015620000e157600080fd5b50620001006004803603620000fa91908101906200099b565b62000487565b604051620001119392919062000ddb565b60405180910390f35b3480156200012757600080fd5b50620001466004803603620001409190810190620009c7565b620005a3565b6040516200015692919062000da0565b60405180910390f35b3480156200016c57600080fd5b506200018b60048036036200018591908101906200099b565b620006c8565b6040516200019a919062000d58565b60405180910390f35b60003373ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614151562000239576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620002309062000e2d565b60405180910390fd5b83838362000246620007db565b620002549392919062000ddb565b604051809103906000f08015801562000271573d6000803e3d6000fd5b50905080600160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050505050565b60003373ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415156200038f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620003869062000e2d565b60405180910390fd5b600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508073ffffffffffffffffffffffffffffffffffffffff1663d65503558585856040518463ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004016200044c9392919062000ddb565b600060405180830381600087803b1580156200046757600080fd5b505af11580156200047c573d6000803e3d6000fd5b505050505050505050565b60608060606000600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508073ffffffffffffffffffffffffffffffffffffffff16633a0f692c6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401600060405180830381600087803b1580156200055557600080fd5b505af11580156200056a573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525062000595919081019062000b90565b935093509350509193909250565b6060806000600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508073ffffffffffffffffffffffffffffffffffffffff166311dd8845856040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040162000661919062000d7c565b600060405180830381600087803b1580156200067c57600080fd5b505af115801562000691573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250620006bc919081019062000b1d565b92509250509250929050565b60606000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508073ffffffffffffffffffffffffffffffffffffffff1663560b318c6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401600060405180830381600087803b1580156200079357600080fd5b505af1158015620007a8573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250620007d3919081019062000ad8565b915050919050565b6040516111368062000fb283390190565b6000620007fa823562000f3b565b905092915050565b600082601f83011215156200081657600080fd5b81516200082d620008278262000e7d565b62000e4f565b9150818183526020840193506020810190508360005b838110156200087757815186016200085c888262000881565b84526020840193506020830192505060018101905062000843565b5050505092915050565b600082601f83011215156200089557600080fd5b8151620008ac620008a68262000ea6565b62000e4f565b91508082526020830160208301858383011115620008c957600080fd5b620008d683828462000f6a565b50505092915050565b600082601f8301121515620008f357600080fd5b81356200090a620009048262000ed3565b62000e4f565b915080825260208301602083018583830111156200092757600080fd5b6200093483828462000f5b565b50505092915050565b600082601f83011215156200095157600080fd5b815162000968620009628262000ed3565b62000e4f565b915080825260208301602083018583830111156200098557600080fd5b6200099283828462000f6a565b50505092915050565b600060208284031215620009ae57600080fd5b6000620009be84828501620007ec565b91505092915050565b60008060408385031215620009db57600080fd5b6000620009eb85828601620007ec565b925050602083013567ffffffffffffffff81111562000a0957600080fd5b62000a1785828601620008df565b9150509250929050565b6000806000806080858703121562000a3857600080fd5b600062000a4887828801620007ec565b945050602085013567ffffffffffffffff81111562000a6657600080fd5b62000a7487828801620008df565b935050604085013567ffffffffffffffff81111562000a9257600080fd5b62000aa087828801620008df565b925050606085013567ffffffffffffffff81111562000abe57600080fd5b62000acc87828801620008df565b91505092959194509250565b60006020828403121562000aeb57600080fd5b600082015167ffffffffffffffff81111562000b0657600080fd5b62000b148482850162000802565b91505092915050565b6000806040838503121562000b3157600080fd5b600083015167ffffffffffffffff81111562000b4c57600080fd5b62000b5a858286016200093d565b925050602083015167ffffffffffffffff81111562000b7857600080fd5b62000b86858286016200093d565b9150509250929050565b60008060006060848603121562000ba657600080fd5b600084015167ffffffffffffffff81111562000bc157600080fd5b62000bcf868287016200093d565b935050602084015167ffffffffffffffff81111562000bed57600080fd5b62000bfb868287016200093d565b925050604084015167ffffffffffffffff81111562000c1957600080fd5b62000c27868287016200093d565b9150509250925092565b600062000c3e8262000f0d565b8084526020840193508360208202850162000c598562000f00565b60005b8481101562000c9857838303885262000c7783835162000ce5565b925062000c848262000f2e565b915060208801975060018101905062000c5c565b508196508694505050505092915050565b600062000cb68262000f23565b80845262000ccc81602086016020860162000f6a565b62000cd78162000fa0565b602085010191505092915050565b600062000cf28262000f18565b80845262000d0881602086016020860162000f6a565b62000d138162000fa0565b602085010191505092915050565b6000601382527f4f6e6c79206f776e65722063616e2063616c6c000000000000000000000000006020830152604082019050919050565b6000602082019050818103600083015262000d74818462000c31565b905092915050565b6000602082019050818103600083015262000d98818462000ca9565b905092915050565b6000604082019050818103600083015262000dbc818562000ca9565b9050818103602083015262000dd2818462000ca9565b90509392505050565b6000606082019050818103600083015262000df7818662000ca9565b9050818103602083015262000e0d818562000ca9565b9050818103604083015262000e23818462000ca9565b9050949350505050565b6000602082019050818103600083015262000e488162000d21565b9050919050565b6000604051905081810181811067ffffffffffffffff8211171562000e7357600080fd5b8060405250919050565b600067ffffffffffffffff82111562000e9557600080fd5b602082029050602081019050919050565b600067ffffffffffffffff82111562000ebe57600080fd5b601f19601f8301169050602081019050919050565b600067ffffffffffffffff82111562000eeb57600080fd5b601f19601f8301169050602081019050919050565b6000602082019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b82818337600083830152505050565b60005b8381101562000f8a57808201518184015260208101905062000f6d565b8381111562000f9a576000848401525b50505050565b6000601f19601f8301169050919050560060806040523480156200001157600080fd5b5060405162001136380380620011368339810180604052620000379190810190620001a0565b826000800190805190602001906200005192919062000093565b5081600060010190805190602001906200006d92919062000093565b5080600060020190805190602001906200008992919062000093565b50505050620002d2565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620000d657805160ff191683800117855562000107565b8280016001018555821562000107579182015b8281111562000106578251825591602001919060010190620000e9565b5b5090506200011691906200011a565b5090565b6200013f91905b808211156200013b57600081600090555060010162000121565b5090565b90565b600082601f83011215156200015657600080fd5b81516200016d62000167826200026f565b62000241565b915080825260208301602083018583830111156200018a57600080fd5b620001978382846200029c565b50505092915050565b600080600060608486031215620001b657600080fd5b600084015167ffffffffffffffff811115620001d157600080fd5b620001df8682870162000142565b935050602084015167ffffffffffffffff811115620001fd57600080fd5b6200020b8682870162000142565b925050604084015167ffffffffffffffff8111156200022957600080fd5b620002378682870162000142565b9150509250925092565b6000604051905081810181811067ffffffffffffffff821117156200026557600080fd5b8060405250919050565b600067ffffffffffffffff8211156200028757600080fd5b601f19601f8301169050602081019050919050565b60005b83811015620002bc5780820151818401526020810190506200029f565b83811115620002cc576000848401525b50505050565b610e5480620002e26000396000f30060806040526004361061006d576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806311dd8845146100725780633a0f692c146100b05780633c995055146100dd578063560b318c1461010a578063d655035514610135575b600080fd5b34801561007e57600080fd5b5061009960048036036100949190810190610a90565b61015e565b6040516100a7929190610c64565b60405180910390f35b3480156100bc57600080fd5b506100c561033f565b6040516100d493929190610c9b565b60405180910390f35b3480156100e957600080fd5b506100f2610531565b60405161010193929190610ce7565b60405180910390f35b34801561011657600080fd5b5061011f610711565b60405161012c9190610c42565b60405180910390f35b34801561014157600080fd5b5061015c60048036036101579190810190610ad1565b6107fa565b005b6060806101696108fb565b6003846040518082805190602001908083835b6020831015156101a1578051825260208201915060208101905060208303925061017c565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020604080519081016040529081600082018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156102785780601f1061024d57610100808354040283529160200191610278565b820191906000526020600020905b81548152906001019060200180831161025b57829003601f168201915b50505050508152602001600182018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561031a5780601f106102ef5761010080835404028352916020019161031a565b820191906000526020600020905b8154815290600101906020018083116102fd57829003601f168201915b5050505050815250509050806000015181602001518191508090509250925050915091565b60608060606000800160006001016000600201828054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156103e75780601f106103bc576101008083540402835291602001916103e7565b820191906000526020600020905b8154815290600101906020018083116103ca57829003601f168201915b50505050509250818054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156104835780601f1061045857610100808354040283529160200191610483565b820191906000526020600020905b81548152906001019060200180831161046657829003601f168201915b50505050509150808054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561051f5780601f106104f45761010080835404028352916020019161051f565b820191906000526020600020905b81548152906001019060200180831161050257829003601f168201915b50505050509050925092509250909192565b6000806000018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156105cb5780601f106105a0576101008083540402835291602001916105cb565b820191906000526020600020905b8154815290600101906020018083116105ae57829003601f168201915b505050505090806001018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156106695780601f1061063e57610100808354040283529160200191610669565b820191906000526020600020905b81548152906001019060200180831161064c57829003601f168201915b505050505090806002018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156107075780601f106106dc57610100808354040283529160200191610707565b820191906000526020600020905b8154815290600101906020018083116106ea57829003601f168201915b5050505050905083565b60606004805480602002602001604051908101604052809291908181526020016000905b828210156107f1578382906000526020600020018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156107dd5780601f106107b2576101008083540402835291602001916107dd565b820191906000526020600020905b8154815290600101906020018083116107c057829003601f168201915b505050505081526020019060010190610735565b50505050905090565b6040805190810160405280838152602001828152506003846040518082805190602001908083835b6020831015156108475780518252602082019150602081019050602083039250610822565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390206000820151816000019080519060200190610896929190610915565b5060208201518160010190805190602001906108b3929190610915565b5090505060048390806001815401808255809150509060018203906000526020600020016000909192909190915090805190602001906108f4929190610995565b5050505050565b604080519081016040528060608152602001606081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061095657805160ff1916838001178555610984565b82800160010185558215610984579182015b82811115610983578251825591602001919060010190610968565b5b5090506109919190610a15565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106109d657805160ff1916838001178555610a04565b82800160010185558215610a04579182015b82811115610a035782518255916020019190600101906109e8565b5b509050610a119190610a15565b5090565b610a3791905b80821115610a33576000816000905550600101610a1b565b5090565b90565b600082601f8301121515610a4d57600080fd5b8135610a60610a5b82610d60565b610d33565b91508082526020830160208301858383011115610a7c57600080fd5b610a87838284610dc7565b50505092915050565b600060208284031215610aa257600080fd5b600082013567ffffffffffffffff811115610abc57600080fd5b610ac884828501610a3a565b91505092915050565b600080600060608486031215610ae657600080fd5b600084013567ffffffffffffffff811115610b0057600080fd5b610b0c86828701610a3a565b935050602084013567ffffffffffffffff811115610b2957600080fd5b610b3586828701610a3a565b925050604084013567ffffffffffffffff811115610b5257600080fd5b610b5e86828701610a3a565b9150509250925092565b6000610b7382610d99565b80845260208401935083602082028501610b8c85610d8c565b60005b84811015610bc5578383038852610ba7838351610c0c565b9250610bb282610dba565b9150602088019750600181019050610b8f565b508196508694505050505092915050565b6000610be182610daf565b808452610bf5816020860160208601610dd6565b610bfe81610e09565b602085010191505092915050565b6000610c1782610da4565b808452610c2b816020860160208601610dd6565b610c3481610e09565b602085010191505092915050565b60006020820190508181036000830152610c5c8184610b68565b905092915050565b60006040820190508181036000830152610c7e8185610bd6565b90508181036020830152610c928184610bd6565b90509392505050565b60006060820190508181036000830152610cb58186610bd6565b90508181036020830152610cc98185610bd6565b90508181036040830152610cdd8184610bd6565b9050949350505050565b60006060820190508181036000830152610d018186610c0c565b90508181036020830152610d158185610c0c565b90508181036040830152610d298184610c0c565b9050949350505050565b6000604051905081810181811067ffffffffffffffff82111715610d5657600080fd5b8060405250919050565b600067ffffffffffffffff821115610d7757600080fd5b601f19601f8301169050602081019050919050565b6000602082019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b82818337600083830152505050565b60005b83811015610df4578082015181840152602081019050610dd9565b83811115610e03576000848401525b50505050565b6000601f19601f83011690509190505600a265627a7a72305820e7f1015f46cf51b05d3da1a94d059060a3f3c61d898d60cb25201ae1b38d02836c6578706572696d656e74616cf50037a265627a7a72305820d294ea8aabeaf4e1264e44fa8d139d23e1e319be14ad581e1b63d68c957ae46c6c6578706572696d656e74616cf50037"

// DeployElectronicSeal deploys a new contract, binding an instance of ElectronicSeal to it.
func DeployElectronicSeal(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ElectronicSeal, error) {
	parsed, err := abi.JSON(strings.NewReader(ElectronicSealABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ElectronicSealBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ElectronicSeal{ElectronicSealCaller: ElectronicSealCaller{contract: contract}, ElectronicSealTransactor: ElectronicSealTransactor{contract: contract}, ElectronicSealFilterer: ElectronicSealFilterer{contract: contract}}, nil
}

func AsyncDeployElectronicSeal(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(ElectronicSealABI))
	if err != nil {
		return nil, err
	}

	tx, err := bind.AsyncDeployContract(auth, handler, parsed, common.FromHex(ElectronicSealBin), backend)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// ElectronicSeal is an auto generated Go binding around a Solidity contract.
type ElectronicSeal struct {
	ElectronicSealCaller     // Read-only binding to the contract
	ElectronicSealTransactor // Write-only binding to the contract
	ElectronicSealFilterer   // Log filterer for contract events
}

// ElectronicSealCaller is an auto generated read-only Go binding around a Solidity contract.
type ElectronicSealCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ElectronicSealTransactor is an auto generated write-only Go binding around a Solidity contract.
type ElectronicSealTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ElectronicSealFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type ElectronicSealFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ElectronicSealSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type ElectronicSealSession struct {
	Contract     *ElectronicSeal   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ElectronicSealCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type ElectronicSealCallerSession struct {
	Contract *ElectronicSealCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ElectronicSealTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type ElectronicSealTransactorSession struct {
	Contract     *ElectronicSealTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ElectronicSealRaw is an auto generated low-level Go binding around a Solidity contract.
type ElectronicSealRaw struct {
	Contract *ElectronicSeal // Generic contract binding to access the raw methods on
}

// ElectronicSealCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type ElectronicSealCallerRaw struct {
	Contract *ElectronicSealCaller // Generic read-only contract binding to access the raw methods on
}

// ElectronicSealTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type ElectronicSealTransactorRaw struct {
	Contract *ElectronicSealTransactor // Generic write-only contract binding to access the raw methods on
}

// NewElectronicSeal creates a new instance of ElectronicSeal, bound to a specific deployed contract.
func NewElectronicSeal(address common.Address, backend bind.ContractBackend) (*ElectronicSeal, error) {
	contract, err := bindElectronicSeal(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ElectronicSeal{ElectronicSealCaller: ElectronicSealCaller{contract: contract}, ElectronicSealTransactor: ElectronicSealTransactor{contract: contract}, ElectronicSealFilterer: ElectronicSealFilterer{contract: contract}}, nil
}

// NewElectronicSealCaller creates a new read-only instance of ElectronicSeal, bound to a specific deployed contract.
func NewElectronicSealCaller(address common.Address, caller bind.ContractCaller) (*ElectronicSealCaller, error) {
	contract, err := bindElectronicSeal(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ElectronicSealCaller{contract: contract}, nil
}

// NewElectronicSealTransactor creates a new write-only instance of ElectronicSeal, bound to a specific deployed contract.
func NewElectronicSealTransactor(address common.Address, transactor bind.ContractTransactor) (*ElectronicSealTransactor, error) {
	contract, err := bindElectronicSeal(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ElectronicSealTransactor{contract: contract}, nil
}

// NewElectronicSealFilterer creates a new log filterer instance of ElectronicSeal, bound to a specific deployed contract.
func NewElectronicSealFilterer(address common.Address, filterer bind.ContractFilterer) (*ElectronicSealFilterer, error) {
	contract, err := bindElectronicSeal(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ElectronicSealFilterer{contract: contract}, nil
}

// bindElectronicSeal binds a generic wrapper to an already deployed contract.
func bindElectronicSeal(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ElectronicSealABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ElectronicSeal *ElectronicSealRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ElectronicSeal.Contract.ElectronicSealCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ElectronicSeal *ElectronicSealRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _ElectronicSeal.Contract.ElectronicSealTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ElectronicSeal *ElectronicSealRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _ElectronicSeal.Contract.ElectronicSealTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ElectronicSeal *ElectronicSealCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ElectronicSeal.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ElectronicSeal *ElectronicSealTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _ElectronicSeal.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ElectronicSeal *ElectronicSealTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _ElectronicSeal.Contract.contract.Transact(opts, method, params...)
}

// GetSealAccount is a free data retrieval call binding the contract method 0x70f2f505.
//
// Solidity: function getSealAccount(address _account) constant returns(string, string, string)
func (_ElectronicSeal *ElectronicSealCaller) GetSealAccount(opts *bind.CallOpts, _account common.Address) (string, string, string, error) {
	var (
		ret0 = new(string)
		ret1 = new(string)
		ret2 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _ElectronicSeal.contract.Call(opts, out, "getSealAccount", _account)
	return *ret0, *ret1, *ret2, err
}

// GetSealAccount is a free data retrieval call binding the contract method 0x70f2f505.
//
// Solidity: function getSealAccount(address _account) constant returns(string, string, string)
func (_ElectronicSeal *ElectronicSealSession) GetSealAccount(_account common.Address) (string, string, string, error) {
	return _ElectronicSeal.Contract.GetSealAccount(&_ElectronicSeal.CallOpts, _account)
}

// GetSealAccount is a free data retrieval call binding the contract method 0x70f2f505.
//
// Solidity: function getSealAccount(address _account) constant returns(string, string, string)
func (_ElectronicSeal *ElectronicSealCallerSession) GetSealAccount(_account common.Address) (string, string, string, error) {
	return _ElectronicSeal.Contract.GetSealAccount(&_ElectronicSeal.CallOpts, _account)
}

// QueryAllHash is a free data retrieval call binding the contract method 0xc547120c.
//
// Solidity: function queryAllHash(address _account) constant returns(string[])
func (_ElectronicSeal *ElectronicSealCaller) QueryAllHash(opts *bind.CallOpts, _account common.Address) ([]string, error) {
	var (
		ret0 = new([]string)
	)
	out := ret0
	err := _ElectronicSeal.contract.Call(opts, out, "queryAllHash", _account)
	return *ret0, err
}

// QueryAllHash is a free data retrieval call binding the contract method 0xc547120c.
//
// Solidity: function queryAllHash(address _account) constant returns(string[])
func (_ElectronicSeal *ElectronicSealSession) QueryAllHash(_account common.Address) ([]string, error) {
	return _ElectronicSeal.Contract.QueryAllHash(&_ElectronicSeal.CallOpts, _account)
}

// QueryAllHash is a free data retrieval call binding the contract method 0xc547120c.
//
// Solidity: function queryAllHash(address _account) constant returns(string[])
func (_ElectronicSeal *ElectronicSealCallerSession) QueryAllHash(_account common.Address) ([]string, error) {
	return _ElectronicSeal.Contract.QueryAllHash(&_ElectronicSeal.CallOpts, _account)
}

// QuerySignature is a free data retrieval call binding the contract method 0xc3602599.
//
// Solidity: function querySignature(address _account, string _hash) constant returns(string, string)
func (_ElectronicSeal *ElectronicSealCaller) QuerySignature(opts *bind.CallOpts, _account common.Address, _hash string) (string, string, error) {
	var (
		ret0 = new(string)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _ElectronicSeal.contract.Call(opts, out, "querySignature", _account, _hash)
	return *ret0, *ret1, err
}

// QuerySignature is a free data retrieval call binding the contract method 0xc3602599.
//
// Solidity: function querySignature(address _account, string _hash) constant returns(string, string)
func (_ElectronicSeal *ElectronicSealSession) QuerySignature(_account common.Address, _hash string) (string, string, error) {
	return _ElectronicSeal.Contract.QuerySignature(&_ElectronicSeal.CallOpts, _account, _hash)
}

// QuerySignature is a free data retrieval call binding the contract method 0xc3602599.
//
// Solidity: function querySignature(address _account, string _hash) constant returns(string, string)
func (_ElectronicSeal *ElectronicSealCallerSession) QuerySignature(_account common.Address, _hash string) (string, string, error) {
	return _ElectronicSeal.Contract.QuerySignature(&_ElectronicSeal.CallOpts, _account, _hash)
}

// AddSealAccount is a paid mutator transaction binding the contract method 0x5c4997ea.
//
// Solidity: function addSealAccount(address _account, string _name, string _cardId, string _datetime) returns()
func (_ElectronicSeal *ElectronicSealTransactor) AddSealAccount(opts *bind.TransactOpts, _account common.Address, _name string, _cardId string, _datetime string) (*types.Transaction, *types.Receipt, error) {
	return _ElectronicSeal.contract.Transact(opts, "addSealAccount", _account, _name, _cardId, _datetime)
}

func (_ElectronicSeal *ElectronicSealTransactor) AsyncAddSealAccount(handler func(*types.Receipt, error), opts *bind.TransactOpts, _account common.Address, _name string, _cardId string, _datetime string) (*types.Transaction, error) {
	return _ElectronicSeal.contract.AsyncTransact(opts, handler, "addSealAccount", _account, _name, _cardId, _datetime)
}

// AddSealAccount is a paid mutator transaction binding the contract method 0x5c4997ea.
//
// Solidity: function addSealAccount(address _account, string _name, string _cardId, string _datetime) returns()
func (_ElectronicSeal *ElectronicSealSession) AddSealAccount(_account common.Address, _name string, _cardId string, _datetime string) (*types.Transaction, *types.Receipt, error) {
	return _ElectronicSeal.Contract.AddSealAccount(&_ElectronicSeal.TransactOpts, _account, _name, _cardId, _datetime)
}

func (_ElectronicSeal *ElectronicSealSession) AsyncAddSealAccount(handler func(*types.Receipt, error), _account common.Address, _name string, _cardId string, _datetime string) (*types.Transaction, error) {
	return _ElectronicSeal.Contract.AsyncAddSealAccount(handler, &_ElectronicSeal.TransactOpts, _account, _name, _cardId, _datetime)
}

// AddSealAccount is a paid mutator transaction binding the contract method 0x5c4997ea.
//
// Solidity: function addSealAccount(address _account, string _name, string _cardId, string _datetime) returns()
func (_ElectronicSeal *ElectronicSealTransactorSession) AddSealAccount(_account common.Address, _name string, _cardId string, _datetime string) (*types.Transaction, *types.Receipt, error) {
	return _ElectronicSeal.Contract.AddSealAccount(&_ElectronicSeal.TransactOpts, _account, _name, _cardId, _datetime)
}

func (_ElectronicSeal *ElectronicSealTransactorSession) AsyncAddSealAccount(handler func(*types.Receipt, error), _account common.Address, _name string, _cardId string, _datetime string) (*types.Transaction, error) {
	return _ElectronicSeal.Contract.AsyncAddSealAccount(handler, &_ElectronicSeal.TransactOpts, _account, _name, _cardId, _datetime)
}

// SealSignature is a paid mutator transaction binding the contract method 0x6d0cfa48.
//
// Solidity: function sealSignature(address _account, string _hash, string _code, string _datetime) returns()
func (_ElectronicSeal *ElectronicSealTransactor) SealSignature(opts *bind.TransactOpts, _account common.Address, _hash string, _code string, _datetime string) (*types.Transaction, *types.Receipt, error) {
	return _ElectronicSeal.contract.Transact(opts, "sealSignature", _account, _hash, _code, _datetime)
}

func (_ElectronicSeal *ElectronicSealTransactor) AsyncSealSignature(handler func(*types.Receipt, error), opts *bind.TransactOpts, _account common.Address, _hash string, _code string, _datetime string) (*types.Transaction, error) {
	return _ElectronicSeal.contract.AsyncTransact(opts, handler, "sealSignature", _account, _hash, _code, _datetime)
}

// SealSignature is a paid mutator transaction binding the contract method 0x6d0cfa48.
//
// Solidity: function sealSignature(address _account, string _hash, string _code, string _datetime) returns()
func (_ElectronicSeal *ElectronicSealSession) SealSignature(_account common.Address, _hash string, _code string, _datetime string) (*types.Transaction, *types.Receipt, error) {
	return _ElectronicSeal.Contract.SealSignature(&_ElectronicSeal.TransactOpts, _account, _hash, _code, _datetime)
}

func (_ElectronicSeal *ElectronicSealSession) AsyncSealSignature(handler func(*types.Receipt, error), _account common.Address, _hash string, _code string, _datetime string) (*types.Transaction, error) {
	return _ElectronicSeal.Contract.AsyncSealSignature(handler, &_ElectronicSeal.TransactOpts, _account, _hash, _code, _datetime)
}

// SealSignature is a paid mutator transaction binding the contract method 0x6d0cfa48.
//
// Solidity: function sealSignature(address _account, string _hash, string _code, string _datetime) returns()
func (_ElectronicSeal *ElectronicSealTransactorSession) SealSignature(_account common.Address, _hash string, _code string, _datetime string) (*types.Transaction, *types.Receipt, error) {
	return _ElectronicSeal.Contract.SealSignature(&_ElectronicSeal.TransactOpts, _account, _hash, _code, _datetime)
}

func (_ElectronicSeal *ElectronicSealTransactorSession) AsyncSealSignature(handler func(*types.Receipt, error), _account common.Address, _hash string, _code string, _datetime string) (*types.Transaction, error) {
	return _ElectronicSeal.Contract.AsyncSealSignature(handler, &_ElectronicSeal.TransactOpts, _account, _hash, _code, _datetime)
}
