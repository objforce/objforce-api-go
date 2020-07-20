package main

import (
	"github.com/k0kubun/pp"
	"github.com/objforce/objforce-api-go"
	"os"
)

var client *objforce.Client

func main() {
	client = objforce.NewClient("", "47.0")
	client.SetDebug(true)
	err := client.Login(os.Getenv("OBJFORCE_USERNAME"), os.Getenv("OBJFORCE_PASSWORD"))
	if err != nil {
		panic(err)
	}
	//deploy()
	deployRecentValidation("0Af2K00000LuQNASA3")
}

func listMetadata() {
	query := []*apiforce.ListMetadataQuery{
		{
			Type: "ApexClass",
		},
	}
	res, err := client.ListMetadata(query)
	if err != nil {
		panic(err)
	}
	pp.Println(res)
}

func readMetadata() {
	res, err := client.ReadMetadata("CustomObject", []string{ "GO__c" })
	if err != nil {
		panic(err)
	}
	pp.Println(res)
}

func createMetadata() {
	request := []objforce.MetadataInterface{
		&objforce.CustomObject{
			FullName: "GO222__c",
			Type: "CustomObject",
			DeploymentStatus: objforce.DeploymentStatusDeployed,
			Description: "これはGOから作ってるよ",
			Label: "GOミラクルオブジェクト",
			NameField: &objforce.CustomField{
				Label: "GO名",
				Length: 80,
				Type: objforce.FieldTypeText,
			},
			PluralLabel: "GOミラクルオブジェクツ",
			SharingModel: objforce.SharingModelReadWrite,
		},
	}
	res, err := client.CreateMetadata(request)
	if err != nil {
		panic(err)
	}
	pp.Println(res)
}

func deleteMetadata() {
	res, err := client.DeleteMetadata("CustomObject", []string{ "GO222__c" })
	if err != nil {
		panic(err)
	}
	pp.Println(res)
}

func retrieve() {
	res, err := client.Retrieve(&objforce.RetrieveRequest{
		ApiVersion: 37.0,
		PackageNames: []string{"CustomObject"},
		SinglePackage: true,
		SpecificFiles: []string{},
		Unpackaged: nil,
	})
	if err != nil {
		panic(err)
	}
	pp.Println(res)
}

func describeMetadata() {
	res, err := client.DescribeMetadata()
	if err != nil {
		panic(err)
	}
	pp.Println(res)
}

func describeValueType() {
	res, err := client.DescribeValueType("{http://soap.sforce.com/2006/04/metadata}ApexClass")
	if err != nil {
		panic(err)
	}
	pp.Println(res)
}

func renameMetadata() {
	res, err := client.RenameMetadata(&objforce.RenameMetadata{
		Type: "CustomObject",
		OldFullName: "GO1__c",
		NewFullName: "GO2__c",
	})
	if err != nil {
		panic(err)
	}
	pp.Println(res)
}

func updateMetadata() {
	res, err := client.UpdateMetadata([]objforce.MetadataInterface{
		&objforce.CustomObject{
			FullName: "GO2__c",
			Type: "CustomObject",
			DeploymentStatus: objforce.DeploymentStatusDeployed,
			Description: "これはGOから作ってるよ 2",
			Label: "GOミラクルオブジェクト",
			NameField: &objforce.CustomField{
				Label: "GO名",
				Length: 80,
				Type: objforce.FieldTypeText,
			},
			PluralLabel: "GOミラクルオブジェクツ",
			SharingModel: objforce.SharingModelReadWrite,
		},
	})
	if err != nil {
		panic(err)
	}
	pp.Println(res)
}

func upsertMetadata() {
	res, err := client.UpsertMetadata([]objforce.MetadataInterface{
		&objforce.CustomObject{
			FullName: "GO1__c",
			Type: "CustomObject",
			DeploymentStatus: objforce.DeploymentStatusDeployed,
			Description: "これはGOから作ってるよ 3",
			Label: "GOミラクルオブジェクト",
			NameField: &objforce.CustomField{
				Label: "GO名",
				Length: 80,
				Type: objforce.FieldTypeText,
			},
			PluralLabel: "GOミラクルオブジェクツ",
			SharingModel: objforce.SharingModelReadWrite,
		},
	})
	if err != nil {
		panic(err)
	}
	pp.Println(res)
}

func deployRecentValidation(validationId string) {
	res, err := client.DeployRecentValidation(validationId)
	if err != nil {
		panic(err)
	}
	pp.Println(res)
}

func deploy() {
	buf := []byte("bytes")
	res, err := client.Deploy(buf, &objforce.DeployOptions{
		CheckOnly: true,
	})
	if err != nil {
		panic(err)
	}
	pp.Println(res)
}
