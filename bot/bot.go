package bot

import (
"fmt"
"log"
"os"
"os/signal"
"strings"
"github.com/bwmarrin/discordgo"
)

var BotToken string

func checkNilErr(e error) {
if e != nil {
	log.Fatal("Error message")
}
}

func Run() {
 // Creating a session
discord, err := discordgo.New("Bot " + BotToken)
checkNilErr(err)

 // Adding a event handler
discord.AddHandler(newMessage)

 // Starting a new Session
discord.Open()
defer discord.Close() // close session, after function termination

 // Keep bot running untill there is NO os interruption (ctrl + C) this will make the bot turn online
fmt.Println("Bot running....")
c := make(chan os.Signal, 1)
signal.Notify(c, os.Interrupt)
<-c
}

func newMessage(discord *discordgo.Session, message *discordgo.MessageCreate) {
// Preventing the bot from responding to its own message
if message.Author.ID == discord.State.User.ID {
	return
}

switch {
	case strings.Contains(message.Content, "Hello"):
	discord.ChannelMessageSend(message.ChannelID, "Hello I am your Personal Bot on your queries about Hewlett Packard Enterprise. How can I help you today?")

	case strings.Contains(message.Content, "How can I join HPE?"):
	discord.ChannelMessageSend(message.ChannelID, "Joining Hewlett Packard Enterprise (HPE) involves going through a standard job application process. You can either apply via job postings on Linkedin, Apply through Career Portal, Job Fairs etc")

	case strings.Contains(message.Content, "What is HPE?"):
	discord.ChannelMessageSend(message.ChannelID, "Hewlett Packard Enterprise (HPE) is a multinational information technology company that focuses on providing a wide range of technology solutions, products, and services to businesses and organizations.")

	case strings.Contains(message.Content, "What are the domains HPE works on?"):
	discord.ChannelMessageSend(message.ChannelID, "Well, HPE works on wide range of domains such as Enterprise Infrastructure, Compute, Storage, Networking, Edge Computing, High Performance Computing, AI and ML, Cybersecurity, Financial Services and much more!")

	case strings.Contains(message.Content, "What are the basic qualifications to join HPE?"):
	discord.ChannelMessageSend(message.ChannelID, "A minimum of 80% in 10th and 12th with no active backlogs with a CGPA above 7 will suffice. You can check more info by visiting our portal https://www.hpe.com/us/en/home.html")

	case strings.Contains(message.Content, "Thank you"):
	discord.ChannelMessageSend(message.ChannelID, "You're welcome!")

	case strings.Contains(message.Content, "Bye"):
	discord.ChannelMessageSend(message.ChannelID, "Bye. Feel free to ask me anything :)")

	default:
	// Default response for unmatched messages
	discord.ChannelMessageSend(message.ChannelID, "Sorry, I'm unable to understand your question. Please try asking something else.")
}
}