package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/BurntSushi/toml"
)

var cmdPtr = flag.String("cmd", "dmenu -i -l 20", "Launcher command to use (dmenu, rofi) including options")
var snipfilePtr = flag.String("snips", "snippets.toml", "Path to snippets file")
var editorPtr = flag.String("editor", "vim", "Text editor")

// Snippets - list of Snippets
type Snippets struct {
	Snippets []SnippetInfo `toml:"snippets"`
}

// SnippetInfo - individual Snippet struct
type SnippetInfo struct {
	Description string   `toml:"description"`
	Command     string   `toml:"command"`
	Tag         []string `toml:"tag"`
	Output      string   `toml:"output"`
}

func expandPath(s string) string {
	if len(s) >= 2 && s[0] == '~' && os.IsPathSeparator(s[1]) {
		if runtime.GOOS == "windows" {
			s = filepath.Join(os.Getenv("USERPROFILE"), s[2:])
		} else {
			s = filepath.Join(os.Getenv("HOME"), s[2:])
		}
	}
	return os.Expand(s, os.Getenv)
}

func isCommandAvailable(name string) bool {
	cmd := exec.Command("/bin/sh", "-c", "command -v "+name)
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

func (snippets *Snippets) load() error {
	if _, err := os.Stat(*snipfilePtr); os.IsNotExist(err) {
		log.Fatal(err.Error())
	}
	if _, err := toml.DecodeFile(*snipfilePtr, snippets); err != nil {
		return fmt.Errorf("Failed to load snippet file. %v", err)
	}
	return nil
}

func (snippets *Snippets) save() error {
	// Save saves the snippets to toml file.
	f, err := os.Create(*snipfilePtr)
	defer f.Close()
	if err != nil {
		return fmt.Errorf("Failed to save snippet file. err: %s", err)
	}
	return toml.NewEncoder(f).Encode(snippets)
}

func (snippets *Snippets) toString() (string, error) {
	// toString returns the contents of toml file.
	var buffer bytes.Buffer
	err := toml.NewEncoder(&buffer).Encode(snippets)
	if err != nil {
		return "", fmt.Errorf("Failed to convert struct to TOML string: %v", err)
	}
	return buffer.String(), nil
}

func main() {
	flag.Parse()
	var snippets Snippets

	if err := snippets.load(); err != nil {
		log.Fatal(err.Error())
	}
	snips, err := snippets.toString()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(snips)
	//for edit := true; edit; {
	//	displayList, m := createMenu(&tasklist, false)
	//	out, _ := display(displayList.String(), *todoPtr)
	//	switch {
	//	case out == "Add Item":
	//		addItem(&tasklist)
	//	case out != "":
	//		t, _ := tasklist.GetTask(m[out])
	//		editItem(t, &tasklist)
	//	default:
	//		edit = false
	//	}
	//}
	//if *archPtr {
	//	archiveDone(&tasklist)
	//}
	//if err := todotxt.WriteToFilename(&tasklist, *todoPtr); err != nil {
	//	log.Fatal(err.Error())
	//}
}
