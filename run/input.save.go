package run

import (
	"blueprintz/recognize"
	"errors"
	"fmt"
	"github.com/gearboxworks/go-status"
	"github.com/gearboxworks/go-status/is"
	"github.com/manifoldco/promptui"
	"strings"
)

//
// Name: Blueprintz for WordPress
//
// Copyright (C) 2019 NewClarity Consulting LLC
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.
//

type pepper struct {
	Name     string
	HeatUnit int
	Peppers  int
}

func _Input() (sts Status) {
	peppers := []pepper{
		{Name: "Bell Pepper", HeatUnit: 0, Peppers: 0},
		{Name: "Banana Pepper", HeatUnit: 100, Peppers: 1},
		{Name: "Poblano", HeatUnit: 1000, Peppers: 2},
		{Name: "Jalapeño", HeatUnit: 3500, Peppers: 3},
		{Name: "Aleppo", HeatUnit: 10000, Peppers: 4},
		{Name: "Tabasco", HeatUnit: 30000, Peppers: 5},
		{Name: "Malagueta", HeatUnit: 50000, Peppers: 6},
		{Name: "Habanero", HeatUnit: 100000, Peppers: 7},
		{Name: "Red Savina Habanero", HeatUnit: 350000, Peppers: 8},
		{Name: "Dragon’s Breath", HeatUnit: 855000, Peppers: 9},
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U0001F336 {{ .Name | cyan }} ({{ .HeatUnit | red }})",
		Inactive: "  {{ .Name | cyan }} ({{ .HeatUnit | red }})",
		Selected: "\U0001F336 {{ .Name | red | cyan }}",
		Details: `
--------- Pepper ----------
{{ "Name:" | faint }}	{{ .Name }}
{{ "Heat Unit:" | faint }}	{{ .HeatUnit }}
{{ "Peppers:" | faint }}	{{ .Peppers }}`,
	}

	searcher := func(input string, index int) bool {
		pepper := peppers[index]
		name := strings.Replace(strings.ToLower(pepper.Name), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}

	prompt := promptui.Select{
		Label:     "Spicy Level",
		Items:     peppers,
		Templates: templates,
		Size:      4,
		Searcher:  searcher,
	}

	i, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return status.Wrap(err)
	}

	fmt.Printf("You choose number %d: %s\n", i+1, peppers[i].Name)

	return nil
}

var prompt promptui.Prompt

func Prompt() (sts Status) {

	for {
		prompt = promptui.Prompt{
			Label:   "Please enter project's repository URL",
			Default: "https://github.com/...",
			Templates: &promptui.PromptTemplates{
				Prompt:  "{{ . }}: ",
				Valid:   "{{ . | bold }}: ",
				Invalid: "{{ . | red }}: ",
				Success: "{{ . }}: ",
			},
		}
		repourl, err := prompt.Run()
		if err != nil {
			if err.Error() == "^C" {
				err = errors.New("user terminated with ^C")
				break
			}
			err = errors.New(fmt.Sprintf("Prompt failed: %v\n", err))
			break
		}
		fmt.Printf(promptui.Styler(promptui.FGMagenta, promptui.FGBold)("Verifying '%s'..."),
			repourl,
		)
		sts = recognize.VerifyUrl(repourl)
		if is.Success(sts) {
			fmt.Println(promptui.Styler(promptui.FGGreen, promptui.FGBold)("\nSuccess!"))
			sts = nil
			break
		}
		fmt.Printf(promptui.Styler(promptui.FGYellow, promptui.FGBold)("\nYour input '%s' does not appear to be a valid URL.\nPlease try again:\n\n"),
			repourl,
		)
	}

	return sts
}
