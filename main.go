package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)
	getAll := getCmd.Bool("all", false, "Get all videos")
	getId := getCmd.String("id", "", "Youtube video ID")

	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addID := addCmd.String("id", "", "Youtube video ID")
	addTitle := addCmd.String("title", "", "Youtube video Title")
	addUrl := addCmd.String("url", "", "Youtube video URL")
	addImageUrl := addCmd.String("imageurl", "", "Youtube video image URL")
	addDesc := addCmd.String("desc", "", "Youtube video Description")

	if len(os.Args) < 2 {
		fmt.Println("Expected 'get' or 'add' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "get":
		HandleGet(getCmd, getAll, getId)
	case "add":
		HandleAdd(addCmd, addID, addTitle, addUrl, addImageUrl, addDesc)
	default:
	}
}

func HandleGet(getCmd *flag.FlagSet, all *bool, id *string) {
	getCmd.Parse(os.Args[2:])

	if *all == false && *id == "" {
		fmt.Println("id is required or specify --all for all videos")
		getCmd.PrintDefaults()
		os.Exit(1)
	}

	if *all {
		videos := getVideos()

		fmt.Printf("ID \t Title \t URL \t ImageURL \t Description \n")

		for _, video := range videos {
			fmt.Printf("%v \t %v \t %v \t %v \t %v \n", video.Id, video.Title, video.Url, video.ImageUrl, video.Description)
		}
	}

	if *id != "" {
		videos := getVideos()
		id := *id
		for _, video := range videos {
			if id == video.Id {
				fmt.Printf("ID \t Title \t URL \t ImageURL \t Description \n")
				fmt.Printf("%v \t %v \t %v \t %v \t %v \n", video.Id, video.Title, video.Url, video.ImageUrl, video.Description)
			}
		}
	}
}

func ValidateVideo(addCmd *flag.FlagSet, id, title, url, imageurl, desc *string) {
	if *id == "" || *title == "" || *url == "" || *imageurl == "" || *desc == "" {
		fmt.Println("all fields are required for adding a video")
		addCmd.PrintDefaults()
		os.Exit(1)
	}
}

func HandleAdd(addCmd *flag.FlagSet, id, title, url, imageurl, desc *string) {
	ValidateVideo(addCmd, id, title, url, imageurl, desc)

	video := video{
		Id:          *id,
		Title:       *title,
		Description: *desc,
		Url:         *url,
		ImageUrl:    *imageurl,
	}

	videos := getVideos()
	videos = append(videos, video)

	saveVideos(videos)
}
