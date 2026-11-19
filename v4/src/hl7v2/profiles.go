package hl7v2

// AdtA01 is admission: PID and PV1 are both required because the visit is the reason for the message.
func AdtA01() Profile {
	return Profile{
		Name:             "ADT_A01",
		MessageType:      "ADT",
		Trigger:          "A01",
		RequiredSegments: []string{"MSH", "EVN", "PID", "PV1"},
		RequiredFields: map[string]string{
			"MSH.10": "Message control ID",
			"PID.3":  "Patient identifier",
			"PID.5.1": "Patient family name",
			"PV1.2":  "Patient class (I/O/E)",
		},
		AllowedValues: map[string][]string{
			"PID.8.1": {"F", "M", "O", "U", "A", "N"},
			"PV1.2":   {"I", "O", "E", "P", "R", "B", "C", "N"},
		},
	}
}

// AdtA08 is "update patient information"; PV1 is often absent, so it is not required.
func AdtA08() Profile {
	return Profile{
		Name:             "ADT_A08",
		MessageType:      "ADT",
		Trigger:          "A08",
		RequiredSegments: []string{"MSH", "EVN", "PID"},
		RequiredFields: map[string]string{
			"MSH.10":  "Message control ID",
			"PID.3":   "Patient identifier",
			"PID.5.1": "Patient family name",
		},
		AllowedValues: map[string][]string{
			"PID.8.1": {"F", "M", "O", "U", "A", "N"},
		},
	}
}

func OruR01() Profile {
	return Profile{
		Name:             "ORU_R01",
		MessageType:      "ORU",
		Trigger:          "R01",
		RequiredSegments: []string{"MSH", "PID", "OBR", "OBX"},
		RequiredFields: map[string]string{
			"MSH.10":  "Message control ID",
			"PID.3":   "Patient identifier",
			"OBR.4.1": "Universal service identifier",
			"OBX.5":   "Observation value",
		},
		AllowedValues: map[string][]string{
			"OBX.11": {"F", "P", "C", "X"},
		},
	}
}

// Profiles maps the manifest keys used by the shared corpus.
var Profiles = map[string]func() Profile{
	"adtA01": AdtA01,
	"adtA08": AdtA08,
	"oruR01": OruR01,
}
