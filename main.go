package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/joho/godotenv"
	"html/template"
	"net/http"
	"os"
)

var (
	APIlistenPath string = ""
	OrgID         string = ""
	PolicyID      string = ""
	BaseAPIID     string = ""
	GatewayURL    string = ""
	AdminSecret   string = ""
	ClientID      string = ""
	RedirectURI   string = ""
)

func index(w http.ResponseWriter, r *http.Request) {
	/*if r.Method != "POST" {
		http.ServeFile(w, r, "tmpl/index.html")
	}*/

	tmplVal := make(map[string]string)

	tmplVal["APIlistenPath"] = APIlistenPath
	tmplVal["OrgID"] = OrgID
	tmplVal["PolicyID"] = PolicyID
	tmplVal["BaseAPIID"] = BaseAPIID
	tmplVal["GatewayURL"] = GatewayURL
	tmplVal["AdminSecret"] = AdminSecret
	tmplVal["ClientID"] = ClientID
	tmplVal["ResponseType"] = "code"
	tmplVal["RedirectURI"] = RedirectURI

	t, _ := template.ParseFiles("tmpl/index.html")
	t.Execute(w, tmplVal)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	tmplVal := make(map[string]string)

	tmplVal["ClientID"] = r.FormValue("client_id")
	tmplVal["ResponseType"] = r.FormValue("response_type")
	tmplVal["RedirectURI"] = r.FormValue("redirect_uri")

	t, _ := template.ParseFiles("tmpl/login.html")
	t.Execute(w, tmplVal)
}

func approvedHandler(w http.ResponseWriter, r *http.Request) {
	var redirect_uri = r.FormValue("redirect_uri")
	var responseType = r.FormValue("response_type")
	var clientId = r.FormValue("client_id")

	thisResponse, rErr := RequestOAuthToken(APIlistenPath,
		redirect_uri, responseType, clientId, "", OrgID, PolicyID, BaseAPIID)

	if rErr != nil {
		log.Error(rErr)
		http.Error(w, "Error!", 500)
	}
	http.Redirect(w, r, thisResponse.RedirectTo, 301)
}

func finalHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "tmpl/final.html")
}

// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(".env"); err != nil {
		log.Print("No .env file found")
	}

	APIlistenPath = getEnv("API_LISTENPATH", "oauth2")
	OrgID = getEnv("ORG_ID", "")
	PolicyID = getEnv("POLICY_ID", "")
	BaseAPIID = getEnv("API_ID", "")
	GatewayURL = getEnv("GATEWAY_URL", "")
	AdminSecret = getEnv("ADMIN_SECRET", "")
	ClientID = getEnv("CLIENT_ID", "")
	RedirectURI = getEnv("REDIRECT_URI", "")
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", loginHandler)
	mux.HandleFunc("/approved", approvedHandler)
	mux.HandleFunc("/final", finalHandler)
	mux.HandleFunc("/", index)
	log.Info("Listening on :8000")
	http.ListenAndServe(":8000", mux)
}
