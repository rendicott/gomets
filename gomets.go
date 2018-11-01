package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

var version string
var cfg config

type tag struct {
	Key, Value string
}

type taglist []*tag

type result struct {
	Metadata interface{}
	Tags     taglist
}

type config struct {
	verbosity bool
}

func getTags(instanceID, region string) (response taglist) {
	response = taglist{}
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)
	if err != nil {
		if cfg.verbosity {
			fmt.Println(err)
		}
		return response
	}
	e := ec2.New(sess)
	query := &ec2.DescribeTagsInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("resource-id"),
				Values: []*string{
					aws.String(instanceID),
				},
			},
		},
	}
	tags, err := e.DescribeTags(query)
	if err != nil {
		if cfg.verbosity {
			fmt.Println(err)
		}
		return response
	}
	for _, t := range tags.Tags {
		newtag := tag{
			Key:   *t.Key,
			Value: *t.Value,
		}
		response = append(response, &newtag)
	}
	return response
}

func main() {
	var enhanced, versionFlag, verbose bool
	flag.BoolVar(&enhanced, "tags", false, "Attempt to describe EC2 tags and include in output.")
	flag.BoolVar(&verbose, "verbose", false, "Verbose output")
	flag.BoolVar(&versionFlag, "version", false, "Prints version and exits")
	flag.Parse()
	if versionFlag {
		fmt.Printf("gomets %s\n", version)
		os.Exit(0)
	}
	cfg.verbosity = verbose
	svc := ec2metadata.New(session.New(), aws.NewConfig())
	id, err := svc.GetInstanceIdentityDocument()
	if err != nil {
		if cfg.verbosity {
			fmt.Println(err)
		}
		panic(err)
	}
	var tags taglist
	if enhanced {
		instanceID := id.InstanceID
		region := id.Region
		tags = getTags(instanceID, region)
	} else {
		tags = taglist{}
	}
	// prepare return
	r := result{}
	r.Metadata = id
	r.Tags = tags
	idBytes, err := json.MarshalIndent(&r, "", "    ")
	if err != nil {
		if cfg.verbosity {
			fmt.Println(err)
		}
		panic(err)
	}
	fmt.Printf("%s\n", string(idBytes))
}
