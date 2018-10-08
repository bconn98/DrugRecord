/**
File: templates
Description: Executes templates and implements the range for html
@author Bryan Conn
@date: 10/7/2018
 */
package utils

import (
	"net/http"
	. "html/template"
	"reflect"
)

var templateFuncs = FuncMap{"rangeStruct": RangeStructer}
var t = New( "t").Funcs(templateFuncs)
var templates = Must(t.ParseGlob("./web/templates/*.html"))

/**
Function: ExecuteTemplate
Description: Executes a html template
@param w The http writer
@param tmpl The template name
@param data The data for the template to use
*/
func ExecuteTemplate(w http.ResponseWriter, tmpl string, data interface{}){
	err := templates.ExecuteTemplate(w, tmpl, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

/**
Function: RangeStructer
Description: Parses the data sent to a template in a manner that
the template can understand.
@param args A list of arguments
@return An array of interfaces
*/
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