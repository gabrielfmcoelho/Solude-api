package domain

import (
	"github.com/golang-jwt/jwt/v4"
)

// JWT has a lot of predefined claims that can be used known as Registered Claims (like Issuer, Subject, Audience, Expiration Time, Not Before, Issued At, JWT ID)
// You can also add custom claims to a JWT token. Custom claims are claims that are not registered in the IANA "JSON Web Token Claims" registry or the reserved claim names defined in the JWT specification.
type JwtCustomClaims struct {
	OrganizationID       uint   `json:"organization_id"`
	OrganizationRoleID   uint   `json:"organization_role_id"`
	OrganizationName     string `json:"organization_name"`
	UserRoleID           uint   `json:"user_role_id"`
	UserID               uint   `json:"user_id"`
	jwt.RegisteredClaims        // userID, user.BioInfo.FirstName, ExpiresAt
}

type JwtCustomRefreshClaims struct {
	OrganizationID       uint   `json:"organization_id"`
	OrganizationRoleID   uint   `json:"organization_role_id"`
	OrganizationName     string `json:"organization_name"`
	UserRoleID           uint   `json:"user_role_id"`
	UserID               uint   `json:"user_id"`
	jwt.RegisteredClaims        // userID, user.BioInfo.FirstName, ExpiresAt
}

// TokenUtil contains the methods to create and validate JWT tokens defined here
// see the use in internal/tokenutil/tokenutil.go
