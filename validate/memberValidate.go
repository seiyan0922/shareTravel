package validate

import (
	"shareTravel/model"

	"github.com/go-playground/validator"
)

func MemberValidater(member *model.Member) (bool, map[string]string) {

	result := make(map[string]string)
	err := validator.New().Struct(member)

	if err != nil {

		errs := err.(validator.ValidationErrors)

		if len(errs) != 0 {
			for i := range errs {
				switch errs[i].Tag() {
				case "required":
					result["Name"] = "名前は必須入力です"
				case "max":
					result["MAX"] = "12文字以内で入力してください。"
				}

			}
		}
		return false, result
	}

	return true, result
}
