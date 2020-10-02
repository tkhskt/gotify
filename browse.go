package gotify

import (
	"github.com/tkhskt/gotify/models/browse"
)

// GetBrowseFeaturedPlaylists : the method for GET https://api.spotify.com/v1/browse/featured-playlists
func (g *Gotify) GetBrowseFeaturedPlaylists() (*browse.FeaturedPlaylists, error) {
	/**
	https://developer.spotify.com/web-api/get-list-featured-playlists/
	*/

	endpoint := "https://api.spotify.com/v1/browse/featured-playlists"

	var browseFeaturedPlaylists *browse.FeaturedPlaylists
	err := g.get(endpoint, g.TokenInfo.GetAccessToken(), &browseFeaturedPlaylists)
	if err != nil {
		return nil, err
	}

	return browseFeaturedPlaylists, nil
}

// GetBrowseNewReleases : the method for GET https://api.spotify.com/v1/browse/new-releases
func (g *Gotify) GetBrowseNewReleases() (*browse.NewReleases, error) {
	/**
	https://developer.spotify.com/web-api/get-list-new-releases/
	*/

	endpoint := "https://api.spotify.com/v1/browse/new-releases"

	var browseNewReleases *browse.NewReleases
	err := g.get(endpoint, g.TokenInfo.GetAccessToken(), browseNewReleases)
	if err != nil {
		return nil, err
	}

	return browseNewReleases, nil
}

// GetBrowseCategories : the method for GET https://api.spotify.com/v1/browse/categories
func (g *Gotify) GetBrowseCategories() (*browse.Categories, error) {
	/**
	https://developer.spotify.com/web-api/get-list-categories/
	*/

	endpoint := "https://api.spotify.com/v1/browse/categories"

	var browseCategories *browse.Categories
	err := g.get(endpoint, g.TokenInfo.GetAccessToken(), &browseCategories)
	if err != nil {
		return nil, err
	}

	return browseCategories, nil
}

// GetBrowseCategory : the method for GET https://api.spotify.com/v1/browse/categories/{category_id}
func (g *Gotify) GetBrowseCategory(categoryID string) (*browse.Category, error) {
	/**
	https://developer.spotify.com/web-api/get-category/
	*/

	endpoint := "https://api.spotify.com/v1/browse/categories/" + categoryID

	var browseCategory *browse.Category
	err := g.get(endpoint, g.TokenInfo.GetAccessToken(), &browseCategory)
	if err != nil {
		return nil, err
	}

	return browseCategory, nil
}

// GetBrowseCategorysPlaylists : the method for GET https://api.spotify.com/v1/browse/categories/{category_id}/playlists
func (g *Gotify) GetBrowseCategorysPlaylists(categoryID string) (*browse.CategoriesPlaylists, error) {
	/**
	https://developer.spotify.com/web-api/get-categorys-playlists/
	*/

	endpoint := "https://api.spotify.com/v1/browse/categories/" + categoryID + "/playlists"

	var browseCategoriesPlaylists *browse.CategoriesPlaylists
	err := g.get(endpoint, g.TokenInfo.GetAccessToken(), &browseCategoriesPlaylists)
	if err != nil {
		return nil, err
	}

	return browseCategoriesPlaylists, nil
}
