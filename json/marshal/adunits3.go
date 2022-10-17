package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type ConfigApp_AdUnit3 struct {
	AdUnitId                            string                                                              `json:"adUnitId,omitempty"`
	Name                                string                                                              `json:"name,omitempty"`
	BannerAdUnitFormatInformation       *ConfigApp_AdUnit3_AdUnit3FormatInformation_BannerInformation       `json:"bannerAdUnitFormatInformation,omitempty"`
	RewardedAdUnitFormatInformation     *ConfigApp_AdUnit3_AdUnit3FormatInformation_RewardedInformation     `json:"rewardedAdUnitFormatInformation,omitempty"`
	InterstitialAdUnitFormatInformation *ConfigApp_AdUnit3_AdUnit3FormatInformation_InterstitialInformation `json:"interstitialAdUnitFormatInformation,omitempty"`
}

type ConfigApp_AdUnit3_AdUnit3FormatInformation_InterstitialInformation struct {
}

type ConfigApp_AdUnit3_AdUnit3FormatInformation_RewardedInformation struct {
}

type ConfigApp_AdUnit3_AdUnit3FormatInformation_BannerInformation struct {
	RefreshRateInMilliseconds int64 `json:"refreshRateInMilliseconds,omitempty"`
}

func main() {
	banner := &ConfigApp_AdUnit3_AdUnit3FormatInformation_BannerInformation{
		RefreshRateInMilliseconds: 10000,
	}

	rewarded := &ConfigApp_AdUnit3_AdUnit3FormatInformation_RewardedInformation{}

	interstitial := &ConfigApp_AdUnit3_AdUnit3FormatInformation_InterstitialInformation{}

	bannerAdUnit := ConfigApp_AdUnit3{
		AdUnitId:                      "myBannerAdUnitId",
		Name:                          "myBannerAdUnitName",
		BannerAdUnitFormatInformation: banner,
	}

	rewardedAdUnit := ConfigApp_AdUnit3{
		AdUnitId:                        "myRewardedAdUnitId",
		Name:                            "myRewardedAdUnitName",
		RewardedAdUnitFormatInformation: rewarded,
	}

	interstitialAdUnit := ConfigApp_AdUnit3{
		AdUnitId:                            "myInterstitialAdUnitId",
		Name:                                "myInterstitialAdUnitName",
		InterstitialAdUnitFormatInformation: interstitial,
	}

	myAdUnits := []ConfigApp_AdUnit3{bannerAdUnit, rewardedAdUnit, interstitialAdUnit}

	data, err := json.Marshal(myAdUnits)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}

	fmt.Printf("%s\n", data)
	fileName := "adunits3.json"
	err = os.WriteFile(fileName, data, 0600)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	loadedAdUnits := make([]ConfigApp_AdUnit3, 0)
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
