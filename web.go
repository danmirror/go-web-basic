package main

import(
	"fmt"
	"net/http"
	"html/template"
	"path"
)
func main(){
	
	http.HandleFunc("/", func(w http.ResponseWriter, r* http.Request){
		var pathfile = path.Join("views","index.html")
		var temp, err = template.ParseFiles(pathfile)
		
		if err != nil {
			http.Error(w, err.Error(),http.StatusInternalServerError)
			return
		}
		var data = map[string]interface{}{
			"title" : "ini judul",
			"body" :"ini body ",
			"message" : "ini go",
		}
		 temp.Execute(w, data)
		// if errr != nil {
		// 	http.Error(w, errr.Error(),http.StatusInternalServerError)
		// 	return
		// }
	})
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))
	
	fmt.Printf("running server")

	http.ListenAndServe(":8080",nil)
}


