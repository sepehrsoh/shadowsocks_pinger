package cli

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

var discordCmd = &cobra.Command{
	Use:   "discord",
	Short: "discord pinger service.",
	Long:  `discord pinger service.`,
	Run: func(cmd *cobra.Command, args []string) {
		discord()
	},
}

func init() {
	rootCmd.AddCommand(discordCmd)
}

func discord() {
	logrus.Infoln("starting discord")
	dg, err := discordgo.New("Bot " + getDiscordToken("Discord_Token"))
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	dg.AddHandler(messageCreate)

	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
	return
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	if m.Content != "ping" {
		return
	}

	channel, err := s.UserChannelCreate(m.Author.ID)
	if err != nil {
		fmt.Println("error creating channel:", err)
		s.ChannelMessageSend(
			m.ChannelID,
			"Something went wrong while sending the DM!",
		)
		return
	}
	_, err = s.ChannelMessageSend(channel.ID, "Pong!")
	if err != nil {
		fmt.Println("error sending DM message:", err)
		s.ChannelMessageSend(
			m.ChannelID,
			"Failed to send you a DM. "+
				"Did you disable DM in your privacy settings?",
		)
	}
}

func getDiscordToken(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		logrus.Panic("Error loading .env file")
	}

	return os.Getenv("token")
}
