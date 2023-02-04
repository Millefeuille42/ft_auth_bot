package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
)

func authCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	fmt.Println(i.Interaction.Member.User.Username)
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					URL:         getLink(i.Interaction.Member.User.ID),
					Type:        "link",
					Title:       "Verification Link",
					Description: "Validate your discord profile with this link",
					Thumbnail: &discordgo.MessageEmbedThumbnail{
						URL: "https://upload.wikimedia.org/wikipedia/commons/thumb/8/8d/42_Logo.svg/1200px-42_Logo.svg.png",
					},
				},
			},
			Flags: discordgo.MessageFlagsEphemeral,
		},
	})
	if err != nil {
		log.Print(err)
	}
}
