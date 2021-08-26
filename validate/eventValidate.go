package validate

import (
	"shareTravel/model"

	"github.com/go-playground/validator"
)

func EventValidater(event *model.Event) (bool, map[string]string) {

	result := make(map[string]string)
	err := validator.New().Struct(event)

	if err != nil {

		errs := err.(validator.ValidationErrors)

		if len(errs) != 0 {
			for i := range errs {
				switch errs[i].StructField() {
				case "Name":
					switch errs[i].Tag() {
					case "required":
						result["Name"] = "名前は必須入力です"
					}
				case "Date":
					switch errs[i].Tag() {
					case "required":
						result["Date"] = "日付は必須入力です"
					}
				}
			}
		}
		return false, result
	}

	return true, result
}
