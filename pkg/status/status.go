package status

import "github.com/kardianos/service"

type Status service.Status

func (s Status) Status(status service.Status) Status {
	return Status(status)
}
