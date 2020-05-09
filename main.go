package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	
	"dataSearch"
)

func main() {
	fileServer := http.FileServer(http.Dir("./qrimg"))
	http.Handle("/img/", http.StripPrefix("/img/", fileServer))
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/url", urlHandler)
	http.ListenAndServe(":8080", nil)
	

}
func urlHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var url UrlLink
	err := decoder.Decode(&url)
	if err != nil {
		panic(err)
	}
	obj:=dataSearch.SearchItem(url.Url)
	fmt.Println(url.Url)
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(obj)
}
type UrlLink struct {
	Url string `json:url`
}





