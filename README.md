# Rainbow-App-Service
Rainbow-app-service is a middle layer between the bot service and the rainbow-api.   

## Run the project
In this project, the bot can be seen as an account in the [Rainbow](https://nftrainbow.cn/). 
In this sense, it is needed to generate the corresponding account and the bot application. 


1. open the config.yaml
````
vim config.yaml
````

2. Replace the `app_id`, `app_secret` with your bot account's
   
3. Replace the `botAddress` with your bot address. This address can get from the [NFTRainbow](https://console.nftrainbow.cn/)

4. Replace the `botToken` with your own bot. The botToken can generate from [Discord Application](https://discord.com/developers/applications/)

5. This project requires the proxy. Replace the `proxy` with your own proxy configuration.

6. Run the project.
````
go run main.go
````