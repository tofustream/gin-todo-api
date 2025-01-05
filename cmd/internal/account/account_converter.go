package account

func fetchedAccountToDTO(fetchedAccount FetchedAccount) FetchedAccountDTO {
	return FetchedAccountDTO{
		ID:        fetchedAccount.ID().Value(),
		Email:     fetchedAccount.Email().Value(),
		CreatedAt: fetchedAccount.CreatedAt(),
		UpdatedAt: fetchedAccount.UpdatedAt(),
	}
}
