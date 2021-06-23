package dao

import (
	"time"

	"github.com/go-openapi/strfmt"
)

const (

	// TeamRoleRoleAdmin captures enum value "admin"
	TeamRoleRoleAdmin string = "admin"

	// TeamRoleRoleMember captures enum value "miembro"
	TeamRoleRoleMember string = "miembro"
)

// User user
//
// userdao User
type User struct {
	// id
	// Identificador SQL
	ID int64

	// email
	// Example: carlos@mail.com
	// Required: true
	// Format: email
	Email *strfmt.Email `json:"email"`

	// pwhash
	// Example: e$ia9s7ATDGba39pakscAKs
	// Required: true
	Pwhash *string `json:"pwhash"`

	// username
	// Example: carlosg72
	// Required: true
	// Pattern: [^@]+
	Username *string `json:"username"`

	// fullname
	// Example: Javier Gatón Herguedas
	// Required: true
	Fullname *string `json:"fullname"`

	// tiporoidl
	// Example: 2
	TipoRolId *int64
}

// Team team
//
// teamdao team
type Team struct {
	// id
	// Identificador SQL
	ID int64

	// description
	// Example: DevTest Team
	Description string `json:"description,omitempty"`

	// solo profesores
	// Example: true
	// Required: true
	SoloProfesores *bool `json:"soloProfesores"`

	// teamname
	// Example: devtestTeam
	// Required: true
	// Pattern: ^[^@ \t\r\n]+$
	Teamname *string `json:"teamname"`
}

// Question question
//
// swagger:model Question
type Question struct {

	// acceso publico no publicada
	// Example: false
	// Required: true
	AccesoPublicoNoPublicada *bool `json:"accesoPublicoNoPublicada"`

	// auto correct
	// Example: true
	// Required: true
	AutoCorrect *bool `json:"autoCorrect"`

	// cantidad favoritos
	// Minimum: 0
	CantidadFavoritos *int64 `json:"cantidadFavoritos,omitempty"`

	// editable
	// Example: false
	// Required: true
	Editable *bool `json:"editable"`

	// eleccion unica
	// Example: false
	EleccionUnica bool `json:"eleccionUnica,omitempty"`

	// estimated time
	// Example: 32600
	// Required: true
	// Minimum: 0
	EstimatedTime *int64 `json:"estimatedTime"`

	// id
	// Example: 1
	ID int64 `json:"id,omitempty"`

	// next Id
	// Example: 4
	NextID int64 `json:"nextId,omitempty"`

	// penalizacion
	// Required: true
	// Maximum: 100
	// Minimum: 0
	Penalizacion *int64 `json:"penalizacion"`

	// posicion
	// Example: 5
	Posicion int64 `json:"posicion,omitempty"`

	// prev Id
	// Example: 1
	PrevID int64 `json:"prevId,omitempty"`

	// question
	// Example: ¿Cual es el lenguaje que tiene un nombre más largo de todos?
	// Required: true
	Question *string `json:"question"`

	// solucion
	// Example: Javadoc
	Solucion string `json:"solucion,omitempty"`

	// tipo pregunta
	// Required: true
	// Enum: [opciones string codigo]
	TipoPregunta *string `json:"tipoPregunta"`

	// title
	// Example: Paralelismo en C
	// Required: true
	Title *string `json:"title"`

	// usuarioid
	// Example: 3
	// Required: true
	Usuarioid int64 `json:"usuarioid"`

	// valor final
	// Required: true
	// Minimum: 0
	ValorFinal *int64 `json:"valorFinal"`
}

