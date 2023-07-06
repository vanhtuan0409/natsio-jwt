package main

import (
	"fmt"

	"github.com/nats-io/jwt/v2"
)

type PermissionFunc func(*jwt.Permissions)

func DefaultPermission() PermissionFunc {
	return func(p *jwt.Permissions) {
		p.Pub.Allow.Add("PUBLIC.>")
		p.Sub.Allow.Add("PUBLIC.>", "_INBOX.>")
	}
}

func AdminPermission() PermissionFunc {
	return func(p *jwt.Permissions) {
		p.Pub.Allow.Add(">")
		p.Sub.Allow.Add(">")
	}
}

func UserPermission(user string) PermissionFunc {
	return CombinePermission(
		DefaultPermission(),
		func(p *jwt.Permissions) {
			p.Sub.Allow.Add(
				fmt.Sprintf("%s.>", user),
			)
		},
	)
}

func CombinePermission(perms ...PermissionFunc) PermissionFunc {
	return func(p *jwt.Permissions) {
		for _, perm := range perms {
			perm(p)
		}
	}
}
