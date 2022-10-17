package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type ConfigApp_AdUnit2 struct {
	AdUnitId                string                                     `json:"adUnitId,omitempty"`
	Name                    string                                     `json:"name,omitempty"`
	AdUnitFormatInformation ConfigApp_AdUnit2_AdUnit2FormatInformation `json:"adUnitFormatInformation,omitempty"`
}

type ConfigApp_AdUnit2_AdUnit2FormatInformation interface {
	isConfigApp_AdUnit2_AdUnit2FormatInformation_AdUnit2FormatInformation()
}

func (*ConfigApp_AdUnit2_AdUnit2FormatInformation_InterstitialInformation) isConfigApp_AdUnit2_AdUnit2FormatInformation_AdUnit2FormatInformation() {
}

func (*ConfigApp_AdUnit2_AdUnit2FormatInformation_RewardedInformation) isConfigApp_AdUnit2_AdUnit2FormatInformation_AdUnit2FormatInformation() {
}

func (*ConfigApp_AdUnit2_AdUnit2FormatInformation_BannerInformation) isConfigApp_AdUnit2_AdUnit2FormatInformation_AdUnit2FormatInformation() {
}

type ConfigApp_AdUnit2_AdUnit2FormatInformation_InterstitialInformation struct {
}

type ConfigApp_AdUnit2_AdUnit2FormatInformation_RewardedInformation struct {
}

type ConfigApp_AdUnit2_AdUnit2FormatInformation_BannerInformation struct {
	RefreshRateInMilliseconds int64 `json:"refreshRateInMilliseconds,omitempty"`
}

func main() {
	banner := &ConfigApp_AdUnit2_AdUnit2FormatInformation_BannerInformation{
		RefreshRateInMilliseconds: 10000,
	}

	rewarded := &ConfigApp_AdUnit2_AdUnit2FormatInformation_RewardedInformation{}

	interstitial := &ConfigApp_AdUnit2_AdUnit2FormatInformation_InterstitialInformation{}

	bannerAdUnit := ConfigApp_AdUnit2{
		AdUnitId:                "myBannerAdUnitId",
		Name:                    "myBannerAdUnitName",
		AdUnitFormatInformation: banner,
	}

	rewardedAdUnit := ConfigApp_AdUnit2{
		AdUnitId:                "myRewardedAdUnitId",
		Name:                    "myRewardedAdUnitName",
		AdUnitFormatInformation: rewarded,
	}

	interstitialAdUnit := ConfigApp_AdUnit2{
		AdUnitId:                "myInterstitialAdUnitId",
		Name:                    "myInterstitialAdUnitName",
		AdUnitFormatInformation: interstitial,
	}

	myAdUnits := []ConfigApp_AdUnit2{bannerAdUnit, rewardedAdUnit, interstitialAdUnit}

	data, err := json.Marshal(myAdUnits)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}

	fmt.Printf("%s\n", data)
	fileName := "adunits2.json"
	err = os.WriteFile(fileName, data, 0600)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	loadedAdUnits := make([]ConfigApp_AdUnit2, 0)
	err = json.Unmarshal(file, &loadedAdUnits)
	if err != nil {
		log.Fatal(err)
	}

	for _, adUnit := range loadedAdUnits {
		fmt.Printf("\n%v\n", adUnit)
		fmt.Printf("%+v\n", adUnit)
		fmt.Printf("%#v\n", adUnit)
	}

	for i, adUnit := range loadedAdUnits {
		if fmt.Sprintf("%+v", adUnit) == fmt.Sprintf("%+v", myAdUnits[i]) {
			fmt.Printf("AdUnit%d loaded successfully\n", i+1)
		} else {
			fmt.Printf("AdUnit%d loaded unsuccessfully\n", i+1)
		}
	}
}
