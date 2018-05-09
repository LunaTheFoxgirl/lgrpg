package main

import "github.com/bwmarrin/discordgo"

type Commander struct {
	Name string
	
}

type Player struct {
	ServerUser discordgo.User
	
}