package main

import (
	"time"
	"strings"
	"os/exec"
	"fmt"
	"io/ioutil"
	"net/http"
)

const serverAddr = ""

func discoverIpAddr() (ipAddr string, err error) {
	cmd := "ip -4 addr show tun0 | grep -oP '(?<=inet\\s)\\d+(\\.\\d+){3}'"
	stdout, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(stdout)), nil
}

func updateServer(ipAddr string) (response string, err error) {
	resp, err := http.Post(fmt.Sprintf("%s/%s", serverAddr, ipAddr), "application.json", nil)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func main() {

	for {
		ipAddr, err := discoverIpAddr()
		if err != nil {
			fmt.Printf("Error retrieving tun0 ip addr. %s\n", err)
			time.Sleep(30 * time.Second)
			continue
		}

		resp, err := updateServer(ipAddr)
		if err != nil {
			fmt.Printf("Error updating server. %s\n", err)
			time.Sleep(30 * time.Second)
			continue;
		}

		fmt.Println(resp)

		time.Sleep(30 * time.Second)
	}
}
