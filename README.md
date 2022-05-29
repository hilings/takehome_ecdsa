# takehome_ecdsa

# start server

> `go run server.go`


# request example

## GetMessage

generate a random message.

for now, it is a dummy deterministic message generator seeded by caller identifier.

to be improved...

* request example

	curl --location --request GET 'http://localhost:1323/get_message' \
		--header 'Public-Key: 0x04cb883652f0337949621acaf4225410b912ca9cde1d12695b67aec1c78ff58e433fd374dc255833ac295df42f459245e80c389d6016aa62555af2e88c97e4b815'
all fields in hex format

* response example

	{
		"message": "0x4773b5629666dd72a586d96529aeaadd8fa33971cd66fee19329e75c869f626d"
	}
all fields in hex format


## Verify

verify if the signed message is indeed signed by the owner of the private key to the address

* request example

	curl --location --request POST 'http://localhost:1323/verify' \
		--header 'Public-Key: 0x04cb883652f0337949621acaf4225410b912ca9cde1d12695b67aec1c78ff58e433fd374dc255833ac295df42f459245e80c389d6016aa62555af2e88c97e4b815' \
		--header 'Content-Type: application/json' \
		--data-raw '{
    		"address": "0x06A85B71941e44B6b7A256042aE0319372282803",
    		"signedMessage": "0x6c2324f62e5f5428c3819753593c6a3ddf7fc631b0a87daf819c55b0ca7d77ea22d44e90d11db9dd06c3e88e89585f438d2b80d48fe7a11c8a4971af6b2a4a5200"
		}'
all fields in hex format

* response example

	{
	    "verified": true
	}



## GenerateKey
generate private / public key pair, and also derive address

this is for experiment convenience only. in reality you probably will not generate keys by calling external APIs

* request example

	curl --location --request POST 'http://localhost:1323/generate_key' \
		--header 'Content-Type: application/json'

* response example

	{
	    "privateKey": "0x33c0f3bd9a866ab7558d6c52a98997f9c32d6774d9511eb1fa84d9deb4b39c43",
	    "publicKey": "0x04cb883652f0337949621acaf4225410b912ca9cde1d12695b67aec1c78ff58e433fd374dc255833ac295df42f459245e80c389d6016aa62555af2e88c97e4b815",
	    "address": "0x06A85B71941e44B6b7A256042aE0319372282803"
	}
all fields in hex format


## Sign
sign message using private key

this is for experiment convenience only. in reality you probably will not sign message by calling external APIs

* request example

	curl --location --request POST 'http://localhost:1323/sign' \
		--header 'Content-Type: application/json' \
		--data-raw '{
    		"message": "0x4773b5629666dd72a586d96529aeaadd8fa33971cd66fee19329e75c869f626d",
    		"privateKey": "0x33c0f3bd9a866ab7558d6c52a98997f9c32d6774d9511eb1fa84d9deb4b39c43"
		}'	
all fields in hex format

* response example

	{
	    "signedMessage": "0x6c2324f62e5f5428c3819753593c6a3ddf7fc631b0a87daf819c55b0ca7d77ea22d44e90d11db9dd06c3e88e89585f438d2b80d48fe7a11c8a4971af6b2a4a5200"
	}
all fields in hex format


# reference
## Echo web framework
https://echo.labstack.com/guide/

## go-ethereum
https://github.com/ethereum/go-ethereum
https://goethereumbook.org/signatures/

