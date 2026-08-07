package subsonic

import (
	"encoding/json"
	"time"
)

// Used for any entity where only a name and an ID appear.
// Eg. OpenSubsonic artists list for an album.
type IDName struct {
	ID   string `xml:"id,attr,omitempty"   json:"id,omitempty"`
	Name string `xml:"name,attr,omitempty" json:"name,omitempty"`
}

type OpenSubsonicExtension struct {
	Name     string `xml:"name,attr"  json:"name"`
	Versions []int  `xml:"versions"   json:"versions"`
}

type openSubsonicExtensions struct {
	OpenSubsonicExtensions []*OpenSubsonicExtension `xml:"openSubsonicExtension,omitempty" json:"openSubsonicExtension,omitempty"`
}

type LyricsList struct {
	StructuredLyrics []*StructuredLyrics `xml:"structuredLyrics,omitempty" json:"structuredLyrics,omitempty"`
}

type StructuredLyrics struct {
	DisplayArtist string      `xml:"displayArtist,attr" json:"displayArtist"`
	DisplayTitle  string      `xml:"displayTitle,attr"  json:"displayTitle"`
	Lang          string      `xml:"lang,attr"          json:"lang"`
	Offset        int         `xml:"offset,attr"        json:"offset"`
	Synced        bool        `xml:"synced,attr"        json:"synced"`
	Lines         []LyricLine `xml:"line,omitempty"     json:"line,omitempty"`
}

type LyricLine struct {
	Start int    `xml:"start,attr"  json:"start"`
	Text  string `xml:",chardata"   json:"value,omitempty"`
	// Navidrome 0.51.0 - 0.52.5 incorrectly returns the lyric line text here
	// This will be removed in the future
	Value string `xml:"value" json:"-"`
}

type ItemDate struct {
	Year  *int `xml:"year,attr,omitempty"  json:"year,omitempty"`
	Month *int `xml:"month,attr,omitempty" json:"month,omitempty"`
	Day   *int `xml:"day,attr,omitempty"   json:"day,omitempty"`
}

type Contributor struct {
	Role   string `xml:"role,attr" json:"role"`
	Artist IDName `xml:"artist"    json:"artist"`
}

type ReplayGain struct {
	TrackGain float64 `xml:"trackGain,omitempty,attr" json:"trackGain,omitempty"`
	AlbumGain float64 `xml:"albumGain,omitempty,attr" json:"albumGain,omitempty"`
	TrackPeak float64 `xml:"trackPeak,omitempty,attr" json:"trackPeak,omitempty"`
	AlbumPeak float64 `xml:"albumPeak,omitempty,attr" json:"albumPeak,omitempty"`
}

type PlayQueueByIndex struct {
	Entries      []*Child  `xml:"entry,omitempty"             json:"entry,omitempty"`
	CurrentIndex int64     `xml:"currentIndex,attr,omitempty" json:"currentIndex,omitempty"`
	Position     int64     `xml:"position,attr,omitempty"     json:"position,omitempty"`
	Username     string    `xml:"username,attr"               json:"username"`
	Changed      time.Time `xml:"changed,attr"                json:"changed"`
	ChangedBy    string    `xml:"changedBy,attr"              json:"changedBy"`
}

type DiscTitle struct {
	Disc     int    `xml:"disc,attr"               json:"disc"`
	Title    string `xml:"title,attr"              json:"title"`
	CoverArt string `xml:"coverArt,attr,omitempty" json:"coverArt,omitempty"`
}

func (t *PlayQueueByIndex) MarshalJSON() ([]byte, error) {
	type T PlayQueueByIndex
	var layout struct {
		*T
		Changed *xsdDateTime `json:"changed"`
	}
	layout.T = (*T)(t)
	layout.Changed = (*xsdDateTime)(&t.Changed)
	return json.Marshal(layout)
}

func (t *PlayQueueByIndex) UnmarshalJSON(data []byte) error {
	type T PlayQueueByIndex
	var overlay struct {
		*T
		Changed *xsdDateTime `json:"changed"`
	}
	overlay.T = (*T)(t)
	overlay.Changed = (*xsdDateTime)(&overlay.T.Changed)
	return json.Unmarshal(data, &overlay)
}
