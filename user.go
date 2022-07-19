package Panacea

type Patient struct {
	//СКОРЕЕ ВСЕГО ВСЕ ТО ЖЕ САМОЕ, ЧТО И В БД
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Telephone  string `json:"telephone"`
	DiplomaNum string `json:"diplomaNum"`
	Rating     int    `json:"rating"`
	PhotoId    int    `json:"photoId"`
	//ревью и блог
}

type Doctor struct {
	//СКОРЕЕ ВСЕГО ВСЕ ТО ЖЕ САМОЕ, ЧТО И В БД
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Telephone  string `json:"telephone"`
	DiplomaNum string `json:"diplomaNum"`
	Rating     int    `json:"rating"`
	PhotoId    int    `json:"photoId"`
	//ревью и блог
}