// Test test
//
// swagger:model Test
type Test struct {

	// acceso publico
	// Example: true
	// Required: true
	AccesoPublico *bool `json:"accesoPublico"`

	// acceso publico no publicado
	// Example: true
	// Required: true
	AccesoPublicoNoPublicado *bool `json:"accesoPublicoNoPublicado"`

	// auto correct
	// Example: true
	// Required: true
	AutoCorrect *bool `json:"autoCorrect"`

	// cantidad favoritos
	// Minimum: 0
	CantidadFavoritos *int64 `json:"cantidadFavoritos,omitempty"`

	// description
	// Example: En este test se evaluaran los conocimientos respecto al lenguaje de programación Java
	// Required: true
	Description *string `json:"description"`

	// editable
	// Example: false
	// Required: true
	Editable *bool `json:"editable"`

	// id
	// Example: 1
	ID int64 `json:"id,omitempty"`

	// Maximum of tries that a user has solving a test. If <1, there is no limit of tries.
	// Required: true
	MaxIntentos *int64 `json:"maxIntentos"`

	// max minutes
	// Example: 60
	// Required: true
	// Minimum: 0
	MaxMinutes *int64 `json:"maxMinutes"`

	// title
	// Example: Test de introduccion a Java
	// Required: true
	Title *string `json:"title"`

	// usuarioid
	// Example: 3
	// Required: true
	Usuarioid int64 `json:"usuarioid"`

	// hora creacion
	// Example: 2021-02-25 14:44:55
	// Format: date-time
	HoraCreacion strfmt.DateTime `json:"horaCreacion,omitempty"`

	// nota maxima
	// Example: 10
	NotaMaxima int64 `json:"notaMaxima,omitempty"`

	// original test ID
	// Example: 15
	OriginalTestID *int64 `json:"originalTestID,omitempty"`

	// tiempo estricto
	// Example: true
	// Required: true
	TiempoEstricto *bool `json:"tiempoEstricto"`

	// visibilidad
	// Required: true
	// Enum: [alEntregar alCorregir manual]
	Visibilidad *string `json:"visibilidad"`
}

// Answer answer
//
// swagger:model Answer
type Answer struct {

	// corregida
	// Example: false
	Corregida bool `json:"corregida,omitempty"`

	// entregado
	// Example: false
	// Required: true
	Entregado *bool `json:"entregado"`

	// finish time
	// Example: 2021-02-25 14:44:55
	// Format: date-time
	FinishTime strfmt.DateTime `json:"finishTime,omitempty"`

	// id
	// Example: 1
	ID int64 `json:"id,omitempty"`

	// puntuacion
	// Example: 8.7
	Puntuacion float64 `json:"puntuacion,omitempty"`

	// startime
	// Example: 2021-02-25 14:44:55
	// Format: date-time
	Startime strfmt.DateTime `json:"startime,omitempty"`

	// testid
	// Example: 343
	Testid int64 `json:"testid,omitempty"`

	// usuarioid
	// Example: 3
	// Required: true
	Usuarioid int64 `json:"usuarioid"`

	// visible para usuario
	// Example: true
	VisibleParaUsuario bool `json:"visibleParaUsuario,omitempty"`
}

// TeamRole team role
//
// swagger:model TeamRole
type TeamRole struct {

	// role
	// Required: true
	// Enum: [admin member]
	Role *string `json:"role"`
}

// Tag tag
//
// swagger:model Tag
type Tag struct {

	// tag
	// Example: Java
	// Required: true
	Tag *string `json:"tag"`
}

// QuestionAnswer question answer
//
// swagger:model QuestionAnswer
type QuestionAnswer struct {

	// corregida
	// Example: true
	// Required: true
	Corregida *bool `json:"corregida"`

	// error compilacion
	// Example: Error line 4 missing ';'
	ErrorCompilacion string `json:"errorCompilacion,omitempty"`

	// estado
	// Enum: [noProbado errorCompilacion ejecutando probado]
	Estado *string `json:"estado,omitempty"`

	// id pregunta
	// Example: 1
	// Required: true
	IDPregunta *int64 `json:"idPregunta"`

	// indices opciones
	IndicesOpciones []int64 `json:"indicesOpciones"`

	// id respuesta
	// Example: 1
	// Required: true
	IDRespuesta *int64 `json:"idRespuesta"`

	// Percentage of the max points given to the answer
	// Example: 1
	// Required: true
	// Maximum: 100
	// Minimum: -100
	Puntuacion *int64 `json:"puntuacion"`

	// respuesta
	// Example: Javadoc
	Respuesta string `json:"respuesta,omitempty"`
}

