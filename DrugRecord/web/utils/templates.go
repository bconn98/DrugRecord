package utils

import (
	"net/http"
	. "html/template"
	"reflect"
)

var templateFuncs = FuncMap{"rangeStruct": RangeStructer}
var t = New( "t").Funcs(templateFuncs)
var templates = Must(t.ParseGlob("./web/templates/*.html"))

func ExecuteTemplate(w http.ResponseWriter, tmpl string, data interface{}){
	err := templates.ExecuteTemplate(w, tmpl, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}


func RangeStructer(args ...interface{}) []interface{} {
	if len(args) == 0 {
		return nil
	}

	v := reflect.ValueOf(args[0])
	if v.Kind() != reflect.Struct {
		return nil
	}

	out := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		out[i] = v.Field(i).Interface()
	}

	return out
}