package main

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/labstack/echo"
)

// Splain contains all of the parsed tokens in the transaction
type Splain struct {
	Tokens []Token
}

// Token contains all the visible fields for each token
type Token struct {
	Hex  string
	Text string
	More string
}

type field int

var (
	NONCE     field = 0
	GAS_PRICE field = 1
	GAS_LIMIT field = 2
	RECIPIENT field = 3
	VALUE     field = 4
	DATA      field = 5
	SIG_V     field = 6
	SIG_R     field = 7
	SIG_S     field = 8
)

func main() {
	// start simple server
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, string(parse(data, false)))
	})

	e.GET("/:tx", func(c echo.Context) error {
		rawTx := c.Param("tx")
		verbose := c.QueryParam("verbose")
		//fmt.Println("verbose", verbose)
		v := false
		if verbose == "true" {
			v = true
		}

		// fetch rawTx from etherscan if it looks like we have a tx hash instead of a raw tx
		if len(rawTx) < 100 {
			rawTx = strings.TrimSpace(etherscanCrawlRaw(rawTx))
		}
		if len(rawTx) < 100 {
			return c.String(http.StatusBadRequest, "")
		}

		return c.String(http.StatusOK, string(parse(rawTx, v)))
	})
	e.Logger.Fatal(e.Start(":8080"))
}

func parse(rawTx string, verbose bool) []byte {
	fmt.Println(rawTx)

	str := strings.TrimPrefix(rawTx, "0x")
	tx := &types.Transaction{}
	buf, err := hex.DecodeString(str)
	if err != nil {
		log.Fatal(err)
	}

	r := bytes.NewReader(buf)
	s := rlp.NewStream(r, 0)
	err = tx.DecodeRLP(s)
	if err != nil {
		log.Fatal(err)
	}

	splain := Splain{}

	// special case for the first rlp node before the nonce
	var tok Token
	prefix := buf[0]
	l := buf[0] - 0xf7
	flen := buf[1 : 1+l]
	tok.Hex = Hex(append([]byte{prefix}, flen...))
	tok.Text = fmt.Sprintf("RLP Prefix. Tells us that this transaction is a list of length 0x%s (%d bytes)", hex.EncodeToString(flen), uint64(flen[0])) // TODO: extend for larger txs
	tok.More = fmt.Sprintf("The first byte (0x%x-0xf7) tells us the length of the length (0x%s) of transaction", prefix, hex.EncodeToString(flen))
	splain.Tokens = append(splain.Tokens, tok)

	// Tokenize transaction fields and their encoding prefixes
	splain.addNode(tx.Nonce(), NONCE, verbose)
	splain.addNode(tx.GasPrice().Bytes(), GAS_PRICE, verbose)
	splain.addNode(tx.Gas(), GAS_LIMIT, verbose)
	// contract creation edgecase
	if tx.To() != nil {
		splain.addNode(tx.To().Bytes(), RECIPIENT, verbose)
	} else {
		splain.addNode([]byte{}, RECIPIENT, verbose)
	}
	splain.addNode(tx.Value().Bytes(), VALUE, verbose)
	splain.addNode(tx.Data(), DATA, verbose)
	sigV, sigR, sigS := tx.RawSignatureValues()
	splain.addNode(sigV.Bytes(), SIG_V, verbose)
	splain.addNode(sigR.Bytes(), SIG_R, verbose)
	splain.addNode(sigS.Bytes(), SIG_S, verbose)

	out, _ := json.MarshalIndent(splain, "", "	")
	return out
}

