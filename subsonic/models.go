package subsonic

/* This file was automatically generated from the xsd schema provided by Subsonic, then manually modified.
 *   http://www.subsonic.org/pages/inc/api/schema/subsonic-rest-api-1.16.1.xsd
 *   xsdgen -o xml.go -pkg subsonic -ns "http://subsonic.org/restapi" subsonic-rest-api-1.16.1.xsd
 * Changes from the original include:
 * - Adding missing name (value of xml element) for each genre
 * - Capitalize "ID" in struct names and add missing ID fields.
 * - Merge *With* variants of structs.
 */

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"time"
)

type PingResponse struct {
	Status        string
	Version       string
	Type          string
	ServerVersion string
	OpenSubsonic  bool
	ServerError   *Error
}

// AlbumID3 is an album that's organized by music file tags.
type AlbumID3 struct {
	ID                  string      `xml:"id,attr"                       json:"id"`
	Song                []*Child    `xml:"song,omitempty"                json:"song,omitempty"`
	Name                string      `xml:"name,attr"                     json:"name"`
	SortName            string      `xml:"sortName,attr,omitempty"       json:"sortName,omitempty"`
	Artist              string      `xml:"artist,attr,omitempty"         json:"artist,omitempty"`
	ArtistID            string      `xml:"artistId,attr,omitempty"       json:"artistId,omitempty"`
	Artists             []IDName    `xml:"artists,omitempty"             json:"artists,omitempty"`
	CoverArt            string      `xml:"coverArt,attr,omitempty"       json:"coverArt,omitempty"`
	SongCount           int         `xml:"songCount,attr"                json:"songCount"`
	Duration            int         `xml:"duration,attr"                 json:"duration"`
	PlayCount           int64       `xml:"playCount,attr,omitempty"      json:"playCount,omitempty"`
	Created             time.Time   `xml:"created,attr"                  json:"created"`
	Starred             time.Time   `xml:"starred,attr,omitempty"        json:"starred,omitempty"`
	Year                int         `xml:"year,attr,omitempty"           json:"year,omitempty"`
	ReleaseDate         *ItemDate   `xml:"releaseDate,omitempty"         json:"releaseDate,omitempty"`
	OriginalReleaseDate *ItemDate   `xml:"originalReleaseDate,omitempty" json:"originalReleaseDate,omitempty"`
	Genre               string      `xml:"genre,attr,omitempty"          json:"genre,omitempty"`
	Genres              []IDName    `xml:"genres,omitempty"              json:"genres,omitempty"`
	ReleaseTypes        []string    `xml:"releaseTypes,omitempty"        json:"releaseTypes,omitempty"`
	IsCompilation       bool        `xml:"isCompilation,attr"            json:"isCompilation,omitempty"`
	DiscTitles          []DiscTitle `xml:"discTitles,omitempty"          json:"discTitles,omitempty"`
}