// Option option
//
// swagger:model Option
type Option struct {

	// correcta
	// Example: false
	Correcta *bool `json:"correcta"`

	// indice
	// Example: 1
	Indice int64 `json:"indice,omitempty"`

	// preguntaid
	// Example: 1
	Preguntaid int64 `json:"preguntaid,omitempty"`

	// texto
	// Example: Esta opcion es la buena
	// Required: true
	Texto *string `json:"texto"`
}

// Mail Token Mail tokens
type MailToken struct {

	// mailtoken
	// Required: true
	Mailtoken *string `json:"mailtoken"`

	Userid int64

	Caducidad time.Time
}

// TipoRol tipo rol
//
// swagger:model TipoRol
type TipoRol struct {

	// true if the users can administrate Answers
	// Required: true
	AdminAnswers *bool `json:"adminAnswers"`

	// true if the users can administrate Server Configuration parameters
	// Required: true
	AdminConfiguration *bool `json:"adminConfiguration"`

	// true if the users can administrate Non Published Questions
	// Required: true
	AdminEQuestions *bool `json:"adminEQuestions"`

	// true if the users can administrate Non Published Tests
	// Required: true
	AdminETests *bool `json:"adminETests"`

	// true if the users can administrate PublishedTests
	// Required: true
	AdminPTests *bool `json:"adminPTests"`

	// true if the users can administrate TipoRol permissions
	// Required: true
	AdminPermissions *bool `json:"adminPermissions"`

	// true if the users can administrate Teams
	// Required: true
	AdminTeams *bool `json:"adminTeams"`

	// true if the users can administrate Users
	// Required: true
	AdminUsers *bool `json:"adminUsers"`

	// true if the users can change anothers user TipoRol if its source priority is less important and target is at most equal
	// Required: true
	ChangeRoles *bool `json:"changeRoles"`

	// id
	// Required: true
	ID *int64 `json:"id"`

	// nombre
	// Example: Moderador
	// Required: true
	Nombre *string `json:"nombre"`

	// Maximum priority is 0, any greater value represents less priority
	// Required: true
	// Minimum: 0
	Prioridad *int64 `json:"prioridad"`

	// rol base
	// Required: true
	// Enum: [administrador profesor estudiante noRegistrado]
	RolBase *string `json:"rolBase"`

	// true if the users can have non published questions
	// Required: true
	TenerEQuestions *bool `json:"tenerEQuestions"`

	// true if the users can have non published questions
	// Required: true
	TenerETests *bool `json:"tenerETests"`

	// true if the users can have published tests
	// Required: true
	TenerPTests *bool `json:"tenerPTests"`

	// true if the users can have teams
	// Required: true
	TenerTeams *bool `json:"tenerTeams"`

	// true if new users will be this TipoRol at first
	// Required: true
	TipoInicial *bool `json:"tipoInicial"`

	// true if the users can see Answers
	// Required: true
	VerAnswers *bool `json:"verAnswers"`

	// true if the users can see Non Published Questions
	// Required: true
	VerEQuestions *bool `json:"verEQuestions"`

	// true if the users can see Non Published Tests
	// Required: true
	VerETests *bool `json:"verETests"`

	// true if the users can see Published Questions
	// Required: true
	VerPQuestions *bool `json:"verPQuestions"`

	// true if the users can see PublishedTests
	// Required: true
	VerPTests *bool `json:"verPTests"`
}

const (
	OrderByNewDate string = "newDate"

	OrderByOldDate string = "oldDate"

	OrderByMoreFav string = "moreFav"

	OrderByLessFav string = "lessFav"

	OrderByMoreTime string = "moreTime"

	OrderByLessTime string = "lessTime"
)

