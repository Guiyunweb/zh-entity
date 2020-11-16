package service

import (
	"shiki-template/model"
	"shiki-template/utils"
	"strings"
)

var java = `/** zh */
private String entity;
 
`

func JavaConversion(dto *model.EntityDTO) string {
	zh := strings.Split(dto.Body, "\n")
	code := ""
	if targetText, err := utils.Translate(dto.Body); err == nil {
		target := strings.Split(targetText, "\n")
		for i := 0; i < len(zh); i++ {
			code += Template(zh[i], target[i])
		}
	}
	return code
}

func Template(zh string, entity string) string {
	template := strings.Replace(java, "zh", zh, -1)
	template = strings.Replace(template, "entity", utils.UderscoreToLowerCamelCase(entity), -1)
	return template
}
