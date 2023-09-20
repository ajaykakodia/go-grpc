package basic

import (
	"log"

	"github.com/ajaykakodia/go-grpc/protogen/basic"
	"google.golang.org/protobuf/encoding/protojson"
)

func BasicUser() {

	user := basic.User{
		Id:       1,
		IsActive: true,
		Username: "Ajay Yadav",
		Password: []byte("Password"),
		Email:    []string{"ajay.yadav@h.com", "baba.jji@palem.com"},
		Gender:   basic.Gender_GENDER_MALE,
		Address: &basic.Address{
			Street:  "Ghanta Street",
			City:    "DO City",
			Country: "India",
		},
	}

	log.Println(&user)
}

func ProtoToJsonUser() {
	p := basic.User{
		Id:       1,
		IsActive: true,
		Username: "Rekha Yadav",
		Password: []byte("Password"),
		Email:    []string{"ajay.yadav@h.com", "baba.jji@palem.com"},
		Gender:   basic.Gender_GENDER_MALE,
	}

	jsonBytes, _ := protojson.Marshal(&p)
	log.Println(string(jsonBytes))
}

// {"id":9, "username":"Rekha Yadav", "is_active":true, "password":"UGFzc12dvcmQ=", "gender":"GENDER_MALE", "email":["ajay.yadav@h.com", "baba.jji@palem.com"]}

func JsonToProtoUser() {
	json := `{"id":10, "username":"Duggu Kakodia", "is_active":true, "gender":"GENDER_MALE", "email":["ajay.yadav@h.com", "baba.jji@palem.com"]}"`

	var usr basic.User

	err := protojson.Unmarshal([]byte(json), &usr)

	if err != nil {
		log.Println("Got Error: ", err.Error())
		return
	}
	log.Println(usr)

}
