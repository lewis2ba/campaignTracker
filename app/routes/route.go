package routes

import (
   "github.com/gorilla/mux"
   "campaignTracker/app/controllers"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/campaigns/", controllers.GetCampaignHandler).Methods("GET")
	r.HandleFunc("/campaigns/", controllers.CreateCampaignHandler).Methods("POST")
	
	return r
}