# takehome_ecdsa

# start server

> `go run server.go`

## Echo web framework
https://echo.labstack.com/guide/

# request example

## GetMessage

> `curl --location --request GET 'http://localhost:1323/get_message'`


## Verify

> `curl --location --request POST 'http://localhost:1323/verify' \
--header 'Content-Type: application/json' \
--data-raw '{
    "address": "some address",
    "signedMessage": "some signed message"
}'`



