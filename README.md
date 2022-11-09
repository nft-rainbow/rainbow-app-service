# Rainbow-App-Service
Rainbow-app-service is a middle layer between the bot service and the rainbow-api.   

## Run
In this sense, it is needed to generate the corresponding account and the bot application. 


1. open the config.yaml
````
vim config.yaml
````

2. Replace the `discordBotToken` with your own discord bot. The botToken can be generated from [Discord Application](https://discord.com/developers/applications/)

3. Replace the `clientId` and `tokenId` with your own dodo bot. The botToken can be generated from [DoDo Application](https://doker.imdodo.com/login)

3. This project requires the proxy. Replace the `proxy` with your own proxy configuration.

4. Run the project.
````
go run main.go
````