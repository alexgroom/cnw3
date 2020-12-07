package main

import (
	"net/http"
	"html/template"
	"log"
	"encoding/json"
)

/**
* Flag for throwing a 503 when enabled
*/
var misbehave = false

func HomePage(w http.ResponseWriter, r *http.Request){

	template := template.Must(template.ParseFiles("templates/homepage.html"))
	  
    err := template.Execute(w, nil) //execute the template
    if err != nil { // if there is an error
		log.Print("template executing error: ", err) //log it
		http.Error(w, err.Error(), http.StatusInternalServerError)
  	}
}

func GetProducts(w http.ResponseWriter, r *http.Request){

	if misbehave {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte("Misbehavior of the Catalog GoLang Service\n"))
	} else {
		products := Products{
			Product{ ItemId: "100000", Name: "Red Fedora", Description: "Official Red Hat Fedora", Price: 34.99},
			Product{ ItemId: "329299", Name: "Quarkus T-shirt", Description: "This updated unisex essential fits like a well-loved favorite, featuring a crew neck, short sleeves and designed with superior combed and ring- spun cotton.", Price: 8.50},
			Product{ ItemId: "329199", Name: "Pronounced Kubernetes", Description: "Kubernetes is changing how enterprises work in the cloud. But one of the biggest questions people have is: How do you pronounce it?", Price: 17.80},
			Product{ ItemId: "165613", Name: "Knit socks", Description: "Your brand will get noticed on these full color knit socks. Imported.", Price: 28.75},
			Product{ ItemId: "165614", Name: "Quarkus H2Go water bottle", Description: "Sporty 16. 9 oz double wall stainless steel thermal bottle with copper vacuum insulation, and threaded insulated lid. Imprinted. Imported.", Price: 6.00},
			Product{ ItemId: "165954", Name: "Patagonia Refugio pack 28L", Description: "Made from 630-denier 100% nylon (50% recycled/50% high-tenacity) plain weave; lined with 200-denier 100% recycled polyester. ...", Price: 24.00},
			Product{ ItemId: "444434", Name: "Red Hat Impact T-shirt", Description: "This 4. 3 ounce, 60% combed ringspun cotton/40% polyester jersey t- shirt features a slightly heathered appearance. The fabric laundered for reduced shrinkage. Next Level brand apparel. Printed.", Price: 106.00},
			Product{ ItemId: "444435", Name: "Quarkus twill cap", Description: "100% cotton chino twill cap with an unstructured, low-profile, six-panel design. The crown measures 3 1/8 and this features a Permacurv visor and a buckle closure with a grommet.", Price: 44.30},
			Product{ ItemId: "444437", Name: "Nanobloc Universal Webcam Cover", Description: "NanoBloc Webcam Cover fits phone, laptop, desktop, Pc, MacBook Pro, iMac, ...", Price: 44.30},
		}
		
		// Define Content-Type = application/json
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(products); err != nil {
			panic(err)
		}
	}
}

func Behave(w http.ResponseWriter, r *http.Request){
	misbehave = false
	log.Print("'misbehave' has been set to 'false'") 
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Next request to / will return 200\n"))
	return
}

func Misbehave(w http.ResponseWriter, r *http.Request){
	misbehave = true
	log.Print("'misbehave' has been set to 'true'")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Next request to / will return a 503\n"))
	return
}