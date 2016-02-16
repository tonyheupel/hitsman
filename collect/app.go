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
	log.Println(fmt.Sprintf("Getting songs from %s...", p.Name))
	doc, err := goquery.NewDocument(p.URL)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(fmt.Sprintf("Processing songs from %s...", p.Name))

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

	log.Println(fmt.Sprintf("Done processing songs from %s...", p.Name))
	results <- songs
}

func getSongsFromProviders(providers []Provider) []Song {
	numProviders := len(providers)
	providerResponses := make(chan []Song, numProviders)

	for _, provider := range providers {
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
	log.Println("Getting list of providers from providers.yml...")
	data, err := ioutil.ReadFile("./providers.yml")
	if err != nil {
		log.Fatal(err)
	}

	var providers []Provider
	if err := yaml.Unmarshal(data, &providers); err != nil {
		log.Fatal(err)
	}

	log.Println(fmt.Sprintf("Found %d providers...", len(providers)))
	songs := getSongsFromProviders(providers)

	filename := "songs.csv" // TODO: Make this configurable

	log.Println(fmt.Sprintf("Writing %d songs to %s...", len(songs), filename))
	writeSongsToFile("songs.csv", songs)
	log.Println("Done!")
}
