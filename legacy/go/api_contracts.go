package esilegacy

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var _ time.Time
var _ = mux.NewRouter

func GetCharactersCharacterIdContracts(w http.ResponseWriter, r *http.Request) {

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
  "acceptor_id" : 0,
  "assignee_id" : 0,
  "availability" : "public",
  "buyout" : 1.000000000001E10,
  "contract_id" : 1,
  "date_accepted" : "2017-06-06T13:12:32Z",
  "date_completed" : "2017-06-06T13:12:32Z",
  "date_expired" : "2017-06-13T13:12:32Z",
  "date_issued" : "2017-06-06T13:12:32Z",
  "days_to_complete" : 0,
  "end_location_id" : 60014719,
  "for_corporation" : true,
  "issuer_corporation_id" : 456,
  "issuer_id" : 123,
  "price" : 1000000.01,
  "reward" : 0.01,
  "start_location_id" : 60014719,
  "status" : "outstanding",
  "type" : "auction",
  "volume" : 0.01
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

func GetCharactersCharacterIdContractsContractIdBids(w http.ResponseWriter, r *http.Request) {

	var (
		localV      interface{}
		err         error
		characterId int32
		contractId  int32
		datasource  string
		token       string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `[ {
  "amount" : 1.23,
  "bid_id" : 1,
  "bidder_id" : 123,
  "date_bid" : "2017-01-01T10:10:10Z"
} ]`
	vars := mux.Vars(r)
	localV, err = processParameters(characterId, vars["character_id"])
	if err != nil {
		errorOut(w, r, err)
		return
	}
	characterId = localV.(int32)
	localV, err = processParameters(contractId, vars["contract_id"])
	if err != nil {
		errorOut(w, r, err)
		return
	}
	contractId = localV.(int32)
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

func GetCharactersCharacterIdContractsContractIdItems(w http.ResponseWriter, r *http.Request) {

	var (
		localV      interface{}
		err         error
		characterId int32
		contractId  int32
		datasource  string
		token       string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `[ {
  "is_included" : true,
  "is_singleton" : false,
  "quantity" : 1,
  "record_id" : 123456,
  "type_id" : 587
} ]`
	vars := mux.Vars(r)
	localV, err = processParameters(characterId, vars["character_id"])
	if err != nil {
		errorOut(w, r, err)
		return
	}
	characterId = localV.(int32)
	localV, err = processParameters(contractId, vars["contract_id"])
	if err != nil {
		errorOut(w, r, err)
		return
	}
	contractId = localV.(int32)
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

func GetContractsPublicBidsContractId(w http.ResponseWriter, r *http.Request) {

	var (
		localV     interface{}
		err        error
		contractId int32
		datasource string
		page       int32
	)
	// shut up warnings
	localV = localV
	err = err

	j := `[ {
  "amount" : 1.23,
  "bid_id" : 1,
  "date_bid" : "2017-01-01T10:10:10Z"
} ]`
	vars := mux.Vars(r)
	localV, err = processParameters(contractId, vars["contract_id"])
	if err != nil {
		errorOut(w, r, err)
		return
	}
	contractId = localV.(int32)
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

func GetContractsPublicItemsContractId(w http.ResponseWriter, r *http.Request) {

	var (
		localV     interface{}
		err        error
		contractId int32
		datasource string
		page       int32
	)
	// shut up warnings
	localV = localV
	err = err

	j := `[ {
  "is_included" : true,
  "item_id" : 123456,
  "quantity" : 1,
  "record_id" : 123456,
  "type_id" : 587
} ]`
	vars := mux.Vars(r)
	localV, err = processParameters(contractId, vars["contract_id"])
	if err != nil {
		errorOut(w, r, err)
		return
	}
	contractId = localV.(int32)
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

func GetContractsPublicRegionId(w http.ResponseWriter, r *http.Request) {

	var (
		localV     interface{}
		err        error
		regionId   int32
		datasource string
		page       int32
	)
	// shut up warnings
	localV = localV
	err = err

	j := `[ {
  "buyout" : 1.000000000001E10,
  "contract_id" : 1,
  "date_expired" : "2017-06-13T13:12:32Z",
  "date_issued" : "2017-06-06T13:12:32Z",
  "days_to_complete" : 0,
  "end_location_id" : 60014719,
  "for_corporation" : true,
  "issuer_corporation_id" : 456,
  "issuer_id" : 123,
  "price" : 1000000.01,
  "reward" : 0.01,
  "start_location_id" : 60014719,
  "type" : "auction",
  "volume" : 0.01
} ]`
	vars := mux.Vars(r)
	localV, err = processParameters(regionId, vars["region_id"])
	if err != nil {
		errorOut(w, r, err)
		return
	}
	regionId = localV.(int32)
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

func GetCorporationsCorporationIdContracts(w http.ResponseWriter, r *http.Request) {

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
  "acceptor_id" : 0,
  "assignee_id" : 0,
  "availability" : "public",
  "buyout" : 1.000000000001E10,
  "contract_id" : 1,
  "date_expired" : "2017-06-13T13:12:32Z",
  "date_issued" : "2017-06-06T13:12:32Z",
  "days_to_complete" : 0,
  "end_location_id" : 60014719,
  "for_corporation" : true,
  "issuer_corporation_id" : 456,
  "issuer_id" : 123,
  "price" : 1000000.01,
  "reward" : 0.01,
  "start_location_id" : 60014719,
  "status" : "outstanding",
  "type" : "auction",
  "volume" : 0.01
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

func GetCorporationsCorporationIdContractsContractIdBids(w http.ResponseWriter, r *http.Request) {

	var (
		localV        interface{}
		err           error
		contractId    int32
		corporationId int32
		datasource    string
		page          int32
		token         string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `[ {
  "amount" : 1.23,
  "bid_id" : 1,
  "bidder_id" : 123,
  "date_bid" : "2017-01-01T10:10:10Z"
} ]`
	vars := mux.Vars(r)
	localV, err = processParameters(contractId, vars["contract_id"])
	if err != nil {
		errorOut(w, r, err)
		return
	}
	contractId = localV.(int32)
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

func GetCorporationsCorporationIdContractsContractIdItems(w http.ResponseWriter, r *http.Request) {

	var (
		localV        interface{}
		err           error
		contractId    int32
		corporationId int32
		datasource    string
		token         string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `[ {
  "is_included" : true,
  "is_singleton" : false,
  "quantity" : 1,
  "record_id" : 123456,
  "type_id" : 587
} ]`
	vars := mux.Vars(r)
	localV, err = processParameters(contractId, vars["contract_id"])
	if err != nil {
		errorOut(w, r, err)
		return
	}
	contractId = localV.(int32)
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
