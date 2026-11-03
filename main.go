package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
)


type MenuNode struct {
	Name    string    `json:"name"`   // Название, которое отображается в Rofi
	IsScript bool     `json:"isScript"`// false — папка, true — скрипт
	Content []string  `json:"content"`// Если папка — список имен детей; если скрипт — [0] элемент содержит команду
}

func initMenu() map[string]MenuNode {
	config, err := os.ReadFile("config.json")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	m := make(map[string]MenuNode)	
	
	if err = json.Unmarshal(config, &m); err != nil {
		log.Fatalf("json pars error: %v", err)
	}

	return m
}



func main() {
	logFile, err := os.OpenFile("/tmp/rofi_debug.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
    	log.SetOutput(logFile) 
    	defer logFile.Close()
	}

	menuMap := initMenu()

	var targetPath string

	if os.Getenv("ROFI_RETV") == "0" {
		targetPath = "/"
	} else {
		targetPath = path.Join(os.Getenv("ROFI_DATA"), os.Args[1])
	}
	
	node, exists := menuMap[targetPath]
	if !exists {
		log.Printf("Не соответствие json и rofi. Проверить имена в json.")
		return
	}

	if !node.IsScript {
		fmt.Printf("\x00prompt\x1f%s\n", targetPath)
		fmt.Printf("\x00clear\x1ftrue\n")
		fmt.Printf("\x00data\x1f%s\n", targetPath)
		for _, value := range node.Content {
			fmt.Println(value)
		}

	} else {
		cmd := exec.Command("bash", "-c", node.Content[0])
		if err := cmd.Start(); err != nil {
			log.Printf("Ошибка в скрипте: %v", err)
		}
		return
	}
}


// сделать оповищение при выполнении скрипта или при невыполнении через notify-send
// доделать нормальное логирование в файл. 
