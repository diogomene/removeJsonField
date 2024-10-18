package forms

import "github.com/charmbracelet/huh"

const OUT_FILE_NAME_PLACEHOLDER = "arquivo.json"

func GetOutputFileName() (string, error) {
	var outputFileName string

	err := huh.NewInput().
		Title("Nome do arquivo de saída").
		Description("Insira o nome do arquivo de saída").
		Value(&outputFileName).
		Placeholder(OUT_FILE_NAME_PLACEHOLDER).
		Run()
	if outputFileName == "" {
		outputFileName = OUT_FILE_NAME_PLACEHOLDER
	}
	return outputFileName, err
}
