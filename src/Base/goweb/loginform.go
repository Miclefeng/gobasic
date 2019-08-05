package main 
import(
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func sayhelloName(w http.ResponseWriter,r *http.Request){
	// 解析url传递的参数，对于POST则解析响应包的主体（request body）
	r.ParseForm()
	// 注意:如果没有调用ParseForm方法，下面无法获取表单的数据
	// 这些信息是输出到服务器端的打印信息
	fmt.Println(r.Form)
	fmt.Println("path",r.URL.Path)
	fmt.Println("scheme",r.URL.Scheme)
	fmt.Println(r.Form["url_log"])
	for k,v := range r.Form{
		fmt.Println("key : ",k)
		fmt.Println("val : ",strings.Join(v,""))
	}
	// 这个写入到w的是输出到客户端的
	fmt.Fprintf(w,"Hello astaxie!")
}

func login(w http.ResponseWriter,r *http.Request){
	fmt.Println("method : ",r.Method)
	if r.Method == "GET"{
		t,_ := template.ParseFiles("login.html")
		t.Execute(w,nil)
	}else{
		// 注意:如果没有调用ParseForm方法，下面无法获取表单的数据
		r.ParseForm()
		// 请求的是登陆数据，那么执行登陆的逻辑判断
		// 如r.Form["username"]也可写成r.FormValue("username")。调用r.FormValue时会自动调用r.ParseForm
		fmt.Println("username : ",r.Form["username"])
		fmt.Println("password : ",r.Form["password"])
	}
}

func main(){
	// 设置访问的路由
	http.HandleFunc("/",sayhelloName)
	http.HandleFunc("/login",login)
	err := http.ListenAndServe(":9090",nil)
	if err != nil{
		log.Fatal("ListenAndServe : ",err)
	}

	
}