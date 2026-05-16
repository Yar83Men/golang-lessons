package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"

	"github.com/fatih/color"
)

func main() {
	file, err := os.Create("file.txt")
	var errorClosingFile error
	defer file.Close()
	if err != nil {
		errorClosingFile = file.Close()
		errorClosingFileHandler(errorClosingFile)
		color.Red(err.Error())
	}

	file.WriteString("Hello\nGolang\nand Java")
	errorClosingFile = file.Close()
	errorClosingFileHandler(errorClosingFile)
	color.Green("Ok %s", file.Name())

	text, err := os.ReadFile("file.txt")
	if err != nil {
		errorClosingFileHandler(err)
	}

	color.Cyan(string(text))

	person := Person{
		Name:  "Ivan",
		Age:   12,
		Email: "ivan@gmail.com",
	}

	person1 := Person{
		Name:  "John",
		Age:   33,
		Email: "john@gmail.com",
	}

	person2 := Person{
		Name:  "Victor",
		Age:   45,
		Email: "victor@gmail.com",
	}

	tagLogin, _ := reflect.TypeOf(&person).Elem().FieldByName("Name")
	color.Magenta(string(tagLogin.Tag))

	file1, err := person.ToBytes()
	if err != nil {
		color.Red("Не удалось преобразовать в JSON")
	}

	var fileErr error
	if !FindFile("my-json.json") {
		fileErr = os.WriteFile("my-json.json", file1, os.FileMode(os.O_RDWR))
	}

	if fileErr != nil {
		color.Red("Ошибка записи JSON %s", fileErr.Error())
	}

	vault := NewVault()

	pesonsSlice := make([]Person, 0)
	pesonsSlice = append(pesonsSlice, person, person1, person2)
	vault.addPersons(pesonsSlice)

	vaultFile, err := vault.ToBytes()
	if err != nil {
		color.Red("Ошибка записи JSON %s", fileErr.Error())
	}

	var vaultCreatingErr error
	if !FindFile("vault.json") {
		vaultCreatingErr = os.WriteFile("vault.json", vaultFile, os.ModeAppend)
	}

	if vaultCreatingErr != nil {
		color.Red("Ошибка записи JSON %s", fileErr.Error())
	}

	var newVault Vault
	readFile, err := ReadFile("vault.json")
	if err != nil {
		color.Red(err.Error())
	}

	unMarshalErr := json.Unmarshal(readFile, &newVault)
	if unMarshalErr != nil {
		color.Red(unMarshalErr.Error())
	}

	fmt.Println(newVault.Persons)
}

func errorClosingFileHandler(err error) {
	if err != nil {
		color.Red(err.Error())
	}
}

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func (person *Person) ToBytes() ([]byte, error) {
	file, err := json.Marshal(person)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func FindFile(name string) bool {
	file, err := os.Stat(name)
	if err != nil {
		color.Red(err.Error())
		return false
	}
	color.Green("Файл %s обнаружен", file.Name())
	return file.Name() != ""
}

func ReadFile(name string) ([]byte, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	return data, nil
}
