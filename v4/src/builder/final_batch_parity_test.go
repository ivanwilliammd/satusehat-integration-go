package builder_test

import (
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/builder"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

func TestCoverageClassAndCostMatchPHPPayload(t *testing.T) {
	b := builder.NewCoverageBuilder()
	b.SetId("cov-1")
	b.SetStatus("active")
	b.SetType(konsep("http://terminology.hl7.org/CodeSystem/v3-ActCode", "EHCPOL", "extended healthcare"))
	b.SetClass("http://terminology.hl7.org/CodeSystem/coverage-class", "group", "G-1", "Grup Karyawan")
	b.SetCostToBeneficiary(50000, "IDR")
	samakan(t, bacaGolden(t, "coverage_full"), b)
}

func TestCoverageEligibilityResponseInsuranceMatchesPHPPayload(t *testing.T) {
	nama := "Rawat inap"
	sistem := "http://terminology.hl7.org/CodeSystem/benefit-category"
	kode := "1"
	tampil := "Medical Care"

	b := builder.NewCoverageEligibilityResponseBuilder()
	b.SetId("cer-1")
	b.SetStatus("active")
	b.SetPurpose([]string{"benefits"})
	b.SetServicedDate("2026-01-05")
	b.SetRequestProvider("Organization/1000000001")
	b.SetPatient("Patient/1000000001")
	b.AddInsurance("Coverage/cov-1", &nama, &sistem, &kode, &tampil)
	samakan(t, bacaGolden(t, "coverage_eligibility_response"), b)
}

func TestDeviceSurfaceMatchesPHPPayload(t *testing.T) {
	b := builder.NewDeviceBuilder()
	b.SetId("dev-1")
	b.SetStatus("active")
	b.SetManufacturer("Philips")
	b.AddDeviceName("Monitor Pasien", "user-friendly-name")
	b.SetOwner(&datatype.Reference{Reference: "Organization/1000000001"})
	b.SetLocation(&datatype.Reference{Reference: "Location/3000000001"})
	b.SetSerialNumber("SN-001")
	samakan(t, bacaGolden(t, "device_full"), b)
}
