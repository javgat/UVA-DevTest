CREATE DATABASE IF NOT EXISTS uva_devtest;
USE uva_devtest;

DROP TABLE IF EXISTS RespuestaString;
DROP TABLE IF EXISTS RespuestaOpcion;
DROP TABLE IF EXISTS RespuestaPregunta;
DROP TABLE IF EXISTS RespuestaExamen;
DROP TABLE IF EXISTS PreguntaEtiqueta;
DROP TABLE IF EXISTS Etiqueta;
DROP TABLE IF EXISTS Opcion;
DROP TABLE IF EXISTS PreguntaEquipo;
DROP TABLE IF EXISTS TestPreguntaEditables;
DROP TABLE IF EXISTS PreguntaCodigo;
DROP TABLE IF EXISTS PreguntaString;
DROP TABLE IF EXISTS PreguntaOpciones;
DROP TABLE IF EXISTS Pregunta;
DROP TABLE IF EXISTS InvitacionTestUsuario;
DROP TABLE IF EXISTS InvitacionTestEquipo;
DROP TABLE IF EXISTS GestionTestEquipo;
DROP TABLE IF EXISTS Test;
DROP TABLE IF EXISTS EquipoUsuario;
DROP TABLE IF EXISTS Equipo;
DROP TABLE IF EXISTS Usuario;


CREATE TABLE Usuario (
  id int(11) NOT NULL AUTO_INCREMENT,
  username varchar(40) COLLATE utf8_unicode_ci NOT NULL,
  email varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  pwhash longtext COLLATE utf8_unicode_ci NOT NULL,
  rol ENUM('administrador', 'profesor', 'estudiante') NOT NULL,
  fullname varchar(200) COLLATE utf8_unicode_ci,
  UNIQUE(username),
  UNIQUE(email),
  PRIMARY KEY (id)
);

CREATE TABLE Equipo (
  id int(11) NOT NULL AUTO_INCREMENT,
  teamname varchar(40) COLLATE utf8_unicode_ci NOT NULL,
  description longtext COLLATE utf8_unicode_ci NOT NULL,
  soloProfesores boolean NOT NULL,
  UNIQUE(teamname),
  PRIMARY KEY(id)
);

CREATE TABLE EquipoUsuario(
  usuarioid int(11) NOT NULL,
  equipoid int(11) NOT NULL,
  rol ENUM('admin', 'miembro') NOT NULL,
  FOREIGN KEY(usuarioid) REFERENCES Usuario(id) ON DELETE CASCADE,
  FOREIGN KEY(equipoid) REFERENCES Equipo(id) ON DELETE CASCADE,
  CONSTRAINT CompKey_EquipoUsuario PRIMARY KEY(usuarioid, equipoid)
);
/* admin pass = admin1 */
INSERT INTO Usuario(username, email, pwhash, rol, fullname) VALUES('admin', 'admin@mail.com', '$2a$14$C0gTluZGQVbau5vcsaB72e0iwiECRIJvCgwNk4cn7IFlEJEMFwuVC', 'administrador', 'admin');

CREATE TABLE Test(
  id int(11) NOT NULL AUTO_INCREMENT,
  title varchar(40) COLLATE utf8_unicode_ci NOT NULL,
  description longtext COLLATE utf8_unicode_ci NOT NULL,
  maxSeconds int(11) NOT NULL,
  accesoPublico boolean NOT NULL,
  editable boolean NOT NULL,
  usuarioid int(11) NOT NULL,
  FOREIGN KEY(usuarioid) REFERENCES Usuario(id),
  PRIMARY KEY(id)
);
CREATE TABLE GestionTestEquipo(
  equipoid int(11) NOT NULL,
  testid int(11) NOT NULL,
  FOREIGN KEY(equipoid) REFERENCES Equipo(id) ON DELETE CASCADE,
  FOREIGN KEY(testid) REFERENCES Test(id) ON DELETE CASCADE,
  CONSTRAINT CompKey_GestionTestEquipo PRIMARY KEY(equipoid, testid)
);
CREATE TABLE InvitacionTestEquipo(
  equipoid int(11) NOT NULL,
  testid int(11) NOT NULL,
  FOREIGN KEY(equipoid) REFERENCES Equipo(id) ON DELETE CASCADE,
  FOREIGN KEY(testid) REFERENCES Test(id) ON DELETE CASCADE,
  CONSTRAINT CompKey_InvitacionTestEquipo PRIMARY KEY(equipoid, testid)
);
CREATE TABLE InvitacionTestUsuario(
  usuarioid int(11) NOT NULL,
  testid int(11) NOT NULL,
  FOREIGN KEY(usuarioid) REFERENCES Usuario(id) ON DELETE CASCADE,
  FOREIGN KEY(testid) REFERENCES Test(id) ON DELETE CASCADE,
  CONSTRAINT CompKey_InvitacionTestUsuario PRIMARY KEY(usuarioid, testid)
);

