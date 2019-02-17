package main

import (
	"testing"
)

func TestContract(t *testing.T) {

	actual := string(parse(contract, false))
	if actual != contractExpected {
		t.Error("Contract tx parsed incorrectly")
	}
}

func TestSimple(t *testing.T) {
	actual := string(parse(simple, false))
	if actual != simpleExpected {
		t.Error("Contract tx parsed incorrectly")
	}
}

var simple = "0xf86b8085012a05f200825208949b0a420cd00b9d75fce4226262789f734046e54987026bf86755a05b8026a06a49585b2e6720633828f7a55e5f98709d9f6f4bfe869c9f5616ce46eb26566aa0751d23163c267e0f141481964100620f3f228da1636fe90129687425d8a8f836"

var simpleExpected = `{
	"Tokens": [
		{
			"Hex": "f86b",
			"Text": "RLP Prefix. Tells us that this transaction is a list of length 0x6b (107 bytes)",
			"More": "The first byte (0xf8-0xf7) tells us the length of the length (0x6b) of transaction"
		},
		{
			"Hex": "80",
			"Text": "Nonce: 0",
			"More": "The nonce is an incrementing sequence number used to prevent message replay"
		},
		{
			"Hex": "85012a05f200",
			"Text": "Gas Price: 5000000000",
			"More": "The price of gas (in wei) that the sender is willing to pay."
		},
		{
			"Hex": "825208",
			"Text": "Gas Limit: 21000",
			"More": "The maximum amount of gas the originator is willing to pay for this transaction."
		},
		{
			"Hex": "949b0a420cd00b9d75fce4226262789f734046e549",
			"Text": "Recipient Address: 0x9b0a420cd00b9d75fce4226262789f734046e549",
			"More": "The address of the user account or contract to interact with"
		},
		{
			"Hex": "87026bf86755a05b",
			"Text": "Value: 681664583147611",
			"More": "The amount of ether (in wei) to send to the recipient address."
		},
		{
			"Hex": "80",
			"Text": "Data: ",
			"More": "Data being sent to a contract function. The first 4 bytes are known as the 'function selector'"
		},
		{
			"Hex": "26",
			"Text": "Signature Prefix Value (v): 26",
			"More": "Indicates both the chainID of the transaction and the parity (odd or even) of the y component of the public key"
		},
		{
			"Hex": "a06a49585b2e6720633828f7a55e5f98709d9f6f4bfe869c9f5616ce46eb26566a",
			"Text": "Signature (r) value: 6a49585b2e6720633828f7a55e5f98709d9f6f4bfe869c9f5616ce46eb26566a",
			"More": "Part of the signature pair (r,s). Represents the X-coordinate of an ephemeral public key created during the ECDSA signing process"
		},
		{
			"Hex": "a0751d23163c267e0f141481964100620f3f228da1636fe90129687425d8a8f836",
			"Text": "Signature (s) value: 751d23163c267e0f141481964100620f3f228da1636fe90129687425d8a8f836",
			"More": "Part of the signature pair (r,s). Generated using the ECDSA signing algorithm"
		}
	]
}`

var contract = "0xf903db82265a8502540be40083045c938080b90386608060405234801561001057600080fd5b50604051602080610366833981016040525160008054600160a060020a03909216600160a060020a0319909216919091179055610314806100526000396000f30060806040526004361061004b5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416633f579f42811461004d578063d2ec4a92146100b6575b005b34801561005957600080fd5b50604080516020600460443581810135601f810184900484028501840190955284845261004b948235600160a060020a03169460248035953695946064949201919081908401838280828437509497506100e79650505050505050565b3480156100c257600080fd5b506100cb6102d9565b60408051600160a060020a039092168252519081900360200190f35b6000809054906101000a9004600160a060020a0316600160a060020a031663c34c08e56040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15801561015257600080fd5b505af1158015610166573d6000803e3d6000fd5b505050506040513d602081101561017c57600080fd5b5051600160a060020a0316331461019257600080fd5b82600160a060020a0316828260405180828051906020019080838360005b838110156101c85781810151838201526020016101b0565b50505050905090810190601f1680156101f55780820380516001836020036101000a031916815260200191505b5091505060006040518083038185875af192505050156102cf577f39f46e1dedea184144e3feaf4e595d78345d9a9d8b43da87912efbe4df3c8a318383836040518084600160a060020a0316600160a060020a0316815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561028e578181015183820152602001610276565b50505050905090810190601f1680156102bb5780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a16102d4565b600080fd5b505050565b600054600160a060020a0316815600a165627a7a72305820e15ee5e2160fdc89ce720dc909c6dc0f003d58418735db64a66e99fd3338afa3002900000000000000000000000098d0c1a1045a3145ea8d06f1db575819c8a7c9bd25a078a6f18a1036ed7e23dd63481fd1cd62064e8cd4b03ee8b0a377c190cb9113a8a0197177c1425a243e39d487c0ac40faa595b03f0dff537555df4256b8d09e7989"

