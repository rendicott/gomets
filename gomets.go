package main

import (
	"fmt"
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
)

func main() {
	svc := ec2metadata.New(session.New(), aws.NewConfig())
	id,err := svc.GetInstanceIdentityDocument()
	if err != nil {
		panic(err)
	}
	idBytes, err := json.MarshalIndent(&id, "", "    ")
        if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", string(idBytes))
}
