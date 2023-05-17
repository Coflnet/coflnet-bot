package discordbot
 
//go:generate go run github.com/ogen-go/ogen/cmd/ogen --package api --target ./schemas/api --clean ./schemas/api/swagger.json
//go:generate go run github.com/ogen-go/ogen/cmd/ogen --package mc_connect --target ./schemas/mc_connect/ --clean ./schemas/mc_connect/swagger.yml
//go:generate go run github.com/ogen-go/ogen/cmd/ogen --package payments --target ./schemas/payments --clean ./schemas/payments/swagger.yml
//go:generate go run github.com/ogen-go/ogen/cmd/ogen --ct-alias text/json=application/json --package chat --target ./schemas/chat --clean ./schemas/chat/swagger.json 
