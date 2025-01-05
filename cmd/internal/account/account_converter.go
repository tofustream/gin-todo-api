package account

func FetchedAccountToDTO(fetchedAccount FetchedAccount) FetchedAccountDTO {
	return FetchedAccountDTO{
		ID:        fetchedAccount.ID().Value(),
		Email:     fetchedAccount.Email().Value(),
		CreatedAt: fetchedAccount.CreatedAt(),
		UpdatedAt: fetchedAccount.UpdatedAt(),
	}
}
