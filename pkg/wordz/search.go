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
// [START tale_gmail]
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
	body, err := apiRequest(word)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(body, &oxford)
	//fmt.Println(res)
	//fmt.Println(string(body))
	s.Stop() // Stop the spinner

	//export OXFORD_APPLICATION_ID="3e0f2d81"
	//export OXFORD_APPLICATION_KEY="»cede5c20ad79e29608e62b3a0d16b2d3"
	fmt.Println("ID: " + oxford.ID)
	fmt.Println("Metadata Operation: " + oxford.Metadata.Operation)
	fmt.Println("Metadata Provider: " + oxford.Metadata.Provider)
	fmt.Println("Metadata Schema: " + oxford.Metadata.Schema)
	fmt.Println("Results ID: " + oxford.Results[0].ID)
	fmt.Println("Results Language: " + oxford.Results[0].Language)
	fmt.Println("Results Type: " + oxford.Results[0].Type)
	fmt.Println("Results Word: " + oxford.Results[0].Word)

	fmt.Println("LexicalEntries: " + oxford.Results[0].LexicalEntries[0].Language)
	fmt.Println("LexicalEntries: " + oxford.Results[0].LexicalEntries[0].Text)

	fmt.Println("LexicalEntries Entries Etymologies: " + oxford.Results[0].LexicalEntries[0].Entries[0].Etymologies[0])

	fmt.Println("LexicalEntries Entries Senses ID: " + oxford.Results[0].LexicalEntries[0].Entries[0].Senses[0].ID)
	//fmt.Println("LexicalEntries Entries Senses EntryID: " + oxford.Results[0].LexicalEntries[0].Entries[0].Senses[0].ThesaurusLinks[0].EntryID)
	//fmt.Println("LexicalEntries Entries Senses SenseID: " + oxford.Results[0].LexicalEntries[0].Entries[0].Senses[0].ThesaurusLinks[0].SenseID)

	//fmt.Println("LexicalEntries Entries Senses Examples Text: " + oxford.Results[0].LexicalEntries[0].Entries[0].Senses[0].Examples[0].Text)
	//fmt.Println("LexicalEntries Entries Senses Examples Text: " + oxford.Results[0].LexicalEntries[0].Entries[0].Senses[0].Examples[1].Text)

	//fmt.Println("LexicalEntries Entries Senses shortDefinitions: " + oxford.Results[0].LexicalEntries[0].Entries[0].Senses[0].ShortDefinitions[0])

	//fmt.Println("LexicalEntries Entries Senses Examples Text: " + oxford.Results[0].LexicalEntries[1].Entries[0].Senses[0].Examples[0].Text)
	//fmt.Println("LexicalEntries Entries Senses shortDefinitions: " + oxford.Results[0].LexicalEntries[1].Entries[0].Senses[0].ShortDefinitions[0])

	//fmt.Println("LexicalEntries LexicalCategory ID: " + oxford.Results[0].LexicalEntries[0].LexicalCategory.ID)
	//fmt.Println("LexicalEntries LexicalCategory Text: " + oxford.Results[0].LexicalEntries[0].LexicalCategory.Text)

	fmt.Println("LexicalEntries Pronunciations AudioFile: " + oxford.Results[0].LexicalEntries[0].Pronunciations[0].AudioFile)
	fmt.Println("LexicalEntries Pronunciations Dialects: " + oxford.Results[0].LexicalEntries[0].Pronunciations[0].Dialects[0])
	playPronunciations(oxford.Results[0].LexicalEntries[0].Pronunciations[0].AudioFile)
	fmt.Println("Word: " + oxford.Word)

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
	fmt.Println(out)
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
