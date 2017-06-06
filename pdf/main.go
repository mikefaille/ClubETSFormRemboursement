package testPdf

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	//rest
	"github.com/mikefaille/testHttp/ressources" //rdsd

	"github.com/signintech/gopdf"
)

type DemandeRemboursement struct {
	IsVoyage          string
	Date              string
	DemandeurActif    Demandeur
	Description       string
	DepensesCourantes ListDepenses
	UBR               int
	Compte            int
	CBS               int
}

type ListDepenses struct {
	Depenses []Depense
}

type Depense struct {
	DescriptionDepense string
	Montant            float64
}

type Demandeur struct {
	CodePerm          string
	Nom               string
	Prenom            string
	ModeRemboursement string
	Addr              Adresse
}

type Adresse struct {
	Rue        string
	Ville      string
	CodePostal string
	Province   string
	Courriel   string
}

func (d ListDepenses) getTotal() (total float64) {

	for _, depense := range d.Depenses {
		total += depense.Montant
	}

	return
}

func GeneratePDF(d DemandeRemboursement) ([]byte, error) {
	var pdfLine = 431.3
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: 612.00, H: 792.00}}) //595.28, 841.89 = A4
	pdf.AddPage()
	fontBytes, err := ressources.Asset("data/DejaVuSans.ttf")
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}
	fontReader := bytes.NewReader(fontBytes)
	err = pdf.AddTTFFontByReader("DejaVuSans", fontReader)
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	// Load image from main executable
	imageBytes, err := ressources.Asset("data/RapportDepenses.png")
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	// Create temps image
	err = ioutil.WriteFile("RapportDepenses.png", imageBytes, 0644)
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}
	// Use image as report backgroung
	pdf.Image("RapportDepenses.png", 0, 0, &gopdf.Rect{W: 612.00, H: 792})
	err = pdf.SetFont("DejaVuSans", "", 9)
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}
	// Remove temp image
	os.Remove("RapportDepenses.png")

	// Autre remboursement
	pdf.SetX(37) //move current location
	pdf.SetY(123)
	pdf.Cell(nil, "X") //print text

	// Étudiant
	pdf.SetX(327.5) //move current location
	pdf.SetY(141.5)
	pdf.Cell(nil, "X") //print text

	pdf.SetX(399) //move current location
	pdf.SetY(142)
	pdf.Cell(nil, d.DemandeurActif.CodePerm) //print text

	pdf.SetX(520) //move current location
	pdf.SetY(142)
	pdf.Cell(nil, d.Date) //print text

	pdf.SetX(60) //move current location
	pdf.SetY(170)
	pdf.Cell(nil, d.DemandeurActif.Nom) //print text

	pdf.SetX(384) //move current location
	pdf.SetY(170)
	pdf.Cell(nil, d.DemandeurActif.Prenom) //print text

	pdf.SetX(81) //move current location
	pdf.SetY(222)
	pdf.Cell(nil, d.DemandeurActif.Addr.Rue) //print text

	pdf.SetX(530) //move current location
	pdf.SetY(222)
	pdf.Cell(nil, d.DemandeurActif.Addr.CodePostal) //print text

	pdf.SetX(81) //move current location
	pdf.SetY(237)
	pdf.Cell(nil, d.DemandeurActif.Addr.Ville) //print text

	pdf.SetX(484) //move current location
	pdf.SetY(237)
	pdf.Cell(nil, d.DemandeurActif.Addr.Province) //print text

	pdf.SetX(81) //move current location
	pdf.SetY(253)
	pdf.Cell(nil, d.DemandeurActif.Addr.Courriel) //print text

	fmt.Println("Is depot needed: ", d.DemandeurActif.ModeRemboursement)
	if d.DemandeurActif.ModeRemboursement == "depot" {
		pdf.SetX(440) //move current location
		pdf.SetY(255)
		pdf.Cell(nil, "X") //print text
	} else if d.DemandeurActif.ModeRemboursement == "cheque" {
		pdf.SetX(487) //move current location
		pdf.SetY(255)
		pdf.Cell(nil, "X") //print text
	}

	pdf.SetX(37) //move current location
	pdf.SetY(285)
	pdf.Cell(nil, d.Description) //print text

	fmt.Println("nb depense to write : ", len(d.DepensesCourantes.Depenses))
	fmt.Println("depense to write : ", d.DepensesCourantes.Depenses)
	for _, depense := range d.DepensesCourantes.Depenses {
		AddDepenceLine(&pdf, depense, &pdfLine)
	}

	pdf.SetX(39) //move current location
	pdf.SetY(601)
	pdf.Cell(nil, strconv.Itoa(d.UBR)) //print text

	pdf.SetX(106) //move current location
	pdf.SetY(601)
	pdf.Cell(nil, strconv.Itoa(d.Compte)) //print text

	TotalDepensesStr := strconv.FormatFloat(d.DepensesCourantes.getTotal(), 'f', 2, 64)
	pdf.SetX(253) //move current location
	pdf.SetY(601)
	rect := gopdf.Rect{W: 85, H: 0}
	pdf.CellWithOption(&rect, TotalDepensesStr, gopdf.CellOption{Align: gopdf.Right, Float: gopdf.Right}) //print text

	pdf.SetX(495) //move current location
	pdf.SetY(557)
	pdf.CellWithOption(&rect, TotalDepensesStr, gopdf.CellOption{Align: gopdf.Right, Float: gopdf.Right}) //print text

	pdf.SetX(495) //move current location
	pdf.SetY(557)
	pdf.CellWithOption(&rect, TotalDepensesStr, gopdf.CellOption{Align: gopdf.Right, Float: gopdf.Right}) //print text

	pdf.SetX(495) //move current location
	pdf.SetY(600)
	pdf.CellWithOption(&rect, TotalDepensesStr, gopdf.CellOption{Align: gopdf.Right, Float: gopdf.Right}) //print text

	pdf.SetX(253) //move current location
	pdf.SetY(674)
	pdf.CellWithOption(&rect, TotalDepensesStr, gopdf.CellOption{Align: gopdf.Right, Float: gopdf.Right}) //print text

	// pdf.SetX(165) //move current location
	// pdf.SetY(600)
	// pdf.Cell(nil, strconv.Itoa(demande.CBS)) //print text
	return pdf.GetBytesPdfReturnErr()

}

func AddDepenceLine(curPdf *gopdf.GoPdf, _depense Depense, line *float64) {
	// Détail des dépenses
	// 1ère ligne
	var localLine float64
	localLine = *line
	curPdf.SetX(36) //move current location
	curPdf.SetY(localLine)
	curPdf.Cell(nil, _depense.DescriptionDepense) //print text
	curPdf.SetX(495)                              //move current location
	montantStr := strconv.FormatFloat(_depense.Montant, 'f', 2, 64)
	rect := gopdf.Rect{W: 85, H: 0}
	curPdf.CellWithOption(&rect, montantStr, gopdf.CellOption{Align: gopdf.Right, Float: gopdf.Right}) //print text
	*line += 15.7

}
