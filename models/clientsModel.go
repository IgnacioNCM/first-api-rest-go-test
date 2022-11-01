package models

// tiene que ir con capital letter (mayusucla) para que funque el modelo
type Client struct {
	Id_cliente            uint `cliente:"primaryKey"`
	Nombre_cliente        string
	Apellido_cliente      string
	Email_cliente         string
	Codigo_cliente        string
	Telefono_cliente      string
	Empresa_cliente       string
	Nota_cliente          string
	Facebook_cliente      string
	Twitter_cliente       string
	Posee_credito         uint
	Credito_maximo        float64
	Estado_cliente        uint
	Bloqueado_comercial   uint
	Id_forma_pago         uint
	Giro_cliente          string
	Ciudad_cliente        string
	Comuna_cliente        string
	Direccion_cliente     string
	Tipo_cliente          uint
	Id_condicion_venta    uint
	Envia_dte             uint
	Id_cliente_prestashop uint
	Cliente_extranjero    uint
	Id_lista_precio       uint
	Total_puntos          float64
	//Ultimo_update_puntos  sql.NullTime
	Acumula_puntos uint
	//Created_at            time.Time
	//Updated_at            time.Time
	Postal_code string
}

//para que busque el nombre en singular en la bd
func (Client) TableName() string { return "cliente" }