var contractExpected = `{
	"Tokens": [
		{
			"Hex": "f903db",
			"Text": "RLP Prefix. Tells us that this transaction is a list of length 0x03db (3 bytes)",
			"More": "The first byte (0xf9-0xf7) tells us the length of the length (0x03db) of transaction"
		},
		{
			"Hex": "82265a",
			"Text": "Nonce: 9818",
			"More": "The nonce is an incrementing sequence number used to prevent message replay"
		},
		{
			"Hex": "8502540be400",
			"Text": "Gas Price: 10000000000",
			"More": "The price of gas (in wei) that the sender is willing to pay."
		},
		{
			"Hex": "83045c93",
			"Text": "Gas Limit: 285843",
			"More": "The maximum amount of gas the originator is willing to pay for this transaction."
		},
		{
			"Hex": "80",
			"Text": "Recipient Address: 0x0",
			"More": "This transaction is a special type of transaction for Contract Creation. Notice the address is the Zero Address 0x0"
		},
		{
			"Hex": "80",
			"Text": "Value: 0",
			"More": "The amount of ether (in wei) to send to the recipient address."
		},
		{
			"Hex": "b90386608060405234801561001057600080fd5b50604051602080610366833981016040525160008054600160a060020a03909216600160a060020a0319909216919091179055610314806100526000396000f30060806040526004361061004b5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416633f579f42811461004d578063d2ec4a92146100b6575b005b34801561005957600080fd5b50604080516020600460443581810135601f810184900484028501840190955284845261004b948235600160a060020a03169460248035953695946064949201919081908401838280828437509497506100e79650505050505050565b3480156100c257600080fd5b506100cb6102d9565b60408051600160a060020a039092168252519081900360200190f35b6000809054906101000a9004600160a060020a0316600160a060020a031663c34c08e56040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15801561015257600080fd5b505af1158015610166573d6000803e3d6000fd5b505050506040513d602081101561017c57600080fd5b5051600160a060020a0316331461019257600080fd5b82600160a060020a0316828260405180828051906020019080838360005b838110156101c85781810151838201526020016101b0565b50505050905090810190601f1680156101f55780820380516001836020036101000a031916815260200191505b5091505060006040518083038185875af192505050156102cf577f39f46e1dedea184144e3feaf4e595d78345d9a9d8b43da87912efbe4df3c8a318383836040518084600160a060020a0316600160a060020a0316815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561028e578181015183820152602001610276565b50505050905090810190601f1680156102bb5780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a16102d4565b600080fd5b505050565b600054600160a060020a0316815600a165627a7a72305820e15ee5e2160fdc89ce720dc909c6dc0f003d58418735db64a66e99fd3338afa3002900000000000000000000000098d0c1a1045a3145ea8d06f1db575819c8a7c9bd",
			"Text": "Data: 608060405234801561001057600080fd5b50604051602080610366833981016040525160008054600160a060020a03909216600160a060020a0319909216919091179055610314806100526000396000f30060806040526004361061004b5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416633f579f42811461004d578063d2ec4a92146100b6575b005b34801561005957600080fd5b50604080516020600460443581810135601f810184900484028501840190955284845261004b948235600160a060020a03169460248035953695946064949201919081908401838280828437509497506100e79650505050505050565b3480156100c257600080fd5b506100cb6102d9565b60408051600160a060020a039092168252519081900360200190f35b6000809054906101000a9004600160a060020a0316600160a060020a031663c34c08e56040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15801561015257600080fd5b505af1158015610166573d6000803e3d6000fd5b505050506040513d602081101561017c57600080fd5b5051600160a060020a0316331461019257600080fd5b82600160a060020a0316828260405180828051906020019080838360005b838110156101c85781810151838201526020016101b0565b50505050905090810190601f1680156101f55780820380516001836020036101000a031916815260200191505b5091505060006040518083038185875af192505050156102cf577f39f46e1dedea184144e3feaf4e595d78345d9a9d8b43da87912efbe4df3c8a318383836040518084600160a060020a0316600160a060020a0316815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561028e578181015183820152602001610276565b50505050905090810190601f1680156102bb5780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a16102d4565b600080fd5b505050565b600054600160a060020a0316815600a165627a7a72305820e15ee5e2160fdc89ce720dc909c6dc0f003d58418735db64a66e99fd3338afa3002900000000000000000000000098d0c1a1045a3145ea8d06f1db575819c8a7c9bd",
			"More": "Data being sent to a contract function. The first 4 bytes are known as the 'function selector'"
		},
		{
			"Hex": "25",
			"Text": "Signature Prefix Value (v): 25",
			"More": "Indicates both the chainID of the transaction and the parity (odd or even) of the y component of the public key"
		},
		{
			"Hex": "a078a6f18a1036ed7e23dd63481fd1cd62064e8cd4b03ee8b0a377c190cb9113a8",
			"Text": "Signature (r) value: 78a6f18a1036ed7e23dd63481fd1cd62064e8cd4b03ee8b0a377c190cb9113a8",
			"More": "Part of the signature pair (r,s). Represents the X-coordinate of an ephemeral public key created during the ECDSA signing process"
		},
		{
			"Hex": "a0197177c1425a243e39d487c0ac40faa595b03f0dff537555df4256b8d09e7989",
			"Text": "Signature (s) value: 197177c1425a243e39d487c0ac40faa595b03f0dff537555df4256b8d09e7989",
			"More": "Part of the signature pair (r,s). Generated using the ECDSA signing algorithm"
		}
	]
}`