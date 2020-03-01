package boxes

import (
	"github.com/google/uuid"
)

func GenID() ProfileID {
	return ProfileID(uuid.New().String())
}

func NewProfile(p ProbeResult, pid ProfileID) {
	CreateProfile(p.GetProfileDir(pid))
}

func NewProfileSetId(p ProbeResult) ProfileID {
	pid := GenID()
	NewProfile(p, pid)
	return pid
}