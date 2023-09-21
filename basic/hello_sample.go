package basic

import (
	"log"

	"github.com/ajaykakodia/go-grpc/protogen/basic"
)

func BasicHello() {
	h := basic.Hello{
		Name: "Ajay Yadav",
	}

	log.Println(&h)

	user := basic.User{
		Id:       1,
		IsActive: true,
		Username: "Ajay Yadav",
		Password: []byte("Password"),
		Email:    []string{"ajay.yadav@h.com", "baba.jji@palem.com"},
		Gender:   basic.Gender_GENDER_MALE,
		Cordinate: &basic.Address_Cordinate{
			Longitude: 0.3434,
			Latitude:  4.5656,
		},
	}

	log.Println(&user)
}