const (
	OrderByAnswerMorePuntuacion string = "morePuntuacion"
	OrderByAnswerLessPuntuacion string = "lessPuntuacion"
	OrderByAnswerMoreDuracion   string = "moreDuracion"
	OrderByAnswerLessDuracion   string = "lessDuracion"
	OrderByAnswerNewStartDate   string = "newStartDate"
	OrderByAnswerOldStartDate   string = "oldStartDate"
)

const (
	TagOrderByFirstAlpha string = "firstAlpha"

	TagOrderByLastAlpha string = "lastAlpha"

	TagOrderByMoreQuestion string = "moreQuestion"

	TagOrderByLessQuestion string = "lessQuestion"

	TagOrderByMoreTest string = "moreTest"

	TagOrderByLessTest string = "lessTest"
)

// CustomizedView customized view
//
// swagger:model CustomizedView
type CustomizedView struct {

	// mensaje inicio
	// Required: true
	MensajeInicio *string `json:"mensajeInicio"`

	// rol base
	// Required: true
	// Enum: [administrador profesor estudiante noRegistrado]
	RolBase *string `json:"rolBase"`
}

// TestPregunta test pregunta
//
// swagger:model TestPregunta
type TestPregunta struct {

	// posicion
	// Example: 5
	// Required: true
	Posicion *int64 `json:"posicion"`

	// valor final
	// Example: 1
	// Required: true
	// Minimum: 0
	ValorFinal *int64 `json:"valorFinal"`
}

// PTestUpdate p test update
//
// swagger:model PTestUpdate
type PTestUpdate struct {

	// acceso publico
	// Example: true
	// Required: true
	AccesoPublico *bool `json:"accesoPublico"`

	// auto correct
	// Example: true
	// Required: true
	AutoCorrect *bool `json:"autoCorrect"`

	// Maximum of tries that a user has solving a test. If <1, there is no limit of tries.
	// Required: true
	MaxIntentos *int64 `json:"maxIntentos"`

	// max minutes
	// Example: 60
	// Required: true
	// Minimum: 0
	MaxMinutes *int64 `json:"maxMinutes"`

	// tiempo estricto
	// Example: true
	// Required: true
	TiempoEstricto *bool `json:"tiempoEstricto"`

	// visibilidad
	// Required: true
	// Enum: [alEntregar alCorregir manual]
	Visibilidad *string `json:"visibilidad"`
}

// Prueba prueba
//
// swagger:model Prueba
type Prueba struct {

	// entrada
	// Example: 3 3 0 0
	// Required: true
	Entrada *string `json:"entrada"`

	// id
	// Example: 1
	ID int64 `json:"id,omitempty"`

	// post entrega
	// Example: true
	// Required: true
	PostEntrega *bool `json:"postEntrega"`

	// preguntaid
	// Example: 1
	Preguntaid int64 `json:"preguntaid,omitempty"`

	// salida
	// Example: 2 2
	// Required: true
	Salida *string `json:"salida"`

	// valor
	// Example: 1
	// Required: true
	// Minimum: 0
	Valor *int64 `json:"valor"`

	// visible
	// Example: true
	// Required: true
	Visible *bool `json:"visible"`
}

// Testing testing
//
// swagger:model Testing
type Testing struct {

	// pruebas superadas
	// Example: 5
	// Required: true
	PruebasSuperadas *int64 `json:"pruebasSuperadas"`

	// pruebas totales
	// Example: 15
	// Required: true
	PruebasTotales *int64 `json:"pruebasTotales"`
}

const (
	EstadoEjecucionCorrecto         = "correcto"
	EstadoEjecucionTiempoExcedido   = "tiempoExcedido"
	EstadoEjecucionErrorRuntime     = "errorRuntime"
	EstadoEjecucionSalidaIncorrecta = "salidaIncorrecta"
)

type Ejecucion struct {
	Pruebaid          *int64
	RespuestaExamenid *int64
	Preguntaid        *int64
	Estado            *string
	SalidaReal        *string
}
