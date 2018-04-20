package esiv2

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var _ time.Time
var _ = mux.NewRouter

func GetUniverseFactions(w http.ResponseWriter, r *http.Request) {

	var (
		localV     interface{}
		err        error
		datasource string
		language   string
		userAgent  string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `[ {
  "faction_id" : 1,
  "name" : "Faction",
  "description" : "blah blah",
  "solar_system_id" : 123,
  "corporation_id" : 456,
  "size_factor" : 1.0,
  "station_count" : 1000,
  "station_system_count" : 100,
  "is_unique" : true
} ]`
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

func GetUniverseStationsStationId(w http.ResponseWriter, r *http.Request) {

	var (
		localV     interface{}
		err        error
		stationId  int32
		datasource string
		userAgent  string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `{
  "station_id" : 60000277,
  "name" : "Jakanerva III - Moon 15 - Prompt Delivery Storage",
  "type_id" : 1531,
  "position" : {
    "x" : 165632286720,
    "y" : 2771804160,
    "z" : -2455331266560
  },
  "system_id" : 30000148,
  "reprocessing_efficiency" : 0.5,
  "reprocessing_stations_take" : 0.05,
  "max_dockable_ship_volume" : 50000000,
  "office_rental_cost" : 10000,
  "services" : [ "courier-missions", "reprocessing-plant", "market", "repair-facilities", "fitting", "news", "storage", "insurance", "docking", "office-rental", "loyalty-point-store", "navy-offices" ],
  "owner" : 1000003,
  "race_id" : 1
}`
	vars := mux.Vars(r)
	localV, err = processParameters(stationId, vars["station_id"])
	if err != nil {
		errorOut(w, r, err)
		return
	}
	stationId = localV.(int32)
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

func GetUniverseSystemKills(w http.ResponseWriter, r *http.Request) {

	var (
		localV     interface{}
		err        error
		datasource string
		userAgent  string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `[ {
  "system_id" : 30002410,
  "ship_kills" : 42,
  "npc_kills" : 0,
  "pod_kills" : 24
} ]`
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

func PostUniverseNames(w http.ResponseWriter, r *http.Request) {

	var (
		localV     interface{}
		err        error
		datasource string
		userAgent  string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `[ {
  "id" : 95465499,
  "name" : "CCP Bartender",
  "category" : "character"
}, {
  "id" : 30000142,
  "name" : "Jita",
  "category" : "solar_system"
} ]`
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