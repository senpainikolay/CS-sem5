package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/senpainikolay/CS-sem5/controllers"
	"github.com/senpainikolay/CS-sem5/models"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB             *gorm.DB
	AuthController controllers.AuthController
	// TODO: randomize it
	oauthStateString = "init-pseudo-random-then-is-the-user-id-uuid"

	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/callback",
		ClientID:     goDotEnvVariable("GOOGLE_CLIENT_ID"),
		ClientSecret: goDotEnvVariable("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
)

func init() {
	var err error
	DB, err = gorm.Open(sqlite.Open("golang.db"), &gorm.Config{})
	DB.AutoMigrate(&models.User{})

	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}
	fmt.Println("Connected Successfully to the Database")

	AuthController = controllers.NewAuthController(DB)
}

func RunServer() {

	r := mux.NewRouter()
	// Basic
	r.HandleFunc("/register", RegisterUser).Methods("POST")
	r.HandleFunc("/login", LogInUser).Methods("POST")
	// OTP
	r.HandleFunc("/generateOTP", GenerateOTP).Methods("POST")
	r.HandleFunc("/verifyOTP", VerifyOTP).Methods("POST")
	r.HandleFunc("/validateOTP", ValidateOTP).Methods("POST")
	r.HandleFunc("/disableOTP", DisableOTP).Methods("POST")

	//Google
	r.HandleFunc("/login-google/{id}", handleGoogleLogin)
	r.HandleFunc("/callback", handleGoogleCallback)

	log.Println("Runining on localhost:8080")
	http.ListenAndServe(":8080", r)
}

func handleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	oauthStateString = vars["id"]
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	content, err := getUserInfo(r.FormValue("state"), r.FormValue("code"))
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	fmt.Fprintf(w, "Content: %s\n  User Updated: %s ", content, AuthController.ValidateOAuth(oauthStateString))
}

func getUserInfo(state string, code string) ([]byte, error) {
	if state != oauthStateString {
		return nil, fmt.Errorf("invalid oauth state")
	}

	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}

	return contents, nil
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	var usrReg models.RegisterUserInput
	err := json.NewDecoder(r.Body).Decode(&usrReg)
	if err != nil {
		log.Fatalln("There was an error decoding the request body into the struct")
	}
	fmt.Fprint(w, AuthController.SignUpUser(usrReg))

}

func GenerateOTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	var otp models.OTPInput
	err := json.NewDecoder(r.Body).Decode(&otp)
	if err != nil {
		log.Fatalln("There was an error decoding the request body into the struct")
	}
	fmt.Fprint(w, AuthController.GenerateOTP(otp))

}
func VerifyOTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	var otp models.OTPInput
	err := json.NewDecoder(r.Body).Decode(&otp)
	if err != nil {
		log.Fatalln("There was an error decoding the request body into the struct")
	}
	fmt.Fprint(w, AuthController.VerifyOTP(otp))

}
func ValidateOTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	var otp models.OTPInput
	err := json.NewDecoder(r.Body).Decode(&otp)
	if err != nil {
		log.Fatalln("There was an error decoding the request body into the struct")
	}
	fmt.Fprint(w, AuthController.ValidateOTP(otp))

}
func DisableOTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	var otp models.OTPInput
	err := json.NewDecoder(r.Body).Decode(&otp)
	if err != nil {
		log.Fatalln("There was an error decoding the request body into the struct")
	}
	fmt.Fprint(w, AuthController.DisableOTP(otp))

}
func LogInUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	var usrLogIn models.LoginUserInput
	err := json.NewDecoder(r.Body).Decode(&usrLogIn)
	if err != nil {
		log.Fatalln("There was an error decoding the request body into the struct")
	}
	fmt.Fprint(w, AuthController.LoginUser(usrLogIn))

}

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
