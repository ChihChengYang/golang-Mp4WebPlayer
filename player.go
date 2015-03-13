package main
//http://127.0.0.1:8080/player 
import (
    "fmt" 
 	"net/http" 
 	"runtime" 
    "html/template" 
    "os"  
)
 
var gStaticMediaPath string = "/html/media/"
var gMP4FilePath string = "/" 

func isExists(path string) bool { 
    _, err := os.Stat(path) 
    if err == nil { 
      return true 
    } 
    return os.IsExist(err) 
} 

func video(w http.ResponseWriter, r *http.Request) { 
    //html/media/test.mp4
    http.ServeFile(w, r,r.URL.Path[1:])     
}

func js(w http.ResponseWriter, r *http.Request) { 
    http.ServeFile(w, r, r.URL.Path[1:])
}
 
//---------------------
type Mp4File struct {
    Name string
}
 

func playerHandler(w http.ResponseWriter, r *http.Request) { 
 
    var mp4FileName string
    mp4FileName = "test"
 
    t, err := template.ParseFiles("html/player.html")  
    if err != nil { 
        http.Error(w, err.Error(),  
        http.StatusInternalServerError)              
        return 
    } 
    t.Execute(w, Mp4File{ Name: mp4FileName }) 
 
} 
//********************************
//   
//********************************
 
func main() {
  	runtime.GOMAXPROCS(runtime.NumCPU()) 
    http.HandleFunc("/player", playerHandler) 
     //Static files
        http.HandleFunc(gStaticMediaPath, video) 
        http.HandleFunc("/html/build/", js)
 	http.ListenAndServe(":8080", nil)    
}

 