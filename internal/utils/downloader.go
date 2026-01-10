package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const gitignoreURL = "https://raw.githubusercontent.com/github/gitignore/main"

type Downloader struct {
	URL string
	Language string

	client *http.Client
	buffer bytes.Buffer
}

// Sets some logical defaults
func NewDownloader() *Downloader {
	return &Downloader{
		URL:      gitignoreURL,
		client:   &http.Client{},
		Language: "go",
		buffer:   bytes.Buffer{},
	}
}

func (d *Downloader) WithClient(client *http.Client) *Downloader {
	d.client = client
	return d
}

func (d *Downloader) Printf(format string, a ...interface{}) {
	d.buffer.WriteString(fmt.Sprintf(format, a...))
}

func (d *Downloader) Download() ([]byte, error) {
	

	url := fmt.Sprintf("%s/%s.gitignore", d.URL, d.CaseLanguage())
	d.Printf("ℹ️ Getting gitignore for %s from %s\n", d.Language, url)
	req, err := http.NewRequestWithContext(context.Background(), "GET", url, nil)
	if err != nil {
		d.Printf("⛔️ Error creating request: %v\n", err)
		return nil, err
	}
	resp, err := d.client.Do(req)
	
	if err != nil {
		d.Printf("⛔️ Error making request: %v\n", err)
		return nil, err
	} 

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		d.Printf("⛔️ Error: received status code %d\n", resp.StatusCode)
		return nil, fmt.Errorf("received status code %d", resp.StatusCode)
	}

	// read the response body
	body, err := io.ReadAll(resp.Body)
	
	if err != nil {
		d.Printf("⛔️ Error reading response body: %v\n", err)
		return nil, err
	}

	if len(body) == 0 {
		d.Printf("⛔️ Didn't get an error for %s. But the file was empty\n", d.Language)
		return nil, fmt.Errorf("gitignore emtpty for %s", d.Language)
	}
	
	return body, nil
}

func (d *Downloader) GetLog() string {
	return d.buffer.String()
}

func (d *Downloader) CaseLanguage() string {
	// Title case the language
	return cases.Title(language.English).String(d.Language)
}

func (d *Downloader) ListLanguages() ([]string, error) {
	url := `https://api.github.com/repos/github/gitignore/git/trees/main?recursive=1`
	d.Printf("ℹ️ Listing languages from %s\n", url)
	req, err := http.NewRequestWithContext(context.Background(), "GET", url, nil)
	if err != nil {
		d.Printf("⛔️ Error creating request: %v\n", err)
		return nil, err
	}
	resp, err := d.client.Do(req)

	if err != nil {
		d.Printf("⛔️ Error making request: %v\n", err)
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		d.Printf("⛔️ Error: received status code %d\n", resp.StatusCode)
		return nil, fmt.Errorf("received status code %d", resp.StatusCode)
	}

	// read the response body
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		d.Printf("⛔️ Error reading response body: %v\n", err)
		return nil, err
	}

	// Parse the JSON response
	var result struct {
		Tree []struct {
			Path string `json:"path"`
		} `json:"tree"`
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		d.Printf("⛔️ Error parsing JSON response: %v\n", err)
		return nil, err
	}

	var languages []string
	for _, item := range result.Tree {
		if strings.HasSuffix(item.Path, ".gitignore") {
			lang := strings.TrimSuffix(item.Path, ".gitignore")
			languages = append(languages, lang)
		}
	}
	return languages, nil

}
