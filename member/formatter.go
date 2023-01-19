package member

type formatterData struct {
	Nama  string `json:"nama"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func FormatterResponsData(member Member, token string) formatterData {
	formatterData := formatterData{
		Nama:  member.Nama,
		Email: member.Email,
		Token: token,
	}

	return formatterData
}
