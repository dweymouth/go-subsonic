// Package subsonic implements an API client library for Subsonic-compatible music streaming servers.
//
// This project handles communication with a remote *sonic server, but does not handle playback of media. The library user should be prepared to do something with the stream of audio in bytes, like decoding and playing that audio on a sound card.
// The list of API endpoints implemented is available on the project's github page.
//
// The API is divided between functions with no suffix, and functions that have a "2" suffix (or "3" in the case of Search3).
// Generally, things with "2" on the end are organized by file tags rather than folder structure. This is how you'd expect most music players to work and is recommended.
// The variants without a suffix organize the library by directory structure; artists are a directory, albums are children of that directory, songs (subsonic.Child) are children of albums.
// This has some disadvantages: possibly duplicating items with identical directory names, treating songs and albums in much the same fashion, and being more difficult to query consistently.
package subsonic

import (
	"crypto/md5"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"path"
	"strings"
)

const (
	defaultAPIVersion = "1.15.0"
	libraryVersion    = "0.0.5"
)

var (
	ErrAuthenticationFailure = errors.New("authentication failure")
)

type Client struct {
	Client              *http.Client
	BaseUrl             string
	User                string
	ClientName          string
	UserAgent           string
	PasswordAuth        bool
	UseJSON             bool
	RequestedAPIVersion string

	openSubsonicExtensions []*OpenSubsonicExtension

	password string
	salt     string
	token    string
}

func generateSalt() string {
	var corpus = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	// length is minimum 6, but let's use ten to start
	b := make([]rune, 10)
	for i := range b {
		b[i] = corpus[rand.Intn(len(corpus))]
	}
	return string(b)
}

