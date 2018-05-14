package main

import (
	"context"
	"github.com/coreos/go-semver/semver"
	"github.com/google/go-github/github"
	"fmt"
)
func bubblesort(releases []*semver.Version,n int) {

	for i:=0;i<n-1;i++ {
		for j:=0;j<n-i-1;j++ {
			if releases[j].LessThan(*releases[j+1])==true {
				var temp *semver.Version//swap the element if j+1>j
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

	var length=0//porduce length for temp(raw array)
	for i:=0;i<len(releases);i++{
		if releases[i].Equal(*minVersion){
			length+=1;continue}//case for minVer= greatest

		if minVersion.LessThan(*releases[i])==true{
			length+=1
		}
	}

	//insert stuff into the temp array
	var temp=make([]*semver.Version, length)
	var index=0//the current position of the temp array
	for i:=0;i<len(releases);i++{
		if releases[i].Equal(*minVersion){
			temp[index]=releases[i]
			index+=1;continue}
		if minVersion.LessThan(*releases[i])==true{
			temp[index]=releases[i]
			index+=1
		}
	}
	//bubblesort
	bubblesort(temp,length)

	//get length of highest patch array
	var Major int64
	var Minor int64
	var resultLength=0
	Major=-1
	Minor=-1

	for i:=0;i<length;i++{
		if temp[i].Major==Major{
			if temp[i].Minor==Minor{
				continue}//if major minor equal then skip
			resultLength+=1
			Major=temp[i].Major//store the value of temp[i].major temp[i].minor
			Minor=temp[i].Minor
			continue
		}
		Major=temp[i].Major//store the value of temp[i].major temp[i].minor
		Minor=temp[i].Minor
		if resultLength==0{resultLength+=1}//base case for resultlength=0
	}
	//insert and sort the array
	var result=make([]*semver.Version, resultLength)
	var position=0
	Major=-1
	Minor=-1
	for i:=0;i<length;i++{
		if temp[i].Major==Major{
			if temp[i].Minor==Minor{
				continue}//case major minor equal then skip
			Major=temp[i].Major
			Minor=temp[i].Minor
			result[position]=temp[i]//insert the result
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

// Here we implement the basics of communicating with github through the library as well as printing the version
// You will need to implement LatestVersions function as well as make this application support the file format outlined in the README
// Please use the format defined by the fmt.Printf line at the bottom, as we will define a passing coding challenge as one that outputs
// the correct information, including this line
func main() {
	// Github
	client := github.NewClient(nil)
	ctx := context.Background()
	opt := &github.ListOptions{PerPage: 10}
	releases, _, err := client.Repositories.ListReleases(ctx, "kubernetes", "kubernetes", opt)
	if err != nil {
		panic(err) // is this really a good way?
	}
	minVersion := semver.New("1.8.0")
	allReleases := make([]*semver.Version, len(releases))
	for i, release := range releases {
		versionString := *release.TagName
		if versionString[0] == 'v' {
			versionString = versionString[1:]
		}
		allReleases[i] = semver.New(versionString)
	}
	versionSlice := LatestVersions(allReleases, minVersion)

	fmt.Printf("latest versions of kubernetes/kubernetes: %s", versionSlice)
}

