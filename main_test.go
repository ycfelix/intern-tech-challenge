
package main

import (
"github.com/coreos/go-semver/semver"
"fmt"
)
func bubblesort(releases []*semver.Version,n int) {

	for i:=0;i<n-1;i++ {
		for j:=0;j<n-i-1;j++ {
			if releases[j].LessThan(*releases[j+1])==true {
				var temp *semver.Version
				temp=releases[j]
				releases[j]=releases[j+1]
				releases[j+1]=temp
			}
		}
	}
}
// LatestVersions returns a sorted slice with the highest version as its first element and the highest version of the smaller minor versions in a descending order
func LatestVersions(releases []*semver.Version, minVersion *semver.Version) []*semver.Version {
	var versionSlice []*semver.Version

	var length=0
	for i:=0;i<len(releases);i++{
		if releases[i].Equal(*minVersion){length+=1;continue}

		if minVersion.LessThan(*releases[i])==true{
			length+=1
		}
	}


	var temp=make([]*semver.Version, length)
	var index=0
	for i:=0;i<len(releases);i++{
		if releases[i].Equal(*minVersion){temp[index]=releases[i]
			index+=1;continue}
		if minVersion.LessThan(*releases[i])==true{
			temp[index]=releases[i]
			index+=1
		}
	}
	bubblesort(temp,length)

	var Major int64
	var Minor int64

	var resultLength=0
	Major=-1
	Minor=-1

	for i:=0;i<length;i++{
		if temp[i].Major==Major{
			if temp[i].Minor==Minor{continue}
			resultLength+=1
			Major=temp[i].Major
			Minor=temp[i].Minor
			continue
		}
		Major=temp[i].Major
		Minor=temp[i].Minor
		if resultLength==0{resultLength+=1}
	}
	//fmt.Printf("%v",resultLength)
	var result=make([]*semver.Version, resultLength)
	var position=0
	Major=-1
	Minor=-1
	for i:=0;i<length;i++{
		if temp[i].Major==Major{
			if temp[i].Minor==Minor{continue}
			Major=temp[i].Major
			Minor=temp[i].Minor
			result[position]=temp[i]
			position+=1
			continue
		}
		Major=temp[i].Major
		Minor=temp[i].Minor
		result[position]=temp[i]
		position+=1
	}

	versionSlice=result

	// This is just an example structure of the code, if you implement this interface, the test cases in main_test.go are very easy to run
	return versionSlice
}

func stringToVersionSlice(stringSlice []string) []*semver.Version {
	versionSlice := make([]*semver.Version, len(stringSlice))
	for i, versionString := range stringSlice {
		versionSlice[i] = semver.New(versionString)
	}
	return versionSlice
}
func main() {
	testCases := []struct {
		versionSlice   []string
		expectedResult []string
		minVersion     *semver.Version
	}{
		{
			versionSlice:   []string{"1.8.11", "1.9.6", "1.10.1", "1.9.5", "1.8.10", "1.10.0", "1.7.14", "1.8.9", "1.9.5"},
			expectedResult: []string{"1.10.1", "1.9.6", "1.8.11"},
			minVersion:     semver.New("1.8.0"),
		},
		{
			versionSlice:   []string{"1.8.11", "1.9.6", "1.10.1", "1.9.5", "1.8.10", "1.10.0", "1.7.14", "1.8.9", "1.9.5"},
			expectedResult: []string{"1.10.1", "1.9.6"},
			minVersion:     semver.New("1.8.12"),
		},
		{
			versionSlice:   []string{"1.10.1", "1.9.5", "1.8.10", "1.10.0", "1.7.14", "1.8.9", "1.9.5"},
			expectedResult: []string{"1.10.1"},
			minVersion:     semver.New("1.10.0"),
		},
		{
			versionSlice:   []string{"2.2.1", "2.2.0"},
			expectedResult: []string{"2.2.1"},
			minVersion:     semver.New("2.2.1"),
		},
		// Implement more relevant test cases here, if you can think of any
	}



for i:=0;i<len(testCases);i++{
		allReleases:= make([]*semver.Version, len(testCases[i].versionSlice))
		minVersion:=testCases[i].minVersion
		allReleases=stringToVersionSlice(testCases[i].versionSlice)
		versionSlice := LatestVersions(allReleases, minVersion)
		fmt.Printf("latest versions of kubernetes/kubernetes: %s", versionSlice)
}
	}