// Authenticate authenticates the current user with a provided password.
// If s.PasswordAuth is false, the password is salted before transmission and requires Subsonic > 1.13.0.
// Returns ErrAuthenticationFailure if the user/pass combo is incorrect,
// or another error type for any other failure reason.
func (s *Client) Authenticate(password string) error {
	if s.PasswordAuth {
		s.password = password
	} else {
		salt := generateSalt()
		h := md5.New()
		_, err := io.WriteString(h, password)
		if err != nil {
			return err
		}
		_, err = io.WriteString(h, salt)
		if err != nil {
			return err
		}
		s.salt = salt
		s.token = fmt.Sprintf("%x", h.Sum(nil))
	}

	// Test authentication
	// Don't use the s.Ping method because that always returns true as long as the servers is up.
	resp, err := s.Request("GET", "ping", nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	subsonicResp, err := s.unmarshalResponse(resp.Body)
	if err != nil {
		return err
	}

	if subsonicResp.Error != nil {
		return ErrAuthenticationFailure
	}

	return nil
}

// Request performs a HTTP request against the Subsonic server as the current user.
// If a nil error is returned, the caller is responsible for closing the response body.
func (s *Client) Request(method string, endpoint string, params url.Values) (*http.Response, error) {
	req, err := s.setupRequest(method, endpoint, params)
	if err != nil {
		return nil, err
	}

	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Get is a convenience interface to issue a GET request and parse the response body (99% of Subsonic API calls)
func (s *Client) Get(endpoint string, params map[string]string) (*Response, error) {
	parameters := url.Values{}
	for k, v := range params {
		parameters.Add(k, v)
	}
	return s.getValues(endpoint, parameters)
}

func (s *Client) setupRequest(method string, endpoint string, params url.Values) (*http.Request, error) {
	baseUrl, err := s.buildRequestURL(endpoint)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, baseUrl.String(), nil)
	if err != nil {
		return nil, err
	}
	if u := s.UserAgent; u != "" {
		req.Header.Set("User-Agent", u)
	}

	q := req.URL.Query()
	s.addDefaultQueryParams(q)

	for key, values := range params {
		for _, val := range values {
			q.Add(key, val)
		}
	}
	req.URL.RawQuery = q.Encode()
	return req, nil
}

func (s *Client) addDefaultQueryParams(params url.Values) {
	if s.UseJSON {
		params.Add("f", "json")
	} else {
		params.Add("f", "xml")
	}
	apiVersion := defaultAPIVersion
	if s.RequestedAPIVersion != "" {
		apiVersion = s.RequestedAPIVersion
	}
	params.Add("v", apiVersion)
	params.Add("c", s.ClientName)
	params.Add("u", s.User)
	if s.PasswordAuth {
		params.Add("p", s.password)
	} else {
		params.Add("t", s.token)
		params.Add("s", s.salt)
	}
}

// getValues is a convenience interface to issue a GET request and parse the response body. It supports multiple values by way of the url.Values argument.
func (s *Client) getValues(endpoint string, params url.Values) (*Response, error) {
	response, err := s.Request(http.MethodGet, endpoint, params)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	parsed, err := s.unmarshalResponse(response.Body)
	if err != nil {
		return nil, err
	}
	if parsed.Error != nil {
		return nil, fmt.Errorf("Error #%d: %s", parsed.Error.Code, parsed.Error.Message)
	}
	//log.Printf("%s: %s\n", endpoint, string(responseBody))
	return parsed, nil
}

// postValues is the same as getValues except it sends the request
// using POST and content type application/x-www-form-urlencoded.
// The server must support the formPost extension.
func (s *Client) postValues(endpoint string, params url.Values) (*Response, error) {
	baseUrl, err := s.buildRequestURL(endpoint)
	if err != nil {
		return nil, err
	}

	s.addDefaultQueryParams(params)
	body := strings.NewReader(params.Encode())
	req, err := http.NewRequest(http.MethodPost, baseUrl.String(), body)
	if err != nil {
		return nil, err
	}
	if u := s.UserAgent; u != "" {
		req.Header.Set("User-Agent", u)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	parsed, err := s.unmarshalResponse(resp.Body)
	if err != nil {
		return nil, err
	}
	if parsed.Error != nil {
		return nil, fmt.Errorf("Error #%d: %s", parsed.Error.Code, parsed.Error.Message)
	}
	return parsed, nil
}

func (s *Client) buildRequestURL(endpoint string) (*url.URL, error) {
	baseUrl, err := url.Parse(s.BaseUrl)
	if err != nil {
		return nil, err
	}
	baseUrl.Path = path.Join(baseUrl.Path, "/rest/", endpoint)
	return baseUrl, nil
}

// unmarshalResponse parses a Subsonic API response body into a Response struct.
// It dispatches to XML or JSON parsing based on the client's UseJSON setting.
func (s *Client) unmarshalResponse(resp io.Reader) (*Response, error) {
	responseBody, err := io.ReadAll(resp)
	if err != nil {
		return nil, err
	}
	if s.UseJSON {
		return unmarshalJSONResponse(responseBody)
	}
	parsed := &Response{}
	if err = xml.Unmarshal(responseBody, parsed); err != nil {
		return nil, err
	}
	return parsed, nil
}

// unmarshalResponse parses an XML Subsonic API response. It is kept as a
// package-level function for use in tests that do not have a Client instance.
func unmarshalResponse(resp io.Reader) (*Response, error) {
	responseBody, err := io.ReadAll(resp)
	if err != nil {
		return nil, err
	}
	parsed := &Response{}
	if err = xml.Unmarshal(responseBody, parsed); err != nil {
		return nil, err
	}
	return parsed, nil
}

func unmarshalJSONResponse(data []byte) (*Response, error) {
	var wrapper struct {
		SubsonicResponse Response `json:"subsonic-response"`
	}
	if err := json.Unmarshal(data, &wrapper); err != nil {
		return nil, err
	}
	return &wrapper.SubsonicResponse, nil
}

// Ping is used to test connectivity with the server. It returns true if the server is up.
// Should generally NOT be called before authenticating as it will be considered an authentication
// by the Subsonic server. (Though this function will still return a nil Go error.)
func (s *Client) Ping() (*PingResponse, error) {
	resp, err := s.Request("GET", "ping", nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	parsed, err := s.unmarshalResponse(resp.Body)
	if err != nil {
		return nil, err
	}
	return &PingResponse{
		Status:        parsed.Status,
		Version:       parsed.Version,
		Type:          parsed.Type,
		ServerVersion: parsed.ServerVersion,
		OpenSubsonic:  parsed.OpenSubsonic,
		ServerError:   parsed.Error,
	}, nil
}

// GetLicense retrieves details about the software license. Subsonic requires a license after a 30-day trial, compatible applications have a perpetually valid license.
func (s *Client) GetLicense() (*License, error) {
	resp, err := s.Get("getLicense", nil)
	if err != nil {
		return nil, err
	}
	return resp.License, nil
}

func (s *Client) getOpenSubsonicExtensions() ([]*OpenSubsonicExtension, error) {
	if s.openSubsonicExtensions != nil {
		return s.openSubsonicExtensions, nil
	}
	ose, err := s.GetOpenSubsonicExtensions()
	s.openSubsonicExtensions = ose
	return ose, err
}
