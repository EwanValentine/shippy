package main

import (
	"log"
	"os"

	pb "github.com/EwanValentine/shippy/user-service/proto/user"
	microclient "github.com/micro/go-micro/client"
	"golang.org/x/net/context"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
)


func main() {

	// cmd.Init()

	// Create new greeter client
	client := pb.NewUserServiceClient("go.micro.srv.user", microclient.DefaultClient)

	service := micro.NewService()

	service.Init(

		micro.Action(func(c *cli.Context) {

			name := "Ewan Valentine"
			email := "ewan.valentine89@gmail.com"
			password := "test123"
			company := "BBC"

			log.Println(name, email, password)

			r, err := client.Create(context.TODO(), &pb.User{
				Name: name,
				Email: email,
				Password: password,
				Company: company,
			})
			if err != nil {
				log.Fatalf("Could not create: %v", err)
			}
			log.Printf("Created: %t", r.User.Id)

			getAll, err := client.GetAll(context.Background(), &pb.Request{})
			if err != nil {
				log.Fatalf("Could not list users: %v", err)
			}
			for _, v := range getAll.Users {
				log.Println(v)
			}

			authResponse, err := client.Auth(context.TODO(), &pb.User{
				Email: email,
				Password: password,
			})

			if err != nil {
				log.Fatalf("Could not authenticate user: %s error: %v\n", email, err)
			}

			log.Printf("Your access token is: %s \n", authResponse.Token)

			// let's just exit because
			os.Exit(0)
		}),
	)

	// Run the server
	if err := service.Run(); err != nil {
		log.Println(err)
	}
}
