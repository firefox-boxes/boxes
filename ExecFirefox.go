package boxes

import "os/exec"

type ProcessID int
type ExecPIDs map[ProfileID]ProcessID
var Exec ExecPIDs = make(ExecPIDs)

func (execPIDs *ExecPIDs) IsRunning(profileID ProfileID) (bool, ProcessID) {
	val, ok := (*execPIDs)[profileID]
	return ok, val
}

func (I *Installation) ExecProfile(profileID ProfileID, pr ProbeResult) {
	cmd := exec.Command(I.Exec, "-profile", pr.GetProfileDir(profileID), "-no-remote")
	cmd.Start()
	Exec[profileID] = ProcessID(cmd.Process.Pid)
	cmd.Wait()
	delete(Exec, profileID)
}

func waitAndDeleteProcessIDLater(cmd *exec.Cmd, profileID ProfileID) {
	cmd.Wait()
	delete(Exec, profileID)
}

func (I *Installation) ExecProfileAndDetach(profileID ProfileID, pr ProbeResult) ProcessID {
	cmd := exec.Command(I.Exec, "-profile", pr.GetProfileDir(profileID), "-no-remote")
	cmd.Start()
	Exec[profileID] = ProcessID(cmd.Process.Pid)
	go waitAndDeleteProcessIDLater(cmd, profileID)
	return Exec[profileID]
}