package main

import (
	"log"
	"os"

	pb "github.com/EwanValentine/shippy/user-service/proto/user"
	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	"golang.org/x/net/context"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
)


func main() {

	cmd.Init()

	// Create new greeter client
	client := pb.NewUserServiceClient("go.micro.srv.user", microclient.DefaultClient)

	service := micro.NewService(
		micro.Flags(
			cli.StringFlag{
				Name:  "name",
				Usage: "You full name",
			},
			cli.StringFlag{
				Name:  "email",
				Usage: "Your email",
			},
			cli.StringFlag{
				Name:  "password",
				Usage: "Your password",
			},
			cli.StringFlag{
				Name: "company",
				Usage: "Your company",
			},
		),
	)

	service.Init(

		micro.Action(func(c *cli.Context) {

			name := c.String("name")
			email := c.String("email")
			password := c.String("password")
			company := c.String("company")

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

			// let's just exit because
			os.Exit(0)
		}),
	)

	// Run the server
	if err := service.Run(); err != nil {
		log.Println(err)
	}
}
