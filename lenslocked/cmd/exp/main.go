package main

import (
	stdctx "context"
	"fmt"

	"github.com/keykibatyr/lenslocked/context"
	"github.com/keykibatyr/lenslocked/models"
)

func main() {
	ctx := stdctx.Background()

	user := models.User{
		Email: "keykibatyr@gmail.com",
	}

	ctx = context.WithUser(ctx, &user)

	retr_user := context.UserValue(ctx)

	fmt.Println(retr_user)

}
