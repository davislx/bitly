// Package bitlinks contains the methods to interact with the Bitlinks in Bitly
package bitlinks

import "encoding/json"

// Bitlink has information about the Bitlink
type Bitlink struct {
	Domain    string     `json:"domain"`
	Title     string     `json:"title"`
	GroupGUID string     `json:"group_guid"`
	Tags      []string   `json:"tags"`
	Deeplinks []Deeplink `json:"deeplinks"`
	LongURL   string     `json:"long_url"`
}

// BitlinkDetails has more in-depth information about the Bitlink
type BitlinkDetails struct {
	References     References `json:"references"`
	Archived       bool       `json:"archived"`
	Tags           []string   `json:"tags"`
	CreatedAt      string     `json:"created_at"`
	Title          string     `json:"title"`
	Deeplinks      []Deeplink `json:"deeplinks"`
	CreatedBy      string     `json:"created_by"`
	LongURL        string     `json:"long_url"`
	ClientID       string     `json:"client_id"`
	CustomBitlinks []string   `json:"custom_bitlinks"`
	Link           string     `json:"link"`
	ID             string     `json:"id"`
}

// Deeplink details
type Deeplink struct {
	Bitlink     string `json:"bitlink,omitempty"`
	InstallURL  string `json:"install_url"`
	Created     string `json:"created,omitempty"`
	AppURIPath  string `json:"app_uri_path"`
	Modified    string `json:"modified,omitempty"`
	InstallType string `json:"install_type"`
	AppGUID     string `json:"app_guid,omitempty"`
	GUID        string `json:"guid,omitempty"`
	OS          string `json:"os,omitempty"`
	AppID       string `json:"app_id,omitempty"`
}

// Link contains the single ID of a Bitlink
type Link struct {
	BitlinkID string `json:"bitlink_id"`
}

// LinkClick contains the number of clicks per date
type LinkClick struct {
	Date   string `json:"date"`
	Clicks int64  `json:"clicks"`
}

// LinkInfo contains information on the Bitlink
type LinkInfo struct {
	LongURL   string `json:"long_url"`
	CreatedAt string `json:"created_at"`
	Link      string `json:"link"`
	ID        string `json:"id"`
}

// Metric contains information on the chosen metric
type Metric struct {
	Clicks int64  `json:"clicks"`
	Value  string `json:"value"`
}

// Metrics is the top level struct returned for all metrics requests
type Metrics struct {
	Units             int64               `json:"units,omitempty"`
	Unit              string              `json:"unit,omitempty"`
	TotalClicks       int64               `json:"total_clicks,omitempty"`
	UnitReference     string              `json:"unit_reference,omitempty"`
	LinkClicks        []LinkClick         `json:"link_clicks,omitempty"`
	Facet             string              `json:"facet,omitempty"`
	Metrics           []Metric            `json:"metrics,omitempty"`
	ReferrersByDomain []ReferrersByDomain `json:"referrers_by_domain,omitempty"`
}

// MetricsRequest is used to generate the metrics request to Bitly
type MetricsRequest struct {
	// A unit of time
	Unit string
	// An integer representing the time units to query data for. pass -1 to return all units of time.
	Units int
	// An ISO-8601 timestamp, indicating the most recent time for which to pull metrics. Will default to current time.
	UnitReference string
	// The quantity of items to be be returned
	Size int
}

// References contains properties generated by Bitly
type References struct {
	Property1 string `json:"property1"`
	Property2 string `json:"property2"`
}

// Referrer contains information about the referrer of a Bitlink click
type Referrer struct {
	Value int64  `json:"value"`
	Key   string `json:"key"`
}

// ReferrersByDomain contains information about the domain of the referrer of a Bitlink click
type ReferrersByDomain struct {
	Referrers []Referrer `json:"Referrers"`
	Network   string     `json:"network"`
}

// ShortenRequest contains the request details to shorten a link using Bitly
type ShortenRequest struct {
	GroupGUID string `json:"group_guid"`
	Domain    string `json:"domain"`
	LongURL   string `json:"long_url"`
}

func (r *Bitlink) marshal() ([]byte, error) {
	return json.Marshal(r)
}

func (r *BitlinkDetails) marshal() ([]byte, error) {
	return json.Marshal(r)
}

func (r *Link) marshal() ([]byte, error) {
	return json.Marshal(r)
}

func (r *ShortenRequest) marshal() ([]byte, error) {
	return json.Marshal(r)
}

func unmarshalBitlinkDetails(data []byte) (BitlinkDetails, error) {
	var r BitlinkDetails
	err := json.Unmarshal(data, &r)
	return r, err
}

func unmarshalLinkInfo(data []byte) (LinkInfo, error) {
	var r LinkInfo
	err := json.Unmarshal(data, &r)
	return r, err
}

func unmarshalMetrics(data []byte) (Metrics, error) {
	var r Metrics
	err := json.Unmarshal(data, &r)
	return r, err
}