package models

type project struct {
	Year 		uint   `json:"year"`
	Site 		string `json:"site"`
	Description string `json:"description"`
}

var portfolio = []project{
	{
		2012,
		"http://DunkelBeer.ru",
		"Промо-сайт темного пива Dunkel от немецкого производителя Löwenbraü выпускаемого в России пивоваренной компанией \"CАН ИнБев\".",
	},
	{
		2012,
		"http://ZopoMobile.ru",
		"Русскоязычный каталог китайских телефонов компании Zopo на базе Android OS и аксессуаров к ним.",
	},
	{
		2012,
		"http://GeekWear.ru",
		"Интернет-магазин брендовой одежды для гиков.",
	},
	{
		2011,
		"http://РоналВарвар.рф",
		"Промо-сайт мультфильма \"Ронал-варвар\" от норвежских режиссеров. Мультфильм о самом нетипичном варваре на Земле, переполненный интересными приключениями и забавными ситуациями.",
	},
	{
		2011,
		"http://TompsonTatoo.ru",
		"Персональный сайт-блог художника-татуировщика Алексея Томпсона из Санкт-Петербурга.",
	},
	{
		2011,
		"http://DaftState.ru",
		"Страничка музыкальных и сануд продюсеров из команды \"DaftState\", работающих в стилях BreakBeat и BigBeat.",
	},
	{
		2011,
		"http://TiltPeople.ru",
		"Сайт сообщества фотографов в стиле Tilt Shif.",
	},
	{
		2011,
		"http://AbsurdGames.ru",
		"Страничка российской команды разработчиков независимых игр с необычной физикой и сюрреалистической графикой.",
	},
}

func GetPortfolio() []project {
	return portfolio
}
