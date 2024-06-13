package output

import (
	"encoding/json"
	"os"
)

func WriteToFileOrStdout(report any, outputFilePath string) error {
	fileData, err := json.Marshal(report)
	if err != nil {
		return err
	}

	if outputFilePath == "" {
		_, err = os.Stdout.Write(fileData)
		if err != nil {
			return err
		}
	} else {
		err = os.WriteFile(outputFilePath, fileData, 0644)
		if err != nil {
			return err
		}
	}
	return nil
}
