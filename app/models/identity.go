package models

import (
	jwt "github.com/dgrijalva/jwt-go"
)

//Tenant represents a tenant
type Tenant struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Subdomain      string `json:"subdomain"`
	Invitation     string `json:"invitation"`
	WelcomeMessage string `json:"welcomeMessage"`
	CNAME          string `json:"cname"`
}

//User represents an user inside our application
type User struct {
	ID        int             `json:"id"`
	Name      string          `json:"name"`
	Email     string          `json:"-"`
	Gravatar  string          `json:"gravatar"`
	Tenant    *Tenant         `json:"-"`
	Role      int             `json:"role"`
	Providers []*UserProvider `json:"-"`
}

var (
	//RoleVisitor is the basic role for every user
	RoleVisitor = 1
	//RoleMember has limited access to administrative console
	RoleMember = 2
	//RoleAdministrator has full access to administrative console
	RoleAdministrator = 3
)

//HasProvider returns true if current user has registered with given provider
func (u *User) HasProvider(provider string) bool {
	for _, p := range u.Providers {
		if p.Name == provider {
			return true
		}
	}
	return false
}

//IsStaff returns true if user has special permissions
func (u *User) IsStaff() bool {
	return u.Role >= RoleMember
}

//UserProvider represents the relashionship between an User and an Authentication provide
type UserProvider struct {
	Name string
	UID  string
}

//FiderClaims represents what goes into JWT tokens
type FiderClaims struct {
	UserID    int    `json:"user/id"`
	UserName  string `json:"user/name"`
	UserEmail string `json:"user/email"`
	jwt.StandardClaims
}

//OAuthClaims represents what goes into temporary OAuth JWT tokens
type OAuthClaims struct {
	OAuthID       string `json:"oauth/id"`
	OAuthProvider string `json:"oauth/provider"`
	OAuthName     string `json:"oauth/name"`
	OAuthEmail    string `json:"oauth/email"`
	jwt.StandardClaims
}

//CreateTenant is the input model used to create a tenant
type CreateTenant struct {
	Token      string `json:"token"`
	Name       string `json:"name"`
	Subdomain  string `json:"subdomain"`
	UserClaims *OAuthClaims
}

//UpdateTenantSettings is the input model used to update tenant settings
type UpdateTenantSettings struct {
	Title          string `json:"title"`
	Invitation     string `json:"invitation"`
	WelcomeMessage string `json:"welcomeMessage"`
	UserClaims     *OAuthClaims
}
