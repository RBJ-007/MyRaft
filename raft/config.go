package raft

import (
	"errors"
	"time"
)

const (
	Stopped      = "stopped"
	Initialized  = "initialized"
	Follower     = "follower"
	Candidate    = "candidate"
	Leader       = "leader"
	Snapshotting = "snapshotting"
)

const (
	MaxLogEntriesPerRequest         = 2000
	NumberOfLogEntriesAfterSnapshot = 200
)

const (
	DefaultHeartbeatInterval = 50 * time.Millisecond
	DefaultElectionTimeout = 150 * time.Millisecond
)

const ElectionTimeoutThresholdPercent = 0.8


var NotLeaderError = errors.New("raft.Server: Not current leader")
var DuplicatePeerError = errors.New("raft.Server: Duplicate peer")
var CommandTimeoutError = errors.New("raft: Command timeout")
var StopError = errors.New("raft: Has been stopped")

//udp包头部代码块
const(
	AddPeerOrder = "1001"
	VoteOrder = "1002"
	VoteBackOrder = "1003"
)
