package controllers
 
import (
	"net/http"
	"encoding/json"
	"fmt"
	"log"
	"campaignTracker/app/models" 
)
var campaigns []models.Campaign

func GetCampaignHandler(w http.ResponseWriter, r *http.Request){
	campaignList, err := json.Marshal(campaigns)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(campaignList))
	w.Write(campaignList)
}

func CreateCampaignHandler(w http.ResponseWriter, r *http.Request){
	campaign := models.Campaign{}
	err := r.ParseForm()
	if err != nil {

	}

	//campaign.title = r.Form.Get("title")
	//campaign.description = []byte(r.Form.Get("description"))

	campaigns = append(campaigns, campaign)

	//Finally, we redirect the user to the original HTMl page
	// (located at `/assets/`), using the http libraries `Redirect` method
	http.Redirect(w, r, "/campaigns/", http.StatusFound)
}