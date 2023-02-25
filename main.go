package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"time"
)

type entry struct {
	path         string
	etype        string
	modifiedDate string
}

const (
	reset  = "\033[0m"
	blue   = "\033[34m"
	green  = "\033[32m"
	yellow = "\033[33m"
)

func main() {
	dir, err := filepath.Abs(".")
	if err != nil {
		panic(err)
	}

	entries := getEntries(dir)

	sort.Slice(entries, func(i, j int) bool {
		return getModifiedTime(entries[i].path).After(getModifiedTime(entries[j].path))
	})

	printEntries(entries)
}

func getEntries(dir string) []entry {
	entries := []entry{}

	dirFile, err := os.Open(dir)
	if err != nil {
		panic(err)
	}
	defer dirFile.Close()

	fileInfos, err := dirFile.Readdir(0)
	if err != nil {
		panic(err)
	}

	for _, fileInfo := range fileInfos {
		if !fileInfo.IsDir() {
			entries = append(entries, entry{
				path:         filepath.Join(dir, fileInfo.Name()),
				etype:        "file",
				modifiedDate: getModifiedTime(filepath.Join(dir, fileInfo.Name())).Format("2006-01-02"),
			})
		} else {
			entries = append(entries, entry{
				path:         filepath.Join(dir, fileInfo.Name()),
				etype:        "dir",
				modifiedDate: getLatestModifiedTime(filepath.Join(dir, fileInfo.Name())).Format("2006-01-02"),
			})
		}
	}

	return entries
}

func getModifiedTime(file string) time.Time {
	info, err := os.Stat(file)
	if err != nil {
		panic(err)
	}
	return info.ModTime()
}

func printEntries(entries []entry) {
	var lastDate string
	for _, entry := range entries {
		if entry.modifiedDate != lastDate {
			fmt.Printf("\n%s----------------%s-----------------%s\n", yellow, entry.modifiedDate, reset)
			lastDate = entry.modifiedDate
		}
		if entry.etype == "dir" {
			printDir(entry.path)
		} else {
			printFile(entry.path)
		}
	}
}

func printFile(file string) {
	fmt.Printf("%s%s%s\n", blue, file, reset)
}

func printDir(dir string) {
	fmt.Printf("%s%-50s %s\n", green, dir, reset)
}

func getLatestModifiedTime(dir string) time.Time {
	var latestTime time.Time

	dirFile, err := os.Open(dir)
	if err != nil {
		panic(err)
	}
	defer dirFile.Close()

	fileInfos, err := dirFile.Readdir(0)
	if err != nil {
		panic(err)
	}

	for _, fileInfo := range fileInfos {
		if !fileInfo.IsDir() {
			modifyTime := fileInfo.ModTime()
			if modifyTime.After(latestTime) {
				latestTime = modifyTime
			}
		}
	}

	return latestTime
}
