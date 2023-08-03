package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const token = "47e2f1ebfa2707" // Token API ipinfo.io

func getIPInfo(ipAddress, token string) (map[string]interface{}, error) {
	url := fmt.Sprintf("https://ipinfo.io/%s/json?token=%s", ipAddress, token)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Failed to make API request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("Failed to read API response: %v", err)
		}

		ipInfo := make(map[string]interface{})
		err = json.Unmarshal(body, &ipInfo)
		if err != nil {
			return nil, fmt.Errorf("Failed to parse API response: %v", err)
		}

		return ipInfo, nil
	}

	return nil, fmt.Errorf("API returned status code %d", resp.StatusCode)
}

func main() {
	fmt.Println("Develop by: LazyySec")

	isList := flag.Bool("l", false, "Check using list. Provide a filename containing a list of IP addresses.")
	host := flag.String("host", "", "Check single IP address. Provide the IP address to check.")
	flag.Parse()

	if *isList {
		filename := flag.Arg(0)
		content, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}

		ipAddresses := strings.Split(string(content), "\n")
		for _, entry := range ipAddresses {
			if entry == "" {
				continue
			}
			ipInfo, err := getIPInfo(entry, token)
			if err != nil {
				fmt.Printf("Failed to retrieve info for IP Address: %s\n", entry)
				continue
			}

			fmt.Printf("IP Address: %s\n", entry)
			fmt.Println("Info:")
			for key, value := range ipInfo {
				fmt.Printf("%s: %v\n", key, value)
			}
			fmt.Println(strings.Repeat("-", 40))
		}
	} else if *host != "" {
		ipInfo, err := getIPInfo(*host, token)
		if err != nil {
			fmt.Printf("Failed to retrieve info for IP Address: %s\n", *host)
			return
		}

		fmt.Printf("IP Address: %s\n", *host)
		fmt.Println("Info:")
		for key, value := range ipInfo {
			fmt.Printf("%s: %v\n", key, value)
		}
		fmt.Println(strings.Repeat("-", 40))
	} else {
		flag.Usage()
	}
}