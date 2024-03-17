/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package kratos

import (
	"net/http"

	oryKratos "github.com/ory/kratos-client-go"
)

type KratosSession struct {
	CurrentUser         *oryKratos.Identity
	CurrentSession      *oryKratos.Session
	CurrentCookie       *http.Cookie
	CurrentSessionToken *string
}

func (s *KratosSession) SetSession(session *oryKratos.Session) {
	s.SetCurrentSession(session)
	s.SetCurrentUser(&session.Identity)
}

func (s *KratosSession) SetCurrentUser(user *oryKratos.Identity) {
	s.CurrentUser = user
}

func (s *KratosSession) SetCurrentSession(session *oryKratos.Session) {
	s.CurrentSession = session
}

func (s *KratosSession) SetCurrentCookie(cookie *http.Cookie) {
	s.CurrentCookie = cookie
}

func (s *KratosSession) SetSessionToken(token *string) {
	s.CurrentSessionToken = token
}
