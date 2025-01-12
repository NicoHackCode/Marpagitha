package domain

type Student struct {
	ID                    string                `bson:"_id,omitempty"`
	Codigo                string                `bson:"codigo"`
	Nombre                string                `bson:"nombre"`
	Curso                 string                `bson:"curso"`
	Documento             Documento             `bson:"documento"`
	ResponsableTributario ResponsableTributario `bson:"responsable_tributario"`
	Padres                Padres                `bson:"padres"`
	Activo                bool                  `bson:"activo"`
}

type Documento struct {
	Tipo   string `bson:"tipo"`
	Numero string `bson:"numero"`
}

type ResponsableTributario struct {
	Nombre    string `bson:"nombre"`
	Cedula    string `bson:"cedula"`
	Correo    string `bson:"correo"`
	Celular   string `bson:"celular"`
	Profesion string `bson:"profesion"`
	Direccion string `bson:"direccion"`
}

type Padres struct {
	Padre DatosPadre `bson:"padre"`
	Madre DatosPadre `bson:"madre"`
}

type DatosPadre struct {
	Nombre    string `bson:"nombre"`
	Cedula    string `bson:"cedula"`
	Correo    string `bson:"correo"`
	Celular   string `bson:"celular"`
	Profesion string `bson:"profesion"`
	Direccion string `bson:"direccion"`
}
