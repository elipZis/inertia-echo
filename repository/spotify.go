package repository

import (
	"errors"
	"github.com/zmb3/spotify"
)

// Definitely not the best to do it like this, but... who cares ^^'
var SpotifyClients = make(map[string]*spotify.Client)

// Save the authenticated client for later requests
func (this *Repository) StoreSpotifyClient(user *spotify.PrivateUser, client *spotify.Client) {
	SpotifyClients[user.ID] = client
}

// Get the user object for the given spotify id, if the user has an active client
func (this *Repository) GetSpotifyUserById(id string) (*spotify.PrivateUser, error) {
	if client, isPresent := SpotifyClients[id]; isPresent {
		return client.CurrentUser()
	}
	return nil, errors.New("spotify.error.no_user")
}

// Get the users recently played track, if the user has an active client
func (this *Repository) GetLastPlayedTracks(id string) ([]spotify.RecentlyPlayedItem, error) {
	if client, isPresent := SpotifyClients[id]; isPresent {
		return client.PlayerRecentlyPlayed()
	}
	return nil, errors.New("spotify.error.no_recent_trackes")
}
