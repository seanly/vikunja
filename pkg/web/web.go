// Vikunja is a to-do list application to facilitate your life.
// Copyright 2018-present Vikunja and contributors. All rights reserved.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public Licensee as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public Licensee for more details.
//
// You should have received a copy of the GNU Affero General Public Licensee
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package web

import (
	"github.com/labstack/echo/v4"
	"xorm.io/xorm"
)

// Rights defines rights methods
type Rights interface {
	CanRead(*xorm.Session, Auth) (bool, int, error)
	CanDelete(*xorm.Session, Auth) (bool, error)
	CanUpdate(*xorm.Session, Auth) (bool, error)
	CanCreate(*xorm.Session, Auth) (bool, error)
}

// CRUDable defines the crud methods
type CRUDable interface {
	Create(*xorm.Session, Auth) error
	ReadOne(*xorm.Session, Auth) error
	ReadAll(s *xorm.Session, auth Auth, search string, page int, perPage int) (result interface{}, resultCount int, numberOfTotalItems int64, err error)
	Update(*xorm.Session, Auth) error
	Delete(*xorm.Session, Auth) error
}

// HTTPErrorProcessor is executed when the defined error is thrown, it will make sure the user sees an appropriate error message and http status code
type HTTPErrorProcessor interface {
	HTTPError() HTTPError
}

// HTTPError holds informations about an http error
type HTTPError struct {
	HTTPCode int    `json:"-"`
	Code     int    `json:"code"`
	Message  string `json:"message"`
}

// Auth defines the authentication interface used to get some auth thing
type Auth interface {
	// Most of the time, we need an ID from the auth object only. Having this method saves the need to cast it.
	GetID() int64
}

// Authprovider is a holder for the implementation of an authprovider by the application
type Authprovider interface {
	GetAuthObject(echo.Context) (Auth, error)
}

// Auths holds the authobject
type Auths struct {
	AuthObject func(echo.Context) (Auth, error)
}