CREATE TABLE Pregunta(
  id int(11) NOT NULL AUTO_INCREMENT,
  title varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  question longtext COLLATE utf8_unicode_ci NOT NULL,
  estimatedTime int(11),
  autoCorrect boolean NOT NULL,
  editable boolean NOT NULL,
  usuarioid int(11) NOT NULL,
  testid int(11),
  FOREIGN KEY(usuarioid) REFERENCES Usuario(id) ON DELETE CASCADE,
  FOREIGN KEY(testid) REFERENCES Test(id),
  PRIMARY KEY(id)
);
CREATE TABLE PreguntaOpciones(
  id int(11) NOT NULL,
  eleccionUnica boolean NOT NULL,
  FOREIGN KEY(id) REFERENCES Pregunta(id) ON DELETE CASCADE,
  PRIMARY KEY(id)
);
CREATE TABLE PreguntaString(
  id int(11) NOT NULL,
  solucion varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  FOREIGN KEY(id) REFERENCES Pregunta(id) ON DELETE CASCADE,
  PRIMARY KEY(id)
);
CREATE TABLE PreguntaCodigo(
  id int(11) NOT NULL,
  FOREIGN KEY(id) REFERENCES Pregunta(id) ON DELETE CASCADE,
  PRIMARY KEY(id)
);
CREATE TABLE TestPreguntaEditables(
  testid int(11) NOT NULL,
  preguntaid int(11) NOT NULL,
  FOREIGN KEY(testid) REFERENCES Test(id) ON DELETE CASCADE,
  FOREIGN KEY(preguntaid) REFERENCES Pregunta(id) ON DELETE CASCADE,
  CONSTRAINT CompKey_TestPreguntaEditables PRIMARY KEY(testid, preguntaid)
);
CREATE TABLE PreguntaEquipo(
  preguntaid int(11) NOT NULL,
  equipoid int(11) NOT NULL,
  FOREIGN KEY(preguntaid) REFERENCES Pregunta(id) ON DELETE CASCADE,
  FOREIGN KEY(equipoid) REFERENCES Equipo(id) ON DELETE CASCADE,
  CONSTRAINT CompKey_PreguntaEquipo PRIMARY KEY(preguntaid, equipoid)
);

CREATE TABLE Opcion(
  indice int(11) NOT NULL,
  texto longtext COLLATE utf8_unicode_ci NOT NULL,
  correcta boolean NOT NULL,
  preguntaOpcionesid int(11) NOT NULL,
  FOREIGN KEY(preguntaOpcionesid) REFERENCES PreguntaOpciones(id) ON DELETE CASCADE,
  CONSTRAINT CompKey_Opcion PRIMARY KEY(preguntaOpcionesid, indice)
);

CREATE TABLE Etiqueta(
  nombre varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  PRIMARY KEY(nombre)
);
CREATE TABLE PreguntaEtiqueta(
  preguntaid int(11) NOT NULL,
  etiquetanombre varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  FOREIGN KEY(preguntaid) REFERENCES Pregunta(id) ON DELETE CASCADE,
  FOREIGN KEY(etiquetanombre) REFERENCES Etiqueta(nombre) ON DELETE CASCADE,
  CONSTRAINT CompKey_PreguntaEtiqueta PRIMARY KEY(preguntaid, etiquetanombre)
);

CREATE TABLE RespuestaExamen(
  id int(11) NOT NULL AUTO_INCREMENT,
  /*startTime, ESTA AUN NO*/
  testid int(11) NOT NULL,
  usuarioid int(11) NOT NULL,
  FOREIGN KEY(testid) REFERENCES Test(id) ON DELETE CASCADE,
  FOREIGN KEY(usuarioid) REFERENCES Usuario(id) ON DELETE CASCADE,
  PRIMARY KEY(id)
);

CREATE TABLE RespuestaPregunta(
  respuestaExamenid int(11) NOT NULL,
  preguntaid int(11) NOT NULL,
  FOREIGN KEY(respuestaExamenid) REFERENCES RespuestaExamen(id) ON DELETE CASCADE,
  FOREIGN KEY(preguntaid) REFERENCES Pregunta(id) ON DELETE CASCADE,
  CONSTRAINT CompKey_RespuestaPregunta PRIMARY KEY(respuestaExamenid, preguntaid)
);
CREATE TABLE RespuestaOpcion(
  respuestaExamenid int(11) NOT NULL,
  preguntaid int(11) NOT NULL,
  opcionindice int(11) NOT NULL,
  opcionPreguntaOpciones int(11) NOT NULL,
  CONSTRAINT fk_RespuestaOpcion_Opcion
      FOREIGN KEY(opcionPreguntaOpciones, opcionindice)
      REFERENCES Opcion(preguntaOpcionesid, indice) ON DELETE CASCADE,
  CONSTRAINT fk_RespuestaOpcion_RespuestaPregunta
      FOREIGN KEY(respuestaExamenid, preguntaid)
      REFERENCES RespuestaPregunta(respuestaExamenid, preguntaid) ON DELETE CASCADE,
  CONSTRAINT CompKey_RespuestaOpcion PRIMARY KEY(respuestaExamenid, preguntaid)
);
CREATE TABLE RespuestaString(
  respuestaExamenid int(11) NOT NULL,
  preguntaid int(11) NOT NULL,
  respuesta varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  CONSTRAINT fk_RespuestaString_RespuestaPregunta
      FOREIGN KEY(respuestaExamenid, preguntaid)
      REFERENCES RespuestaPregunta(respuestaExamenid, preguntaid) ON DELETE CASCADE,
  CONSTRAINT CompKey_RespuestaString PRIMARY KEY(respuestaExamenid, preguntaid)
);
