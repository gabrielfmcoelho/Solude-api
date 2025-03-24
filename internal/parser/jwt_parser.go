package parser

import (
	"time"

	"github.com/gabrielfmcoelho/platform-core/domain"
	jwt "github.com/golang-jwt/jwt/v4"
)

// parse User to JWTCustomClaims
func ToJwtCustomClaims(u *domain.User, expireTime time.Time) *domain.JwtCustomClaims {
	return &domain.JwtCustomClaims{
		UserID:             u.ID,
		UserRoleID:         u.RoleID,
		OrganizationID:     u.OrganizationID,
		OrganizationRoleID: u.Organization.RoleID,
		OrganizationName:   u.Organization.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   u.Email,
			ExpiresAt: jwt.NewNumericDate(expireTime),
		},
	}
}

// parse User to JWTCustomRefreshClaims
func ToJwtCustomRefreshClaims(u *domain.User, expireTime time.Time) *domain.JwtCustomRefreshClaims {
	return &domain.JwtCustomRefreshClaims{
		UserID:             u.ID,
		UserRoleID:         u.RoleID,
		OrganizationID:     u.OrganizationID,
		OrganizationRoleID: u.Organization.RoleID,
		OrganizationName:   u.Organization.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   u.Email,
			ExpiresAt: jwt.NewNumericDate(expireTime),
		},
	}
}