func (s *Splain) addNode(val interface{}, f field, verbose bool) {

	enc, err := rlp.EncodeToBytes(val)
	if err != nil {
		log.Fatal(err)
	}
	i := 0
	if verbose {
		i = addRLPNode(s, enc)
	}

	// add the value node skipping however long the prefix was
	var tok Token
	tok.Hex = Hex(enc[i:])

	// construct the explanatory text
	var txt, more string
	switch f {
	case NONCE:
		txt, more = nonceInfo(val, verbose)
	case GAS_PRICE:
		txt, more = gasPriceInfo(val, verbose)
	case GAS_LIMIT:
		txt, more = gasLimitInfo(val, verbose)
	case RECIPIENT:
		txt, more = recipientInfo(val, verbose)
	case VALUE:
		txt, more = valueInfo(val)
	case DATA:
		txt, more = dataInfo(val, verbose)
	case SIG_V:
		txt, more = sigVInfo(val)
	case SIG_R:
		txt, more = sigRInfo(val)
	case SIG_S:
		txt, more = sigSInfo(val)

	default:
		txt = "NOT IMPLEMENTED"
		more = "Not IMPLEMENTED"

	}
	tok.Text = txt
	tok.More = more

	// Edgcase for when the prefix tells us the data length of the next argument is zero
	// we don't want to add a node for no data
	//if len(tok.Hex) > 0 {
	s.Tokens = append(s.Tokens, tok)
	//}

}

func nonceInfo(val interface{}, verbose bool) (string, string) {

	i, _ := val.(uint64)
	txt := fmt.Sprintf("Nonce: %d", i)
	more := shortNonce
	if verbose {
		more = verboseNonce
	}

	return txt, more
}

var verboseNonce = "The nonce is a sequence number issued my the transaction creator used to prevent message replay. The nonce of each transaction of an account must be exactly 1 greater than the previous nonce used. The Ethereum yellow paper defines the nonce as 'A scalar value equal to the number of transactions sent from this address or, in the case of accounts with associated code, the number of contract-creations made by this account"

var shortNonce = "The nonce is an incrementing sequence number used to prevent message replay"

func gasPriceInfo(val interface{}, verbose bool) (string, string) {
	buf, _ := val.([]byte)
	i := big.NewInt(0).SetBytes(buf)

	txt := fmt.Sprintf("Gas Price: %s", i.String())
	more := shortGasPrice
	if verbose {
		more = verboseGasPrice
	}
	return txt, more
}

var shortGasPrice = "The price of gas (in wei) that the sender is willing to pay."
var verboseGasPrice = "The price of gas (in wei) that the sender is willing to pay. Gas is purchased with ether and serves to protect the limited resources of the network (computation, memory, and storage). The amount of ether spent for gas can be calculated by multiplying the Gas Price by the amount of gas consumed in the transaction (21000 gas for a standard transaction)"

func gasLimitInfo(val interface{}, verbose bool) (string, string) {
	i := val.(uint64)
	txt := fmt.Sprintf("Gas Limit: %d", i)
	more := shortGasLimit
	if verbose {
		more = verboseGasLimit
	}

	return txt, more
}

var shortGasLimit = "The maximum amount of gas the originator is willing to pay for this transaction."
var verboseGasLimit = "The maximum amount of gas the originator is willing to pay for this transaction. The amount of gas consumed depends on how much computation your transaction requires."

func recipientInfo(val interface{}, verbose bool) (string, string) {
	addrBytes := val.([]byte)
	//if len(addrBytes) == 0 || (len(addrBytes) == 1 && addrBytes[0] == 0x0) {
	if len(addrBytes) == 0 {
		txt := fmt.Sprintf("Recipient Address: 0x0")
		more := "This transaction is a special type of transaction for Contract Creation. Notice the address is the Zero Address 0x0"
		return txt, more
	}
	txt := fmt.Sprintf("Recipient Address: 0x%s", hex.EncodeToString(addrBytes))
	more := shortTo
	if verbose {
		more = verboseTo
	}

	return txt, more
}

var shortTo = "The address of the user account or contract to interact with"
var verboseTo = `An ethereum address is generated with the following steps
1. Generate a public key by multiplying the private key 'k' by the Ethereum generator point G. The public key is the concatenated x + y coordinate of the result of this multiplication
2. Take the Keccak-256 hash of that public key 
3. Take the last 20 bytes of that hash and encode to hexidecimal.`

func valueInfo(val interface{}) (string, string) {
	buf, _ := val.([]byte)
	i := big.NewInt(0).SetBytes(buf)

	txt := fmt.Sprintf("Value: %s", i.String())
	more := "The amount of ether (in wei) to send to the recipient address."
	return txt, more
}

