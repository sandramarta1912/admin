package main

import (
	"time"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"github.com/gorilla/context"
	"html/template"
)

type partner struct {
	Id      string        `db:"id"`
	IsSsp   bool          `db:"is_ssp"`
	IsDsp   bool          `db:"is_dsp"`
	Name    string        `db:"name"`
	Timeout time.Duration `db:"timeout"`
	URL     string        `db:"url"`
	Method  string        `db:"method"`
}

type partnerCollection struct {
	Partners []partner
}

type partnersData struct {
	Random int
	Manager Manager
	Partners partnerCollection
}

type partnerData struct {
	Random int
	Manager Manager
	Partner partner
}

type data struct {
	Random int
	Manager Manager
}

var partnerTemplates = template.Must(template.ParseGlob("tpl/*"))

func AddPartnerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		managerContext := context.Get(r, "manager")
		randomContext := context.Get(r, "random")

		var random int
		var manager Manager
		switch r := randomContext.(type) {
		case int:
			random = r
		}
		switch u := managerContext.(type) {
		case Manager:
			manager = u
		}
		execTmpl(w, partnerData{random, manager, partner{ "" , false, false , "", 0 , "", ""}}, "./tpl/add_partner.html")

	}

	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			panic(err)
		}

		id := r.PostFormValue("id")
		sspcheck := r.PostFormValue("is_ssp")
		ssp := false
		if sspcheck != "" {
			ssp = true
		}
		dspcheck := r.PostFormValue("is_dsp")
		dsp := false
		if dspcheck != "" {
			dsp = true
		}
		name := r.PostFormValue("name")
		timeoutStr := r.PostFormValue("timeout")
		timeout, _ := strconv.Atoi(timeoutStr)
		url := r.PostFormValue("url")
		method := r.PostFormValue("method")
		partner := partner{id, ssp, dsp, name, time.Duration(timeout), url, method}

		var createPartnerQuery = "INSERT Partners SET id=?,is_ssp=?,is_dsp=?,name=?,timeout=?, url=?, method=?"
		_, err = ds.MySql.Exec(createPartnerQuery, partner.Id, partner.IsSsp, partner.IsDsp, partner.Name, partner.Timeout, partner.URL, partner.Method)
		if err != nil {
			panic(err)
		} else {
			http.Redirect(w, r, "/partners", http.StatusSeeOther)

		}
	}
}

func ListPartnersHandler(w http.ResponseWriter, r *http.Request) {
	managerContext := context.Get(r, "manager")
	randomContext := context.Get(r, "random")


	var partners partnerCollection
	var selectAllPartnersQuery = "SELECT * FROM Partners"
	err := ds.MySql.Select(&partners.Partners, selectAllPartnersQuery)
	if err != nil {
		panic(err)
	}

	var random int
	var manager Manager
	switch r := randomContext.(type) {
	case int:
		random = r

	}
	switch u := managerContext.(type) {
	case Manager:
		manager = u

	}
	execTmpl(w, partnersData{random, manager, partners}, "./tpl/all_partners.html")

}
func DeletePartnerHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	deletePartnerQuery := "DELETE FROM Partners WHERE id=?"
	_, err := ds.MySql.Exec(deletePartnerQuery, id)
	if err != nil {
		panic(err)
	} else {
		http.Redirect(w, r, "/partners", http.StatusSeeOther)
	}
}

func EditPartnerHandler(w http.ResponseWriter, r *http.Request) {
	randomContext := context.Get(r, "random")
	managerContext := context.Get(r, "manager")
	id := mux.Vars(r)["id"]

	selectAPartnerQuery := "SELECT * FROM Partners WHERE id=?"
	var p partner
	err := ds.MySql.Get(&p, selectAPartnerQuery, id)
	if r.Method == http.MethodGet {
		var random int
		var manager Manager
		switch r := randomContext.(type) {
		case int:
			random = r

		}
		switch u := managerContext.(type) {
		case Manager:
			manager = u

		}
		execTmpl(w, partnerData{random, manager, p}, "./tpl/add_partner.html")
	}

	if r.Method == http.MethodPost {
		err = r.ParseForm()
		if err != nil {
			panic(err)
		}
		sspcheck := r.PostFormValue("is_ssp")
		ssp := false
		if sspcheck != "" {
			ssp = true
		}
		dspcheck := r.PostFormValue("is_dsp")

		dsp := false
		if dspcheck != "" {
			dsp = true
		}
		name := r.PostFormValue("name")
		timeoutStr := r.PostFormValue("timeout")
		timeout, _ := strconv.Atoi(timeoutStr)
		url := r.PostFormValue("url")
		method := r.PostFormValue("method")

		updatePartnerQuery := "UPDATE Partners SET is_ssp=?, is_dsp=?, name=?, timeout=?, url=?, method=? WHERE id=?"
		_, err = ds.MySql.Exec(updatePartnerQuery, ssp, dsp, name, timeout, url, method, id)
		if err != nil {
			panic(err)
			return
		}
		http.Redirect(w, r, "/partners", http.StatusSeeOther)
	}

}