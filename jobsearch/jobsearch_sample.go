package jobsearch

import (
	"log"

	"github.com/ajaykakodia/go-grpc/protogen/basic"
	"github.com/ajaykakodia/go-grpc/protogen/dummy"
	"github.com/ajaykakodia/go-grpc/protogen/jobsearch"
	"google.golang.org/protobuf/encoding/protojson"
)

func JobSearchSoftware() {
	js := jobsearch.JobSoftware{
		JobSoftwareId: 777,
		Application: &basic.Application{
			Version:   "1.0.0",
			Name:      "The Amazing Proto",
			Platforms: []string{"Mac", "Windows", "Linux"},
		},
	}

	jsonBytes, _ := protojson.Marshal(&js)
	log.Printf("Software: %s\n", string(jsonBytes))
}

func JobSearchCandidate() {
	jc := jobsearch.JobCandidate{
		JobCandidateId: 888,
		Application: &dummy.Application{
			ApplicationId:       887,
			ApplicationFullName: "Ajay Rao Ji",
			Phone:               "8080808008",
			Email:               "ajay.kakodia@gmail.com",
		},
	}

	jsonBytes, _ := protojson.Marshal(&jc)
	log.Println("Candidate: ", string(jsonBytes))
}
