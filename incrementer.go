package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	VersionFile string
)

func init() {
	flag.StringVar(&VersionFile, "versionFile", "version.go", "Version File")
	flag.Parse()
}

func main() {
	var (
		major        int
		intermediate int
		minor        int
	)
	if FileContent, err := os.ReadFile(VersionFile); err != nil {
		log.Fatal(err)
	} else {
		lines := strings.Split(string(FileContent), "\n")
		for idx, line := range lines {
			if strings.HasPrefix(line, "const version = ") {
				if num, err := fmt.Sscanf(line, `const version = "%d.%d.%d"`, &major, &intermediate, &minor); err != nil {
					log.Fatal(err)
				} else {
					if num != 3 {
						log.Fatal("Failed to find 3 numbers on the version line.")
					} else {
						lines[idx] = fmt.Sprintf(`const version = "%d.%d.%d"`, major, intermediate, minor+1)
					}
				}
			}
		}
		outString := strings.Join(lines, "\n")
		if err := os.WriteFile(VersionFile, []byte(outString), 0666); err != nil {
			log.Fatal(err)
		}
	}
}
