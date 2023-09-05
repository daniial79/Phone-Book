package service

import (
	"github.com/daniial79/Phone-Book/src/core"
	"github.com/daniial79/Phone-Book/src/dto"
	"github.com/daniial79/Phone-Book/src/errs"
)

// ContactDefaultService primary actor
type ContactDefaultService struct {
	repo core.ContactRepository
}

func NewContactDefaultService(repository core.ContactRepository) ContactDefaultService {
	return ContactDefaultService{repo: repository}
}

func (s ContactDefaultService) NewContact(request dto.NewContactRequest) (*dto.ContactResponse, *errs.AppError) {
	coreTypedObject := new(core.Contact)

	coreTypedObject.Id = ""
	coreTypedObject.FirstName = request.FirstName
	coreTypedObject.LastName = request.LastName

	for _, number := range request.PhoneNumbers {
		coreTypedObject.PhoneNumbers = append(coreTypedObject.PhoneNumbers, core.Number{
			Id:          "",
			ContactId:   "",
			PhoneNumber: number.Number,
			Label:       number.Label,
		})
	}

	for _, email := range request.Emails {
		coreTypedObject.Emails = append(coreTypedObject.Emails, core.Email{
			Id:        "",
			ContactId: "",
			Address:   email.Address,
		})
	}

	createdContact, err := s.repo.CreateContact(coreTypedObject)
	if err != nil {
		return nil, err
	}

	response := createdContact.ToContactResponseDto()
	return response, nil
}
