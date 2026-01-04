package utils

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"text/template"

	"github.com/knyazev-ro/perturabo/templates"

	"github.com/fatih/color"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func Contains[T any](arr []T, needle func(T) bool) int {
	for i := range arr {
		if needle(arr[i]) {
			return i
		}
	}
	return -1
}

func Filter[T any](ss []T, callback func(T) bool) (ret []T) {
	for _, s := range ss {
		if callback(s) {
			ret = append(ret, s)
		}
	}
	return
}

func IsInt(s string) bool {
	_, err := strconv.ParseInt(s, 10, 64)
	return err == nil
}

func IsFloat(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func ValidateName(module string) (string, error) {

	module = strings.TrimSpace(module)
	module = strings.ToLower(module)
	println(module)
	pattern := `^[0-9]{4}(?:_[a-z0-9]+)*$`
	re := regexp.MustCompile(pattern)
	if !re.MatchString(module) {
		return "", os.ErrInvalid
	}

	return module, nil
}

func Normalize(name string) (string, string, string, error) {

	name, err := ValidateName(name)
	if err != nil {
		fmt.Println("Filename is invalid")
		return "", "", "", err
	}

	splitName := strings.Split(name, "_")
	number := "_" + splitName[0]
	splitName = splitName[1:]
	for i, s := range splitName {
		splitName[i] = cases.Title(language.English).String(s)
	}
	data := map[string]string{
		"Name":    strings.Join(splitName, "") + number,
		"NameVar": strings.ToLower(splitName[0]) + strings.Join(splitName[1:], "") + number,
	}

	return data["Name"], data["NameVar"], name, nil
}

func ParseTemplate(templatePath string, outputFilePath string, data interface{}, args []string) (string, error) {

	isForce := Contains(args, func(x string) bool {
		return x == "--force"
	}) >= 0

	//check if output directory exists, if it exists then nothing
	if _, err := os.Stat(outputFilePath); !os.IsNotExist(err) && !isForce {
		WarningPrintln("Warning: output file already exists:", outputFilePath)
		return outputFilePath, os.ErrExist
	}

	tmpl, err := template.ParseFS(templates.TemplatesFS, templatePath)
	if err != nil {
		ErrorPrintln("Error parsing template:", err.Error())
		return "", err
	}

	out, err := os.Create(outputFilePath)
	if err != nil {
		ErrorPrintln("Error creating:", err.Error())
		return "", err
	}

	defer out.Close()
	err = tmpl.Execute(out, data)
	if err != nil {
		ErrorPrintln("Error executing template:", err.Error())
		return "", err
	}

	return outputFilePath, nil
}

func ErrorPrintln(a ...any) {
	c := color.New(color.BgRed)
	c.Println(a...)
}

func SuccessPrintln(a ...any) {
	c := color.New(color.BgGreen)
	c.Println(a...)
}

func WarningPrintln(a ...any) {
	c := color.New(color.BgYellow)
	c.Println(a...)
}
