package main

import (
    "fmt"
    "github.com/hyperledger/aries-framework-go/pkg/framework/aries"
    "github.com/hyperledger/aries-framework-go/pkg/client/didexchange"
)

func main(){
    fmt.Println("Hello")
    framework, _ := aries.New()
    ctx, _ := framework.Context()
    client, _ := didexchange.New(ctx)
    invitation, _ := client.CreateInvitation("test")
    fmt.Println(invitation.Invitation.ID)
    fmt.Println(invitation.Invitation.Label)
    fmt.Println(invitation.Invitation.DID)
    fmt.Println(invitation.Invitation.Type)
}
