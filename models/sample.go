// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Sample sample
// swagger:model Sample
type Sample struct {

	// cf plasma location
	CfPlasmaLocation *string `json:"cf_plasma_location,omitempty"`

	// comments
	Comments *string `json:"comments,omitempty"`

	// container id
	ContainerID *string `json:"container_id,omitempty"`

	// container name
	ContainerName *string `json:"container_name,omitempty"`

	// container type
	ContainerType *string `json:"container_type,omitempty"`

	// container well
	ContainerWell *string `json:"container_well,omitempty"`

	// copath num
	CopathNum *string `json:"copath_num,omitempty"`

	// date collected
	DateCollected *string `json:"date_collected,omitempty"`

	// date received
	DateReceived *string `json:"date_received,omitempty"`

	// date submitted
	DateSubmitted *string `json:"date_submitted,omitempty"`

	// delta ct value
	DeltaCtValue *float32 `json:"delta_ct_Value,omitempty"`

	// dna concentration
	DnaConcentration *float32 `json:"dna_concentration,omitempty"`

	// dna extraction date
	DnaExtractionDate *string `json:"dna_extraction_date,omitempty"`

	// dna location
	DnaLocation *string `json:"dna_location,omitempty"`

	// dna quality
	DnaQuality *string `json:"dna_quality,omitempty"`

	// dna quality by rnase p
	DnaQualityByRnaseP *float32 `json:"dna_quality_by_rnase_p,omitempty"`

	// dna sample barcode
	DnaSampleBarcode *string `json:"dna_sample_barcode,omitempty"`

	// dna volume
	DnaVolume *float32 `json:"dna_volume,omitempty"`

	// facility
	Facility *string `json:"facility,omitempty"`

	// ffpe qc date
	FfpeQcDate *string `json:"ffpe_qc_date,omitempty"`

	// h e slide location
	HESlideLocation *string `json:"h_e_slide_location,omitempty"`

	// has sample files
	HasSampleFiles *bool `json:"has_sample_files,omitempty"`

	// historical diagnosis
	HistoricalDiagnosis *string `json:"historical_diagnosis,omitempty"`

	// material received
	MaterialReceived *string `json:"material_received,omitempty"`

	// material received num
	MaterialReceivedNum *string `json:"material_received_num,omitempty"`

	// material received other
	MaterialReceivedOther *string `json:"material_received_other,omitempty"`

	// name of requestor
	NameOfRequestor *string `json:"name_of_requestor,omitempty"`

	// non uhn id
	NonUhnID *string `json:"non_uhn_id,omitempty"`

	// other identifier
	OtherIdentifier *string `json:"other_identifier,omitempty"`

	// pb bm location
	PbBmLocation *string `json:"pb_bm_location,omitempty"`

	// plasma location
	PlasmaLocation *string `json:"plasma_location,omitempty"`

	// reviewed by
	ReviewedBy *string `json:"reviewed_by,omitempty"`

	// rna concentration
	RnaConcentration *float32 `json:"rna_concentration,omitempty"`

	// rna extraction date
	RnaExtractionDate *string `json:"rna_extraction_date,omitempty"`

	// rna location
	RnaLocation *string `json:"rna_location,omitempty"`

	// rna lysate location
	RnaLysateLocation *string `json:"rna_lysate_location,omitempty"`

	// rna quality
	RnaQuality *float32 `json:"rna_quality,omitempty"`

	// rna volume
	RnaVolume *float32 `json:"rna_volume,omitempty"`

	// rnase p date
	RnasePDate *string `json:"rnase_p_date,omitempty"`

	// sample id
	SampleID *string `json:"sample_id,omitempty"`

	// sample name
	SampleName *string `json:"sample_name,omitempty"`

	// sample size
	SampleSize *string `json:"sample_size,omitempty"`

	// sample type
	SampleType *string `json:"sample_type,omitempty"`

	// se num
	SeNum *string `json:"se_num,omitempty"`

	// study id
	StudyID *string `json:"study_id,omitempty"`

	// surgical num
	SurgicalNum *string `json:"surgical_num,omitempty"`

	// test requested
	TestRequested *string `json:"test_requested,omitempty"`

	// tumor percnt of circled
	TumorPercntOfCircled *float32 `json:"tumor_percnt_of_circled,omitempty"`

	// tumor percnt of total
	TumorPercntOfTotal *float32 `json:"tumor_percnt_of_total,omitempty"`

	// tumor site
	TumorSite *string `json:"tumor_site,omitempty"`

	// volume of blood marrow
	VolumeOfBloodMarrow *float32 `json:"volume_of_blood_marrow,omitempty"`

	// wbc location
	WbcLocation *string `json:"wbc_location,omitempty"`
}

// Validate validates this sample
func (m *Sample) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *Sample) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Sample) UnmarshalBinary(b []byte) error {
	var res Sample
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}