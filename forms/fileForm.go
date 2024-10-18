package forms

import (
	"github.com/charmbracelet/huh"
)

func GetJsonFilePath(jsonFilePath *string) error {
	var useFileExplorer bool
	huh.NewConfirm().
		Title("Utilizar explorador de arquivo?").
		Description("Deseja utilizar o explorador de arquivo? Caso não, deverá inserir o caminho do arquivo").
		Affirmative("Sim, utilizar explorador").
		Negative("Não, inserir caminho").
		Value(&useFileExplorer).
		Run()

	if useFileExplorer {
		return huh.NewForm(
			huh.NewGroup(
				huh.NewFilePicker().
					Value(jsonFilePath).
					Description("Selecione o arquivo JSON que deseja modificar").
					Title("Arquivo JSON").
					AllowedTypes([]string{".json"}).
					ShowSize(true).
					ShowPermissions(false).
					Height(7),
			),
		).WithHeight(8).Run()
	}
	return huh.NewInput().
		Title("Insira o caminho do arquivo").
		Description("Informe o arquivo do arquivo de input").
		Placeholder("C:\\...\\").
		Value(jsonFilePath).
		Run()
}
