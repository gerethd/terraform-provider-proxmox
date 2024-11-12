package proxmox_client

type QemuResponse struct { //some of the optional fields in the spec will not appear if they do not differ from the default value
	Data struct {
		Acpi             int                    `json:"acpi"`
		AutoStart        int                    `json:"autostart"`
		Bios             string                 `json:"bios"`
		CpuLimit         string                 `json:"cpulimit"`
		HostStartupOrder string                 `json:"startup"`
		Kvm              int                    `json:"kvm"`
		Tags             string                 `json:"tags"`
		Memory           string                 `json:"memory"`
		Name             string                 `json:"name"`
		Cpu              string                 `json:"cpu"`
		OnBoot           int                    `json:"onboot"`
		Description      string                 `json:"description"`
		OsType           string                 `json:"ostype"`
		VmGenId          string                 `json:"vmgenid"`
		Sockets          int                    `json:"sockets"`
		Meta             string                 `json:"meta"`
		ScsiHw           string                 `json:"scsihw"`
		Agent            string                 `json:"agent"`
		Boot             string                 `json:"boot"`
		Cores            int                    `json:"cores"`
		Numa             int                    `json:"numa"`
		Nameserver       string                 `json:"nameserver"`
		CloudInitUpgrade int                    `json:"ciupgrade"`
		Protection       int                    `json:"protection"`
		SshKeys          string                 `json:"sshKeys"`
		OtherFields      map[string]interface{} `json:"-"` //skip this key
	} `json:"data"`
}

type TaskStatus struct {
	Data struct {
		Tokenid    string `json:"tokenid"`
		Upid       string `json:"upid"`
		Pstart     int    `json:"pstart"`
		User       string `json:"user"`
		Pid        int    `json:"pid"`
		Status     string `json:"status"`
		Node       string `json:"node"`
		Id         string `json:"id"`
		Starttime  int    `json:"starttime"`
		Exitstatus string `json:"exitstatus"`
		Type       string `json:"type"`
	} `json:"data"`
}

type TaskCreationResponse struct {
	Upid string `json:"data"`
}

type NodeListResponse struct {
	Data []struct {
		Status         string  `json:"status"`
		Maxmem         int64   `json:"maxmem"`
		Level          string  `json:"level"`
		Id             string  `json:"id"`
		Type           string  `json:"type"`
		Maxcpu         int     `json:"maxcpu"`
		Maxdisk        int64   `json:"maxdisk"`
		Node           string  `json:"node"`
		Uptime         int     `json:"uptime"`
		Cpu            float64 `json:"cpu"`
		Disk           int64   `json:"disk"`
		Mem            int64   `json:"mem"`
		SslFingerprint string  `json:"ssl_fingerprint"`
	} `json:"data"`
}

type SdnZoneResponse struct {
	Data struct {
		Digest string `json:"digest"`
		Zone   string `json:"zone"`
		Ipam   string `json:"ipam"`
		Nodes  string `json:"nodes"`
		Peers  string `json:"peers"`
		Type   string `json:"type"`
	} `json:"data"`
}