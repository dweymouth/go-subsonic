package subsonic

type OpenSubsonicExtension struct {
	Name     string `xml:"name,attr"`
	Versions []int  `xml:"versions"`
}

type openSubsonicExtensions struct {
	OpenSubsonicExtensions []*OpenSubsonicExtension `xml:"openSubsonicExtension,omitempty"`
}

type LyricsList struct {
	StructuredLyrics []*StructuredLyrics `xml:"structuredLyrics,omitempty"`
}

type StructuredLyrics struct {
	DisplayArtist string      `xml:"displayArtist,attr"`
	DisplayTitle  string      `xml:"displayTitle,attr"`
	Lang          string      `xml:"lang,attr"`
	Offset        int         `xml:"offset,attr"`
	Synced        bool        `xml:"synced,attr"`
	Lines         []LyricLine `xml:"line,omitempty"`
}

type LyricLine struct {
	Start int    `xml:"start,attr"`
	Text  string `xml:",chardata"`
	// Navidrome 0.51.0 - 0.52.5 incorrecty returns the lyric line text here
	// This will be removed in the future
	Value string `xml:"value"`
}

type ItemDate struct {
	Year  *int `xml:"year,attr,omitempty"`
	Month *int `xml:"month,attr,omitempty"`
	Date  *int `xml:"date,attr,omitempty"`
}
