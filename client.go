package Panacea

type Client struct {
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
