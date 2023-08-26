package filereader

import "os"

func ReadTextFile(fileName string) (string, error) {
	contents, err := os.ReadFile(fileName)
	if err != nil {
		// handle error
		// couldn't read the file
		return "", err
	} else {
		return string(contents), nil
	}
}
