package linux

import (
	"bytes"
	"os/exec"
	"strings"
)

//HCIAdapterInfo contains details for an adapter
type HCIAdapterInfo struct {
	Enabled bool
	Address string
	Type    string
	Bus     string
}

// HCIConfig an hciconfig command wrapper
type HCIConfig struct {
}

func getAdapterStatus(adapterID string) (HCIAdapterInfo, error) {

	cfg := HCIAdapterInfo{}

	cmd := exec.Command("hciconfig", adapterID)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return HCIAdapterInfo{}, err
	}

	s := strings.Replace(out.String()[6:], "\t", "", -1)
	lines := strings.Split(s, "\n")
	// var parts []string
	for i, line := range lines {
		if i > 2 {
			break
		}
		if i == 2 {
			pp := strings.Split(line, " ")
			cfg.Enabled = (pp[0] == "UP")
			continue
		}

		subparts := strings.Split(line, "  ")
		for _, subpart := range subparts {
			pp := strings.Split(subpart, ": ")
			switch pp[0] {
			case "Type":
				cfg.Type = pp[1]
				continue
			case "Bus":
				cfg.Bus = pp[1]
				continue
			case "BD Address":
				cfg.Address = pp[1]
				continue
			}
		}
	}
	// log.Printf("%v", strings.Join(parts, ","))

	return cfg, nil
}

// Up Turn on an HCI device
func (h *HCIConfig) Up(adapterID string) (HCIAdapterInfo, error) {
	cmd := exec.Command("hciconfig", adapterID, "up")
	err := cmd.Run()
	if err != nil {
		return HCIAdapterInfo{}, err
	}
	return getAdapterStatus(adapterID)
}

// Down Turn down an HCI device
func (h *HCIConfig) Down(adapterID string) (HCIAdapterInfo, error) {
	cmd := exec.Command("hciconfig", adapterID, "down")
	err := cmd.Run()
	if err != nil {
		return HCIAdapterInfo{}, err
	}
	return getAdapterStatus(adapterID)
}
