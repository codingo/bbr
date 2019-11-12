package main

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
)

// Command line flags
var (
	target       = flag.String("t", "", "Target domain e.g www.google.com (required)")
	username     = flag.String("u", "", "Username (required)")
	outputFile   = flag.String("o", "", "Output file name. (optional)")
	templateFile = flag.String("r", "", "Template file to process. (required)")
)

// whoIs return the whois output of the _target_
func whoIs() ([]byte, error) {
	out, _ := exec.Command("whois", *target).Output()
	return out, nil
}

// nameServers return output of "dig NS @8.8.8.8 +trace"
func nameServers() ([]byte, error) {
	out, err := exec.Command("dig", "NS", "@8.8.8.8", "trace").Output()
	return out, err
}

// digTarget return the output of the dig command for target
func digTarget() ([]byte, error) {
	out, err := exec.Command("dig", *target).Output()
	return out, err
}

// sha256Username returns the SHA256 value of username
func sha256Username() []byte {
	hash := sha256.New()
	hash.Write([]byte(*username))
	md := hash.Sum(nil)
	return []byte(fmt.Sprintf("%x", md))
}

// curlTarget returns the curl output for target
func curlTarget() ([]byte, error) {
        out, err := exec.Command("curl", *target).Output()
        return out, err
}

func validateFlags() error {
	var err []string

	if *target == "" {
		err = append(err, "Target is required.")
	}
	if *username == "" {
		err = append(err, "Username is required.")
	}
	if *templateFile == "" {
		err = append(err, "Template file path is required.")
	}
	if len(err) > 0 {
		return errors.New(strings.Join(err, "\n"))
	}
	return nil

}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	flag.Parse()

	err := validateFlags()
	checkError(err)

	// read template file
	content, err := ioutil.ReadFile(*templateFile)
	checkError(err)

	whoIsTarget, err := whoIs()
	checkError(err)

	digTarget, err := digTarget()
	checkError(err)

	nameServerTarget, err := nameServers()
	checkError(err)

        curlTarget, err := curlTarget()
        checkError(err)

	md := sha256Username()

	content = bytes.Replace(content, []byte("_whois_"), whoIsTarget, -1)
	content = bytes.Replace(content, []byte("_target_"), []byte(*target), -1)
	content = bytes.Replace(content, []byte("_dig_"), digTarget, -1)
	content = bytes.Replace(content, []byte("_sha_"), md, -1)
	content = bytes.Replace(content, []byte("_username_"), []byte(*username), -1)
	content = bytes.Replace(content, []byte("_nameservers_"), nameServerTarget, -1)
        content = bytes.Replace(content, []byte("_curl"), curlTarget, -1)

	fmt.Printf("%s", content)

	// If output flag is set
	if *outputFile != "" {
		ioutil.WriteFile(*outputFile, content, 0677)
	}
}