func (t *AlbumID3) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T AlbumID3
	var layout struct {
		*T
		Created *xsdDateTime `xml:"created,attr"`
		Starred *xsdDateTime `xml:"starred,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.Created = (*xsdDateTime)(&layout.T.Created)
	layout.Starred = (*xsdDateTime)(&layout.T.Starred)
	return e.EncodeElement(layout, start)
}
func (t *AlbumID3) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T AlbumID3
	var overlay struct {
		*T
		Created *xsdDateTime `xml:"created,attr"`
		Starred *xsdDateTime `xml:"starred,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Created = (*xsdDateTime)(&overlay.T.Created)
	overlay.Starred = (*xsdDateTime)(&overlay.T.Starred)
	return d.DecodeElement(&overlay, &start)
}
func (t *AlbumID3) MarshalJSON() ([]byte, error) {
	type T AlbumID3
	var layout struct {
		*T
		Created *xsdDateTime `json:"created"`
		Starred *xsdDateTime `json:"starred,omitempty"`
	}
	layout.T = (*T)(t)
	layout.Created = (*xsdDateTime)(&t.Created)
	if !t.Starred.IsZero() {
		layout.Starred = (*xsdDateTime)(&t.Starred)
	}
	return json.Marshal(layout)
}
func (t *AlbumID3) UnmarshalJSON(data []byte) error {
	type T AlbumID3
	var overlay struct {
		*T
		Created *xsdDateTime `json:"created"`
		Starred *xsdDateTime `json:"starred,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Created = (*xsdDateTime)(&overlay.T.Created)
	overlay.Starred = (*xsdDateTime)(&overlay.T.Starred)
	return json.Unmarshal(data, &overlay)
}

// AlbumInfo is a collection of notes and links describing an album.
type AlbumInfo struct {
	Notes          string `xml:"notes,omitempty"          json:"notes,omitempty"`
	MusicBrainzID  string `xml:"musicBrainzId,omitempty"  json:"musicBrainzId,omitempty"`
	LastFmUrl      string `xml:"lastFmUrl,omitempty"      json:"lastFmUrl,omitempty"`
	SmallImageUrl  string `xml:"smallImageUrl,omitempty"  json:"smallImageUrl,omitempty"`
	MediumImageUrl string `xml:"mediumImageUrl,omitempty" json:"mediumImageUrl,omitempty"`
	LargeImageUrl  string `xml:"largeImageUrl,omitempty"  json:"largeImageUrl,omitempty"`
}

type albumList struct {
	Album []*Child `xml:"album,omitempty" json:"album,omitempty"`
}

type albumList2 struct {
	Album []*AlbumID3 `xml:"album,omitempty" json:"album,omitempty"`
}

// Artist is an artist from the server, organized in the folders pattern.
type Artist struct {
	ID             string    `xml:"id,attr"                      json:"id"`
	Name           string    `xml:"name,attr"                    json:"name"`
	ArtistImageUrl string    `xml:"artistImageUrl,attr,omitempty" json:"artistImageUrl,omitempty"`
	Starred        time.Time `xml:"starred,attr,omitempty"       json:"starred,omitempty"`
	UserRating     int       `xml:"userRating,attr,omitempty"    json:"userRating,omitempty"`
	AverageRating  float64   `xml:"averageRating,attr,omitempty" json:"averageRating,omitempty"`
}

func (t *Artist) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T Artist
	var layout struct {
		*T
		Starred *xsdDateTime `xml:"starred,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.Starred = (*xsdDateTime)(&layout.T.Starred)
	return e.EncodeElement(layout, start)
}
func (t *Artist) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T Artist
	var overlay struct {
		*T
		Starred *xsdDateTime `xml:"starred,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Starred = (*xsdDateTime)(&overlay.T.Starred)
	return d.DecodeElement(&overlay, &start)
}
func (t *Artist) MarshalJSON() ([]byte, error) {
	type T Artist
	var layout struct {
		*T
		Starred *xsdDateTime `json:"starred,omitempty"`
	}
	layout.T = (*T)(t)
	if !t.Starred.IsZero() {
		layout.Starred = (*xsdDateTime)(&t.Starred)
	}
	return json.Marshal(layout)
}
func (t *Artist) UnmarshalJSON(data []byte) error {
	type T Artist
	var overlay struct {
		*T
		Starred *xsdDateTime `json:"starred,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Starred = (*xsdDateTime)(&overlay.T.Starred)
	return json.Unmarshal(data, &overlay)
}

// ArtistID3 is an artist from the server, organized by ID3 tag.
type ArtistID3 struct {
	ID             string      `xml:"id,attr"                       json:"id"`
	Album          []*AlbumID3 `xml:"album,omitempty"               json:"album,omitempty"`
	Name           string      `xml:"name,attr"                     json:"name"`
	CoverArt       string      `xml:"coverArt,attr,omitempty"       json:"coverArt,omitempty"`
	ArtistImageUrl string      `xml:"artistImageUrl,attr,omitempty" json:"artistImageUrl,omitempty"`
	AlbumCount     int         `xml:"albumCount,attr"               json:"albumCount"`
	Starred        time.Time   `xml:"starred,attr,omitempty"        json:"starred,omitempty"`
	SortName       string      `xml:"sortName,attr,omitempty"       json:"sortName,omitempty"`
	MusicBrainzId  string      `xml:"musicBrainzId,attr,omitempty"  json:"musicBrainzId,omitempty"`
}

func (t *ArtistID3) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T ArtistID3
	var layout struct {
		*T
		Starred *xsdDateTime `xml:"starred,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.Starred = (*xsdDateTime)(&layout.T.Starred)
	return e.EncodeElement(layout, start)
}
func (t *ArtistID3) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T ArtistID3
	var overlay struct {
		*T
		Starred *xsdDateTime `xml:"starred,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Starred = (*xsdDateTime)(&overlay.T.Starred)
	return d.DecodeElement(&overlay, &start)
}
func (t *ArtistID3) MarshalJSON() ([]byte, error) {
	type T ArtistID3
	var layout struct {
		*T
		Starred *xsdDateTime `json:"starred,omitempty"`
	}
	layout.T = (*T)(t)
	if !t.Starred.IsZero() {
		layout.Starred = (*xsdDateTime)(&t.Starred)
	}
	return json.Marshal(layout)
}
func (t *ArtistID3) UnmarshalJSON(data []byte) error {
	type T ArtistID3
	var overlay struct {
		*T
		Starred *xsdDateTime `json:"starred,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Starred = (*xsdDateTime)(&overlay.T.Starred)
	return json.Unmarshal(data, &overlay)
}

// ArtistInfo is all auxillary information about an artist from GetArtistInfo.
type ArtistInfo struct {
	SimilarArtist  []*Artist `xml:"similarArtist,omitempty"  json:"similarArtist,omitempty"`
	Biography      string    `xml:"biography,omitempty"      json:"biography,omitempty"`
	MusicBrainzID  string    `xml:"musicBrainzId,omitempty"  json:"musicBrainzId,omitempty"`
	LastFmUrl      string    `xml:"lastFmUrl,omitempty"      json:"lastFmUrl,omitempty"`
	SmallImageUrl  string    `xml:"smallImageUrl,omitempty"  json:"smallImageUrl,omitempty"`
	MediumImageUrl string    `xml:"mediumImageUrl,omitempty" json:"mediumImageUrl,omitempty"`
	LargeImageUrl  string    `xml:"largeImageUrl,omitempty"  json:"largeImageUrl,omitempty"`
}

// ArtistInfo2 is all auxillary information about an artist from GetArtistInfo2, with similar artists organized by ID3 tags.
type ArtistInfo2 struct {
	SimilarArtist  []*ArtistID3 `xml:"similarArtist,omitempty"  json:"similarArtist,omitempty"`
	Biography      string       `xml:"biography,omitempty"      json:"biography,omitempty"`
	MusicBrainzID  string       `xml:"musicBrainzId,omitempty"  json:"musicBrainzId,omitempty"`
	LastFmUrl      string       `xml:"lastFmUrl,omitempty"      json:"lastFmUrl,omitempty"`
	SmallImageUrl  string       `xml:"smallImageUrl,omitempty"  json:"smallImageUrl,omitempty"`
	MediumImageUrl string       `xml:"mediumImageUrl,omitempty" json:"mediumImageUrl,omitempty"`
	LargeImageUrl  string       `xml:"largeImageUrl,omitempty"  json:"largeImageUrl,omitempty"`
}

// ArtistsID3 is an index of every artist on the server organized by ID3 tag, from getArtists.
type ArtistsID3 struct {
	Index           []*IndexID3 `xml:"index,omitempty"       json:"index,omitempty"`
	IgnoredArticles string      `xml:"ignoredArticles,attr"  json:"ignoredArticles"`
}

type Bookmark struct {
	Entry    *Child    `xml:"entry"              json:"entry"`
	Position int64     `xml:"position,attr"      json:"position"`
	Username string    `xml:"username,attr"      json:"username"`
	Comment  string    `xml:"comment,attr,omitempty" json:"comment,omitempty"`
	Created  time.Time `xml:"created,attr"       json:"created"`
	Changed  time.Time `xml:"changed,attr"       json:"changed"`
}

func (t *Bookmark) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T Bookmark
	var layout struct {
		*T
		Created *xsdDateTime `xml:"created,attr"`
		Changed *xsdDateTime `xml:"changed,attr"`
	}
	layout.T = (*T)(t)
	layout.Created = (*xsdDateTime)(&layout.T.Created)
	layout.Changed = (*xsdDateTime)(&layout.T.Changed)
	return e.EncodeElement(layout, start)
}
func (t *Bookmark) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T Bookmark
	var overlay struct {
		*T
		Created *xsdDateTime `xml:"created,attr"`
		Changed *xsdDateTime `xml:"changed,attr"`
	}
	overlay.T = (*T)(t)
	overlay.Created = (*xsdDateTime)(&overlay.T.Created)
	overlay.Changed = (*xsdDateTime)(&overlay.T.Changed)
	return d.DecodeElement(&overlay, &start)
}
func (t *Bookmark) MarshalJSON() ([]byte, error) {
	type T Bookmark
	var layout struct {
		*T
		Created *xsdDateTime `json:"created"`
		Changed *xsdDateTime `json:"changed"`
	}
	layout.T = (*T)(t)
	layout.Created = (*xsdDateTime)(&t.Created)
	layout.Changed = (*xsdDateTime)(&t.Changed)
	return json.Marshal(layout)
}
func (t *Bookmark) UnmarshalJSON(data []byte) error {
	type T Bookmark
	var overlay struct {
		*T
		Created *xsdDateTime `json:"created"`
		Changed *xsdDateTime `json:"changed"`
	}
	overlay.T = (*T)(t)
	overlay.Created = (*xsdDateTime)(&overlay.T.Created)
	overlay.Changed = (*xsdDateTime)(&overlay.T.Changed)
	return json.Unmarshal(data, &overlay)
}

type bookmarks struct {
	Bookmark []*Bookmark `xml:"bookmark,omitempty" json:"bookmark,omitempty"`
}

type ChatMessage struct {
	Username string `xml:"username,attr" json:"username"`
	Time     int64  `xml:"time,attr"     json:"time"`
	Message  string `xml:"message,attr"  json:"message"`
}

type chatMessages struct {
	ChatMessage []*ChatMessage `xml:"chatMessage,omitempty" json:"chatMessage,omitempty"`
}

// Child is a song, or a generic entry in the hierarchical directory structure of the database.
// You can tell if Child is used as a song contextually based on what it was returned by, or if the IsDir boolean was set to true.
type Child struct {
	ID                    string        `xml:"id,attr"                            json:"id"`
	Parent                string        `xml:"parent,attr,omitempty"              json:"parent,omitempty"`
	IsDir                 bool          `xml:"isDir,attr"                         json:"isDir"`
	Title                 string        `xml:"title,attr"                         json:"title"`
	Album                 string        `xml:"album,attr,omitempty"               json:"album,omitempty"`
	Artist                string        `xml:"artist,attr,omitempty"              json:"artist,omitempty"`
	Artists               []IDName      `xml:"artists,omitempty"                  json:"artists,omitempty"`
	AlbumArtists          []IDName      `xml:"albumArtists,omitempty"             json:"albumArtists,omitempty"`
	DisplayArtist         string        `xml:"displayArtist,attr,omitempty"       json:"displayArtist,omitempty"`
	DisplayAlbumArtist    string        `xml:"displayAlbumArtist,attr,omitempty"  json:"displayAlbumArtist,omitempty"`
	Contributors          []Contributor `xml:"contributors,omitempty"             json:"contributors,omitempty"`
	DisplayComposer       string        `xml:"displayComposer,attr,omitempty"     json:"displayComposer,omitempty"`
	Track                 int           `xml:"track,attr,omitempty"               json:"track,omitempty"`
	Year                  int           `xml:"year,attr,omitempty"                json:"year,omitempty"`
	Genre                 string        `xml:"genre,attr,omitempty"               json:"genre,omitempty"`
	Genres                []IDName      `xml:"genres,omitempty"                   json:"genres,omitempty"`
	Comment               string        `xml:"comment,attr,omitempty"             json:"comment,omitempty"`
	BPM                   int           `xml:"bpm,attr"                           json:"bpm,omitempty"`
	MusicBrainzID         string        `xml:"musicBrainzId,attr,omitempty"       json:"musicBrainzId,omitempty"`
	CoverArt              string        `xml:"coverArt,attr,omitempty"            json:"coverArt,omitempty"`
	Size                  int64         `xml:"size,attr,omitempty"                json:"size,omitempty"`
	ContentType           string        `xml:"contentType,attr,omitempty"         json:"contentType,omitempty"`
	Suffix                string        `xml:"suffix,attr,omitempty"              json:"suffix,omitempty"`
	TranscodedContentType string        `xml:"transcodedContentType,attr,omitempty" json:"transcodedContentType,omitempty"`
	TranscodedSuffix      string        `xml:"transcodedSuffix,attr,omitempty"    json:"transcodedSuffix,omitempty"`
	Duration              int           `xml:"duration,attr,omitempty"            json:"duration,omitempty"`
	BitRate               int           `xml:"bitRate,attr,omitempty"             json:"bitRate,omitempty"`
	BitDepth              int           `xml:"bitDepth,attr,omitempty"            json:"bitDepth,omitempty"`
	SamplingRate          int           `xml:"samplingRate,attr,omitempty"        json:"samplingRate,omitempty"`
	ChannelCount          int           `xml:"channelCount,attr,omitempty"        json:"channelCount,omitempty"`
	Path                  string        `xml:"path,attr,omitempty"                json:"path,omitempty"`
	IsVideo               bool          `xml:"isVideo,attr,omitempty"             json:"isVideo,omitempty"`
	UserRating            int           `xml:"userRating,attr,omitempty"          json:"userRating,omitempty"`
	AverageRating         float64       `xml:"averageRating,attr,omitempty"       json:"averageRating,omitempty"`
	PlayCount             int64         `xml:"playCount,attr,omitempty"           json:"playCount,omitempty"`
	DiscNumber            int           `xml:"discNumber,attr,omitempty"          json:"discNumber,omitempty"`
	Created               time.Time     `xml:"created,attr,omitempty"             json:"created,omitempty"`
	Starred               time.Time     `xml:"starred,attr,omitempty"             json:"starred,omitempty"`
	Played                time.Time     `xml:"played,attr,omitempty"              json:"played,omitempty"`
	AlbumID               string        `xml:"albumId,attr,omitempty"             json:"albumId,omitempty"`
	ArtistID              string        `xml:"artistId,attr,omitempty"            json:"artistId,omitempty"`
	Type                  string        `xml:"type,attr,omitempty"                json:"type,omitempty"`
	BookmarkPosition      int64         `xml:"bookmarkPosition,attr,omitempty"    json:"bookmarkPosition,omitempty"`
	OriginalWidth         int           `xml:"originalWidth,attr,omitempty"       json:"originalWidth,omitempty"`
	OriginalHeight        int           `xml:"originalHeight,attr,omitempty"      json:"originalHeight,omitempty"`
	ReplayGain            *ReplayGain   `xml:"replayGain,omitempty"               json:"replayGain,omitempty"`
}

func (t *Child) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T Child
	var layout struct {
		*T
		Created *xsdDateTime `xml:"created,attr,omitempty"`
		Starred *xsdDateTime `xml:"starred,attr,omitempty"`
		Played  *xsdDateTime `xml:"played,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.Created = (*xsdDateTime)(&layout.T.Created)
	layout.Starred = (*xsdDateTime)(&layout.T.Starred)
	layout.Played = (*xsdDateTime)(&layout.T.Played)
	return e.EncodeElement(layout, start)
}
func (t *Child) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T Child
	var overlay struct {
		*T
		Created *xsdDateTime `xml:"created,attr,omitempty"`
		Starred *xsdDateTime `xml:"starred,attr,omitempty"`
		Played  *xsdDateTime `xml:"played,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Created = (*xsdDateTime)(&overlay.T.Created)
	overlay.Starred = (*xsdDateTime)(&overlay.T.Starred)
	overlay.Played = (*xsdDateTime)(&overlay.T.Played)
	return d.DecodeElement(&overlay, &start)
}
func (t *Child) MarshalJSON() ([]byte, error) {
	type T Child
	var layout struct {
		*T
		Created *xsdDateTime `json:"created,omitempty"`
		Starred *xsdDateTime `json:"starred,omitempty"`
		Played  *xsdDateTime `json:"played,omitempty"`
	}
	layout.T = (*T)(t)
	if !t.Created.IsZero() {
		layout.Created = (*xsdDateTime)(&t.Created)
	}
	if !t.Starred.IsZero() {
		layout.Starred = (*xsdDateTime)(&t.Starred)
	}
	if !t.Played.IsZero() {
		layout.Played = (*xsdDateTime)(&t.Played)
	}
	return json.Marshal(layout)
}
func (t *Child) UnmarshalJSON(data []byte) error {
	type T Child
	var overlay struct {
		*T
		Created *xsdDateTime `json:"created,omitempty"`
		Starred *xsdDateTime `json:"starred,omitempty"`
		Played  *xsdDateTime `json:"played,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Created = (*xsdDateTime)(&overlay.T.Created)
	overlay.Starred = (*xsdDateTime)(&overlay.T.Starred)
	overlay.Played = (*xsdDateTime)(&overlay.T.Played)
	return json.Unmarshal(data, &overlay)
}

// Directory is an entry in the hierarchical folder structure organization of the server database.
type Directory struct {
	ID            string    `xml:"id,attr"                     json:"id"`
	Child         []*Child  `xml:"child,omitempty"             json:"child,omitempty"`
	Parent        string    `xml:"parent,attr,omitempty"       json:"parent,omitempty"`
	Name          string    `xml:"name,attr"                   json:"name"`
	Starred       time.Time `xml:"starred,attr,omitempty"      json:"starred,omitempty"`
	UserRating    int       `xml:"userRating,attr,omitempty"   json:"userRating,omitempty"`
	AverageRating float64   `xml:"averageRating,attr,omitempty" json:"averageRating,omitempty"`
	PlayCount     int64     `xml:"playCount,attr,omitempty"    json:"playCount,omitempty"`
}

func (t *Directory) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T Directory
	var layout struct {
		*T
		Starred *xsdDateTime `xml:"starred,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.Starred = (*xsdDateTime)(&layout.T.Starred)
	return e.EncodeElement(layout, start)
}
func (t *Directory) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T Directory
	var overlay struct {
		*T
		Starred *xsdDateTime `xml:"starred,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Starred = (*xsdDateTime)(&overlay.T.Starred)
	return d.DecodeElement(&overlay, &start)
}
func (t *Directory) MarshalJSON() ([]byte, error) {
	type T Directory
	var layout struct {
		*T
		Starred *xsdDateTime `json:"starred,omitempty"`
	}
	layout.T = (*T)(t)
	if !t.Starred.IsZero() {
		layout.Starred = (*xsdDateTime)(&t.Starred)
	}
	return json.Marshal(layout)
}
func (t *Directory) UnmarshalJSON(data []byte) error {
	type T Directory
	var overlay struct {
		*T
		Starred *xsdDateTime `json:"starred,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Starred = (*xsdDateTime)(&overlay.T.Starred)
	return json.Unmarshal(data, &overlay)
}

type Error struct {
	Code    int    `xml:"code,attr"           json:"code"`
	Message string `xml:"message,attr,omitempty" json:"message,omitempty"`
}

// Genre is a style tag for a collection of songs and albums.
type Genre struct {
	Name       string `xml:",chardata"        json:"value"`
	SongCount  int    `xml:"songCount,attr"   json:"songCount"`
	AlbumCount int    `xml:"albumCount,attr"  json:"albumCount"`
}

type genres struct {
	Genre []*Genre `xml:"genre,omitempty" json:"genre,omitempty"`
}

// Index is a collection of artists that begin with the same first letter, along with that letter or category.
type Index struct {
	Artist []*Artist `xml:"artist,omitempty" json:"artist,omitempty"`
	Name   string    `xml:"name,attr"        json:"name"`
}

// Index is a collection of artists by ID3 tag that begin with the same first letter, along with that letter or category.
type IndexID3 struct {
	Artist []*ArtistID3 `xml:"artist,omitempty" json:"artist,omitempty"`
	Name   string       `xml:"name,attr"        json:"name"`
}

// Indexes is the full index of the database, returned by getIndex.
// It contains some Index structs for each letter of the DB, plus Child entries for individual tracks.
type Indexes struct {
	Shortcut        []*Artist `xml:"shortcut,omitempty"      json:"shortcut,omitempty"`
	Index           []*Index  `xml:"index,omitempty"         json:"index,omitempty"`
	Child           []*Child  `xml:"child,omitempty"         json:"child,omitempty"`
	LastModified    int64     `xml:"lastModified,attr"       json:"lastModified"`
	IgnoredArticles string    `xml:"ignoredArticles,attr"    json:"ignoredArticles"`
}

type InternetRadioStation struct {
	Name        string `xml:"name,attr"                  json:"name"`
	StreamUrl   string `xml:"streamUrl,attr"             json:"streamUrl"`
	HomePageUrl string `xml:"homePageUrl,attr,omitempty" json:"homePageUrl,omitempty"`
	CoverArt    string `xml:"coverArt,attr,omitempty"    json:"coverArt,omitempty"`
}

type internetRadioStations struct {
	InternetRadioStation []*InternetRadioStation `xml:"internetRadioStation,omitempty" json:"internetRadioStation,omitempty"`
}

type JukeboxPlaylist struct {
	Entry        []*Child `xml:"entry,omitempty"       json:"entry,omitempty"`
	CurrentIndex int      `xml:"currentIndex,attr"     json:"currentIndex"`
	Playing      bool     `xml:"playing,attr"          json:"playing"`
	Gain         float32  `xml:"gain,attr"             json:"gain"`
	Position     int      `xml:"position,attr,omitempty" json:"position,omitempty"`
}

type JukeboxStatus struct {
	CurrentIndex int     `xml:"currentIndex,attr"     json:"currentIndex"`
	Playing      bool    `xml:"playing,attr"          json:"playing"`
	Gain         float32 `xml:"gain,attr"             json:"gain"`
	Position     int     `xml:"position,attr,omitempty" json:"position,omitempty"`
}

// License contains information about the Subsonic server's license validity and contact information in the case of a trial subscription.
type License struct {
	Valid          bool      `xml:"valid,attr"                      json:"valid"`
	Email          string    `xml:"email,attr,omitempty"            json:"email,omitempty"`
	LicenseExpires time.Time `xml:"licenseExpires,attr,omitempty"   json:"licenseExpires,omitempty"`
	TrialExpires   time.Time `xml:"trialExpires,attr,omitempty"     json:"trialExpires,omitempty"`
}

func (t *License) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T License
	var layout struct {
		*T
		LicenseExpires *xsdDateTime `xml:"licenseExpires,attr,omitempty"`
		TrialExpires   *xsdDateTime `xml:"trialExpires,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.LicenseExpires = (*xsdDateTime)(&layout.T.LicenseExpires)
	layout.TrialExpires = (*xsdDateTime)(&layout.T.TrialExpires)
	return e.EncodeElement(layout, start)
}
func (t *License) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T License
	var overlay struct {
		*T
		LicenseExpires *xsdDateTime `xml:"licenseExpires,attr,omitempty"`
		TrialExpires   *xsdDateTime `xml:"trialExpires,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.LicenseExpires = (*xsdDateTime)(&overlay.T.LicenseExpires)
	overlay.TrialExpires = (*xsdDateTime)(&overlay.T.TrialExpires)
	return d.DecodeElement(&overlay, &start)
}
func (t *License) MarshalJSON() ([]byte, error) {
	type T License
	var layout struct {
		*T
		LicenseExpires *xsdDateTime `json:"licenseExpires,omitempty"`
		TrialExpires   *xsdDateTime `json:"trialExpires,omitempty"`
	}
	layout.T = (*T)(t)
	if !t.LicenseExpires.IsZero() {
		layout.LicenseExpires = (*xsdDateTime)(&t.LicenseExpires)
	}
	if !t.TrialExpires.IsZero() {
		layout.TrialExpires = (*xsdDateTime)(&t.TrialExpires)
	}
	return json.Marshal(layout)
}
func (t *License) UnmarshalJSON(data []byte) error {
	type T License
	var overlay struct {
		*T
		LicenseExpires *xsdDateTime `json:"licenseExpires,omitempty"`
		TrialExpires   *xsdDateTime `json:"trialExpires,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.LicenseExpires = (*xsdDateTime)(&overlay.T.LicenseExpires)
	overlay.TrialExpires = (*xsdDateTime)(&overlay.T.TrialExpires)
	return json.Unmarshal(data, &overlay)
}

type Lyrics struct {
	Artist string `xml:"artist,attr,omitempty" json:"artist,omitempty"`
	Title  string `xml:"title,attr,omitempty"  json:"title,omitempty"`
	Text   string `xml:",chardata"             json:"value,omitempty"`
}

// MusicFolder is a representation of a source of music files added to the server. It is identified primarily by the numeric ID.
type MusicFolder struct {
	ID   int    `xml:"id,attr"           json:"id"`
	Name string `xml:"name,attr,omitempty" json:"name,omitempty"`
}

type musicFolders struct {
	MusicFolder []*MusicFolder `xml:"musicFolder,omitempty" json:"musicFolder,omitempty"`
}

type newestPodcasts struct {
	Episode []*PodcastEpisode `xml:"episode,omitempty" json:"episode,omitempty"`
}

type nowPlaying struct {
	Entry []*NowPlayingEntry `xml:"entry,omitempty" json:"entry,omitempty"`
}

// NowPlayingEntry is one individual stream coming from the server along with information about who was streaming it.
type NowPlayingEntry struct {
	Username              string    `xml:"username,attr"                      json:"username"`
	MinutesAgo            int       `xml:"minutesAgo,attr"                    json:"minutesAgo"`
	PlayerID              int       `xml:"playerId,attr"                      json:"playerId"`
	PlayerName            string    `xml:"playerName,attr,omitempty"          json:"playerName,omitempty"`
	Parent                string    `xml:"parent,attr,omitempty"              json:"parent,omitempty"`
	IsDir                 bool      `xml:"isDir,attr"                         json:"isDir"`
	Title                 string    `xml:"title,attr"                         json:"title"`
	Album                 string    `xml:"album,attr,omitempty"               json:"album,omitempty"`
	Artist                string    `xml:"artist,attr,omitempty"              json:"artist,omitempty"`
	Track                 int       `xml:"track,attr,omitempty"               json:"track,omitempty"`
	Year                  int       `xml:"year,attr,omitempty"                json:"year,omitempty"`
	Genre                 string    `xml:"genre,attr,omitempty"               json:"genre,omitempty"`
	CoverArt              string    `xml:"coverArt,attr,omitempty"            json:"coverArt,omitempty"`
	Size                  int64     `xml:"size,attr,omitempty"                json:"size,omitempty"`
	ContentType           string    `xml:"contentType,attr,omitempty"         json:"contentType,omitempty"`
	Suffix                string    `xml:"suffix,attr,omitempty"              json:"suffix,omitempty"`
	TranscodedContentType string    `xml:"transcodedContentType,attr,omitempty" json:"transcodedContentType,omitempty"`
	TranscodedSuffix      string    `xml:"transcodedSuffix,attr,omitempty"    json:"transcodedSuffix,omitempty"`
	Duration              int       `xml:"duration,attr,omitempty"            json:"duration,omitempty"`
	BitRate               int       `xml:"bitRate,attr,omitempty"             json:"bitRate,omitempty"`
	Path                  string    `xml:"path,attr,omitempty"                json:"path,omitempty"`
	IsVideo               bool      `xml:"isVideo,attr,omitempty"             json:"isVideo,omitempty"`
	UserRating            int       `xml:"userRating,attr,omitempty"          json:"userRating,omitempty"`
	AverageRating         float64   `xml:"averageRating,attr,omitempty"       json:"averageRating,omitempty"`
	PlayCount             int64     `xml:"playCount,attr,omitempty"           json:"playCount,omitempty"`
	DiscNumber            int       `xml:"discNumber,attr,omitempty"          json:"discNumber,omitempty"`
	Created               time.Time `xml:"created,attr,omitempty"             json:"created,omitempty"`
	Starred               time.Time `xml:"starred,attr,omitempty"             json:"starred,omitempty"`
	AlbumID               string    `xml:"albumId,attr,omitempty"             json:"albumId,omitempty"`
	ArtistID              string    `xml:"artistId,attr,omitempty"            json:"artistId,omitempty"`
	Type                  string    `xml:"type,attr,omitempty"                json:"type,omitempty"`
	BookmarkPosition      int64     `xml:"bookmarkPosition,attr,omitempty"    json:"bookmarkPosition,omitempty"`
	OriginalWidth         int       `xml:"originalWidth,attr,omitempty"       json:"originalWidth,omitempty"`
	OriginalHeight        int       `xml:"originalHeight,attr,omitempty"      json:"originalHeight,omitempty"`
	// OpenSubsonic playbackReport extension fields
	State        string  `xml:"state,attr,omitempty"        json:"state,omitempty"`
	PositionMs   int64   `xml:"positionMs,attr,omitempty"   json:"positionMs,omitempty"`
	PlaybackRate float64 `xml:"playbackRate,attr,omitempty" json:"playbackRate,omitempty"`
}

func (t *NowPlayingEntry) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T NowPlayingEntry
	var layout struct {
		*T
		Created *xsdDateTime `xml:"created,attr,omitempty"`
		Starred *xsdDateTime `xml:"starred,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.Created = (*xsdDateTime)(&layout.T.Created)
	layout.Starred = (*xsdDateTime)(&layout.T.Starred)
	return e.EncodeElement(layout, start)
}
func (t *NowPlayingEntry) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T NowPlayingEntry
	var overlay struct {
		*T
		Created *xsdDateTime `xml:"created,attr,omitempty"`
		Starred *xsdDateTime `xml:"starred,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Created = (*xsdDateTime)(&overlay.T.Created)
	overlay.Starred = (*xsdDateTime)(&overlay.T.Starred)
	return d.DecodeElement(&overlay, &start)
}
func (t *NowPlayingEntry) MarshalJSON() ([]byte, error) {
	type T NowPlayingEntry
	var layout struct {
		*T
		Created *xsdDateTime `json:"created,omitempty"`
		Starred *xsdDateTime `json:"starred,omitempty"`
	}
	layout.T = (*T)(t)
	if !t.Created.IsZero() {
		layout.Created = (*xsdDateTime)(&t.Created)
	}
	if !t.Starred.IsZero() {
		layout.Starred = (*xsdDateTime)(&t.Starred)
	}
	return json.Marshal(layout)
}
func (t *NowPlayingEntry) UnmarshalJSON(data []byte) error {
	type T NowPlayingEntry
	var overlay struct {
		*T
		Created *xsdDateTime `json:"created,omitempty"`
		Starred *xsdDateTime `json:"starred,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Created = (*xsdDateTime)(&overlay.T.Created)
	overlay.Starred = (*xsdDateTime)(&overlay.T.Starred)
	return json.Unmarshal(data, &overlay)
}

type PlayQueue struct {
	Entries   []*Child  `xml:"entry,omitempty"         json:"entry,omitempty"`
	Current   string    `xml:"current,attr,omitempty"  json:"current,omitempty"`
	Position  int64     `xml:"position,attr,omitempty" json:"position,omitempty"`
	Username  string    `xml:"username,attr"           json:"username"`
	Changed   time.Time `xml:"changed,attr"            json:"changed"`
	ChangedBy string    `xml:"changedBy,attr"          json:"changedBy"`
}

func (t *PlayQueue) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PlayQueue
	var layout struct {
		*T
		Changed *xsdDateTime `xml:"changed,attr"`
	}
	layout.T = (*T)(t)
	layout.Changed = (*xsdDateTime)(&layout.T.Changed)
	return e.EncodeElement(layout, start)
}
func (t *PlayQueue) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PlayQueue
	var overlay struct {
		*T
		Changed *xsdDateTime `xml:"changed,attr"`
	}
	overlay.T = (*T)(t)
	overlay.Changed = (*xsdDateTime)(&overlay.T.Changed)
	return d.DecodeElement(&overlay, &start)
}
func (t *PlayQueue) MarshalJSON() ([]byte, error) {
	type T PlayQueue
	var layout struct {
		*T
		Changed *xsdDateTime `json:"changed"`
	}
	layout.T = (*T)(t)
	layout.Changed = (*xsdDateTime)(&t.Changed)
	return json.Marshal(layout)
}
func (t *PlayQueue) UnmarshalJSON(data []byte) error {
	type T PlayQueue
	var overlay struct {
		*T
		Changed *xsdDateTime `json:"changed"`
	}
	overlay.T = (*T)(t)
	overlay.Changed = (*xsdDateTime)(&overlay.T.Changed)
	return json.Unmarshal(data, &overlay)
}

// Playlist is a collection of songs with metadata like a name, comment, and information about the total duration of the playlist.
type Playlist struct {
	ID          string    `xml:"id,attr"              json:"id"`
	Entry       []*Child  `xml:"entry,omitempty"      json:"entry,omitempty"`
	AllowedUser []string  `xml:"allowedUser,omitempty" json:"allowedUser,omitempty"`
	Name        string    `xml:"name,attr"            json:"name"`
	Comment     string    `xml:"comment,attr,omitempty" json:"comment,omitempty"`
	Owner       string    `xml:"owner,attr,omitempty"  json:"owner,omitempty"`
	Public      bool      `xml:"public,attr,omitempty" json:"public,omitempty"`
	SongCount   int       `xml:"songCount,attr"       json:"songCount"`
	Duration    int       `xml:"duration,attr"        json:"duration"`
	Created     time.Time `xml:"created,attr"         json:"created"`
	Changed     time.Time `xml:"changed,attr,omitempty" json:"changed,omitempty"`
	CoverArt    string    `xml:"coverArt,attr,omitempty" json:"coverArt,omitempty"`
}

func (t *Playlist) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T Playlist
	var layout struct {
		*T
		Created *xsdDateTime `xml:"created,attr"`
		Changed *xsdDateTime `xml:"changed,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.Created = (*xsdDateTime)(&layout.T.Created)
	layout.Changed = (*xsdDateTime)(&layout.T.Changed)
	return e.EncodeElement(layout, start)
}
func (t *Playlist) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T Playlist
	var overlay struct {
		*T
		Created *xsdDateTime `xml:"created,attr"`
		Changed *xsdDateTime `xml:"changed,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Created = (*xsdDateTime)(&overlay.T.Created)
	overlay.Changed = (*xsdDateTime)(&overlay.T.Changed)
	return d.DecodeElement(&overlay, &start)
}
func (t *Playlist) MarshalJSON() ([]byte, error) {
	type T Playlist
	var layout struct {
		*T
		Created *xsdDateTime `json:"created"`
		Changed *xsdDateTime `json:"changed,omitempty"`
	}
	layout.T = (*T)(t)
	layout.Created = (*xsdDateTime)(&t.Created)
	if !t.Changed.IsZero() {
		layout.Changed = (*xsdDateTime)(&t.Changed)
	}
	return json.Marshal(layout)
}
func (t *Playlist) UnmarshalJSON(data []byte) error {
	type T Playlist
	var overlay struct {
		*T
		Created *xsdDateTime `json:"created"`
		Changed *xsdDateTime `json:"changed,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Created = (*xsdDateTime)(&overlay.T.Created)
	overlay.Changed = (*xsdDateTime)(&overlay.T.Changed)
	return json.Unmarshal(data, &overlay)
}

type playlists struct {
	Playlist []*Playlist `xml:"playlist,omitempty" json:"playlist,omitempty"`
}

type PodcastChannel struct {
	Episode          []*PodcastEpisode `xml:"episode,omitempty"              json:"episode,omitempty"`
	Url              string            `xml:"url,attr"                       json:"url"`
	Title            string            `xml:"title,attr,omitempty"           json:"title,omitempty"`
	Description      string            `xml:"description,attr,omitempty"     json:"description,omitempty"`
	CoverArt         string            `xml:"coverArt,attr,omitempty"        json:"coverArt,omitempty"`
	OriginalImageUrl string            `xml:"originalImageUrl,attr,omitempty" json:"originalImageUrl,omitempty"`
	Status           string            `xml:"status,attr"                    json:"status"`
	ErrorMessage     string            `xml:"errorMessage,attr,omitempty"    json:"errorMessage,omitempty"`
}

type PodcastEpisode struct {
	StreamID              string    `xml:"streamId,attr,omitempty"            json:"streamId,omitempty"`
	ChannelID             string    `xml:"channelId,attr"                     json:"channelId"`
	Description           string    `xml:"description,attr,omitempty"         json:"description,omitempty"`
	Status                string    `xml:"status,attr"                        json:"status"`
	PublishDate           time.Time `xml:"publishDate,attr,omitempty"         json:"publishDate,omitempty"`
	Parent                string    `xml:"parent,attr,omitempty"              json:"parent,omitempty"`
	IsDir                 bool      `xml:"isDir,attr"                         json:"isDir"`
	Title                 string    `xml:"title,attr"                         json:"title"`
	Album                 string    `xml:"album,attr,omitempty"               json:"album,omitempty"`
	Artist                string    `xml:"artist,attr,omitempty"              json:"artist,omitempty"`
	Track                 int       `xml:"track,attr,omitempty"               json:"track,omitempty"`
	Year                  int       `xml:"year,attr,omitempty"                json:"year,omitempty"`
	Genre                 string    `xml:"genre,attr,omitempty"               json:"genre,omitempty"`
	CoverArt              string    `xml:"coverArt,attr,omitempty"            json:"coverArt,omitempty"`
	Size                  int64     `xml:"size,attr,omitempty"                json:"size,omitempty"`
	ContentType           string    `xml:"contentType,attr,omitempty"         json:"contentType,omitempty"`
	Suffix                string    `xml:"suffix,attr,omitempty"              json:"suffix,omitempty"`
	TranscodedContentType string    `xml:"transcodedContentType,attr,omitempty" json:"transcodedContentType,omitempty"`
	TranscodedSuffix      string    `xml:"transcodedSuffix,attr,omitempty"    json:"transcodedSuffix,omitempty"`
	Duration              int       `xml:"duration,attr,omitempty"            json:"duration,omitempty"`
	BitRate               int       `xml:"bitRate,attr,omitempty"             json:"bitRate,omitempty"`
	Path                  string    `xml:"path,attr,omitempty"                json:"path,omitempty"`
	IsVideo               bool      `xml:"isVideo,attr,omitempty"             json:"isVideo,omitempty"`
	UserRating            int       `xml:"userRating,attr,omitempty"          json:"userRating,omitempty"`
	AverageRating         float64   `xml:"averageRating,attr,omitempty"       json:"averageRating,omitempty"`
	PlayCount             int64     `xml:"playCount,attr,omitempty"           json:"playCount,omitempty"`
	DiscNumber            int       `xml:"discNumber,attr,omitempty"          json:"discNumber,omitempty"`
	Created               time.Time `xml:"created,attr,omitempty"             json:"created,omitempty"`
	Starred               time.Time `xml:"starred,attr,omitempty"             json:"starred,omitempty"`
	AlbumID               string    `xml:"albumId,attr,omitempty"             json:"albumId,omitempty"`
	ArtistID              string    `xml:"artistId,attr,omitempty"            json:"artistId,omitempty"`
	Type                  string    `xml:"type,attr,omitempty"                json:"type,omitempty"`
	BookmarkPosition      int64     `xml:"bookmarkPosition,attr,omitempty"    json:"bookmarkPosition,omitempty"`
	OriginalWidth         int       `xml:"originalWidth,attr,omitempty"       json:"originalWidth,omitempty"`
	OriginalHeight        int       `xml:"originalHeight,attr,omitempty"      json:"originalHeight,omitempty"`
}

func (t *PodcastEpisode) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PodcastEpisode
	var layout struct {
		*T
		PublishDate *xsdDateTime `xml:"publishDate,attr,omitempty"`
		Created     *xsdDateTime `xml:"created,attr,omitempty"`
		Starred     *xsdDateTime `xml:"starred,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.PublishDate = (*xsdDateTime)(&layout.T.PublishDate)
	layout.Created = (*xsdDateTime)(&layout.T.Created)
	layout.Starred = (*xsdDateTime)(&layout.T.Starred)
	return e.EncodeElement(layout, start)
}
func (t *PodcastEpisode) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PodcastEpisode
	var overlay struct {
		*T
		PublishDate *xsdDateTime `xml:"publishDate,attr,omitempty"`
		Created     *xsdDateTime `xml:"created,attr,omitempty"`
		Starred     *xsdDateTime `xml:"starred,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.PublishDate = (*xsdDateTime)(&overlay.T.PublishDate)
	overlay.Created = (*xsdDateTime)(&overlay.T.Created)
	overlay.Starred = (*xsdDateTime)(&overlay.T.Starred)
	return d.DecodeElement(&overlay, &start)
}
func (t *PodcastEpisode) MarshalJSON() ([]byte, error) {
	type T PodcastEpisode
	var layout struct {
		*T
		PublishDate *xsdDateTime `json:"publishDate,omitempty"`
		Created     *xsdDateTime `json:"created,omitempty"`
		Starred     *xsdDateTime `json:"starred,omitempty"`
	}
	layout.T = (*T)(t)
	if !t.PublishDate.IsZero() {
		layout.PublishDate = (*xsdDateTime)(&t.PublishDate)
	}
	if !t.Created.IsZero() {
		layout.Created = (*xsdDateTime)(&t.Created)
	}
	if !t.Starred.IsZero() {
		layout.Starred = (*xsdDateTime)(&t.Starred)
	}
	return json.Marshal(layout)
}
func (t *PodcastEpisode) UnmarshalJSON(data []byte) error {
	type T PodcastEpisode
	var overlay struct {
		*T
		PublishDate *xsdDateTime `json:"publishDate,omitempty"`
		Created     *xsdDateTime `json:"created,omitempty"`
		Starred     *xsdDateTime `json:"starred,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.PublishDate = (*xsdDateTime)(&overlay.T.PublishDate)
	overlay.Created = (*xsdDateTime)(&overlay.T.Created)
	overlay.Starred = (*xsdDateTime)(&overlay.T.Starred)
	return json.Unmarshal(data, &overlay)
}

type podcasts struct {
	Channel []*PodcastChannel `xml:"channel,omitempty" json:"channel,omitempty"`
}

// Response is the main target for unmarshalling data from the API - everything within the "subsonic-response" key
type Response struct {
	License                *License                 `xml:"license"               json:"license,omitempty"`
	MusicFolders           *musicFolders            `xml:"musicFolders"          json:"musicFolders,omitempty"`
	Indexes                *Indexes                 `xml:"indexes"               json:"indexes,omitempty"`
	Directory              *Directory               `xml:"directory"             json:"directory,omitempty"`
	Genres                 *genres                  `xml:"genres"                json:"genres,omitempty"`
	Artists                *ArtistsID3              `xml:"artists"               json:"artists,omitempty"`
	Artist                 *ArtistID3               `xml:"artist"                json:"artist,omitempty"`
	Album                  *AlbumID3                `xml:"album"                 json:"album,omitempty"`
	Song                   *Child                   `xml:"song"                  json:"song,omitempty"`
	NowPlaying             *nowPlaying              `xml:"nowPlaying"            json:"nowPlaying,omitempty"`
	SearchResult2          *SearchResult2           `xml:"searchResult2"         json:"searchResult2,omitempty"`
	SearchResult3          *SearchResult3           `xml:"searchResult3"         json:"searchResult3,omitempty"`
	Playlists              *playlists               `xml:"playlists"             json:"playlists,omitempty"`
	Playlist               *Playlist                `xml:"playlist"              json:"playlist,omitempty"`
	JukeboxStatus          *JukeboxStatus           `xml:"jukeboxStatus"         json:"jukeboxStatus,omitempty"`
	JukeboxPlaylist        *JukeboxPlaylist         `xml:"jukeboxPlaylist"       json:"jukeboxPlaylist,omitempty"`
	Users                  *users                   `xml:"users"                 json:"users,omitempty"`
	User                   *User                    `xml:"user"                  json:"user,omitempty"`
	ChatMessages           *chatMessages            `xml:"chatMessages"          json:"chatMessages,omitempty"`
	AlbumList              *albumList               `xml:"albumList"             json:"albumList,omitempty"`
	AlbumList2             *albumList2              `xml:"albumList2"            json:"albumList2,omitempty"`
	RandomSongs            *songs                   `xml:"randomSongs"           json:"randomSongs,omitempty"`
	SongsByGenre           *songs                   `xml:"songsByGenre"          json:"songsByGenre,omitempty"`
	Lyrics                 *Lyrics                  `xml:"lyrics"                json:"lyrics,omitempty"`
	Podcasts               *podcasts                `xml:"podcasts"              json:"podcasts,omitempty"`
	NewestPodcasts         *newestPodcasts          `xml:"newestPodcasts"        json:"newestPodcasts,omitempty"`
	InternetRadioStations  *internetRadioStations   `xml:"internetRadioStations" json:"internetRadioStations,omitempty"`
	Bookmarks              *bookmarks               `xml:"bookmarks"             json:"bookmarks,omitempty"`
	PlayQueue              *PlayQueue               `xml:"playQueue"             json:"playQueue,omitempty"`
	PlayQueueByIndex       *PlayQueueByIndex        `xml:"playQueueByIndex"      json:"playQueueByIndex,omitempty"`
	Shares                 *shares                  `xml:"shares"                json:"shares,omitempty"`
	Starred                *Starred                 `xml:"starred"               json:"starred,omitempty"`
	Starred2               *Starred2                `xml:"starred2"              json:"starred2,omitempty"`
	AlbumInfo              *AlbumInfo               `xml:"albumInfo"             json:"albumInfo,omitempty"`
	ArtistInfo             *ArtistInfo              `xml:"artistInfo"            json:"artistInfo,omitempty"`
	ArtistInfo2            *ArtistInfo2             `xml:"artistInfo2"           json:"artistInfo2,omitempty"`
	SimilarSongs           *similarSongs            `xml:"similarSongs"          json:"similarSongs,omitempty"`
	SimilarSongs2          *similarSongs2           `xml:"similarSongs2"         json:"similarSongs2,omitempty"`
	TopSongs               *topSongs                `xml:"topSongs"              json:"topSongs,omitempty"`
	ScanStatus             *ScanStatus              `xml:"scanStatus"            json:"scanStatus,omitempty"`
	Error                  *Error                   `xml:"error"                 json:"error,omitempty"`
	Status                 string                   `xml:"status,attr"           json:"status"`
	Version                string                   `xml:"version,attr"          json:"version"`
	Type                   string                   `xml:"type,attr"             json:"type"`
	ServerVersion          string                   `xml:"serverVersion,attr"    json:"serverVersion"`
	OpenSubsonic           bool                     `xml:"openSubsonic,attr"     json:"openSubsonic"`
	OpenSubsonicExtensions []*OpenSubsonicExtension `xml:"openSubsonicExtensions" json:"openSubsonicExtensions,omitempty"`
	LyricsList             *LyricsList              `xml:"lyricsList"            json:"lyricsList,omitempty"`
}

type ScanStatus struct {
	Scanning bool  `xml:"scanning,attr"       json:"scanning"`
	Count    int64 `xml:"count,attr,omitempty" json:"count,omitempty"`
}

// SearchResult2 is a collection of songs, albums, and artists related to a query.
type SearchResult2 struct {
	Artist []*Artist `xml:"artist,omitempty" json:"artist,omitempty"`
	Album  []*Child  `xml:"album,omitempty"  json:"album,omitempty"`
	Song   []*Child  `xml:"song,omitempty"   json:"song,omitempty"`
}

// SearchResult3 is a collection of songs, albums, and artists related to a query.
type SearchResult3 struct {
	Artist []*ArtistID3 `xml:"artist,omitempty" json:"artist,omitempty"`
	Album  []*AlbumID3  `xml:"album,omitempty"  json:"album,omitempty"`
	Song   []*Child     `xml:"song,omitempty"   json:"song,omitempty"`
}

type Share struct {
	ID          string    `xml:"id,attr"                    json:"id"`
	Entry       []*Child  `xml:"entry,omitempty"            json:"entry,omitempty"`
	Url         string    `xml:"url,attr"                   json:"url"`
	Description string    `xml:"description,attr,omitempty" json:"description,omitempty"`
	Username    string    `xml:"username,attr"              json:"username"`
	Created     time.Time `xml:"created,attr"               json:"created"`
	Expires     time.Time `xml:"expires,attr,omitempty"     json:"expires,omitempty"`
	LastVisited time.Time `xml:"lastVisited,attr,omitempty" json:"lastVisited,omitempty"`
	VisitCount  int       `xml:"visitCount,attr"            json:"visitCount"`
}

func (t *Share) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T Share
	var layout struct {
		*T
		Created     *xsdDateTime `xml:"created,attr"`
		Expires     *xsdDateTime `xml:"expires,attr,omitempty"`
		LastVisited *xsdDateTime `xml:"lastVisited,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.Created = (*xsdDateTime)(&layout.T.Created)
	layout.Expires = (*xsdDateTime)(&layout.T.Expires)
	layout.LastVisited = (*xsdDateTime)(&layout.T.LastVisited)
	return e.EncodeElement(layout, start)
}
func (t *Share) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T Share
	var overlay struct {
		*T
		Created     *xsdDateTime `xml:"created,attr"`
		Expires     *xsdDateTime `xml:"expires,attr,omitempty"`
		LastVisited *xsdDateTime `xml:"lastVisited,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Created = (*xsdDateTime)(&overlay.T.Created)
	overlay.Expires = (*xsdDateTime)(&overlay.T.Expires)
	overlay.LastVisited = (*xsdDateTime)(&overlay.T.LastVisited)
	return d.DecodeElement(&overlay, &start)
}
func (t *Share) MarshalJSON() ([]byte, error) {
	type T Share
	var layout struct {
		*T
		Created     *xsdDateTime `json:"created"`
		Expires     *xsdDateTime `json:"expires,omitempty"`
		LastVisited *xsdDateTime `json:"lastVisited,omitempty"`
	}
	layout.T = (*T)(t)
	layout.Created = (*xsdDateTime)(&t.Created)
	if !t.Expires.IsZero() {
		layout.Expires = (*xsdDateTime)(&t.Expires)
	}
	if !t.LastVisited.IsZero() {
		layout.LastVisited = (*xsdDateTime)(&t.LastVisited)
	}
	return json.Marshal(layout)
}
func (t *Share) UnmarshalJSON(data []byte) error {
	type T Share
	var overlay struct {
		*T
		Created     *xsdDateTime `json:"created"`
		Expires     *xsdDateTime `json:"expires,omitempty"`
		LastVisited *xsdDateTime `json:"lastVisited,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Created = (*xsdDateTime)(&overlay.T.Created)
	overlay.Expires = (*xsdDateTime)(&overlay.T.Expires)
	overlay.LastVisited = (*xsdDateTime)(&overlay.T.LastVisited)
	return json.Unmarshal(data, &overlay)
}

type shares struct {
	Share []*Share `xml:"share,omitempty" json:"share,omitempty"`
}

type similarSongs struct {
	Song []*Child `xml:"song,omitempty" json:"song,omitempty"`
}

type similarSongs2 struct {
	Song []*Child `xml:"song,omitempty" json:"song,omitempty"`
}

type songs struct {
	Song []*Child `xml:"song,omitempty" json:"song,omitempty"`
}

// Starred is a collection of songs, albums, and artists annotated by a user as starred.
type Starred struct {
	Artist []*Artist `xml:"artist,omitempty" json:"artist,omitempty"`
	Album  []*Child  `xml:"album,omitempty"  json:"album,omitempty"`
	Song   []*Child  `xml:"song,omitempty"   json:"song,omitempty"`
}

// Starred2 is a collection of songs, albums, and artists organized by ID3 tags annotated by a user as starred.
type Starred2 struct {
	Artist []*ArtistID3 `xml:"artist,omitempty" json:"artist,omitempty"`
	Album  []*AlbumID3  `xml:"album,omitempty"  json:"album,omitempty"`
	Song   []*Child     `xml:"song,omitempty"   json:"song,omitempty"`
}

type topSongs struct {
	Song []*Child `xml:"song,omitempty" json:"song,omitempty"`
}

type User struct {
	Folder              []int     `xml:"folder,omitempty"              json:"folder,omitempty"`
	Username            string    `xml:"username,attr"                 json:"username"`
	Email               string    `xml:"email,attr,omitempty"          json:"email,omitempty"`
	ScrobblingEnabled   bool      `xml:"scrobblingEnabled,attr"        json:"scrobblingEnabled"`
	MaxBitRate          int       `xml:"maxBitRate,attr,omitempty"     json:"maxBitRate,omitempty"`
	AdminRole           bool      `xml:"adminRole,attr"                json:"adminRole"`
	SettingsRole        bool      `xml:"settingsRole,attr"             json:"settingsRole"`
	DownloadRole        bool      `xml:"downloadRole,attr"             json:"downloadRole"`
	UploadRole          bool      `xml:"uploadRole,attr"               json:"uploadRole"`
	PlaylistRole        bool      `xml:"playlistRole,attr"             json:"playlistRole"`
	CoverArtRole        bool      `xml:"coverArtRole,attr"             json:"coverArtRole"`
	CommentRole         bool      `xml:"commentRole,attr"              json:"commentRole"`
	PodcastRole         bool      `xml:"podcastRole,attr"              json:"podcastRole"`
	StreamRole          bool      `xml:"streamRole,attr"               json:"streamRole"`
	JukeboxRole         bool      `xml:"jukeboxRole,attr"              json:"jukeboxRole"`
	ShareRole           bool      `xml:"shareRole,attr"                json:"shareRole"`
	VideoConversionRole bool      `xml:"videoConversionRole,attr"      json:"videoConversionRole"`
	AvatarLastChanged   time.Time `xml:"avatarLastChanged,attr,omitempty" json:"avatarLastChanged,omitempty"`
}

func (t *User) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T User
	var layout struct {
		*T
		AvatarLastChanged *xsdDateTime `xml:"avatarLastChanged,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.AvatarLastChanged = (*xsdDateTime)(&layout.T.AvatarLastChanged)
	return e.EncodeElement(layout, start)
}
func (t *User) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T User
	var overlay struct {
		*T
		AvatarLastChanged *xsdDateTime `xml:"avatarLastChanged,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.AvatarLastChanged = (*xsdDateTime)(&overlay.T.AvatarLastChanged)
	return d.DecodeElement(&overlay, &start)
}
func (t *User) MarshalJSON() ([]byte, error) {
	type T User
	var layout struct {
		*T
		AvatarLastChanged *xsdDateTime `json:"avatarLastChanged,omitempty"`
	}
	layout.T = (*T)(t)
	if !t.AvatarLastChanged.IsZero() {
		layout.AvatarLastChanged = (*xsdDateTime)(&t.AvatarLastChanged)
	}
	return json.Marshal(layout)
}
func (t *User) UnmarshalJSON(data []byte) error {
	type T User
	var overlay struct {
		*T
		AvatarLastChanged *xsdDateTime `json:"avatarLastChanged,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.AvatarLastChanged = (*xsdDateTime)(&overlay.T.AvatarLastChanged)
	return json.Unmarshal(data, &overlay)
}

type users struct {
	User []*User `xml:"user,omitempty" json:"user,omitempty"`
}

type xsdDateTime time.Time

func (t *xsdDateTime) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006-01-02T15:04:05.999999999")
}
func (t xsdDateTime) MarshalText() ([]byte, error) {
	return []byte((time.Time)(t).Format("2006-01-02T15:04:05.999999999")), nil
}
func (t xsdDateTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdDateTime) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}
func _unmarshalTime(text []byte, t *time.Time, format string) (err error) {
	s := string(bytes.TrimSpace(text))
	if s == "" {
		*t = time.Time{}
		return nil
	}
	*t, err = time.Parse(format, s)
	if _, ok := err.(*time.ParseError); ok {
		*t, err = time.Parse(format+"Z07:00", s)
	}
	return err
}
