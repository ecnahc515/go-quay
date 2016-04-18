package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*NewPrototype Description of a new prototype

swagger:model NewPrototype
*/
type NewPrototype struct {

	/* activating user
	 */
	ActivatingUser *NewPrototypeActivatingUser `json:"activating_user,omitempty"`

	/* delegate

	Required: true
	*/
	Delegate *NewPrototypeDelegate `json:"delegate"`

	/* Role that should be applied to the delegate

	Required: true
	*/
	Role *string `json:"role"`
}

// Validate validates this new prototype
func (m *NewPrototype) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActivatingUser(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDelegate(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRole(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewPrototype) validateActivatingUser(formats strfmt.Registry) error {

	if swag.IsZero(m.ActivatingUser) { // not required
		return nil
	}

	if m.ActivatingUser != nil {

		if err := m.ActivatingUser.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *NewPrototype) validateDelegate(formats strfmt.Registry) error {

	if m.Delegate != nil {

		if err := m.Delegate.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

var newPrototypeTypeRolePropEnum []interface{}

// prop value enum
func (m *NewPrototype) validateRoleEnum(path, location string, value string) error {
	if newPrototypeTypeRolePropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["read","write","admin"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			newPrototypeTypeRolePropEnum = append(newPrototypeTypeRolePropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, newPrototypeTypeRolePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *NewPrototype) validateRole(formats strfmt.Registry) error {

	if err := validate.Required("role", "body", m.Role); err != nil {
		return err
	}

	// value enum
	if err := m.validateRoleEnum("role", "body", *m.Role); err != nil {
		return err
	}

	return nil
}

/*NewPrototypeActivatingUser Repository creating user to whom the rule should apply

swagger:model NewPrototypeActivatingUser
*/
type NewPrototypeActivatingUser struct {

	/* The username for the activating_user

	Required: true
	*/
	Name *string `json:"name"`
}

// Validate validates this new prototype activating user
func (m *NewPrototypeActivatingUser) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewPrototypeActivatingUser) validateName(formats strfmt.Registry) error {

	if err := validate.Required("activating_user"+"."+"name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

/*NewPrototypeDelegate Information about the user or team to which the rule grants access

swagger:model NewPrototypeDelegate
*/
type NewPrototypeDelegate struct {

	/* Whether the delegate is a user or a team

	Required: true
	*/
	Kind *string `json:"kind"`

	/* The name for the delegate team or user

	Required: true
	*/
	Name *string `json:"name"`
}

// Validate validates this new prototype delegate
func (m *NewPrototypeDelegate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKind(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var newPrototypeDelegateTypeKindPropEnum []interface{}

// prop value enum
func (m *NewPrototypeDelegate) validateKindEnum(path, location string, value string) error {
	if newPrototypeDelegateTypeKindPropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["user","team"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			newPrototypeDelegateTypeKindPropEnum = append(newPrototypeDelegateTypeKindPropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, newPrototypeDelegateTypeKindPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *NewPrototypeDelegate) validateKind(formats strfmt.Registry) error {

	if err := validate.Required("delegate"+"."+"kind", "body", m.Kind); err != nil {
		return err
	}

	// value enum
	if err := m.validateKindEnum("delegate"+"."+"kind", "body", *m.Kind); err != nil {
		return err
	}

	return nil
}

func (m *NewPrototypeDelegate) validateName(formats strfmt.Registry) error {

	if err := validate.Required("delegate"+"."+"name", "body", m.Name); err != nil {
		return err
	}

	return nil
}
