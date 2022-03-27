package main

func main() {
	s3 := S3Info{
		AwsProfileName: "<profile name>",
		AwsS3Region:    "<region>",
		AwsSecretKey:   "",
		AwsAccessKey:   "",
		BucketName:     "<bucket name>",
	}

	//err := s3.SetConfigByKey()
	//err = s3.SetS3ConfigByKey()
	err := s3.SetConfigByDefault()
	if err != nil {
		panic(err)
	}

	//s3.GetBucketList()
	//s3.CreateBucket("bucket name",types.BucketLocationConstraintApNortheast2)
	//s3.GetItems("")
	//test := make(map[string]interface{})
	//test["name"] = "main"
	//doc, _ := json.Marshal(test)
}
