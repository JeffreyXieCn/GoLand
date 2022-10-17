package main

import (
	"fmt"
	"log"
	"os"

	"github.com/golang/protobuf/proto"

	"protobuf/api"
)

func main() {
	// generate protobuf go version with below command:
	//➜  protobuf git:(main) ✗ docker run --rm -v /Users/jeffrey.xie/Git/GitHub/GoLand/protobuf/api/:/api/ namely/protoc-all:1.27_0 -d /api/ -o /api/ -l go

	appInit := &api.AppInitData{
		Id:                 "12345",
		CoreOrganizationId: "67890",
		PublisherAppId:     "12345",
		Platform:           api.Platform_ANDROID,
		Networks: []*api.AppInitData_NetworkParam{
			{
				Name:              api.AdNetworkName_UNITY,
				Type:              api.UsageType_TRADITIONAL,
				AdapterParameters: map[string]string{"placementId": "myPlacementId"},
			},
		},
		CreatedAt: 1666033584371246000,
		Version:   "1234567890",
		Coppa:     false,
	}

	fmt.Printf("Original appInit:\n%+v\n", appInit)

	bytes, err := proto.Marshal(appInit)
	if err != nil {
		log.Fatalf("ProtoBuf marshaling failed: %s", err)
	}

	fmt.Printf("marshalled bytes:[%s]\n", bytes)
	fileName := "appinitdata.bin"
	err = os.WriteFile(fileName, bytes, 0600)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	decodedAppInit := &api.AppInitData{}
	err = proto.Unmarshal(file, decodedAppInit)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Decoded appInit:\n%+v\n", decodedAppInit)
	if fmt.Sprintf("%+v", appInit) == fmt.Sprintf("%+v", decodedAppInit) {
		fmt.Printf("Decoded protobuf binary successfully")
	} else {
		fmt.Printf("Decoded protobuf binary unsuccessfully")
	}
}
