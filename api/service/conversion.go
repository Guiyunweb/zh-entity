package service

import (
	"shiki-template/model"
	"shiki-template/utils"
	"strings"
)

var java = `/** zh */
private String entity;
 
`

var sql = "ADD COLUMN `${entity}` VARCHAR(255) COMMENT '${zh}'"

func JavaConversion(dto *model.EntityDTO) string {
	zh := strings.Split(dto.Body, "\n")
	code := ""
	if targetText, err := utils.Translate(dto.Body); err == nil {
		target := strings.Split(targetText, "\n")
		for i := 0; i < len(zh); i++ {
			code += JavaTemplate(zh[i], target[i])
		}
	}
	return code
}

func SQLConversion(dto *model.EntityDTO) string {
	zh := strings.Split(dto.Body, "\n")
	code := "ALTER TABLE [表名]\n"
	if targetText, err := utils.Translate(dto.Body); err == nil {
		target := strings.Split(targetText, "\n")
		for i := 0; i < len(zh); i++ {
			code += SQLTemplate(zh[i], target[i],dto.Type)
			if i==len(zh) -1 {
				code += ";\n"
			}else {
				code += ",\n"
			}
		}
	}
	return code
}

func JavaTemplate(zh string, entity string) string {
	template := strings.Replace(java, "zh", zh, -1)
	template = strings.Replace(template, "entity", utils.UderscoreToLowerCamelCase(entity), -1)
	return template
}

func SQLTemplate(zh string, entity string,types string) string {
	template := strings.Replace(sql, "${zh}", zh, -1)
	if types == "1" {
		template = strings.Replace(template, "${entity}", utils.WordToLowerUdnderscore(entity), -1)
	}else {
		template = strings.Replace(template, "${entity}", utils.WordToUpperUdnderscore(entity), -1)
	}
	return template
}
