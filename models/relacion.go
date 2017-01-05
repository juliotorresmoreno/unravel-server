package models

import "time"

// Relacion modelo de usuario
type Relacion struct {
	Id                   uint   `xorm:"bigint not null autoincr pk" json:"id"`
	Usuario              string `xorm:"varchar(100) not null unique index" valid:"required" json:"usuario"`
	Email                string `xorm:"varchar(200)" valid:"email" json:"email"`
	PermisoEmail         string `xorm:"varchar(20)" json:"permiso_email"`
	NacimientoDia        string `xorm:"varchar(2)" json:"nacimiento_dia"`
	NacimientoMes        string `xorm:"varchar(2)" json:"nacimiento_mes"`
	PermisoNacimientoDia string `xorm:"varchar(20)" json:"permiso_nacimiento_dia"`
	NacimientoAno        string `xorm:"varchar(4)" json:"nacimiento_ano"`
	PermisoNacimientoAno string `xorm:"varchar(20)" json:"permiso_nacimiento_ano"`
	Sexo                 string `xorm:"varchar(4)" json:"sexo"`
	PermisoSexo          string `xorm:"varchar(20)" json:"permiso_sexo"`

	NacimientoPais          string `xorm:"varchar(200)" json:"nacimiento_pais"`
	PermisoNacimientoPais   string `xorm:"varchar(20)" json:"permiso_nacimiento_pais"`
	NacimientoCiudad        string `xorm:"varchar(200)" json:"nacimiento_ciudad"`
	PermisoNacimientoCiudad string `xorm:"varchar(20)" json:"permiso_nacimiento_ciudad"`
	ResidenciaPais          string `xorm:"varchar(200)" json:"residencia_pais"`
	PermisoResidenciaPais   string `xorm:"varchar(20)" json:"permiso_residencia_pais"`
	ResidenciaCiudad        string `xorm:"varchar(200)" json:"residencia_ciudad"`
	PermisoResidenciaCiudad string `xorm:"varchar(20)" json:"permiso_residencia_ciudad"`
	Direccion               string `xorm:"varchar(20)" json:"direccion"`
	PermisoDireccion        string `xorm:"varchar(20)" json:"permiso_direccion"`
	Telefono                string `xorm:"varchar(20)" json:"telefono"`
	PermisoTelefono         string `xorm:"varchar(20)" json:"permiso_telefono"`
	Celular                 string `xorm:"varchar(20)" json:"celular"`
	PermisoCelular          string `xorm:"varchar(20)" json:"permiso_celular"`

	Personalidad               string `xorm:"text" json:"personalidad"`
	PermisoPersonalidad        string `xorm:"varchar(20)" json:"permiso_personalidad"`
	Intereses                  string `xorm:"text" json:"intereses"`
	PermisoIntereses           string `xorm:"varchar(20)" json:"permiso_intereses"`
	Series                     string `xorm:"text" json:"series"`
	PermisoSeries              string `xorm:"varchar(20)" json:"permiso_series"`
	Musica                     string `xorm:"text" json:"musica"`
	PermisoMusica              string `xorm:"varchar(20)" json:"permiso_musica"`
	CreenciasReligiosas        string `xorm:"text" json:"creencias_religiosas"`
	PermisoCreenciasReligiosas string `xorm:"varchar(20)" json:"permiso_creencias_religiosas"`
	CreenciasPoliticas         string `xorm:"text" json:"creencias_politicas"`
	PermisoCreenciasPoliticas  string `xorm:"varchar(20)" json:"permiso_creencias_politicas"`

	CreateAt time.Time `xorm:"created" json:"create_at"`
	UpdateAt time.Time `xorm:"updated" json:"update_at"`
}

// TableName establece el nombre de la tabla que usara el modelo
func (u Relacion) TableName() string {
	return "relacion"
}

func init() {
	orm.Sync2(new(Profile))
}
