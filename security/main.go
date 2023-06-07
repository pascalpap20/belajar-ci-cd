package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/pquerna/otp/hotp"
	"golang.org/x/crypto/bcrypt"
	"net/url"
	"time"
)

func main() {
	person := Person{
		Name:  "Jhon Doe",
		Age:   30,
		Email: "jhondoe@gmail.com",
	}

	encoded, err := json.Marshal(person)
	if err != nil {
		fmt.Println("error encding JSON:", err)
	}
	fmt.Println(string(encoded))

	decodingJson(encoded)
	encodingUrl()
	hash()
	integrityHash()

	handler()

	//	jwt
	generateToken()

	hotpImplementation()
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

func decodingJson(encoded []byte) {
	var person Person

	err := json.Unmarshal(encoded, &person)
	if err != nil {
		fmt.Println("err decoding json:", err)
		return
	}

	fmt.Println(person)
}

func encodingUrl() {
	originalURL := "https://example.com/search?q=hello world&category=books"
	encodedURL := url.QueryEscape(originalURL)

	fmt.Println("Original URL:", originalURL)
	fmt.Println("Encoded URL:", encodedURL)
}

func hash() {
	data := "Hello, world!"

	// Membuat objek hash dari algoritma SHA-256
	hash := sha256.New()

	// Mengupdate hash dengan data yang ingin di-hash
	hash.Write([]byte(data))

	// Mengambil nilai hash sebagai array byte
	hashBytes := hash.Sum(nil)

	// Mengubah array byte menjadi representasi heksadesimal
	hashString := hex.EncodeToString(hashBytes)

	fmt.Println("Data:", data)
	fmt.Println("Hash:", hashString)
}

func integrityHash() {
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

func generateToken() {
	// Inisialisasi klaim-klaim yang ingin Anda sertakan dalam token
	claims := jwt.MapClaims{
		"sub":  "1234567890",
		"name": "John Doe",
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(time.Hour * 1).Unix(),
	}

	// Tandatangani token dengan kunci rahasia
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("secret-key"))
	if err != nil {
		// Penanganan kesalahan
		fmt.Println(err)
	}

	// Gunakan signedToken seperti yang Anda butuhkan
	fmt.Println(signedToken)
	verifyToken(signedToken)
}

func verifyToken(tok string) {
	// Token yang diterima
	receivedToken := tok

	// Verifikasi token dengan kunci rahasia
	token, err := jwt.Parse(receivedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte("secret-key"), nil
	})
	if err != nil {
		// Penanganan kesalahan
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Token valid, akses klaim-klaim yang ada
		fmt.Println(claims["sub"], claims["name"])
	} else {
		// Token tidak valid
	}

}

func hotpImplementation() {
	// Membuat kunci rahasia baru untuk HOTP
	key, err := hotp.Generate(hotp.GenerateOpts{
		Issuer:      "MyApp",
		AccountName: "user@example.com",
	})
	if err != nil {
		fmt.Println("Gagal menghasilkan kunci rahasia:", err)
		return
	}

	// Menghasilkan kode sandi HOTP dengan counter 0
	code, _ := hotp.GenerateCode(key.Secret(), 0)
	fmt.Println("Kode Sandi HOTP:", code)

	// Memverifikasi kode sandi HOTP dengan counter 0
	valid, _ := hotp.ValidateCustom(code, 0, key.Secret(), hotp.ValidateOpts{Digits: 6})
	if valid {
		fmt.Println("Kode sandi HOTP valid.")
	} else {
		fmt.Println("Kode sandi HOTP tidak valid.")
	}

}
