package esiv3

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var _ time.Time
var _ = mux.NewRouter

func GetUniverseSystemsSystemId(w http.ResponseWriter, r *http.Request) {

	var (
		localV     interface{}
		err        error
		systemId   int32
		datasource string
		language   string
		userAgent  string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `{
  "system_id" : 30000003,
  "name" : "Akpivem",
  "position" : {
    "x" : -91174141133075340,
    "y" : 43938227486247170,
    "z" : -56482824383339900
  },
  "security_status" : 0.8462923765,
  "constellation_id" : 20000001,
  "planets" : [ {
    "planet_id" : 40000041,
    "moons" : [ 40000042 ]
  }, {
    "planet_id" : 40000043
  } ],
  "stargates" : [ 50000342 ],
  "star_id" : 40000040,
  "security_class" : "B"
}`
	vars := mux.Vars(r)
	localV, err = processParameters(systemId, vars["system_id"])
	if err != nil {
		errorOut(w, r, err)
		return
	}
	systemId = localV.(int32)
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
	if r.Form.Get("language") != "" {
		localV, err = processParameters(language, r.Form.Get("language"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		language = localV.(string)
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

func GetUniverseTypesTypeId(w http.ResponseWriter, r *http.Request) {

	var (
		localV     interface{}
		err        error
		typeId     int32
		datasource string
		language   string
		userAgent  string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `{
  "type_id" : 587,
  "name" : "Rifter",
  "description" : "The Rifter is a...",
  "published" : true,
  "group_id" : 25
}`
	vars := mux.Vars(r)
	localV, err = processParameters(typeId, vars["type_id"])
	if err != nil {
		errorOut(w, r, err)
		return
	}
	typeId = localV.(int32)
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
	if r.Form.Get("language") != "" {
		localV, err = processParameters(language, r.Form.Get("language"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		language = localV.(string)
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