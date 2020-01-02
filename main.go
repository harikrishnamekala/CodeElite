package main

//Importing Standard Packages
import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	//Message for Server port
	fmt.Print("Started Server at 8080")
	//------------------------------------------------------------------
	//http.Handle("/", http.FileServer(http.Dir("node_modules")))
	//router := httprouter.New()
	//resourcesRouter()
	//router.GET("/", showIndex)
	//router.GET("/", http.FileServer(http.Dir(".")))
	//router.ServeFiles("./node_modules", http.Dir("./node_modules"))
	//router.POST("/executecode", executecode)
	//------------------------------------------------------------------

	//Handler for Routes for the Requests

	// http.HandleFunc("/", showIndex)
	// http.Handle("/node_modules/", http.StripPrefix("/node_modules/", http.FileServer(http.Dir("node_modules"))))

	// http.HandleFunc("/executecode", executecode)

	// http.ListenAndServe(":8080", nil)

	r := gin.Default()
	r.GET("/ping", Pong)
	r.Run()

}

func Pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

/*
 There is no Inbuild Function to Convert []string to []byte so I had to write
 my own one
*/
// const maxInt32 = 1<<(32-1) - 1

// func writeLen(b []byte, l int) []byte {
// 	if 0 > l || l > maxInt32 {
// 		panic("writeLen: invalid length")
// 	}
// 	var lb [4]byte
// 	binary.BigEndian.PutUint32(lb[:], uint32(l))
// 	return append(b, lb[:]...)
// }
// func readLen(b []byte) ([]byte, int) {
// 	if len(b) < 4 {
// 		panic("readLen: invalid length")
// 	}
// 	l := binary.BigEndian.Uint32(b)
// 	if l > maxInt32 {
// 		panic("readLen: invalid length")
// 	}
// 	return b[4:], int(l)
// }
// func Encode(s []string) []byte {
// 	var b []byte
// 	b = writeLen(b, len(s))
// 	for _, ss := range s {
// 		b = writeLen(b, len(ss))
// 		b = append(b, ss...)
// 	}
// 	return b
// }

/*
 To Generate the Random Indexes to Handle Concurrent Requests of Compilations
*/

// const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// func init() {
// 	rand.Seed(time.Now().UnixNano())
// }

// func RandStringBytes(n int) string {
// 	b := make([]byte, n)
// 	for i := range b {
// 		b[i] = letterBytes[rand.Intn(len(letterBytes))]
// 	}
// 	return string(b)
// }

// /*
// THe Handler Method to Handle the Post of Student's Code
// */
// func executecode(w http.ResponseWriter, r *http.Request) {
// 	//Parse the Received Form in the Request Object
// 	r.ParseForm()
// 	//Getting Values based on the Form Attributes
// 	code_str := r.Form["scode"]
// 	language_id_arr := r.Form["Programming_language"]
// 	language_id, err := strconv.Atoi(language_id_arr[0])
// 	if err != nil {
// 		panic(err)
// 	}

// 	language_id -= 1
// 	//In Order to Write the Code Files
// 	//code := []byte(code_str[0])
// 	input_str := r.Form["scode_input"]
// 	code := code_str[0]
// 	input := input_str[0]

// 	fmt.Println(input)
// 	fmt.Println(code_str)
// 	fmt.Println(language_id)

// 	templateObjVal := controller.Runcode(language_id, code, input)

// 	problemtem, err := template.ParseFiles("./views/index.html")
// 	if err != nil {
// 		panic(err)
// 	}

// 	err = problemtem.Execute(w, templateObjVal)
// 	if err != nil {
// 		panic(err)

// 	}
// }

/*-------------------------------------------------------------
  	randfolder := RandStringBytes(12)
  	path := "./controller/vol/" + randfolder

  	if _, err := os.Stat(path); os.IsNotExist(err) {
  		os.Mkdir(path, 0777)
  	}
  	_, copyHostCfile := exec.Command("/bin/bash", "-c", "cp ./controller/vol/main.c "+path).Output()
  	if copyHostCfile != nil {
  		panic(copyHostCfile)
  	}
  	_, copyHostShellfile := exec.Command("/bin/bash", "-c", "cp ./controller/vol/compile.sh "+path).Output()
  	if copyHostShellfile != nil {
  		panic(copyHostShellfile)
  	}
  	_, copyHostInputfile := exec.Command("/bin/bash", "-c", "cp ./controller/vol/input.txt "+path).Output()
  	if copyHostInputfile != nil {
  		panic(copyHostInputfile)
  	}
  	fmt.Println("Created Host ENV")
  	err := ioutil.WriteFile("./controller/vol/"+randfolder+"/main.c", code, 0777)
  	if err != nil {
  		panic(err)
  	}
  	err = ioutil.WriteFile("./controller/vol/"+randfolder+"/input.txt", input, 0777)
  	if err != nil {
  		panic(err)
  	}
  	runner.Runcode(path, randfolder)

  	//time.Sleep(time.Second * 10)
  	output, err := ioutil.ReadFile("./controller/vol/" + randfolder + "/data.txt")
  	if err != nil {
  		panic(err)
  	}
  	output_errors, err := ioutil.ReadFile("./controller/vol/" + randfolder + "/errors.txt")
  	if err != nil {
  		panic(err)
  	}
  	templ_output := string(output) + string(output_errors)
  	//fmt.Fprintf(w, "%s", string(code))
  	//fmt.Fprintf(w, "%s", ps.ByName("code"))
  	templ, err := template.ParseFiles("views/index.html")

  	if err != nil {
  		panic(err)
  	}
  	templ_output_obj := OutCode{
  		Code:   code_str[0],
  		Output: templ_output,
  	}
  	err = templ.Execute(w, templ_output_obj)
  }

*/

// func showIndex(w http.ResponseWriter, r *http.Request) {

// 	templ, err := template.ParseFiles("views/index.html")
// 	//static_html, err := ioutil.ReadFile("views/index.html")

// 	if err != nil {
// 		panic(err)
// 	}

// 	dataObj := new(controller.OutputTeplStr)

// 	err = templ.Execute(w, &dataObj)
// 	if err != nil {
// 		panic(err)
// 	}
// 	//fmt.Fprintf(w, "%s", static_html)

// }

/*func resourcesRouter() {
	searchDir := "./node_modules"

	fileList := []string{}
	_ = filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, path)
		return nil
	})
	//router := httprouter.New()

	for _, file := range fileList[4:] {

		fmt.Println(file)

		//router.ServeFiles("./"+file, http.Dir(searchDir))
	}
	fmt.Print("------------------------------------------------------------Statin Glb")
	arr, err := filepath.Glob("node_modules/*")
	if err != nil {
		panic(err)
	}
	for _, data := range arr {
		fmt.Println(data)
		router.ServeFiles("./"+data, http.Dir(searchDir))
	}

}*/
