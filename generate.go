package coflnet_bot

//go:generate oapi-codegen -package=chatgen -generate=types,client,spec -o=internal/gen/chat/client.go ./openapi/chat.json
//go:generate oapi-codegen -package=paymentgen -generate=types,client,spec -o=internal/gen/payment/client.go ./openapi/payment.json
//go:generate oapi-codegen -package=mcconnectgen -generate=types,client,spec -o=internal/gen/mcconnect/client.go ./openapi/mc-connect.json
//go:generate oapi-codegen -package=proxygen -generate=types,client,spec -o=internal/gen/proxy/client.go ./openapi/proxy.json
//go:generate oapi-codegen -package=playernamegen -generate=types,client,spec -o=internal/gen/playername/client.go ./openapi/playername.json
