// File generated by Gopher Sauce
// DO NOT EDIT!!
package templates

import (
	"github.com/thestrukture/IDE/types"
	"log"
)

// Template path
var templateIDNavCustom = "tmpl/ui/navbars.tmpl"

//
// Renders HTML of template
// NavCustom with struct types.Navbars
func NavCustom(d types.Navbars) string {
	return netbNavCustom(d)
}

// Render template with JSON string as
// data.
func netNavCustom(args ...interface{}) string {

	// Get data from JSON
	var d = netcNavCustom(args...)
	return netbNavCustom(d)

}

// template render function
func netbNavCustom(d types.Navbars) string {
	localid := templateIDNavCustom
	name := "NavCustom"
	defer templateRecovery(name, localid, &d)

	// render and return template result
	return executeTemplate(name, localid, &d)
}

// Unmarshal a json string to the template's struct
// type
func netcNavCustom(args ...interface{}) (d types.Navbars) {

	if len(args) > 0 {
		jsonData := args[0].(string)
		err := parseJSON(jsonData, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	}

	return
}
