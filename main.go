package main

import (
	"net/http"
   "campaignTracker/app/routes"

)


func main(){
	//var campaign Campaign
	//campaign.title = "Cthulu's Wrath"
	//campaign.description = []byte("this is a description")
	//fmt.Println(campaign.title)
	routes := routes.NewRouter()
	http.ListenAndServe(":8080", routes)
}