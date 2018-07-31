package utils

import (
	"net/http"

	"github.com/blastbeatsandcode/blastbeatsandcode-website/models"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

// Create a new cookie store with a key to sign our cookies
// In this way, creating a fake cookie will not allow access
var store = sessions.NewCookieStore([]byte("s0m3t1ngt0t4llys3cr3t"))

func init() {
	store.MaxAge(3600 * 6) // Cookie expires after 6 hours
}

/*	HandleAccess returns true if user is authorized to access */
func HandleAccess(r *http.Request) bool {
	// Get connection to database
	db := GetDB()
	defer db.Close()

	// Get session and username from request
	session, _ := store.Get(r, "session")
	username := session.Values["username"]

	// Query junction on permissions and user table
	// TODO: Add logic for regular users, not just Administrators
	var user models.User
	db.Where("username = ?", username).First(&user)

	return user.IsAdmin
}

/* GetCurrentUser returns the current user */
func GetCurrentUser(r *http.Request) interface{} {
	// Get session and username from request
	session, _ := store.Get(r, "session")
	username := session.Values["username"]

	return username
}

/* Check login returns an error if the user could not log in */
func CheckLogin(username string, password string) error {
	// check if username exists in database
	db := GetDB()
	defer db.Close()

	// Get user hash
	user := models.User{}
	db.Where("username = ?", username).Find(&user)

	// Compare the hash and the password for authentication
	err := bcrypt.CompareHashAndPassword(user.Password, []byte(password))
	return err
}

/* GetStore returns the current cookiestore */
func GetStore() *sessions.CookieStore {
	return store
}

/* Returns a list of all admins in database */
func GetAdmins() []int {
	db := GetDB()
	defer db.Close()

	// Get users with 9111 project number
	var admins []int
	rows, err := db.Raw("SELECT user_id FROM users WHERE is_admin = 1").Rows()

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var nextID int
		rows.Scan(&nextID)
		admins = append(admins, nextID)
	}

	return admins
}

/* Handle logging the user out */
func Logout(w http.ResponseWriter, r *http.Request) {
	clearSession(r, w)
}

/* Clear session cookie */
func clearSession(r *http.Request, w http.ResponseWriter) {
	session, _ := store.Get(r, "session")
	session.Values["username"] = ""

	session.Save(r, w)
}
