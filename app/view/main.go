package view

import "github.com/andyzhou/tackle/json"

/*
 * main page view
 */

type HomePageView struct {
	BaseView
}

type MainPageView struct {
	BaseView
}

type Video2GifHomePageView struct {
	BaseView
}

type Video2GifListPageView struct {
	FilesInfo        []*json.Video2GifFileJson `json:"FilesInfo"`
	ListMoreSwitcher bool                      `json:"ListMoreSwitcher"`
	BaseView
}
