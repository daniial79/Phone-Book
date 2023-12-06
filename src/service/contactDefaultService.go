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

func NewContactDefaultService(repo core.ContactRepositoryDb) ContactDefaultService {
	return ContactDefaultService{repo}
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

func (s ContactDefaultService) GetContacts() ([]dto.NewContactResponse, *errs.AppError) {
	coreTypedContacts, err := s.repo.GetAllContacts()
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

func (s ContactDefaultService) GetContactCredentials(cId string) (*dto.ContactCredentialsResponse, *errs.AppError) {
	coreTypedNumbers, err := s.repo.GetContactNumbers(cId)
	if err != nil {
		return nil, err
	}

	coreTypedEmails, err := s.repo.GetContactEmails(cId)
	if err != nil {
		return nil, err
	}

	phoneNumberResponse := make([]dto.PhoneNumberResponse, 0)
	emailResponse := make([]dto.EmailResponse, 0)

	for _, n := range coreTypedNumbers {
		phoneNumberResponse = append(phoneNumberResponse,
			n.ToPhoneNumberResponseDto(),
		)
	}

	for _, e := range coreTypedEmails {
		emailResponse = append(emailResponse,
			e.ToEmailResponseDto(),
		)
	}

	response := dto.ContactCredentialsResponse{
		PhoneNumbers: phoneNumberResponse,
		Emails:       emailResponse,
	}

	return &response, nil
}

func (s ContactDefaultService) DeleteEmailFromContact(cId, eId string) (*dto.NoContentResponse, *errs.AppError) {
	if err := s.repo.DeleteContactEmail(cId, eId); err != nil {
		return nil, err
	}

	response := dto.NoContentResponse{}

	return &response, nil
}

func (s ContactDefaultService) DeletePhoneNumberFromContact(cId, eId string) (*dto.NoContentResponse, *errs.AppError) {
	if err := s.repo.DeleteContactPhoneNumber(cId, eId); err != nil {
		return nil, err
	}

	response := dto.NoContentResponse{}

	return &response, nil
}

func (s ContactDefaultService) DeleteContact(cId string) (*dto.NoContentResponse, *errs.AppError) {
	if err := s.repo.DeleteContact(cId); err != nil {
		return nil, err
	}
	response := dto.NoContentResponse{}

	return &response, nil
}

func (s ContactDefaultService) UpdateContactNumber(cId, nId string, request dto.UpdateContactNumberRequest) (*dto.UpdateContactNumberResponse, *errs.AppError) {
	coreTypedNumber := core.Number{
		Id:          nId,
		ContactId:   cId,
		PhoneNumber: request.NewPhoneNumber,
		Label:       request.NewLabel,
	}

	updatedCoreTypedNumber, err := s.repo.UpdateContactPhoneNumber(coreTypedNumber)
	if err != nil {
		return nil, err
	}

	response := updatedCoreTypedNumber.ToUpdatedNumberResponseDto()

	return &response, nil
}
