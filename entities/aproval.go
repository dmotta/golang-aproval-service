package entities

type Proyecto struct {
	IdProyecto    uint    `json:"idProyecto"`
	IdCliente     uint    `json:"idCliente"`
	Fecha         string  `json:"fecha"`
	Descripcion   string  `json:"descripcion"`
	Importe       float64 `json:"importe"`
	IdEstado      string  `json:"idEstado"`
	IdUsuario     uint    `json:"idUsuario"`
	FechaCreacion string  `json:"fechaCreacion"`
	Activo        bool    `json:"activo"`
}
