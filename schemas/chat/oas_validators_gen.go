// Code generated by ogen, DO NOT EDIT.

package chat

import (
	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/validate"
)

func (s *APIChatInternalClientPostApplicationJSON) Validate() error {
	alias := (*ClientThing)(s)
	if err := alias.Validate(); err != nil {
		return err
	}
	return nil
}

func (s *APIChatInternalClientPostApplicationJSONOK) Validate() error {
	alias := (*CientCreationResponse)(s)
	if err := alias.Validate(); err != nil {
		return err
	}
	return nil
}

func (s *APIChatInternalClientPostTextJSON) Validate() error {
	alias := (*ClientThing)(s)
	if err := alias.Validate(); err != nil {
		return err
	}
	return nil
}

func (s *APIChatInternalClientPostTextJSONOK) Validate() error {
	alias := (*CientCreationResponse)(s)
	if err := alias.Validate(); err != nil {
		return err
	}
	return nil
}

func (s *APIChatMuteDeleteApplicationJSON) Validate() error {
	alias := (*UnMute)(s)
	if err := alias.Validate(); err != nil {
		return err
	}
	return nil
}

func (s *APIChatMuteDeleteApplicationJSONOK) Validate() error {
	alias := (*UnMute)(s)
	if err := alias.Validate(); err != nil {
		return err
	}
	return nil
}

func (s *APIChatMuteDeleteTextJSON) Validate() error {
	alias := (*UnMute)(s)
	if err := alias.Validate(); err != nil {
		return err
	}
	return nil
}

func (s *APIChatMuteDeleteTextJSONOK) Validate() error {
	alias := (*UnMute)(s)
	if err := alias.Validate(); err != nil {
		return err
	}
	return nil
}

func (s *APIChatMutePostApplicationJSON) Validate() error {
	alias := (*Mute)(s)
	if err := alias.Validate(); err != nil {
		return err
	}
	return nil
}

func (s *APIChatMutePostApplicationJSONOK) Validate() error {
	alias := (*Mute)(s)
	if err := alias.Validate(); err != nil {
		return err
	}
	return nil
}

func (s *APIChatMutePostTextJSON) Validate() error {
	alias := (*Mute)(s)
	if err := alias.Validate(); err != nil {
		return err
	}
	return nil
}

func (s *APIChatMutePostTextJSONOK) Validate() error {
	alias := (*Mute)(s)
	if err := alias.Validate(); err != nil {
		return err
	}
	return nil
}

func (s *APIChatSendPostApplicationJSON) Validate() error {
	alias := (*ChatMessage)(s)
	if err := alias.Validate(); err != nil {
		return err
	}
	return nil
}

func (s *APIChatSendPostApplicationJSONOK) Validate() error {
	alias := (*ChatMessage)(s)
	if err := alias.Validate(); err != nil {
		return err
	}
	return nil
}

func (s *APIChatSendPostTextJSON) Validate() error {
	alias := (*ChatMessage)(s)
	if err := alias.Validate(); err != nil {
		return err
	}
	return nil
}

func (s *APIChatSendPostTextJSONOK) Validate() error {
	alias := (*ChatMessage)(s)
	if err := alias.Validate(); err != nil {
		return err
	}
	return nil
}

func (s *ChatMessage) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if value, ok := s.UUID.Get(); ok {
			if err := func() error {
				if err := (validate.String{
					MinLength:    0,
					MinLengthSet: true,
					MaxLength:    32,
					MaxLengthSet: true,
					Email:        false,
					Hostname:     false,
					Regex:        nil,
				}).Validate(string(value)); err != nil {
					return errors.Wrap(err, "string")
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "uuid",
			Error: err,
		})
	}
	if err := func() error {
		if value, ok := s.Name.Get(); ok {
			if err := func() error {
				if err := (validate.String{
					MinLength:    0,
					MinLengthSet: false,
					MaxLength:    16,
					MaxLengthSet: true,
					Email:        false,
					Hostname:     false,
					Regex:        nil,
				}).Validate(string(value)); err != nil {
					return errors.Wrap(err, "string")
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "name",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *CientCreationResponse) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if value, ok := s.Client.Get(); ok {
			if err := func() error {
				if err := value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "client",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *ClientThing) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if value, ok := s.Name.Get(); ok {
			if err := func() error {
				if err := (validate.String{
					MinLength:    0,
					MinLengthSet: false,
					MaxLength:    32,
					MaxLengthSet: true,
					Email:        false,
					Hostname:     false,
					Regex:        nil,
				}).Validate(string(value)); err != nil {
					return errors.Wrap(err, "string")
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "name",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *Mute) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if value, ok := s.UUID.Get(); ok {
			if err := func() error {
				if err := (validate.String{
					MinLength:    0,
					MinLengthSet: true,
					MaxLength:    32,
					MaxLengthSet: true,
					Email:        false,
					Hostname:     false,
					Regex:        nil,
				}).Validate(string(value)); err != nil {
					return errors.Wrap(err, "string")
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "uuid",
			Error: err,
		})
	}
	if err := func() error {
		if value, ok := s.Muter.Get(); ok {
			if err := func() error {
				if err := (validate.String{
					MinLength:    0,
					MinLengthSet: true,
					MaxLength:    32,
					MaxLengthSet: true,
					Email:        false,
					Hostname:     false,
					Regex:        nil,
				}).Validate(string(value)); err != nil {
					return errors.Wrap(err, "string")
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "muter",
			Error: err,
		})
	}
	if err := func() error {
		if value, ok := s.UnMuter.Get(); ok {
			if err := func() error {
				if err := (validate.String{
					MinLength:    0,
					MinLengthSet: true,
					MaxLength:    32,
					MaxLengthSet: true,
					Email:        false,
					Hostname:     false,
					Regex:        nil,
				}).Validate(string(value)); err != nil {
					return errors.Wrap(err, "string")
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "unMuter",
			Error: err,
		})
	}
	if err := func() error {
		if value, ok := s.Status.Get(); ok {
			if err := func() error {
				if err := value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "status",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s MuteStatus) Validate() error {
	switch s {
	case 0:
		return nil
	case 1:
		return nil
	case 2:
		return nil
	case 4:
		return nil
	case 8:
		return nil
	case 16:
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s *UnMute) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if value, ok := s.UUID.Get(); ok {
			if err := func() error {
				if err := (validate.String{
					MinLength:    0,
					MinLengthSet: true,
					MaxLength:    32,
					MaxLengthSet: true,
					Email:        false,
					Hostname:     false,
					Regex:        nil,
				}).Validate(string(value)); err != nil {
					return errors.Wrap(err, "string")
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "uuid",
			Error: err,
		})
	}
	if err := func() error {
		if value, ok := s.UnMuter.Get(); ok {
			if err := func() error {
				if err := (validate.String{
					MinLength:    0,
					MinLengthSet: true,
					MaxLength:    32,
					MaxLengthSet: true,
					Email:        false,
					Hostname:     false,
					Regex:        nil,
				}).Validate(string(value)); err != nil {
					return errors.Wrap(err, "string")
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "unMuter",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
