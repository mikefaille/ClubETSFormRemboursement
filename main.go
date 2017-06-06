package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/mikefaille/testHttp/form"
	"github.com/mikefaille/testHttp/pdf"
)

func init() {

}

func generatePDFHandler(w http.ResponseWriter, r *http.Request) {
	var demande testPdf.DemandeRemboursement
	if r.Method == "GET" {
		tmpl := template.Must(template.New("form").Parse(form.Form))
		if err := tmpl.ExecuteTemplate(w, "form", nil); err != nil {
			fmt.Println(err)

		}

	} else {

		var ceDemandeur testPdf.Demandeur
		var cetteAdresse testPdf.Adresse

		err := r.ParseForm()

		if err != nil {
			fmt.Println(err)
		}

		decoder := schema.NewDecoder()
		err = decoder.Decode(&ceDemandeur, r.PostForm)
		err = decoder.Decode(&cetteAdresse, r.PostForm)
		ceDemandeur.Addr = cetteAdresse
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Voici le demandeur :", ceDemandeur)

		var depenses testPdf.ListDepenses
		err = decoder.Decode(&depenses, r.PostForm)
		demande.DepensesCourantes = depenses
		fmt.Println("ma dépense: ", depenses)
		demande.DemandeurActif = ceDemandeur

		err = decoder.Decode(&demande, r.PostForm)

		pdfBytes, err := testPdf.GeneratePDF(demande)
		if err != nil {
			fmt.Println(err)
		}

		w.Header().Set("Content-Type", "application/pdf")
		w.Write(pdfBytes)

	}

}

func main() {

	http.HandleFunc("/pdf/", generatePDFHandler)

	http.ListenAndServe(":8080", nil)
}
