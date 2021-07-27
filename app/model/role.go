package model

type Role struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type Roles []*Role

func (r Roles) UserRoles(userId int64) []UserRole {
	var userRoles []UserRole
	for _, role := range r {
		if role == nil {
			continue
		}
		userRoles = append(userRoles, UserRole{
			UserId: userId,
			RoleId: role.Id,
		})
	}
	return userRoles
}

func (r Roles) HasRole(names ...string) bool {
	for _, role := range r {
		if role == nil {
			continue
		}
		for _, name := range names {
			if role.Name == name {
				return true
			}
		}
	}
	return false
}

type UserRole struct {
	UserId int64 `json:"user_id,omitempty"`
	RoleId int64 `json:"role_id,omitempty"`
}
