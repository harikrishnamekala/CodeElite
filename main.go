package main

//Importing Standard Packages
import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func main() {
	fmt.Print("Started Server at 8080")
	//http.Handle("/", http.FileServer(http.Dir("node_modules")))
	router := httprouter.New()
	router.GET("/", showIndex)
	//router.GET("/", http.FileServer(http.Dir(".")))
	router.GET("/execute/:code", executecode)
	//http.HandleFunc("/execute", executecode)
	http.ListenAndServe(":8080", router)

}

/*func FileServe(rw http.ResponseWriter, r *http.Request) {
	http.FileServer(http.Dir("views"))
}*/
func executecode(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	code := []byte(ps.ByName("code"))
	err := ioutil.WriteFile("./controller/vol/main.c", code, 0777)
	if err != nil {
		panic(err)
	}
	fmt.Print(code)
	//runner.Runcode()
	//output, err := ioutil.ReadFile("./controller/vol/data.txt")]

	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "%s", string(code))
	fmt.Fprintf(w, "%s", ps.ByName("code"))
}
func showIndex(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	templ, err := template.New("IndexPage").ParseFiles("views/index.html")
	//static_html, err := ioutil.ReadFile("views/index.html")
	if err != nil {
		panic(err)
	}
	err = templ.Execute(os.Stdout, nil)
	//fmt.Fprintf(w, "%s", static_html)

}
