/**
 * @license
 * Copyright Google Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package wordz

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/gookit/color"
)

// Pronunciations is JSON the response
type Pronunciations struct {
	AudioFile        string   `json:"audioFile"`
	Dialects         []string `json:"dialects"`
	PhoneticNotation string   `json:"phoneticNotation"`
	PhoneticSpelling string   `json:"phoneticSpelling"`
}

// LexicalCategory is JSON the response
type LexicalCategory struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

// ThesaurusLinks is JSON the response
type ThesaurusLinks struct {
	EntryID string `json:"entry_id"`
	SenseID string `json:"sense_id"`
}

// Examples is JSON the response
type Examples struct {
	Text string `json:"text"`
}

// Senses is JSON the response
type Senses struct {
	ID               string           `json:"id"`
	Definitions      []string         `json:"definitions"`
	Examples         []Examples       `json:"examples"`
	ShortDefinitions []string         `json:"shortDefinitions"`
	ThesaurusLinks   []ThesaurusLinks `json:"thesaurusLinks"`
}

// Entries is JSON the response
type Entries struct {
	Etymologies []string `json:"etymologies"`
	Senses      []Senses `json:"senses"`
}

// LexicalEntries is JSON the response
type LexicalEntries struct {
	Entries         []Entries        `json:"entries"`
	Language        string           `json:"language"`
	LexicalCategory LexicalCategory  `json:"lexicalCategory"`
	Pronunciations  []Pronunciations `json:"pronunciations"`
	Text            string           `json:"text"`
}

// Results is JSON the response
type Results struct {
	ID             string           `json:"id"`
	Language       string           `json:"language"`
	LexicalEntries []LexicalEntries `json:"lexicalEntries"`
	Type           string           `json:"type"`
	Word           string           `json:"word"`
}

// Metadata is JSON the response
type Metadata struct {
	Operation string `json:"operation"`
	Provider  string `json:"provider"`
	Schema    string `json:"schema"`
}

// OxfordDictionariesAPIResponse is JSON the response
type OxfordDictionariesAPIResponse struct {
	ID       string    `json:"id"`
	Metadata Metadata  `json:"metadata"`
	Results  []Results `json:"results"`
	Word     string    `json:"word"`
}

// SearchWord will list the labels used in gmail account
func SearchWord(word string) {
	var oxford OxfordDictionariesAPIResponse
	// Build our new spinner
	s := spinner.New(spinner.CharSets[39], 100*time.Millisecond)
	s.Prefix = "Loading.. .. .. "       // Prefix text before the spinner
	s.Suffix = "    ¯\\_(ツ)_/¯"         // Append text after the spinner
	s.Color("bgBlack", "bold", "fgRed") // Set the spinner color to a bold red
	s.Start()                           // Start the spinner
	//time.Sleep(4 * time.Second)         // Run for some time to simulate work
	fmt.Println("\n Word: ", word)
	fmt.Println("")

	body, err := apiRequest(word)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(body, &oxford)
	s.Stop()
	if oxford.ID == "" {
		fmt.Println("No Result..")
		os.Exit(0)
	}
	red := color.FgRed.Render
	green := color.FgGreen.Render
	gray := color.FgGray.Render
	voiceActivate := os.Getenv("OXFORD_VOICE_ACTIVATE")

	for _, results := range oxford.Results {
		for _, lexicalEntries := range results.LexicalEntries {
			color.Println("<warning> -->>>>> " + results.Word + " </> " + green(lexicalEntries.LexicalCategory.Text))

			for _, pronunciations := range lexicalEntries.Pronunciations {
				fmt.Printf("\t Notation: %s  Phonetic: %s \n", red(pronunciations.PhoneticNotation), green(pronunciations.PhoneticSpelling))
				if voiceActivate == "1" {
					playPronunciations(pronunciations.AudioFile)
				}

				for i := 0; len(pronunciations.Dialects) > i; i++ {
					fmt.Println("\t Dialects: " + gray(pronunciations.Dialects[i]))
				}
			}

			for _, entries := range lexicalEntries.Entries {
				/*for i := 0; len(entries.Etymologies) > i; i++ {
					fmt.Println("\t \t Etymologies: " + entries.Etymologies[i])
				}*/
				for _, senses := range entries.Senses {
					for i := 0; len(senses.ShortDefinitions) > i; i++ {
						fmt.Print("\n")
						if i == 0 {
							color.Light.Printf("\t %s ", "Short Definitions: ")
						}
						fmt.Println(" " + green(senses.ShortDefinitions[i]))
					}
					for i := 0; len(senses.Definitions) > i; i++ {
						color.Light.Printf("\t %s ", "Definitions: ")
						color.Println(" <error> " + senses.Definitions[i] + " </> ")
					}
					for i := 0; len(senses.Examples) > i; i++ {
						if i == 0 {
							fmt.Println("\t \t Example sentence: ")
						}
						color.Warn.Println("\t \t \t - " + senses.Examples[i].Text)
					}
					for i := 0; len(entries.Etymologies) > i; i++ {
						if i == 0 {
							fmt.Println("\t \t Etymologies: ")
						}
						color.Warn.Println("\t \t \t - " + entries.Etymologies[i])
					}

				}
			}
			fmt.Print("\n \n")

		}
	}
}

func apiRequest(word string) ([]byte, error) {
	url := "https://od-api.oxforddictionaries.com:443/api/v2/entries/en-gb/" + word

	req, _ := http.NewRequest("GET", url, nil)
	id := os.Getenv("OXFORD_APPLICATION_ID")
	key := os.Getenv("OXFORD_APPLICATION_KEY")
	req.Header.Add("app_id", id)
	req.Header.Add("app_key", key)
	req.Header.Add("Host", "od-api.oxforddictionaries.com:443")
	req.Header.Add("Accept-Encoding", "gzip, deflate")
	req.Header.Add("Connection", "keep-alive")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Unable to get details: %v", err)
		return nil, err
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return body, nil
}

func playPronunciations(fileURL string) {
	// Dowload the MP# from online and save it in temporary file
	downloadFile("/tmp/oxforddictionaries-play.mp3", fileURL)

	f, err := os.Open("/tmp/oxforddictionaries-play.mp3")
	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done

}

func downloadFile(filepath string, url string) (err error) {

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
