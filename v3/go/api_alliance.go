package esiv3

import (
	"net/http"
	"github.com/gorilla/mux"
	"time"
)

var _ time.Time
var _ = mux.NewRouter

func GetAlliancesAllianceId(w http.ResponseWriter, r *http.Request) {

	var (
		localV interface{}
		err error
		allianceId int32
		datasource string
		userAgent string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `{
  "creator_corporation_id" : 45678,
  "creator_id" : 12345,
  "date_founded" : "2016-06-26T21:00:00Z",
  "executor_corporation_id" : 98356193,
  "name" : "C C P Alliance",
  "ticker" : "<C C P>"
}`
	vars := mux.Vars(r)
	localV, err = processParameters(allianceId, vars["alliance_id"])
	if err != nil {
		errorOut(w, r, err)
		return
	}
	allianceId = localV.(int32)
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
	if r.Form.Get("userAgent") != "" {
		localV, err = processParameters(userAgent, r.Form.Get("user_agent"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		userAgent = localV.(string)
	}

	if r.Form.Get("page") != "" {
		var (
			localPage int32 
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


