package main

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

)

// Command line flags
var (
	target       = flag.String("t", "", "Target domain e.g www.google.com (required) ")
	username     = flag.String("u", "", "Username ")
	outputFile   = flag.String("o", "", "Output file name. (optional)")
	templateFile = flag.String("r", "", "Template file to process.")
	program      = flag.String("p", "", "Program name. ")
	researcher   = flag.String("re", "", "Researcher name. ")
)

// whoIs return the whois output of the target
func whoIs() ([]byte, error) {
	out, _ := exec.Command("whois", *target).Output()
	return out, nil
}

// nameServers return output of "dig NS @8.8.8.8 -trace"
func nameServers() ([]byte, error) {
	out, err := exec.Command("dig", "NS", "@8.8.8.8", "trace").Output()
	return out, err
}

// digTarget return the output of the dig command for target
func digTarget() ([]byte, error) {
	out, err := exec.Command("dig", *target).Output()
	return out, err
}

// dig-TXT-Target return the output of the dig command for target
func digTXTTarget() ([]byte, error) {
	out, err := exec.Command("dig", "TXT", *target).Output()
	return out, err
}

// curlTarget return the output of the dig command for target
func curlTarget() ([]byte, error) {
	out, err := exec.Command("curl", *target).Output()
	return out, err
}

// sha256Username returns the SHA256 value of username
func sha256Username() []byte {
	hash := sha256.New()
	hash.Write([]byte(*username))
	md := hash.Sum(nil)
	return []byte(fmt.Sprintf("%x", md))
}

// validateFlags validates if all the required flags are set
func validateFlags() error {
	var err []string

	if *target == "" {
		err = append(err, "Target is required.")
	}
	if *templateFile == "" {
		err = append(err, "Template file path is required.")
	}
	if len(err) > 0 {
		return errors.New(strings.Join(err, "\n"))
	}
	return nil

}

func inputFlags(content []byte) {
	reader := bufio.NewReader(os.Stdin)
	if *username == "" && bytes.Contains(content, []byte("_username_")) {
		fmt.Print("A _username_ is used in this template, please specify: ")
		*username, _ = reader.ReadString('\n')
		*username = strings.TrimSpace(*username)
	}
	if *researcher == "" && bytes.Contains(content, []byte("_researcher_")) {
		fmt.Print("A _researcher_ is used in this template, please specify: ")
		*researcher, _ = reader.ReadString('\n')
		*researcher = strings.TrimSpace(*researcher)

	}
	if *program == "" && bytes.Contains(content, []byte("_program_")) {
		fmt.Print("A _program_ is used in this template, please specify: ")
		*program, _ = reader.ReadString('\n')
		*program = strings.TrimSpace(*program)
	}
}

// checkError if error is not nil and exiting if error found
func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}


func main() {
	flag.Parse()

	err := validateFlags()
	checkError(err)

	// Reading template file
	content, err := ioutil.ReadFile(*templateFile)
	checkError(err)

	inputFlags(content)

	// Replacing the values
	content = bytes.Replace(content, []byte("_target_"), []byte(*target), -1)
	content = bytes.Replace(content, []byte("_username_"), []byte(*username), -1)
	content = bytes.Replace(content, []byte("_program_"), []byte(*program), -1)
	content = bytes.Replace(content, []byte("_researcher_"), []byte(*researcher), -1)

	if bytes.Contains(content, []byte("_whois_")) {
		whoIsTarget, err := whoIs()
		checkError(err)
		content = bytes.Replace(content, []byte("_whois_"), whoIsTarget, -1)
	}

	if bytes.Contains(content, []byte("_dig_")) {
		digTarget, err := digTarget()
		checkError(err)
		content = bytes.Replace(content, []byte("_dig_"), digTarget, -1)
	}
	
	if bytes.Contains(content, []byte("_dig-txt_")) {
		digTXTTarget, err := digTXTTarget()
		checkError(err)
		content = bytes.Replace(content, []byte("_dig-txt_"), digTXTTarget, -1)
	}

	if bytes.Contains(content, []byte("_nameservers_")) {
		nameServerTarget, err := nameServers()
		checkError(err)
		content = bytes.Replace(content, []byte("_nameservers_"), nameServerTarget, -1)
	}

	if bytes.Contains(content, []byte("_curl_")) {
		curl, err := curlTarget()
		checkError(err)
		content = bytes.Replace(content, []byte("_curl_"), curl, -1)
	}

	if bytes.Contains(content, []byte("_sha_")) {
		md := sha256Username()
		content = bytes.Replace(content, []byte("_sha_"), md, -1)
	}
	// If output flag is set
	if *outputFile != "" {
		ioutil.WriteFile(*outputFile, content, 0677)
		fmt.Printf("File saved to %v", *outputFile)
	} else {
		// Printing the data to console
		fmt.Printf("%s\n", content)

	}
}
