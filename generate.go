package coflnet_bot

//go:generate oapi-codegen -package=chatgen -generate=types,client,spec -o=gen/chat/client.go ./openapi/chat.json
//go:generate oapi-codegen -package=paymentgen -generate=types,client,spec -o=gen/payment/client.go ./openapi/payment.json
//go:generate oapi-codegen -package=mcconnectgen -generate=types,client,spec -o=gen/mcconnect/client.go ./openapi/mc-connect.json
//go:generate oapi-codegen -package=proxygen -generate=types,client,spec -o=gen/proxy/client.go ./openapi/proxy.json
