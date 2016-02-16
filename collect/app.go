package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v2"

	"github.com/PuerkitoBio/goquery"
)

type Song struct {
	Title        string
	Artist       string
	ProviderName string
}

type Provider struct {
	Name             string `yaml:"name"`
	URL              string `yaml:"url"`
	ItemRootSelector string `yaml:"root_selector"`
	TitleSelector    string `yaml:"title_selector"`
	ArtistSelector   string `yaml:"artist_selector"`
}

func (p Provider) getSongs(results chan []Song) {
	doc, err := goquery.NewDocument(p.URL)
	if err != nil {
		log.Fatal(err)
	}

	var songs []Song
	doc.Find(p.ItemRootSelector).Each(func(i int, s *goquery.Selection) {
		artist := s.Find(p.ArtistSelector).Text()
		title := s.Find(p.TitleSelector).Text()
		songs = append(songs, Song{
			Title:        strings.TrimSpace(title),
			Artist:       strings.TrimSpace(artist),
			ProviderName: strings.TrimSpace(p.Name),
		})
	})

	results <- songs
}

func getSongsFromProviders(providers []Provider) []Song {
	numProviders := len(providers)
	providerResponses := make(chan []Song, numProviders)

	for _, provider := range providers {
		fmt.Println("Getting songs for provider: ", provider)
		go provider.getSongs(providerResponses)
	}

	var results = make([]Song, 0)
	for i := 0; i < numProviders; i++ {
		response := <-providerResponses
		results = append(results, response...)
	}

	return results
}

func writeSongsToFile(filename string, songs []Song) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	// close fo on exit and check for its returned error
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	// make a write buffer
	w := bufio.NewWriter(file)
	for _, song := range songs {
		w.WriteString(fmt.Sprintf("\"%s\",\"%s\",\"%s\"\n", song.Title, song.Artist, song.ProviderName))
	}

	w.Flush()
}

func main() {
	data, err := ioutil.ReadFile("./providers.yml")
	if err != nil {
		log.Fatal(err)
	}

	var providers []Provider
	if err := yaml.Unmarshal(data, &providers); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Providers:", providers)
	songs := getSongsFromProviders(providers)

	writeSongsToFile("songs.csv", songs)
	fmt.Printf("%+v\n", songs)
}
