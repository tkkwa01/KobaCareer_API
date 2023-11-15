package validator

import (
	"KobaCareer_API/domain"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type IInternshipValidator interface {
	InternshipValidate(internship domain.Internships) error
}

type internshipValidator struct{}

func NewInternshipValidator() IInternshipValidator {
	return &internshipValidator{}
}

func (iv *internshipValidator) InternshipValidate(internship domain.Internships) error {
	return validation.ValidateStruct(&internship,
		validation.Field(
			&internship.Company,
			validation.Required.Error("company is required"),
			validation.RuneLength(1, 30).Error("limited max 30 char"),
		),
		validation.Field(
			&internship.Title,
			validation.Required.Error("title is required"),
			validation.RuneLength(1, 30).Error("limited max 30 char"),
		),
		validation.Field(
			&internship.Salary,
			validation.Required.Error("salary is required"),
		),
		validation.Field(
			&internship.Period,
			validation.Required.Error("period is required"),
			validation.RuneLength(1, 30).Error("limited max 30 char"),
		),
		validation.Field(
			&internship.Selection,
			validation.Required.Error("selection is required"),
			validation.RuneLength(1, 10).Error("limited max 10 char"),
		),
		validation.Field(
			&internship.Deadline,
			validation.Required.Error("deadline is required"),
			validation.RuneLength(1, 30).Error("limited max 30 char"),
		),
		validation.Field(
			&internship.Contributor,
			validation.Required.Error("contributor is required"),
			validation.RuneLength(1, 30).Error("limited max 30 char"),
		),
		validation.Field(
			&internship.Detail,
			validation.Required.Error("detail is required"),
			validation.RuneLength(1, 1000).Error("limited max 1000 char"),
		),
		validation.Field(
			&internship.FutureJob,
			validation.Required.Error("future_job is required"),
			validation.RuneLength(1, 1000).Error("limited max 1000 char"),
		),

		validation.Field(
			&internship.Flow,
			validation.Required.Error("flow is required"),
			validation.RuneLength(1, 1000).Error("limited max 1000 char"),
		),

		validation.Field(
			&internship.Method,
			validation.Required.Error("method is required"),
			validation.RuneLength(1, 100).Error("limited max 100 char"),
		),
	)
}
