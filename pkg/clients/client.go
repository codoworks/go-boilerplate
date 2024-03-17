/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package clients

type IClient interface {
	Name() string
	Configure(v any)
}
