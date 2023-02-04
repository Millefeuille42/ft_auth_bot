package main

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"net/http"
	"os"
)

var (
	bot     *discordgo.Session
	states  = make(map[string]string)
	command = &discordgo.ApplicationCommand{
		Name:        "authenticate",
		Description: "Authenticate with 42",
	}
)

func getAuth(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	if params.Get("state") == "" || params.Get("code") == "" {
		w.WriteHeader(400)
		_, err := w.Write([]byte("Missing parameters"))
		if err != nil {
			log.Println(err)
		}
		return
	}
	if _, ok := states[params.Get("state")]; !ok {
		w.WriteHeader(400)
		_, err := w.Write([]byte("Unknown state"))
		if err != nil {
			log.Println(err)
		}
		return
	}

	err := bot.GuildMemberRoleAdd(os.Getenv("GUILD_ID"),
		states[params.Get("state")],
		os.Getenv("ROLE_ID"))
	if err != nil {
		log.Print(err)
		w.WriteHeader(500)
		_, err = w.Write([]byte("Unable to assign role, please try again later"))
		if err != nil {
			log.Println(err)
		}
		delete(states, params.Get("state"))
		return
	}
	delete(states, params.Get("state"))

	w.WriteHeader(200)
	_, err = w.Write([]byte("Authentication successful, you can now close this window"))
	if err != nil {
		log.Println(err)
	}
}

func setUpBot() *discordgo.Session {
	discordBot, err := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	discordBot.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})
	discordBot.AddHandler(authCommand)

	err = discordBot.Open()
	if err != nil {
		log.Fatal(err)
	}

	_, err = discordBot.ApplicationCommandCreate(discordBot.State.User.ID, os.Getenv("GUILD_ID"), command)
	if err != nil {
		log.Fatal(err)
	}

	SetUpCloseHandler(discordBot)
	return discordBot
}

func main() {
	bot = setUpBot()

	http.HandleFunc("/auth", getAuth)

	log.Fatal(http.ListenAndServe(":9000", nil))
}
