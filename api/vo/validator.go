package vo

import (
	"github.com/martini-contrib/binding"
	"net/http"
	"strings"
)

func (af ArticleForm) Validate(errors binding.Errors, req *http.Request) binding.Errors {
	if strings.TrimSpace(af.Content) == "" {
		errors = append(errors, binding.Error{
			FieldNames:     []string{"content"},
			Classification: "EmptyError",
			Message:        "content can`t be empty",
		})
	}
	return errors
}
