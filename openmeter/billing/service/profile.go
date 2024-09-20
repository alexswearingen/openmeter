package service

import (
	"context"
	"fmt"

	"github.com/openmeterio/openmeter/openmeter/billing"
)

var _ billing.ProfileService = (*Service)(nil)

func (s *Service) CreateProfile(ctx context.Context, input billing.CreateProfileInput) (*billing.Profile, error) {
	input = input.WithDefaults()

	if err := input.Validate(); err != nil {
		return nil, billing.ValidationError{
			Err: err,
		}
	}

	return billing.WithTx(ctx, s.adapter, func(ctx context.Context, adapter billing.TxAdapter) (*billing.Profile, error) {
		// Given that we have multiple constraints let's validate those here for better error reporting
		if input.Default {
			defaultProfile, err := adapter.GetDefaultProfile(ctx, billing.GetDefaultProfileInput{
				Namespace: input.Namespace,
			})
			if err != nil {
				return nil, err
			}

			if defaultProfile != nil {
				return nil, billing.ValidationError{
					Err: fmt.Errorf("%w [id=%s,key=%s]", billing.ErrDefaultProfileAlreadyExists, defaultProfile.ID, defaultProfile.Key),
				}
			}
		}

		// Let's check for key uniqueness
		existingProfileWithKey, err := adapter.GetProfileByKey(ctx, billing.GetProfileByKeyAdapterInput{
			Namespace: input.Namespace,
			Key:       input.Key,
		})
		if err != nil {
			return nil, err
		}

		if existingProfileWithKey != nil {
			return nil, billing.ValidationError{
				Err: fmt.Errorf("%w [id=%s]", billing.ErrProfileWithKeyAlreadyExists, existingProfileWithKey.ID),
			}
		}

		profile, err := adapter.CreateProfile(ctx, input)
		if err != nil {
			return nil, err
		}

		if err := profile.Validate(); err != nil {
			return nil, billing.ValidationError{
				Err: fmt.Errorf("error validating profile: %w", err),
			}
		}

		return profile, nil
	})
}

func (s *Service) GetDefaultProfile(ctx context.Context, input billing.GetDefaultProfileInput) (*billing.Profile, error) {
	if err := input.Validate(); err != nil {
		return nil, billing.ValidationError{
			Err: err,
		}
	}

	return s.adapter.GetDefaultProfile(ctx, billing.GetDefaultProfileInput{
		Namespace: input.Namespace,
	})
}

func (s *Service) GetProfileByKeyOrID(ctx context.Context, input billing.GetProfileByKeyOrIDInput) (*billing.Profile, error) {
	if err := input.Validate(); err != nil {
		return nil, billing.ValidationError{
			Err: err,
		}
	}
	return s.adapter.GetProfileByKeyOrID(ctx, input)
}

func (s *Service) DeleteProfileByKeyOrID(ctx context.Context, input billing.DeleteProfileByKeyOrIDInput) error {
	if err := input.Validate(); err != nil {
		return billing.ValidationError{
			Err: err,
		}
	}

	return billing.WithTxNoValue(ctx, s.adapter, func(ctx context.Context, adapter billing.TxAdapter) error {
		profile, err := s.adapter.GetProfileByKeyOrID(ctx, billing.GetProfileByKeyOrIDInput(input))
		if err != nil {
			return err
		}

		if profile == nil {
			return billing.ValidationError{
				Err: fmt.Errorf("%w [id=%s]", billing.ErrProfileNotFound, input.IDOrKey),
			}
		}

		if profile.DeletedAt != nil {
			return billing.ValidationError{
				Err: fmt.Errorf("%w [id=%s]", billing.ErrProfileAlreadyDeleted, profile.ID),
			}
		}

		return adapter.DeleteProfile(ctx, billing.DeleteProfileAdapterInput{
			Namespace: input.Namespace,
			ID:        profile.ID,
		})
	})
}

func (s *Service) UpdateProfile(ctx context.Context, input billing.UpdateProfileInput) (*billing.Profile, error) {
	if err := input.Validate(); err != nil {
		return nil, billing.ValidationError{
			Err: err,
		}
	}

	return billing.WithTx(ctx, s.adapter, func(ctx context.Context, adapter billing.TxAdapter) (*billing.Profile, error) {
		profile, err := adapter.GetProfileByID(ctx, billing.GetProfileByIDAdapterInput{
			Namespace: input.Namespace,
			ID:        input.ID,
		})
		if err != nil {
			return nil, err
		}

		if profile == nil {
			return nil, billing.ValidationError{
				Err: fmt.Errorf("%w [id=%s]", billing.ErrProfileNotFound, input.ID),
			}
		}

		if profile.DeletedAt != nil {
			return nil, billing.ValidationError{
				Err: fmt.Errorf("%w [id=%s]", billing.ErrProfileAlreadyDeleted, input.ID),
			}
		}

		if profile.UpdatedAt != input.UpdatedAt {
			return nil, billing.UpdateAfterDeleteError{
				Err: fmt.Errorf("%w [id=%s]", billing.ErrProfileUpdateAfterDelete, input.ID),
			}
		}

		if !profile.Default && input.Default {
			defaultProfile, err := adapter.GetDefaultProfile(ctx, billing.GetDefaultProfileInput{
				Namespace: input.Namespace,
			})
			if err != nil {
				return nil, err
			}

			if defaultProfile != nil && defaultProfile.ID != input.ID {
				return nil, billing.ValidationError{
					Err: fmt.Errorf("%w [id=%s]", billing.ErrDefaultProfileAlreadyExists, defaultProfile.ID),
				}
			}
		}

		profile, err = adapter.UpdateProfile(ctx, billing.UpdateProfileAdapterInput{
			TargetState:      billing.Profile(input),
			WorkflowConfigID: profile.WorkflowConfig.ID,
		})
		if err != nil {
			return nil, err
		}

		if err := profile.Validate(); err != nil {
			return nil, billing.ValidationError{
				Err: fmt.Errorf("error validating profile: %w", err),
			}
		}

		return profile, nil
	})
}
