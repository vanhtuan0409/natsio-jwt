package main

import (
	"fmt"
	"os"
	"time"

	"github.com/nats-io/jwt/v2"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nkeys"
)

const (
	NatServer = "nats://127.0.0.1:4222"
	// NatServer = "nats://nats.gondor.svc.kube:4222"
	Subject    = "time"
	AccountKey = "/home/tuan/.local/share/nats/nsc/keys/keys/A/AC/AACTO6KATGZMTM7F6MCRBA6CK7O5VNLBOOKY5BCDTQFUK6PSO6APRICH.nk"
)

func main() {
	accountKP, err := readAccountKp()
	if err != nil {
		panic(err)
	}
	userJwt, userSeed, err := generateJwt(accountKP, "gondor", AdminPermission())
	if err != nil {
		panic(err)
	}

	nc, err := nats.Connect(
		NatServer,
		nats.Name("server"),
		nats.UserJWTAndSeed(userJwt, userSeed),
	)
	if err != nil {
		panic(err)
	}
	defer nc.Drain()

	for {
		msg := time.Now().Format(time.RFC3339)
		if err := nc.Publish(Subject, []byte(msg)); err != nil {
			fmt.Printf("Failed to publish msg. ERR: %+v\n", err)
		}
		fmt.Println("Published 1 msg")
		time.Sleep(time.Second)
	}
}

func readAccountKp() (nkeys.KeyPair, error) {
	keyData, err := os.ReadFile(AccountKey)
	if err != nil {
		panic(err)
	}
	return nkeys.FromSeed(keyData)
}

func generateJwt(accountKP nkeys.KeyPair, subject string, pf PermissionFunc) (string, string, error) {
	userKP, err := nkeys.CreateUser()
	if err != nil {
		return "", "", err
	}
	userPub, err := userKP.PublicKey()
	if err != nil {
		return "", "", err
	}
	userSeed, err := userKP.Seed()
	if err != nil {
		return "", "", err
	}

	claims := jwt.NewUserClaims(userPub)
	claims.Name = subject
	pf(&claims.Permissions)

	userJwt, err := claims.Encode(accountKP)
	if err != nil {
		return "", "", err
	}
	return userJwt, string(userSeed), nil
}
