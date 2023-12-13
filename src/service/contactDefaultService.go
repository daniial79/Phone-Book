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

func (s ContactDefaultService) NewContact(requestBody dto.NewContactRequest) (*dto.NewContactResponse, *errs.AppError) {
	coreTypedObject := new(core.Contact)

	if !requestBody.IsValid() {
		return nil, errs.NewUnProcessableErr("Unprocessable request")
	}

	coreTypedObject.Id = ""
	coreTypedObject.FirstName = requestBody.FirstName
	coreTypedObject.LastName = requestBody.LastName

	for _, number := range requestBody.PhoneNumbers {
		coreTypedObject.PhoneNumbers = append(coreTypedObject.PhoneNumbers, core.Number{
			Id:          "",
			ContactId:   "",
			PhoneNumber: number.Number,
			Label:       number.Label,
		})
	}

	for _, email := range requestBody.Emails {
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

func (s ContactDefaultService) AddNewNumbers(requestBody []dto.AddNumberRequest, contactId string) ([]dto.AddNumberResponse, *errs.AppError) {
	for _, r := range requestBody {
		if !r.IsValid() {
			return nil, errs.NewUnProcessableErr("Unprocessable request")
		}
	}

	coreTypedNumbers := make([]core.Number, len(requestBody))

	for i, numberRequest := range requestBody {
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

func (s ContactDefaultService) AddNewEmails(requestBody []dto.AddEmailRequest, contactId string) ([]dto.AddEmailResponse, *errs.AppError) {
	coreTypedEmails := make([]core.Email, len(requestBody))

	for _, r := range requestBody {
		if !r.IsValid() {
			return nil, errs.NewUnProcessableErr("Unprocessable request")
		}
	}

	for i, email := range requestBody {
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

func (s ContactDefaultService) UpdateContactNumber(cId, nId string, requestBody dto.UpdateNumberRequest) (*dto.UpdateNumberResponse, *errs.AppError) {

	if !requestBody.IsValid() {
		return nil, errs.NewUnProcessableErr("Unprocessable request")
	}

	coreTypedNumber := core.Number{
		Id:          nId,
		ContactId:   cId,
		PhoneNumber: requestBody.NewPhoneNumber,
		Label:       requestBody.NewLabel,
	}

	updatedCoreTypedNumber, err := s.repo.UpdateContactPhoneNumber(coreTypedNumber)
	if err != nil {
		return nil, err
	}

	response := updatedCoreTypedNumber.ToUpdatedNumberResponseDto()

	return &response, nil
}

func (s ContactDefaultService) UpdateContactEmail(cId, eId string, requestBody dto.UpdateEmailRequest) (*dto.UpdateEmailResponse, *errs.AppError) {

	if !requestBody.IsValid() {
		return nil, errs.NewUnProcessableErr("Unprocessable request")
	}

	coreTypedEmail := core.Email{
		Id:        eId,
		ContactId: cId,
		Address:   requestBody.NewAddress,
	}

	updatedCoreTypedEmail, err := s.repo.UpdateContactEmail(coreTypedEmail)
	if err != nil {
		return nil, err
	}

	response := updatedCoreTypedEmail.ToUpdatedEmailResponseDto()
	return &response, nil
}

func (s ContactDefaultService) UpdateContact(cId string, requestBody dto.UpdateContactRequest) (*dto.UpdateContactResponse, *errs.AppError) {

	if !requestBody.IsValid() {
		return nil, errs.NewUnProcessableErr("Unprocessable request")
	}

	coreTypedContact := core.Contact{
		Id:        cId,
		FirstName: requestBody.NewFirstName,
		LastName:  requestBody.NewLastName,
	}

	updatedContactCoredType, err := s.repo.UpdateContact(coreTypedContact)
	if err != nil {
		return nil, err
	}

	response := updatedContactCoredType.ToUpdatedContactResponseDto()
	return &response, nil
}
