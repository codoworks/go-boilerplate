/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package kratos

import (
	"context"
	"errors"
	"net/http"
	"reflect"

	"github.com/codoworks/go-boilerplate/pkg/config/features"

	oryKratos "github.com/ory/kratos-client-go"
)

var kratosClient *Kratos

type Kratos struct {
	name    string
	config  features.KratosConfig
	Public  *oryKratos.APIClient
	admin   *oryKratos.APIClient
	Session *KratosSession
}

func (k *Kratos) Name() string {
	return k.name
}

func (k *Kratos) Configure(v any) {
	k.config = v.(reflect.Value).Interface().(features.KratosConfig)
}

func (k *Kratos) ValidateSession(r *http.Request) (*oryKratos.Session, error) {
	cookie, err := r.Cookie("ory_session_playground")
	if err != nil {
		return nil, err
	}
	if cookie == nil {
		return nil, errors.New("no session found in cookie")
	}
	resp, _, err := k.Public.FrontendApi.ToSession(context.Background()).Cookie(cookie.String()).Execute()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (k *Kratos) CreateIdentity(user map[string]interface{}, password string) (*oryKratos.Identity, error) {
	newIdentity := *oryKratos.NewCreateIdentityBody("default", user)
	response, _, err := k.admin.IdentityApi.
		CreateIdentity(context.Background()).
		CreateIdentityBody(newIdentity).
		Execute()
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (k *Kratos) GetAllIdentity() ([]oryKratos.Identity, error) {
	identities, _, err := k.admin.IdentityApi.
		ListIdentities(context.Background()).
		Execute()
	return identities, err
}

func (k *Kratos) GetIdentity(id string) (*oryKratos.Identity, error) {
	identity, _, err := k.admin.IdentityApi.
		GetIdentity(context.Background(), id).
		Execute()
	return identity, err
}

func (k *Kratos) DeleteIdentity(id string) error {
	_, err := k.admin.IdentityApi.
		DeleteIdentity(context.Background(), id).
		Execute()
	return err
}
