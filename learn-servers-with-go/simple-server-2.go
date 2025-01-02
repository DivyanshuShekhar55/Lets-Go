// in this tutorial we will see a more modern and clean way for handling routes and various other things

package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type (
	api struct {
		addr string
	}

	// see getting JSON data part for this
	// If we want to have things being exportable in Go, the names must start with Uppercase
	// but unmarshalling is type-insensitive, so incoming json will match name with Name
	User struct {
		id   int
		Name string

		// remember case insensitive so Username will still be matched to username from json
		UserName string
		Email    string

		// here the field name of struct and the key name in incoming json are different
		// so we can specify that the 'address' key from json be matched to 'Addr' field here
		Addr Address `json:"address"`

		// "-" means that don't include phone in the response
		// can be used with passwords, like don't show the password in returned response
		Phone   string `json:"-"`
		Website string

		// omitempty means don't include this field in the response, if the field value is empty
		// if value is present then show me
		Company Company `json:"company,omitempty"`
	}

	Address struct {
		Street    string
		Suite     string
		City      string
		Zipcode   string
		Geography Geo `json:"geo"`
	}

	Geo struct {
		Lat string
		Lng string
	}

	Company struct {
		Name        string
		CatchPhrase string
		BS          string
	}
)

func (a *api) getHomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home Page GET"))
}

func (a *api) getUserHandler(w http.ResponseWriter, r *http.Request) {

	// lets set the headers
	// 'Add' methods appends to values associated with the key
	// 'Set' method clears all values associated with key and sets a new one
	w.Header().Set("Content-Type", "application/json")

	// next try to send and receive JSON data
	// in Go we have two terms : Marshal and Unmarshal which are analogous to serialise and deserialise in other languages like JS
	// json.Unmarshal([]byte(input), &target) this is the syntax
	// input is the json data and target is the Go data-structure in which form we want our JSON
	// the data structure : map[string]any can be used anywhere, but it's not a good practise
	// it is better to use a 'struct', that has fields matching to the values of the incoming json
	// read more here at : https://betterstack.com/community/guides/scaling-go/json-in-go/
	// if your API interactions involve potentially large JSON payloads or streaming data, using the Encode and Decode methods from the encoding/json package is more efficient and advisable
	// For simpler cases with smaller datasets where you can afford to load everything into memory, using Marshal and Unmarshal will work just fine

	// lets get data from : https://jsonplaceholder.typicode.com/users/1

	resp, err := http.Get("https://jsonplaceholder.typicode.com/users/1")
	if err != nil {
		http.Error(w, "Failed to get user", http.StatusInternalServerError)
		return
	}

	// ensure after all is done we close the response body
	defer resp.Body.Close()

	// check if request was successful
	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Failed to get user", resp.StatusCode)
		return
	}

	// read the data, on success response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response body", http.StatusInternalServerError)
		return
	}

	// decode the json body to proper struct - the User struct
	var user User
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, "Failed to decode JSON data", http.StatusInternalServerError)
		return
	}

	// encode (convert back to JSON) and send as response to browser
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, "Failed to encode json data", http.StatusInternalServerError)
		return
	}

}

func main() {
	api := &api{addr: ":8080"}

	// mux is just the fancy name for router functionality
	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    api.addr,
		Handler: mux,
	}

	mux.HandleFunc("GET /", api.getHomeHandler)
	mux.HandleFunc("GET /users", api.getUserHandler)

	srv.ListenAndServe()
}
