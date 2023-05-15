package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"

	plistlib "howett.net/plist"
)

func main() {
	home, _ := os.UserHomeDir()
	plistFile := filepath.Join(home, "Library/Preferences/com.navicat.NavicatPremium.plist")
	supportFileDir := filepath.Join(home, "Library/Application Support/PremiumSoft CyberTech/Navicat CC/Navicat Premium")
	plistData, _ := os.ReadFile(plistFile)
	plist := make(map[string]interface{})
	plistlib.Unmarshal(plistData, &plist)
	trialKey := ""
	re := regexp.MustCompile(`[0-9A-Z]{32}`)
	for k := range plist {
		if matched := re.MatchString(k); matched {
			trialKey = k
			break
		}
	}
	if trialKey != "" {
		_, err := exec.Command("/usr/libexec/PlistBuddy", "-c", "Delete "+trialKey, plistFile).Output()
		if err != nil {
			fmt.Printf("delete trial key %s in %s failed: %s\n", trialKey, plistFile, err)
		} else {
			fmt.Printf("delete trial key %s in %s\n", trialKey, plistFile)
		}
	}
	re = regexp.MustCompile(`\.[0-9A-Z]{32}`)
	supportFiles, _ := os.ReadDir(supportFileDir)
	for _, entry := range supportFiles {
		if matched := re.MatchString(entry.Name()); matched {
			supportFile := filepath.Join(supportFileDir, entry.Name())
			err := os.Remove(supportFile)
			if err != nil {
				fmt.Printf("remove support file %s failed: %s\n", supportFile, err)
			} else {
				fmt.Printf("remove support file %s\n", supportFile)
			}
		}
	}
	fmt.Println("reset trial ok.")
}
