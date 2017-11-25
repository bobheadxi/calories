package scheduler

import (
	"io/ioutil"
	"strings"
	"time"
)

var zoneDirs = []string{
	// Update path according to your OS
	"/usr/share/zoneinfo/",
	"/usr/share/lib/zoneinfo/",
	"/usr/lib/locale/TZ/",
}

var zoneDir string

// Scheduler :
type Scheduler struct {
}

//NewScheduler : constructor
func (scheduler *Scheduler) NewScheduler() *Scheduler {
	s := Scheduler{}
	go s.Start()
	return &s
}

// Start : start and check every hour for scheduled summaries
func (scheduler *Scheduler) Start() {

	t := time.NewTicker(time.Hour * 1)
	for {
		Do()
		<-t.C
	}
}

// Do : send user the scheduled summaries if time matches
func Do() {
	// get all regions with different timezone
	var allRegions []string
	for _, zoneDir = range zoneDirs {
		allRegions = append(allRegions, ReadFile("")...)
	}

	for _, region := range allRegions {
		loc, err := time.LoadLocation(region)
		if err != nil {
			// handle error
		}

		//check the time in that location (loc is a string representing for ex "America/New_York")
		t := time.Now().In(loc)
		if t.Equal(10:00) {
			// incomplete: get users according to their timezone and send them a message
		}
	}

}

// ReadFile : return a slice of all the timezones, where each string is for example "America/New_York"
func ReadFile(path string) []string {
	var regions []string
	files, _ := ioutil.ReadDir(zoneDir + path)
	for _, f := range files {
		if f.Name() != strings.ToUpper(f.Name()[:1])+f.Name()[1:] {
			continue
		}
		if f.IsDir() {
			ReadFile(path + "/" + f.Name())
		} else {
			regions = append(regions, f.Name())
		}
	}
	return regions
}
