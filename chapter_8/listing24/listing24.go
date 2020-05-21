// this sample program demonstrates how to decode a JSON response using the json package and NewDecoder function
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type (
	// gResult maps to the result document received from the search
	gResult struct {
		GsearchResultClass string 'json:"GsearchResultClass"'
		UnescapedURL       string 'json:"unescapedUrl"'
		URL                string 'json:"url"'
		VisibleURL         string 'json:"visibleUrl"'
		CacheURL           string 'json:"cacheUrl"'
		Title              string 'json:"title"'
		TitleNoFormatting  string 'json:"titleNoFormatting"'
		Content            string 'json:"content"'
	}

	// gResponse contains the top level document
	gResponse struct {
		ResponceData struct {
			Results []gResult 'json:"results"'
		} 'json:"responseData"'
	}
)

func main() { 
	uri := "http://ajax.googleapis.com/ajax/services/search/web?v=1.0&rsz=8&q=golang"

	// issue the search against Google
	resp, err := http.Get(uri)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}
	defer resp.Body.Close()

	// decode the JSON response into our struct type
	var gr gResponse
	err = json.NewDecoder(resp.Body).Decode(&gr)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	fmt.Println(gr)
}