package validator

import (
	"fmt"
	dgo "github.com/bwmarrin/discordgo"
)

func IsAdmin(s *dgo.Session, guildID string, userID string) bool {
	check := false
	uRoles := getUserRole(s, guildID, userID)
	for _, uRole := range uRoles {
		if getRoleName(s, guildID, uRole) == "Admin" {
			check = true
			break
		}
	}
	return check
}

func getRoleName(s *dgo.Session, guildID string, roleID string) string {
	roles, err := s.GuildRoles(guildID)
	if err != nil {
		panic(err)
	}
	var result string
	for _, role := range roles {
		if role.ID == roleID {
			result = role.Name
		}
	}
	fmt.Println(result)
	return result
}

func getUserRole(s *dgo.Session, guildID string, userID string) []string {
	result, err := s.GuildMember(guildID, userID)
	if err != nil {
		panic(err)
	}
	return result.Roles
}
