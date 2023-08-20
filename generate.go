package discordbot

//go:generate oapi-codegen --package=chat -generate=types,client -o ./codegen/chat/chat.gen.go ./codegen/chat/swagger.json
//go:generate oapi-codegen --package=mcconnect -generate=types,client -o ./codegen/mcconnect/mcconnect.gen.go ./codegen/mcconnect/swagger.json
