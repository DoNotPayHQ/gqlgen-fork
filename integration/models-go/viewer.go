package models

import "github.com/DoNotPayHQ/gqlgen-fork/integration/remote_api"

type Viewer struct {
	User *remote_api.User
}
