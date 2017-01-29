package models

import "time"

// Profile modelo de usuario
type Profile struct {
	Id                         uint      `xorm:"bigint not null autoincr pk" json:"id"`
	Usuario                    string    `xorm:"varchar(100) not null unique index" valid:"required" json:"usuario"`
	Email                      string    `xorm:"varchar(200)" valid:"email" json:"email"`
	PermisoEmail               string    `xorm:"varchar(20)" json:"permiso_email" valid:"matches(^(private|friends|public)$)"`
	NacimientoDia              string    `xorm:"varchar(2)" json:"nacimiento_dia"`
	NacimientoMes              string    `xorm:"varchar(2)" json:"nacimiento_mes"`
	PermisoNacimientoDia       string    `xorm:"varchar(20)" json:"permiso_nacimiento_dia" valid:"matches(^(private|friends|public)$)"`
	NacimientoAno              string    `xorm:"varchar(4)" json:"nacimiento_ano"`
	PermisoNacimientoAno       string    `xorm:"varchar(20)" json:"permiso_nacimiento_ano" valid:"matches(^(private|friends|public)$)"`
	Sexo                       string    `xorm:"varchar(4)" json:"sexo"`
	PermisoSexo                string    `xorm:"varchar(20)" json:"permiso_sexo" valid:"matches(^(private|friends|public)$)"`
	NacimientoPais             string    `xorm:"varchar(200)" json:"nacimiento_pais"`
	PermisoNacimientoPais      string    `xorm:"varchar(20)" json:"permiso_nacimiento_pais" valid:"matches(^(private|friends|public)$)"`
	NacimientoCiudad           string    `xorm:"varchar(200)" json:"nacimiento_ciudad"`
	PermisoNacimientoCiudad    string    `xorm:"varchar(20)" json:"permiso_nacimiento_ciudad" valid:"matches(^(private|friends|public)$)"`
	ResidenciaPais             string    `xorm:"varchar(200)" json:"residencia_pais"`
	PermisoResidenciaPais      string    `xorm:"varchar(20)" json:"permiso_residencia_pais" valid:"matches(^(private|friends|public)$)"`
	ResidenciaCiudad           string    `xorm:"varchar(200)" json:"residencia_ciudad"`
	PermisoResidenciaCiudad    string    `xorm:"varchar(20)" json:"permiso_residencia_ciudad" valid:"matches(^(private|friends|public)$)"`
	Direccion                  string    `xorm:"varchar(20)" json:"direccion"`
	PermisoDireccion           string    `xorm:"varchar(20)" json:"permiso_direccion" valid:"matches(^(private|friends|public)$)"`
	Telefono                   string    `xorm:"varchar(20)" json:"telefono"`
	PermisoTelefono            string    `xorm:"varchar(20)" json:"permiso_telefono" valid:"matches(^(private|friends|public)$)"`
	Celular                    string    `xorm:"varchar(20)" json:"celular"`
	PermisoCelular             string    `xorm:"varchar(20)" json:"permiso_celular" valid:"matches(^(private|friends|public)$)"`
	Personalidad               string    `xorm:"text" json:"personalidad"`
	PermisoPersonalidad        string    `xorm:"varchar(20)" json:"permiso_personalidad" valid:"matches(^(private|friends|public)$)"`
	Intereses                  string    `xorm:"text" json:"intereses"`
	PermisoIntereses           string    `xorm:"varchar(20)" json:"permiso_intereses" valid:"matches(^(private|friends|public)$)"`
	Series                     string    `xorm:"text" json:"series"`
	PermisoSeries              string    `xorm:"varchar(20)" json:"permiso_series" valid:"matches(^(private|friends|public)$)"`
	Musica                     string    `xorm:"text" json:"musica"`
	PermisoMusica              string    `xorm:"varchar(20)" json:"permiso_musica" valid:"matches(^(private|friends|public)$)"`
	CreenciasReligiosas        string    `xorm:"text" json:"creencias_religiosas"`
	PermisoCreenciasReligiosas string    `xorm:"varchar(20)" json:"permiso_creencias_religiosas" valid:"matches(^(private|friends|public)$)"`
	CreenciasPoliticas         string    `xorm:"text" json:"creencias_politicas"`
	PermisoCreenciasPoliticas  string    `xorm:"varchar(20)" json:"permiso_creencias_politicas" valid:"matches(^(private|friends|public)$)"`
	CreateAt                   time.Time `xorm:"created" json:"create_at"`
	UpdateAt                   time.Time `xorm:"updated" json:"update_at"`
}

// TableName establece el nombre de la tabla que usara el modelo
func (self Profile) TableName() string {
	return "profile"
}

func init() {
	orm.Sync2(new(Profile))
}
