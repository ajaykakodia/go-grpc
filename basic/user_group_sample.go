package basic

import (
	"log"

	"github.com/ajaykakodia/go-grpc/protogen/basic"
	"google.golang.org/protobuf/encoding/protojson"
)

func UserGroupCall() {
	usrGrp := basic.UserGroup{
		GroupId:   "1",
		GroupName: "Admin",
		Roles:     []string{"Admin", "USer"},
		Users: []*basic.User{
			{
				Id:       1,
				Username: "Raka",
				IsActive: true,
				Password: []byte("test"),
				Gender:   basic.Gender_GENDER_FEMALE,
				Email:    []string{"daku@t.com"},
				Address: &basic.Address{
					Street:  "Test",
					City:    "Data",
					Country: "India",
				},
			},
			{
				Id:       2,
				Username: "Duggu",
				IsActive: true,
				Password: []byte("test"),
				Gender:   basic.Gender_GENDER_FEMALE,
				Email:    []string{"daku@t.com"},
				Address: &basic.Address{
					Street:  "Test",
					City:    "Data",
					Country: "India",
				},
			},
		},
		Description: "This is my First Group",
	}

	grp, _ := protojson.Marshal(&usrGrp)
	log.Println(string(grp))
}
