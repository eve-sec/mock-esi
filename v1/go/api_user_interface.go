package esiv1

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var _ time.Time
var _ = mux.NewRouter

func PostUiOpenwindowContract(w http.ResponseWriter, r *http.Request) {

	var (
		localV     interface{}
		err        error
		contractId int32
		datasource string
		token      string
	)
	// shut up warnings
	localV = localV
	err = err

	j := ``
	if err := r.ParseForm(); err != nil {
		errorOut(w, r, err)
		return
	}
	localV, err = processParameters(contractId, r.Form.Get("contract_id"))
	if err != nil {
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
			w.Header().Set("Content-Type", "")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("[]"))
			return
		}
	}

	w.Header().Set("Content-Type", "")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(j))
}

func PostUiOpenwindowInformation(w http.ResponseWriter, r *http.Request) {

	var (
		localV     interface{}
		err        error
		targetId   int32
		datasource string
		token      string
	)
	// shut up warnings
	localV = localV
	err = err

	j := ``
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
	localV, err = processParameters(targetId, r.Form.Get("target_id"))
	if err != nil {
		errorOut(w, r, err)
		return
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
			w.Header().Set("Content-Type", "")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("[]"))
			return
		}
	}

	w.Header().Set("Content-Type", "")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(j))
}

func PostUiOpenwindowMarketdetails(w http.ResponseWriter, r *http.Request) {

	var (
		localV     interface{}
		err        error
		typeId     int32
		datasource string
		token      string
	)
	// shut up warnings
	localV = localV
	err = err

	j := ``
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
	localV, err = processParameters(typeId, r.Form.Get("type_id"))
	if err != nil {
		errorOut(w, r, err)
		return
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
			w.Header().Set("Content-Type", "")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("[]"))
			return
		}
	}

	w.Header().Set("Content-Type", "")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(j))
}

func PostUiOpenwindowNewmail(w http.ResponseWriter, r *http.Request) {

	var (
		localV     interface{}
		err        error
		datasource string
		token      string
	)
	// shut up warnings
	localV = localV
	err = err

	j := ``
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
			w.Header().Set("Content-Type", "")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("[]"))
			return
		}
	}

	w.Header().Set("Content-Type", "")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(j))
}
