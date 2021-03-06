// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2018 Datadog, Inc.

package config

import (
	"time"

	"github.com/spf13/viper"
)

// ViperConfig is a global variable for the viper configuration
// TODO move it to the cmd package and make it private see https://github.com/DataDog/pupernetes/issues/40
var ViperConfig = viper.New()

const (
	// JobTypeKey is the key for daemon types
	JobTypeKey = "job-type"

	// JobSystemd is the value to daemonise in a systemd unit .service
	JobSystemd = "systemd"

	// JobForeground is the value to daemonise the current process
	JobForeground = "fg"

	defaultAPIAddress = "127.0.0.1:8989"

	// CRIContainerd is a container runtime engine
	CRIContainerd = "containerd"
)

func init() {
	ViperConfig.SetDefault("version", false)

	ViperConfig.SetDefault("skip-binaries-version", false)
	ViperConfig.SetDefault("hyperkube-version", "1.10.7")
	ViperConfig.SetDefault("vault-version", "0.9.5")
	ViperConfig.SetDefault("etcd-version", "3.1.19")
	ViperConfig.SetDefault("cni-version", "0.7.0")
	ViperConfig.SetDefault("containerd-version", "1.1.3")
	ViperConfig.SetDefault("runc-version", "1.0.0-rc5")

	ViperConfig.SetDefault("container-runtime", "docker")

	ViperConfig.SetDefault("download-timeout", time.Minute*30)

	ViperConfig.SetDefault("kubernetes-cluster-ip-range", "192.168.254.0/24")
	ViperConfig.SetDefault("pod-ip-range", "192.168.253.0/24")
	ViperConfig.SetDefault("bind-address", defaultAPIAddress)
	ViperConfig.SetDefault("api-address", defaultAPIAddress)
	ViperConfig.SetDefault("kubelet-root-dir", "/var/lib/p8s-kubelet")
	ViperConfig.SetDefault("systemd-unit-prefix", "p8s-")

	ViperConfig.SetDefault("kubectl-link", "")
	ViperConfig.SetDefault("vault-root-token", "")
	ViperConfig.SetDefault("vault-listen-address", "127.0.0.1:8201")

	ViperConfig.SetDefault("clean", "etcd,kubelet,logs,mounts,iptables")
	ViperConfig.SetDefault("keep", "")
	ViperConfig.SetDefault("drain", "all")
	ViperConfig.SetDefault("skip-probes", false)
	ViperConfig.SetDefault("gc", time.Second*60)

	// The supported job-type are "fg" and "systemd"
	ViperConfig.SetDefault(JobTypeKey, JobForeground)

	ViperConfig.SetDefault("systemd-job-name", "pupernetes")

	ViperConfig.SetDefault("apply", false)

	ViperConfig.SetDefault("logging-since", time.Minute*5)
	ViperConfig.SetDefault("unit-to-watch", "pupernetes.service")
	ViperConfig.SetDefault("wait-timeout", time.Minute*15)
	ViperConfig.SetDefault("client-timeout", time.Minute*1)
	ViperConfig.SetDefault("kubeconfig-path", "")
	ViperConfig.SetDefault("dns-queries", []string{"coredns.kube-system.svc.cluster.local."})
	ViperConfig.SetDefault("dns-check", false)
}
