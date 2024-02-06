package cgroups

import (
	"github.com/containerd/cgroups"
)

// CgroupsManager defines the interface for managing cgroups.
type CgroupsManager interface {
	// AddProcess adds a process with the given PID to the cgroup.
	AddProcess(pid int) error

	// Set sets the cgroup resources to the specified configuration.
	Set(config *cgroups.Config) error

	CPULimit(ns uint64) error
	MemoryLimit(ns uint64) error
	// Delete removes the cgroup.
	Delete() error
}

// NewCgroupsManager creates a new cgroups manager for the given subsystem.
func NewCgroupsManager(subsystem string, path string) (CgroupsManager, error) {
	// Create a new cgroups manager for the specified subsystem and path.
	cgroup, err := cgroups.New(subsystem, cgroups.StaticPath(path))
	if err != nil {
		return nil, err
	}

	return &cgroupsManager{
		cgroup: cgroup,
	}, nil
}

// cgroupsManager is an implementation of the CgroupsManager interface.
type cgroupsManager struct {
	cgroup cgroups.Cgroup
}

// AddProcess adds a process with the given PID to the cgroup.
func (m *cgroupsManager) AddProcess(pid int) error {
	return m.cgroup.Add(cgroups.Process{Pid: pid})
}

// Set sets the cgroup resources to the specified configuration.
func (m *cgroupsManager) Set(config *cgroups.Config) error {
	return m.cgroup.Set(config)
}

// Delete removes the cgroup.
func (m *cgroupsManager) Delete() error {
	return m.cgroup.Delete()
}
