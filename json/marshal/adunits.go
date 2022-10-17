package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type ConfigApp_AdUnit struct {
	AdUnitId                string                                    `json:"adUnitId,omitempty"`
	Name                    string                                    `json:"name,omitempty"`
	AdUnitFormatInformation *ConfigApp_AdUnit_AdUnitFormatInformation `json:"adUnitFormatInformation,omitempty"`
}

type ConfigApp_AdUnit_AdUnitFormatInformation struct {
	// Types that are valid to be assigned to AdUnitFormatInformation:
	//	*ConfigApp_AdUnit_AdUnitFormatInformation_InterstitialInformation_
	//	*ConfigApp_AdUnit_AdUnitFormatInformation_RewardedInformation_
	//	*ConfigApp_AdUnit_AdUnitFormatInformation_BannerInformation_
	AdUnitFormatInformation isConfigApp_AdUnit_AdUnitFormatInformation_AdUnitFormatInformation
}

type isConfigApp_AdUnit_AdUnitFormatInformation_AdUnitFormatInformation interface {
	isConfigApp_AdUnit_AdUnitFormatInformation_AdUnitFormatInformation()
}

type ConfigApp_AdUnit_AdUnitFormatInformation_InterstitialInformation_ struct {
	InterstitialInformation *ConfigApp_AdUnit_AdUnitFormatInformation_InterstitialInformation
}

type ConfigApp_AdUnit_AdUnitFormatInformation_RewardedInformation_ struct {
	RewardedInformation *ConfigApp_AdUnit_AdUnitFormatInformation_RewardedInformation
}

type ConfigApp_AdUnit_AdUnitFormatInformation_BannerInformation_ struct {
	BannerInformation *ConfigApp_AdUnit_AdUnitFormatInformation_BannerInformation
}

func (*ConfigApp_AdUnit_AdUnitFormatInformation_InterstitialInformation_) isConfigApp_AdUnit_AdUnitFormatInformation_AdUnitFormatInformation() {
}

func (*ConfigApp_AdUnit_AdUnitFormatInformation_RewardedInformation_) isConfigApp_AdUnit_AdUnitFormatInformation_AdUnitFormatInformation() {
}

func (*ConfigApp_AdUnit_AdUnitFormatInformation_BannerInformation_) isConfigApp_AdUnit_AdUnitFormatInformation_AdUnitFormatInformation() {
}

func (m *ConfigApp_AdUnit_AdUnitFormatInformation) GetAdUnitFormatInformation() isConfigApp_AdUnit_AdUnitFormatInformation_AdUnitFormatInformation {
	if m != nil {
		return m.AdUnitFormatInformation
	}
	return nil
}

func (m *ConfigApp_AdUnit_AdUnitFormatInformation) GetInterstitialInformation() *ConfigApp_AdUnit_AdUnitFormatInformation_InterstitialInformation {
	if x, ok := m.GetAdUnitFormatInformation().(*ConfigApp_AdUnit_AdUnitFormatInformation_InterstitialInformation_); ok {
		return x.InterstitialInformation
	}
	return nil
}

func (m *ConfigApp_AdUnit_AdUnitFormatInformation) GetRewardedInformation() *ConfigApp_AdUnit_AdUnitFormatInformation_RewardedInformation {
	if x, ok := m.GetAdUnitFormatInformation().(*ConfigApp_AdUnit_AdUnitFormatInformation_RewardedInformation_); ok {
		return x.RewardedInformation
	}
	return nil
}

func (m *ConfigApp_AdUnit_AdUnitFormatInformation) GetBannerInformation() *ConfigApp_AdUnit_AdUnitFormatInformation_BannerInformation {
	if x, ok := m.GetAdUnitFormatInformation().(*ConfigApp_AdUnit_AdUnitFormatInformation_BannerInformation_); ok {
		return x.BannerInformation
	}
	return nil
}

type ConfigApp_AdUnit_AdUnitFormatInformation_InterstitialInformation struct {
}

type ConfigApp_AdUnit_AdUnitFormatInformation_RewardedInformation struct {
}

type ConfigApp_AdUnit_AdUnitFormatInformation_BannerInformation struct {
	RefreshRateInMilliseconds int64 `json:"refreshRateInMilliseconds,omitempty"`
}

func main() {
	banner := &ConfigApp_AdUnit_AdUnitFormatInformation_BannerInformation_{
		&ConfigApp_AdUnit_AdUnitFormatInformation_BannerInformation{
			RefreshRateInMilliseconds: 10000,
		},
	}

	rewarded := &ConfigApp_AdUnit_AdUnitFormatInformation_RewardedInformation_{
		&ConfigApp_AdUnit_AdUnitFormatInformation_RewardedInformation{},
	}

	bannerAdUnit := ConfigApp_AdUnit{
		AdUnitId: "myAdUnitId",
		Name:     "myAdUnitName",
		AdUnitFormatInformation: &ConfigApp_AdUnit_AdUnitFormatInformation{
			AdUnitFormatInformation: banner},
	}

	rewardedAdUnit := ConfigApp_AdUnit{
		AdUnitId: "myRewardedAdUnitId",
		Name:     "myRewardedAdUnitName",
		AdUnitFormatInformation: &ConfigApp_AdUnit_AdUnitFormatInformation{
			AdUnitFormatInformation: rewarded},
	}

	myAdUnits := []ConfigApp_AdUnit{bannerAdUnit, rewardedAdUnit}

	data, err := json.Marshal(myAdUnits)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}

	fmt.Printf("%s\n", data)
	fileName := "adunits.json"
	err = os.WriteFile(fileName, data, 0600)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	loadedAdUnits := make([]ConfigApp_AdUnit, 0)
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
