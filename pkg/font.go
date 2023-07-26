package pkg

import (
	"errors"
	"log"
	"os"
	"strings"
)

func Font(banner, text string) (string, int, error) {
	if len(text) == 0 {
		return "", 1, nil
	}

	if banner != "Standard" && banner != "Shadow" && banner != "Thinkertoy" {
		return "", 2, errors.New("Error: " + banner + " not exits")
	}

	for i := 0; i < len(text); i++ {
		if (text[i] < 32 || text[i] > 126) && text[i] != 10 && text[i] != 9 && text[i] != 13 {
			return "", 1, errors.New("Error: the text contains not ASCII Table characters")
		}
	}

	file, err := os.ReadFile("./pkg/banners/" + strings.ToLower(banner) + ".txt")
	if err != nil {
		log.Println(err.Error())
		return "", 2, err
	}

	fileHash := MD5(string(file))
	if !CheckHash(fileHash, banner) {
		return "", 2, errors.New("Error: hash code of banner has been changed")
	}

	return Logic(banner, text, file), 1, nil
}
