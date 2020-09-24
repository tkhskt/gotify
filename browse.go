package gotify

import (
	"encoding/json"
	"github.com/tkhskt/gotify/extensions"
	"github.com/tkhskt/gotify/models"
)

// GetBrowseFeaturedPlaylists : the method for GET https://api.spotify.com/v1/browse/featured-playlists
func (g *Gotify) GetBrowseFeaturedPlaylists() (*models.BrowseFeaturedPlaylists, error) {
	/**
	https://developer.spotify.com/web-api/get-list-featured-playlists/
	*/

	endpoint := "https://api.spotify.com/v1/browse/featured-playlists"

	res, err := extensions.GetRequest(endpoint, g.TokenInfo.GetAccessToken())
	if err != nil {
		return nil, err
	}

	browseFeaturedPlaylists := new(models.BrowseFeaturedPlaylists)

	err = json.Unmarshal(res, browseFeaturedPlaylists)
	if err != nil {
		return nil, err
	}
	return browseFeaturedPlaylists, nil
}

// GetBrowseNewReleases : the method for GET https://api.spotify.com/v1/browse/new-releases
func (g *Gotify) GetBrowseNewReleases() (*models.BrowseNewReleases, error) {
	/**
	https://developer.spotify.com/web-api/get-list-new-releases/
	*/

	endpoint := "https://api.spotify.com/v1/browse/new-releases"

	res, err := extensions.GetRequest(endpoint, g.TokenInfo.GetAccessToken())
	if err != nil {
		return nil, err
	}

	browseNewReleases := new(models.BrowseNewReleases)

	err = json.Unmarshal(res, browseNewReleases)
	if err != nil {
		return nil, err
	}
	return browseNewReleases, nil
}

// GetBrowseCategories : the method for GET https://api.spotify.com/v1/browse/categories
func (g *Gotify) GetBrowseCategories() (*models.BrowseCategories, error) {
	/**
	https://developer.spotify.com/web-api/get-list-categories/
	*/

	endpoint := "https://api.spotify.com/v1/browse/categories"

	res, err := extensions.GetRequest(endpoint, g.TokenInfo.GetAccessToken())
	if err != nil {
		return nil, err
	}

	browseCategories := new(models.BrowseCategories)

	err = json.Unmarshal(res, browseCategories)
	if err != nil {
		return nil, err
	}
	return browseCategories, nil
}

// GetBrowseCategory : the method for GET https://api.spotify.com/v1/browse/categories/{category_id}
func (g *Gotify) GetBrowseCategory(categoryID string) (*models.BrowseCategory, error) {
	/**
	https://developer.spotify.com/web-api/get-category/
	*/

	endpoint := "https://api.spotify.com/v1/browse/categories/" + categoryID

	res, err := extensions.GetRequest(endpoint, g.TokenInfo.GetAccessToken())
	if err != nil {
		return nil, err
	}

	browseCategory := new(models.BrowseCategory)

	err = json.Unmarshal(res, browseCategory)
	if err != nil {
		return nil, err
	}
	return browseCategory, nil
}

// GetBrowseCategorysPlaylists : the method for GET https://api.spotify.com/v1/browse/categories/{category_id}/playlists
func (g *Gotify) GetBrowseCategorysPlaylists(categoryID string) (*models.BrowseCategorysPlaylists, error) {
	/**
	https://developer.spotify.com/web-api/get-categorys-playlists/
	*/

	endpoint := "https://api.spotify.com/v1/browse/categories/" + categoryID + "/playlists"

	res, err := extensions.GetRequest(endpoint, g.TokenInfo.GetAccessToken())
	if err != nil {
		return nil, err
	}

	browseCategorysPlaylists := new(models.BrowseCategorysPlaylists)

	err = json.Unmarshal(res, browseCategorysPlaylists)
	if err != nil {
		return nil, err
	}
	return browseCategorysPlaylists, nil
}
