package main

import (
	"encoding/json"
	"fmt"
	Rt "github.com/dancannon/gorethink"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"strings"
	"time"
)

// Map - string - interface
type Mst map[string]interface{} 

// Global session for database connect
var sessionArray []*Rt.Session


/***************************************************************************************************************************************
 *
 *   Title        : Connection to DB
 *   Initialisation Service
 *   Date         : 2017-18-12
 *   Description  : Initialization DB Connect
 *   Author       : Savchenko Arthur
 *   ☎           : 8-097-0000000007
 *
 ****************************************************************************************************************************************/
func init() {
	session, err := Rt.Connect(Rt.ConnectOpts{Address: IpPort, Database: DatabaseName, AuthKey: DatabaseKey})
	Err(err, "No connection to database.")
	session.SetMaxOpenConns(200)
	session.SetMaxIdleConns(200)
	sessionArray = append(sessionArray, session)
}


/********************************************************************************
 Main procedure
********************************************************************************/
func main() {

	// Route
	// Api data
	http.HandleFunc("/api/data/insert/",   api_data_insert)       // Добавление данных

	err := http.ListenAndServe(Ipport, nil)
	Err(err, "Error start service.")
}



/***************************************************************
  Author      Savchenko Arthur
  Company     
  Description this
  Time        Add array documents to database
  Title       api_data_insert
  Url         /api/data/insert/
 ****************************************************************/
func api_data_insert(w http.ResponseWriter, r *http.Request) {

	// Init variables
	var m []Mst
	d := r.Header.Get("Db") // Контроль секретного ключа для передачи
	t := r.Header.Get("Tb") // Контроль ИД структуры

	// Один раз при вставке
	merge := Mst{"Info": "Ins", "Unixid":  "Created": NowTime(), "Title":"Test insert"}

	// Контроль конфликтных ситуаций при вставке с одинаковыми ID
	// Замена документа при вставке с неуникальным
	// конфликтом первичного ключа  (ID)
	Conflictrule := Rt.InsertOpts{Conflict: "replace", Durability: "soft"}
  
        // Body
	reads, errbody := ioutil.ReadAll(r.Body)
	Err(errbody, "Error read body.")
	defer r.Body.Close()

	// Документы
	errj := json.Unmarshal([]byte(reads), &m)
	Err(errj, "Error unmarshaling document.")

	// Добавление документа c добавлением дополнительных полей
	err := Rt.DB("test").Table("work").Insert(Rt.Expr(m), Conflictrule).Exec(sessionArray[0])

	Err(err, "Not insert data")
  
        // Контроль возвращаемых записей
	// Recinsert  := Rec.Inserted
	// fmt.Fprintln(w,"Adding records to database ",Recinsert)
}

/***************************************************************
  Description Error
  Time        11-12-2017
  Title       Error exeption
 ****************************************************************/
func Err(Er error, Txt string) {
	if Er != nil {
		log.Println("ERROR : "+Txt, "Description : "+Er.Error())
		return
	}
}


