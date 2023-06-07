package security

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"net/url"
)

func main() {
	decodingJson()
}

type Person struct {
	Name  string
	Age   int
	Email string
}

type User struct {
	Username string
	Password string
	Salt     string
}

func decodingJson() {
	person := Person{
		Name:  "John Doe",
		Age:   30,
		Email: "johndoe@gmail.com",
	}

	encdoded, err := json.Marshal(person)
	if err != nil {
		fmt.Println("error encoding JSON:", err)
	}
	fmt.Println(string(encdoded))
}

func encodeUrl() {
	originalURL := "https://example.com/search?q=hello world&category=books"
	encodedURL := url.QueryEscape(originalURL)

	fmt.Println("Original URL:", originalURL)
	fmt.Println("Encoded URL:", encodedURL)
}

func compareHas() {
	// Data awal
	data := "Hello, world!"

	// Hash data awal
	initialHash := generateHash(data)

	// Simulasikan perubahan data
	modifiedData := "Hello, modified!"

	// Hash data yang diubah
	modifiedHash := generateHash(modifiedData)

	// Verifikasi integritas data
	isValid := verifyIntegrity(data, initialHash)
	fmt.Println("Data integrity:", isValid) // Output: true

	isValid = verifyIntegrity(modifiedData, initialHash)
	fmt.Println("Data integrity:", isValid) // Output: false

	isValid = verifyIntegrity(modifiedData, modifiedHash)
	fmt.Println("Data integrity:", isValid) // Output: true

}

func generateHash(data string) string {
	// Membuat objek hash dari algoritma SHA-256
	hash := sha256.New()

	// Mengupdate hash dengan data yang ingin di-hash
	hash.Write([]byte(data))

	// Mengambil nilai hash sebagai array byte
	hashBytes := hash.Sum(nil)

	// Mengubah array byte menjadi representasi heksadesimal
	hashString := fmt.Sprintf("%x", hashBytes)

	return hashString
}

func verifyIntegrity(data, expectedHash string) bool {
	// Menghasilkan hash dari data yang diberikan
	hash := generateHash(data)

	// Membandingkan hash yang dihasilkan dengan hash yang diharapkan
	return hash == expectedHash
}

func handler() {
	// Contoh data pengguna yang didaftarkan
	registeredUser := User{
		Username: "john_doe",
		Password: "password123",
	}

	// Proses pendaftaran pengguna
	RegisterUser(&registeredUser)

	// Contoh proses login
	loginUsername := "john_doe"
	loginPassword := "password123"

	// Verifikasi login
	isValid := VerifyLogin(loginUsername, loginPassword, registeredUser)
	fmt.Println("Login valid:", isValid) // Output: true

	// Contoh login dengan password yang salah
	invalidPassword := "wrongpassword"
	isValid = VerifyLogin(loginUsername, invalidPassword, registeredUser)
	fmt.Println("Login valid:", isValid) // Output: false

}

func RegisterUser(user *User) {
	// Generate salt
	salt := generateSalt()

	// Combine password and salt
	passwordWithSalt := []byte(user.Password + salt)

	// Hash the password + salt combination
	hashedPassword, _ := bcrypt.GenerateFromPassword(passwordWithSalt, bcrypt.DefaultCost)

	// Update user data
	user.Password = string(hashedPassword)
	user.Salt = salt
}

func VerifyLogin(username, password string, user User) bool {
	// Combine password and salt
	passwordWithSalt := []byte(password + user.Salt)

	// Hash the password + salt combination
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), passwordWithSalt)

	return err == nil
}

func generateSalt() string {
	// Generate random salt using cryptographic randomness
	salt := make([]byte, 16)
	rand.Read(salt)

	return hex.EncodeToString(salt)
}
