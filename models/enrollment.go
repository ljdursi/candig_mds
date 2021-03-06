// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Enrollment enrollment
// swagger:model Enrollment
type Enrollment struct {

	// age at enrollment
	AgeAtEnrollment *string `json:"age_at_enrollment,omitempty" db:"age_at_enrollment"`

	// cross enrollment
	CrossEnrollment *string `json:"cross_enrollment,omitempty" db:"cross_enrollment"`

	// eligibility at enrollment
	EligibilityAtEnrollment *string `json:"eligibility_at_enrollment,omitempty" db:"eligibility_at_enrollment"`

	// enrollment approval date
	EnrollmentApprovalDate *string `json:"enrollment_approval_date,omitempty" db:"enrollment_approval_date"`

	// enrollment institution
	EnrollmentInstitution *string `json:"enrollment_institution,omitempty" db:"enrollment_institution"`

	// other personalized medicine study id
	OtherPersonalizedMedicineStudyID *string `json:"other_personalized_medicine_study_id,omitempty" db:"other_personalized_medicine_study_id"`

	// other personalized medicine study name
	OtherPersonalizedMedicineStudyName *string `json:"other_personalized_medicine_study_name,omitempty" db:"other_personalized_medicine_study_name"`

	// sample id
	SampleID *string `json:"sample_id,omitempty" db:"sample_id"`

	// primary oncologist contact
	PrimaryOncologistContact *string `json:"primary_oncologist_contact,omitempty" db:"primary_oncologist_contact"`

	// primary oncologist name
	PrimaryOncologistName *string `json:"primary_oncologist_name,omitempty" db:"primary_oncologist_name"`

	// referring physician contact
	ReferringPhysicianContact *string `json:"referring_physician_contact,omitempty" db:"referring_physician_contact"`

	// referring physician name
	ReferringPhysicianName *string `json:"referring_physician_name,omitempty" db:"referring_physician_name"`

	// status at enrollment
	StatusAtEnrollment *string `json:"status_at_enrollment,omitempty" db:"status_at_enrollment"`

	// summary of id request
	SummaryOfIDRequest *string `json:"summary_of_id_request,omitempty" db:"summary_of_id_request"`

	// treating centre name
	TreatingCentreName *string `json:"treating_centre_name,omitempty" db:"treating_centre_name"`

	// treating centre province
	TreatingCentreProvince *string `json:"treating_centre_province,omitempty" db:"treating_centre_province"`

	// hash induced neoplasm details
	Hash *string `json:"hash,omitempty" db:"hash"`
}

// Validate validates this enrollment
func (m *Enrollment) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *Enrollment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Enrollment) UnmarshalBinary(b []byte) error {
	var res Enrollment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
