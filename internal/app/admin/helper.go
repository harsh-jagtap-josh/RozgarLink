package admin

import "github.com/harsh-jagtap-josh/RozgarLink/internal/repo"

func MapAdminServiceToRepo(adminData Admin) repo.Admin {
	return repo.Admin(adminData)
}
