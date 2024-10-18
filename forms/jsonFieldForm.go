package forms

import (
	"github.com/charmbracelet/huh"
	"github.com/diogomene/removeJsonField/entities"
)

func DeleteJsonFieldsForm(jsonFile entities.JsonFile) ([]string, error) {

	var options []huh.Option[string]
	for _, field := range jsonFile.Fields {
		options = append(options, huh.NewOption(field, field))
	}

	var selected []string

	err := huh.NewForm(
		huh.NewGroup(
			huh.NewMultiSelect[string]().
				Title("Campos a serem removidos").
				Description("Selecione os campos que deseja remover do arquivo json").
				Options(options...).
				Value(&selected),
		),
	).WithHeight(8).Run()
	if err != nil {
		return nil, err
	}

	return selected, nil
}
