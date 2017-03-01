# iop-content-address-network-ipns-gen

[![Build Status](https://travis-ci.org/DeCentral-Budapest/ipns-gen.svg?branch=master)](https://travis-ci.org/DeCentral-Budapest/ipns-gen)
[![standard-readme compliant](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)

> A tool to generate and sign IPNS records on one machine, which can be then published on the IPFS network with
`ipfs name upload` on another machine without sharing the private key of the IPNS record.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Contribute](#contribute)
- [License](#license)

## Install

```
go get -d github.com/DeCentral-Budapest/ipns-gen
make install
```

Installation goes faster if you are running an IPFS daemon on localhost. If not, dependencies will be downloaded from
http://gateway.ipfs.io:8080/ that is a bit crowded every now and then.

## Usage

First, you'll need to generate a publish key with [`ipfs-key`](https://github.com/DeCentral-Budapest/ipfs-key).
Once you have one, you need to extract the public key from that key:

```
$ ipns-gen -key=key.priv pub
z4XTTM1jxcwnF4smRV3orZYBvGci92wLZqbovXg81uQprp4u7f
```

This will be used as a `--key` option to `ipfs name upload`. Whenever you want to update the IPNS record
that belongs to this key, you need to generate and sign a new record that has a higher sequence number
than all records published to that key.

```
$ ipns-gen -key=key.priv -seq=1 /ipfs/QmbR6SFFFZ5ik1urY5pViF21MPEfPqfu2TRWGN7GFiDQBa
UCjQvaXBmcy9RbWJSNlNGRkZaNWlrMXVyWTVwVmlGMjFNUEVmUHFmdTJUUldHTjdHRmlEUUJhEkCqK-0pHHkZvGXUl0tlvhYeAY8jAHYPMUfrmQDTlztaTKClUib19Ct81SqMvRbdUK5nhRvR6wstRVic-Q-OhVELGAAiHjIwMTctMDEtMTJUMTc6MzQ6MzYuMzExODI0MzM0WigBMIDgpZa7EQ==
```

After that, you can call the following command on an IPFS node without sharing the private key:

```
$ ipfs name upload --key z4XTTM1jxcwnF4smRV3orZYBvGci92wLZqbovXg81uQprp4u7f UCjQvaXBmcy...
```

Whenever you are in doubt what is the ID of a key, you can display it with the following command:

```
$ ipns-gen -key=key.priv id
Qmb1z6tLB1pf56vyeLh4GwGoBmJjeHWFC5kzuydo5WaifT
```

Based on that information, you know that any IPFS node will resolve this ID to the path you specified in the record before:

```
$ ipfs name resolve /ipns/Qmb1z6tLB1pf56vyeLh4GwGoBmJjeHWFC5kzuydo5WaifT
/ipfs/QmbR6SFFFZ5ik1urY5pViF21MPEfPqfu2TRWGN7GFiDQBa
```

## Contribute

PRs accepted.

## License

[MIT](LICENSE) © 2016 [Fermat Foundation](http://www.fermat.org/)
