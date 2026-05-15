package savefile

import (
	"log"
	"os"
)

// BkpMemory create a txt file to save memory
func BkpMemory(file, memory string) error {

	f, err := os.Create(file)
	defer f.Close()
	if err != nil {
		log.Println(err)
	}
	b := []byte(memory)
	f.Write(b)
	
	return nil
}
