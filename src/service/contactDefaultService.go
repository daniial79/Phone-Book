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

func (s ContactDefaultService) NewContact(request dto.NewContactRequest) (*dto.NewContactResponse, *errs.AppError) {
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

func (s ContactDefaultService) AddNewNumbers(request []dto.AddNumberRequest, contactId string) ([]dto.AddNumberResponse, *errs.AppError) {
	coreTypedNumbers := make([]core.Number, len(request))

	for i, numberRequest := range request {
		coreTypedNumbers[i] = core.Number{
			Id:          "",
			ContactId:   contactId,
			PhoneNumber: numberRequest.Number,
			Label:       numberRequest.Label,
		}
	}

	addedNumbers, err := s.repo.AddNewNumber(coreTypedNumbers)
	if err != nil {
		return nil, err
	}

	response := make([]dto.AddNumberResponse, len(addedNumbers))
	for i, number := range addedNumbers {
		response[i] = number.ToAddNumberResponseDto()
	}

	return response, nil
}

func (s ContactDefaultService) GetContacts(options map[string]string) ([]dto.NewContactResponse, *errs.AppError) {
	coreTypedContacts, err := s.repo.GetAllContacts(options)
	if err != nil {
		return nil, err
	}

	response := make([]dto.NewContactResponse, len(coreTypedContacts))
	for i, contact := range coreTypedContacts {
		response[i] = dto.NewContactResponse{
			Id:        contact.Id,
			FirstName: contact.FirstName,
			LastName:  contact.LastName,
		}
	}

	return response, nil
}

func (s ContactDefaultService) AddNewEmails(request []dto.AddEmailRequest, contactId string) ([]dto.AddEmailResponse, *errs.AppError) {
	coreTypedEmails := make([]core.Email, len(request))

	for i, email := range request {
		coreTypedEmails[i] = core.Email{
			Id:        "",
			ContactId: contactId,
			Address:   email.Address,
		}
	}

	addedEmails, err := s.repo.AddNewEmails(coreTypedEmails)
	if err != nil {
		return nil, err
	}

	response := make([]dto.AddEmailResponse, len(coreTypedEmails))
	for i, email := range addedEmails {
		response[i] = email.ToAddEmailResponseDto()
	}

	return response, nil
}
