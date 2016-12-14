package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/ipfs/go-ipfs/namesys"
	"github.com/ipfs/go-ipfs/path"

	multibase "gx/ipfs/QmShp7G5GEsLVZ52imm6VP4nukpc5ipdHbscrxJMNasmSd/go-multibase"
	crypto "gx/ipfs/QmfWDLQjGjVe4fr5CoztYW2DYYjRysMJrFe1RCsXLPTf46/go-libp2p-crypto"
	peer "gx/ipfs/QmfMmLGoKzCHDN7cGgk64PJr4iipzidDRME8HABSJqvmhC/go-libp2p-peer"
)

var seq *uint64
var reol *time.Duration
var ttl *time.Duration

func main() {
	keyFile := flag.String("key", "", "Filename of the private key")
	seq = flag.Uint64("seq", 0, "Sequence number greater than any previously published value")
	reol = flag.Duration("reol", time.Hour*24*30, "How long the record should be valid on the IPFS network")
	ttl = flag.Duration("ttl", time.Minute*10, "How long the record can be cached")

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