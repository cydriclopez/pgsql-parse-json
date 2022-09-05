package treedata

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"webserv/common"
)

// ***********************************************************
// This JsonData should agree with Angular-side interface
type JsonData struct {
	Data string `json:"data"`
} // *********************************************************
// Angular-side interface in src/app/services/nodeservice.ts
// export interface JsonData {
//     data:   string;
// }
// ***********************************************************

// Small-case non-exported local identifier
type tData struct {
	Jdata JsonData
	Pgx   *common.Pgx
}

// Constructor pattern using factory method
func TData(pgx *common.Pgx) *tData {
	t := new(tData)
	t.Pgx = pgx
	return t
}

// Controller for url "/api/postjsonstring"
func (t *tData) PostJsonData(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		var err error
		var unmarshalTypeError *json.UnmarshalTypeError

		decoder := json.NewDecoder(r.Body)
		if err = decoder.Decode(&t.Jdata); err != nil {

			if errors.As(err, &unmarshalTypeError) {
				jsonResponse(w, http.StatusBadRequest, "Error wrong data type: "+unmarshalTypeError.Field)

			} else {
				jsonResponse(w, http.StatusBadRequest, "Error: "+err.Error())
			}

			return
		}

		// Save json data to db
		if err = t.saveJsonData(); err != nil {
			errmsg := "Postgresql exec error:" + err.Error()
			log.Print(errmsg)
			jsonResponse(w, http.StatusInternalServerError, errmsg)
			return
		}

		jsonResponse(w, http.StatusOK, "Success")
		return
	}

	log.Print("http.NotFound")
	http.NotFound(w, r)
}

// Save json data to db
func (t *tData) saveJsonData() error {
	// Print the data from the client
	log.Println("jsonData:", t.Jdata.Data)

	// SQL statement to call the stored-function
	sql := "select tree_insert($1)"

	// Call the Postgresql stored-function
	if _, err := t.Pgx.Con.Exec(t.Pgx.Ctx, sql, t.Jdata.Data); err != nil {
		return err
	}
	return nil
}

func jsonResponse(w http.ResponseWriter, statusCode int, errorMsg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	// For production, use generic "Bad request or data error".
	// Detailed error message is not advised in production.
	w.Write([]byte(fmt.Sprintf(`{"message": "%s"}`, errorMsg) + "\n"))
}
