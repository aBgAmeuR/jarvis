package identity

import "slices"

const AdminGroup = "admin"

type User struct {
	Name   string
	Groups []string
}

func (u User) IsAdmin() bool { return slices.Contains(u.Groups, AdminGroup) }