func dataInfo(val interface{}, verbose bool) (string, string) {
	buf, _ := val.([]byte)

	txt := fmt.Sprintf("Data: %s", hex.EncodeToString(buf))
	more := shortData
	if verbose {
		more = verboseData
	}
	return txt, more
}

var verboseData = "Data being sent to a contract function. The first 4 bytes are known as the 'function selector'. The remaining data represents arguments to the chosen function"
var shortData = "Data being sent to a contract function. The first 4 bytes are known as the 'function selector'"

func sigVInfo(val interface{}) (string, string) {
	buf, _ := val.([]byte)

	txt := fmt.Sprintf("Signature Prefix Value (v): %s", hex.EncodeToString(buf))
	more := "Indicates both the chainID of the transaction and the parity (odd or even) of the y component of the public key"
	return txt, more
}

func sigRInfo(val interface{}) (string, string) {
	buf, _ := val.([]byte)

	txt := fmt.Sprintf("Signature (r) value: %s", hex.EncodeToString(buf))
	more := "Part of the signature pair (r,s). Represents the X-coordinate of an ephemeral public key created during the ECDSA signing process"
	return txt, more
}

func sigSInfo(val interface{}) (string, string) {
	buf, _ := val.([]byte)

	txt := fmt.Sprintf("Signature (s) value: %s", hex.EncodeToString(buf))
	more := "Part of the signature pair (r,s). Generated using the ECDSA signing algorithm"
	return txt, more
}

// if there is a rlp length prefix add a node for it, else do nothing.
// Return how many bytes the prefix took
func addRLPNode(s *Splain, enc []byte) int {
	length := len(enc)
	if length == 0 {
		log.Fatal("Unable to decode length of val:", enc)
	}

	var node Token

	prefix := enc[0]
	// This is a single byte value that is its own rlp encoding so no node to add
	if prefix <= 0x7F {
		return 0
	}
	// "string" value of length 0-55
	if prefix <= 0xB7 && length > int(prefix-0x80) {
		node.Hex = Hex([]byte{prefix})
		node.Text = fmt.Sprintf("RLP Length Prefix. The next field is an RLP 'string' of length 0x%x - 0x80", prefix)
		node.More = ""
		s.Tokens = append(s.Tokens, node)
		return 1
	}
	// "string" value of length > 55
	if prefix < 0xC0 {
		// prefix tells us the length of the length of the field
		l := prefix - 0xb7
		flen := enc[1 : 1+l]
		node.Hex = Hex(append([]byte{prefix}, flen...))
		node.Text = fmt.Sprintf("RLP Length Prefix. The next field is an RLP 'string' of length 0x%s", hex.EncodeToString(flen))
		node.More = fmt.Sprintf("The first byte (0x%x-0x80) tells us the length of the length (0x%s) of the next field", prefix, hex.EncodeToString(flen))
		s.Tokens = append(s.Tokens, node)
		return 1 + len(flen)

	}

	log.Fatal("Not Implemented")

	return -1
}

// Hex how do i fix my linter plx halp
func Hex(b []byte) string {
	return hex.EncodeToString(b)
}

func rlpExplain(buf []byte) string {

	return ""
}

var data = "0xf89182032d8504a817c80082fe90940b95993a39a363d99280ac950f5e4536ab5c5566871550f7dca70000a41a6952300000000000000000000000001b46d8845f5a30447f182ac925c7da8b65a0124a26a0df820a48d3a6cd4e986b00a601138a1a7d0969334edd1ec1e2f6ad3c6890a468a0573ca6ccd5dc1eab646aa996c8fa7c6f1ec3256d2d051e0f2a0a04e0066025b6"

// example txhash = 0x9a7c62249dc4d4df8ce424c256fe4e57c06fb8b45101b43384db00a1d73799b5
//var data = "0xf8aa0185012a05f2008327c50e9435fb136cbadbc168910b66a9f7c40b03e4bd467f80b8441e9a695000000000000000000000000035fb136cbadbc168910b66a9f7c40b03e4bd467f000000000000000000000000000000000000000000000000000000003b9aca0026a00320143282b77654f3eedf2c6d384346a4be52c902f6603227f8f0220d30aa35a076ea8a4947327f33e149ec928efd6efa9e49aafe89a189abae7aad599c5feef2"
