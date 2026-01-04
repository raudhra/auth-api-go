package handlers
import(
	"encoding/json"
	"net/http"

	"go-auth-api/db"
	"go-auth-api/utils"
)

type Credentials struct{
	Email string `json:"email"`
	Password string `json:"password"`
}

func Register(w http.ResponseWriter, r *http.Request){
	var creds Credentials
	json.NewDecoder(r.Body).Decode(&creds)
}

hash, _ := utils.HashPassword(creds.Password)

_, err := db.DB.Exec(
	"INSERT INTO users (email,password) VALUES ($1,$2)", creds.Email,hash
)

if err != nil{
	http.Error(w, "User Exits", 400)
	return
}

w.Write([]byte("User Registered"))

func Login(w http.ResponseWriter, r *http.Request){
	var cred Credentials
	json.NewDecoder(r.Body).Decode(&creds)
	var id int
	var hash string
	err := db.Db.QueryRow(
		"SELECT id, password FROM user WHERE email=$1", creds.Email
	).Scan(&id,&hash)
}
if err!=nil{
	http.Error(w,"Invalid Credentials", 401)
	return
}
if utils.CheckPassword(hash, creds.Password) != nil{
	http.Error(w,"Invalid Credentials", 401)
	return
}
token, _ := utils.GenerateJWT(id)
json.NewEncoder(w).Encode(map[string]string{
	"token" : token
})
func Profile(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Protected profile data"))
}