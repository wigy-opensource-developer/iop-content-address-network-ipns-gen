# ipns-gen

A tool to generate and sign IPNS records on one machine and then be published on the IPFS network with
`ipfs name upload` on another machine without sharing the private key of the IPNS record.

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

## Installation

```
go get github.com/DeCentral-Budapest/ipns-gen
```

### License
MIT
