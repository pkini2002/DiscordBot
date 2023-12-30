## Discord Bot

A personalized discord bot developer using Discord Developer Profile that can be seamlessly integrated into any discord server that you own and acts like a personal chatbot to know more about Hewlett Packard Enterprise (HPE)

## Techstacks and Dependencies Used:

- Golang
- <a href="https://discord.com/developers/applications"> Discord Developer Portal </a>
- Dependency for DiscordGo:
  
  ```
  https://www.github.com/bwmarrin/discordgo
  ```

## Steps to create your own Discord Bot

- Navigate to <a href="https://discord.com/developers/applications"> Discord Developer Portal </a>
- Click on the `New Application button` and provide any name to your application (You can't keep a name that has the word Discord in it)

> [!NOTE]
>
> The application is not the Discord Bot
  
- On the left menu, select `Bot`
- Uncheck the public bot check box
- Click the `Reset Token button`, to generate a Token ID.
- Go to `URL Generator` under `OAuth2` on the left, and provide the scope as `Bot`, and permissions necessary as per the bot you want to develop. You can set permissions like Administrator privileges, send messages, send messages in threads, React to messages, and almost all other activities that you can perform on discord
- Copy the URL generated at the bottom, after the permissions are set.
- Go to that URL and add the bot to any of your servers, after all these steps you should have a Token ID and Bot(offline) in your server.

> [!IMPORTANT]
> After you have your file structure and the code ready do the following:
> Initialize the module
> ```
> go mod init example.com/hello_world_bot
> ```
> Add the discord Go dependency
> ```
> go get github.com/bwmarrin/discordgo
> ```

### Running the Project

This command will make the Bot online and can respond to your requests

```
go run main.go
```

### Screenshots

<img width="709" alt="Screenshot 2023-12-30 161829" src="https://github.com/pkini2002/AR-Ecommerce-App/assets/84091455/b2cc1231-b467-4a0a-bd12-0db8abb176c9">
<img width="718" alt="Screenshot 2023-12-30 161903" src="https://github.com/pkini2002/AR-Ecommerce-App/assets/84091455/6f2a4241-fb4f-4680-b72e-ef9f397460fc">
