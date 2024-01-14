package runtime

import "taskup/db"

type Runtime struct {
	DbConn db.Client
}
