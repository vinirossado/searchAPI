package models

type Destination struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Country     string      `json:"country"`
	places      int         `json:"places"`
	Population  int         `json:"population"`
	Language    string      `json:"language"`
	Capital     string      `json:"capital"`
	CountryCode string      `json:"countryCode"`
	Currency    Currency    `json:"currency"`
	Type        string      `json:"type"`
	Iso         string      `json:"iso"`
	Gmt         int         `json:"gmt"`
	Description interface{} `json:"description"`
}

// Destination(
//     topImage:
//         'https://www.brasilnaitalia.net/wp-content/uploads/2014/01/veneza-viagem.jpg',
//     image:
//         'https://media.tacdn.com/media/attractions-splice-spp-674x446/07/93/21/8e.jpg',
//     name: 'Veneza',
//     country: 'Italia',
//     description:
//         'Veneza, a capital da região de Vêneto, no norte da Itália, é formada por mais de 100 pequenas ilhas em uma lagoa no Mar Adriático. A cidade não tem estradas, apenas canais (como a via Grand Canal), repletos de palácios góticos e renascentistas. Na praça central, Piazza San Marco, ficam a Basílica de São Marcos, coberta de mosaicos bizantinos, e o campanário, com vista para os telhados vermelhos da cidade.',
//     places: 12,
//     currency: 'Euro',
//     type: 'Popular',
//   ),
