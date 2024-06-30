package file_manager

import (
	"bufio"
	"encoding/json"
	"os"
)

type FileManager struct{
	InputFilePath string
	OutputFilePath string
}

func (io *FileManager) ReadLines()([]string, error){
	file, err := os.Open(io.InputFilePath)
	
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan(){
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	
	if err != nil{
		file.Close()
		return nil, err
	}

	file.Close()
	
	return lines, nil

}

func (io *FileManager) WriteLines(data interface{})(error){
	file, err := os.Create(io.OutputFilePath)

	if err != nil  {
		return err
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")

	err =  encoder.Encode(data)

	if err != nil{
		file.Close()
		return err
	}

	file.Close()
	return nil 
}


func New(inputFilePath string, outputFilePath string) (*FileManager){
	return &FileManager{
		InputFilePath: inputFilePath,
		OutputFilePath: outputFilePath,
	}
}