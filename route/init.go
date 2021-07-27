package route

import (
	"github.com/fobus1289/marshrudka/router"
	"v2/route/admin"
	"v2/route/auth"
	"v2/route/branch"
	"v2/route/company"
)

func Init(drive *router.Drive) {
	auth.Init(drive)
	admin.User(drive)
	admin.Company(drive)
	admin.Branch(drive)
	company.Branch(drive)
	company.User(drive)
	branch.User(drive)
}
