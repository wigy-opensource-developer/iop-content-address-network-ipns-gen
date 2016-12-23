package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	multibase "gx/ipfs/QmShp7G5GEsLVZ52imm6VP4nukpc5ipdHbscrxJMNasmSd/go-multibase"
	"gx/ipfs/QmeQ9Ktc14W93VvghbktL5yjAFf4UErPgjb1mBK5pX5kEu/go-ipfs/namesys"
	"gx/ipfs/QmeQ9Ktc14W93VvghbktL5yjAFf4UErPgjb1mBK5pX5kEu/go-ipfs/path"
	peer "gx/ipfs/QmfMmLGoKzCHDN7cGgk64PJr4iipzidDRME8HABSJqvmhC/go-libp2p-peer"
	crypto "gx/ipfs/QmfWDLQjGjVe4fr5CoztYW2DYYjRysMJrFe1RCsXLPTf46/go-libp2p-crypto"
)

var keyFile *string = flag.String("key", "", "Filename of the private key")
var seq *uint64 = flag.Uint64("seq", 0, "Sequence number greater than any previously published value")
var reol *time.Duration = flag.Duration("reol", time.Hour*24*30, "How long the record should be valid on the IPFS network")
var ttl *time.Duration = flag.Duration("ttl", time.Minute*10, "How long the record can be cached")

func main() {
	flag.Parse()

	arg := flag.Arg(0)

	//fmt.Printf("%s %d %v %v '%s'\n", *keyFile, *seq, *reol, *ttl, arg)

	sk, err := loadPrivKey(*keyFile)
	if err != nil {
		fatal(err)
	}

	switch arg {
	case "pub":
		showPubKey(sk)
	case "id":
		showKeyId(sk)
	default:
		generateRecord(sk, arg)
	}
}

func fatal(i interface{}) {
	fmt.Fprintln(os.Stderr, i)
	os.Exit(1)
}

func loadPrivKey(fileName string) (crypto.PrivKey, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	return crypto.UnmarshalPrivateKey(data)
}

func showPubKey(sk crypto.PrivKey) {
	pk := sk.GetPublic()
	pubKeyBytes, err := pk.Bytes()
	if err != nil {
		fatal(err)
	}

	pubKeyString, err := multibase.Encode(multibase.Base58BTC, pubKeyBytes)
	if err != nil {
		fatal(err)
	}

	fmt.Println(pubKeyString)
}

func showKeyId(sk crypto.PrivKey) {
	id, err := peer.IDFromPrivateKey(sk)
	if err != nil {
		fatal(err)
	}

	fmt.Println(id.Pretty())
}

func generateRecord(sk crypto.PrivKey, pathString string) {
	eol := time.Now().Add(*reol)

	newPath, err := path.ParsePath(pathString)
	if err != nil {
		fatal(err)
	}

	bytes, err := namesys.CreateEntry(sk, newPath, *seq, eol, *ttl)
	if err != nil {
		fatal(err)
	}

	output, err := multibase.Encode(multibase.Base64urlPad, bytes)
	if err != nil {
		fatal(err)
	}

	fmt.Println(output)
}
