package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/google/uuid"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func getLink(id string) string {
	state := strings.Replace(uuid.NewString(), "-", "", -1)
	states[state] = id

	fmt.Println(state, "=", id)
	url := os.Getenv("REDIR_URI") + "&state=" + state

	return url
}

func SetUpCloseHandler(session *discordgo.Session) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		commands, err := session.ApplicationCommands(session.State.User.ID, os.Getenv("GUILD_ID"))
		if err != nil {
			log.Print(err)
		} else {
			for _, comm := range commands {
				_ = session.ApplicationCommandDelete(session.State.User.ID, os.Getenv("GUILD_ID"), comm.ID)
			}
		}
		_ = session.Close()
		os.Exit(0)
	}()
}
