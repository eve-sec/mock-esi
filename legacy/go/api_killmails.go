package esilegacy

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var _ time.Time
var _ = mux.NewRouter

func GetCharactersCharacterIdKillmailsRecent(w http.ResponseWriter, r *http.Request) {

	var (
		localV      interface{}
		err         error
		characterId int32
		datasource  string
		page        int32
		token       string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `[ {
  "killmail_hash" : "8eef5e8fb6b88fe3407c489df33822b2e3b57a5e",
  "killmail_id" : 2
}, {
  "killmail_hash" : "b41ccb498ece33d64019f64c0db392aa3aa701fb",
  "killmail_id" : 1
} ]`
	vars := mux.Vars(r)
	localV, err = processParameters(characterId, vars["character_id"])
	if err != nil {
		errorOut(w, r, err)
		return
	}
	characterId = localV.(int32)
	if err := r.ParseForm(); err != nil {
		errorOut(w, r, err)
		return
	}
	if r.Form.Get("datasource") != "" {
		localV, err = processParameters(datasource, r.Form.Get("datasource"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		datasource = localV.(string)
	}
	if r.Form.Get("page") != "" {
		localV, err = processParameters(page, r.Form.Get("page"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		page = localV.(int32)
	}
	if r.Form.Get("token") != "" {
		localV, err = processParameters(token, r.Form.Get("token"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		token = localV.(string)
	}

	if r.Form.Get("page") != "" {
		var (
			localPage    int32
			localIntPage interface{}
		)
		localIntPage, err := processParameters(localPage, r.Form.Get("page"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		localPage = localIntPage.(int32)
		if localPage > 1 {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("[]"))
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(j))
}

func GetCorporationsCorporationIdKillmailsRecent(w http.ResponseWriter, r *http.Request) {

	var (
		localV        interface{}
		err           error
		corporationId int32
		datasource    string
		page          int32
		token         string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `[ {
  "killmail_hash" : "8eef5e8fb6b88fe3407c489df33822b2e3b57a5e",
  "killmail_id" : 2
}, {
  "killmail_hash" : "b41ccb498ece33d64019f64c0db392aa3aa701fb",
  "killmail_id" : 1
} ]`
	vars := mux.Vars(r)
	localV, err = processParameters(corporationId, vars["corporation_id"])
	if err != nil {
		errorOut(w, r, err)
		return
	}
	corporationId = localV.(int32)
	if err := r.ParseForm(); err != nil {
		errorOut(w, r, err)
		return
	}
	if r.Form.Get("datasource") != "" {
		localV, err = processParameters(datasource, r.Form.Get("datasource"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		datasource = localV.(string)
	}
	if r.Form.Get("page") != "" {
		localV, err = processParameters(page, r.Form.Get("page"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		page = localV.(int32)
	}
	if r.Form.Get("token") != "" {
		localV, err = processParameters(token, r.Form.Get("token"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		token = localV.(string)
	}

	if r.Form.Get("page") != "" {
		var (
			localPage    int32
			localIntPage interface{}
		)
		localIntPage, err := processParameters(localPage, r.Form.Get("page"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		localPage = localIntPage.(int32)
		if localPage > 1 {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("[]"))
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(j))
}

func GetKillmailsKillmailIdKillmailHash(w http.ResponseWriter, r *http.Request) {

	var (
		localV       interface{}
		err          error
		killmailHash string
		killmailId   int32
		datasource   string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `{
  "attackers" : [ {
    "character_id" : 95810944,
    "corporation_id" : 1000179,
    "damage_done" : 5745,
    "faction_id" : 500003,
    "final_blow" : true,
    "security_status" : -0.3,
    "ship_type_id" : 17841,
    "weapon_type_id" : 3074
  } ],
  "killmail_id" : 56733821,
  "killmail_time" : "2016-10-22T17:13:36Z",
  "solar_system_id" : 30002976,
  "victim" : {
    "alliance_id" : 621338554,
    "character_id" : 92796241,
    "corporation_id" : 841363671,
    "damage_taken" : 5745,
    "items" : [ {
      "flag" : 20,
      "item_type_id" : 5973,
      "quantity_dropped" : 1,
      "singleton" : 0
    } ],
    "position" : {
      "x" : 4.521866005694748E11,
      "y" : 1.4670496149090222E11,
      "z" : 1.0951459653254477E11
    },
    "ship_type_id" : 17812
  }
}`
	vars := mux.Vars(r)
	localV, err = processParameters(killmailHash, vars["killmail_hash"])
	if err != nil {
		errorOut(w, r, err)
		return
	}
	killmailHash = localV.(string)
	localV, err = processParameters(killmailId, vars["killmail_id"])
	if err != nil {
		errorOut(w, r, err)
		return
	}
	killmailId = localV.(int32)
	if err := r.ParseForm(); err != nil {
		errorOut(w, r, err)
		return
	}
	if r.Form.Get("datasource") != "" {
		localV, err = processParameters(datasource, r.Form.Get("datasource"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		datasource = localV.(string)
	}

	if r.Form.Get("page") != "" {
		var (
			localPage    int32
			localIntPage interface{}
		)
		localIntPage, err := processParameters(localPage, r.Form.Get("page"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		localPage = localIntPage.(int32)
		if localPage > 1 {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("[]"))
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(j))
}
