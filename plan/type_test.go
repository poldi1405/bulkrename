package plan

import (
	"os"
	"testing"

	j "gitlab.com/poldi1405/bulkrename/plan/jobdescriptor"
)

func TestExecuteRemove(t *testing.T) {
	_, err := os.Create("1")
	if err != nil {
		t.Skipf(err.Error())
	}
	defer os.Remove("1")

	var p Plan
	p.jobs = []j.JobDescriptor{{
		Action:     -1,
		SourcePath: "1",
	}}

	errOcc, descs, errs := p.Execute()
	if errOcc || len(descs) != 0 || len(errs) != 0 {
		t.Fail()
	}
}

func TestExecuteRemoveFails(t *testing.T) {
	var p Plan
	p.jobs = []j.JobDescriptor{{
		Action:     -1,
		SourcePath: "1",
	}}

	err, descs, errs := p.Execute()
	if !err && len(descs) == 0 && len(errs) == 0 {
		t.Fail()
	}
}

func TestExecuteRename(t *testing.T) {
	_, err := os.Create("1")
	if err != nil {
		t.Skipf(err.Error())
	}
	defer os.Remove("1")

	var p Plan
	p.jobs = []j.JobDescriptor{{
		Action:     1,
		SourcePath: "1",
		DstPath:    "2",
	}}

	errOcc, descs, errs := p.Execute()
	if errOcc || len(descs) != 0 || len(errs) != 0 {
		t.Fail()
	}

	err = os.Remove("2")
	if err != nil {
		t.Fail()
	}
}

func TestExecuteRenameFails(t *testing.T) {
	var p Plan
	p.jobs = []j.JobDescriptor{{
		Action:     1,
		SourcePath: "1",
	}}

	err, descs, errs := p.Execute()
	if !err && len(descs) == 0 && len(errs) == 0 {
		t.Fail()
	}
}

func TestExecuteMkdir(t *testing.T) {
	var p Plan
	p.jobs = []j.JobDescriptor{{
		Action:  2,
		DstPath: "1",
	}}

	p.Execute()
	errOcc, descs, errs := p.Execute()
	if errOcc || len(descs) != 0 || len(errs) != 0 {
		t.Fail()
	}

	err := os.Remove("1")
	if err != nil {
		t.Error(err)
	}
}

func TestExecuteMkdirFails(t *testing.T) {
	_, err := os.Create("1")
	if err != nil {
		t.Skipf(err.Error())
	}
	defer os.Remove("1")

	var p Plan
	p.jobs = []j.JobDescriptor{{
		Action:  2,
		DstPath: "1",
	}}

	errOcc, descs, errs := p.Execute()
	if !errOcc && len(descs) == 0 && len(errs) == 0 {
		t.Fail()
	}
}
