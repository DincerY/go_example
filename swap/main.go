package main

import (
	"fmt"
	"os/exec"
)

func main() {

	cmd := exec.Command("cmd", "/C", "dir")

	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}

	fmt.Println("Çıktı:")
	fmt.Println(string(output))
}
