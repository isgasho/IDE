// File generated by Gopher Sauce
// DO NOT EDIT!!
package handlers

import (
	methods "github.com/thestrukture/IDE/api/methods"
	"net/http"

	"github.com/gorilla/sessions"
)

//
func GETApiKanban(w http.ResponseWriter, r *http.Request, session *sessions.Session) (response string, callmet bool) {

	pkgName := r.FormValue("pkg")
	response = mResponse(methods.GetKanBan(pkgName))

	callmet = true
	return
}
