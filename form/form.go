package form

const Form = `
<h1>Générateur de formulaire de remboursement</h1>
{{.Name}}
<form action="/pdf/" method="POST">
    Description:<textarea name="Description"></textarea><br>
    Date au format yyyy-mm-jj :<textarea cols="10" id="" name="Date" rows="1"></textarea><br/>
    CodePerm  :<textarea cols="30" id="" name="CodePerm" rows="1"></textarea><br/>
    Nom       :<textarea cols="30" id="" name="Nom" rows="1"></textarea><br/>
    Prenom    :<textarea cols="30" id="" name="Prenom" rows="1"></textarea><br/>
    Quel est le mode de remboursement souhaité ? : </br>
    <input type="radio" name="ModeRemboursement" value="depot">Dépôt<br>
    <input type="radio" name="ModeRemboursement" value="cheque">Chèque<br>

    <br/>
    <b>Adresse du demandeur</b><br/>

    Rue       :<textarea cols="10" id="" name="Rue" rows="1"></textarea><br/>
    Ville     :<textarea cols="10" id="" name="Ville" rows="1"></textarea><br/>
    CodePostal:<textarea cols="10" id="" name="CodePostal" rows="1"></textarea><br/>
    Province  :<textarea cols="10" id="" name="Province" rows="1"></textarea><br/>
    Courriel  :<textarea cols="10" id="" name="Courriel" rows="1"></textarea><br/>
    <hr>

    Description Achat 1:<textarea cols="10" id="" name="Depenses.0.DescriptionDepense" rows="1" row="1"></textarea><br/>
    Montant    :<textarea cols="10" id="" name="Depenses.0.Montant" rows="1"></textarea><br/>
    Description Achat  2:<textarea cols="10" id="" name="Depenses.1.DescriptionDepense" rows="1" row="1"></textarea><br/>
    Montant 2   :<textarea cols="10" id="" name="Depenses.1.Montant" rows="1"></textarea><br/>
    Description Achat  3:<textarea cols="10" id="" name="Depenses.2.DescriptionDepense" rows="1" row="1"></textarea><br/>
    Montant 3   :<textarea cols="10" id="" name="Depenses.2.Montant" rows="1"></textarea><br/>
   Description Achat  4:<textarea cols="10" id="" name="Depenses.3.DescriptionDepense" rows="1" row="1"></textarea><br/>
    Montant 4   :<textarea cols="10" id="" name="Depenses.3.Montant" rows="1"></textarea><br/>
    Description Achat  5:<textarea cols="10" id="" name="Depenses.4.DescriptionDepense" rows="1" row="1"></textarea><br/>
    Montant 5   :<textarea cols="10" id="" name="Depenses.4.Montant" rows="1"></textarea><br/>
    Description Achat  6:<textarea cols="10" id="" name="Depenses.5.DescriptionDepense" rows="1" row="1"></textarea><br/>
    Montant 6   :<textarea cols="10" id="" name="Depenses.5.Montant" rows="1"></textarea><br/>
    Description Achat  7:<textarea cols="10" id="" name="Depenses.6.DescriptionDepense" rows="1" row="1"></textarea><br/>
    Montant 7   :<textarea cols="10" id="" name="Depenses.6.Montant" rows="1"></textarea><br/>
    Description Achat  8:<textarea cols="10" id="" name="Depenses.7.DescriptionDepense" rows="1" row="1"></textarea><br/>
    Montant 8   :<textarea cols="10" id="" name="Depenses.7.Montant" rows="1"></textarea><br/>


    <br/>
    <hr>
    <br/>
    UBR :<textarea cols="10" id="" name="UBR" rows="1"></textarea><br/>
    No de compte comptable :<textarea cols="10" id="" name="Compte" rows="1"></textarea><br/>
    <br/>

    <input type="submit" value="Save">
</form>`
