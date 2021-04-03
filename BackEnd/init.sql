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
DROP TABLE IF EXISTS TestPregunta;
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
  CONSTRAINT PRIMARY KEY(usuarioid, equipoid)
);

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
  CONSTRAINT PRIMARY KEY(equipoid, testid)
);
CREATE TABLE InvitacionTestEquipo(
  equipoid int(11) NOT NULL,
  testid int(11) NOT NULL,
  FOREIGN KEY(equipoid) REFERENCES Equipo(id) ON DELETE CASCADE,
  FOREIGN KEY(testid) REFERENCES Test(id) ON DELETE CASCADE,
  CONSTRAINT PRIMARY KEY(equipoid, testid)
);
CREATE TABLE InvitacionTestUsuario(
  usuarioid int(11) NOT NULL,
  testid int(11) NOT NULL,
  FOREIGN KEY(usuarioid) REFERENCES Usuario(id) ON DELETE CASCADE,
  FOREIGN KEY(testid) REFERENCES Test(id) ON DELETE CASCADE,
  CONSTRAINT PRIMARY KEY(usuarioid, testid)
);

CREATE TABLE Pregunta(
  id int(11) NOT NULL AUTO_INCREMENT,
  title varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  question longtext COLLATE utf8_unicode_ci NOT NULL,
  estimatedTime int(11),
  autoCorrect boolean NOT NULL,
  editable boolean NOT NULL,
  usuarioid int(11) NOT NULL,
  eleccionUnica boolean,
  solucion varchar(100) COLLATE utf8_unicode_ci,
  FOREIGN KEY(usuarioid) REFERENCES Usuario(id) ON DELETE CASCADE,
  PRIMARY KEY(id)
);

CREATE TABLE TestPregunta(
  testid int(11) NOT NULL,
  preguntaid int(11) NOT NULL,
  valorFinal int(11) NOT NULL,
  FOREIGN KEY(testid) REFERENCES Test(id) ON DELETE CASCADE,
  FOREIGN KEY(preguntaid) REFERENCES Pregunta(id) ON DELETE CASCADE,
  CONSTRAINT PRIMARY KEY(testid, preguntaid)
);
CREATE TABLE PreguntaEquipo(
  preguntaid int(11) NOT NULL,
  equipoid int(11) NOT NULL,
  FOREIGN KEY(preguntaid) REFERENCES Pregunta(id) ON DELETE CASCADE,
  FOREIGN KEY(equipoid) REFERENCES Equipo(id) ON DELETE CASCADE,
  CONSTRAINT PRIMARY KEY(preguntaid, equipoid)
);

CREATE TABLE Opcion(
  indice int(11) NOT NULL,
  texto longtext COLLATE utf8_unicode_ci NOT NULL,
  correcta boolean NOT NULL,
  preguntaid int(11) NOT NULL,
  FOREIGN KEY(preguntaid) REFERENCES Pregunta(id) ON DELETE CASCADE,
  CONSTRAINT PRIMARY KEY(preguntaid, indice)
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
  CONSTRAINT PRIMARY KEY(preguntaid, etiquetanombre)
);

CREATE TABLE RespuestaExamen(
  id int(11) NOT NULL AUTO_INCREMENT,
  startTime DateTime NOT NULL,
  finished boolean NOT NULL,
  testid int(11) NOT NULL,
  usuarioid int(11) NOT NULL,
  FOREIGN KEY(testid) REFERENCES Test(id) ON DELETE CASCADE,
  FOREIGN KEY(usuarioid) REFERENCES Usuario(id) ON DELETE CASCADE,
  PRIMARY KEY(id)
);

CREATE TABLE RespuestaPregunta(
  respuestaExamenid int(11) NOT NULL,
  preguntaid int(11) NOT NULL,
  puntuacion int(11),
  corregida boolean NOT NULL,
  FOREIGN KEY(respuestaExamenid) REFERENCES RespuestaExamen(id) ON DELETE CASCADE,
  FOREIGN KEY(preguntaid) REFERENCES Pregunta(id) ON DELETE CASCADE,
  CONSTRAINT PRIMARY KEY(respuestaExamenid, preguntaid)
);
CREATE TABLE RespuestaOpcion(
  respuestaExamenid int(11) NOT NULL,
  preguntaid int(11) NOT NULL,
  opcionindice int(11) NOT NULL,
  CONSTRAINT fk_RespuestaOpcion_Opcion
      FOREIGN KEY(preguntaid, opcionindice)
      REFERENCES Opcion(preguntaid, indice) ON DELETE CASCADE,
  CONSTRAINT fk_RespuestaOpcion_RespuestaPregunta
      FOREIGN KEY(respuestaExamenid, preguntaid)
      REFERENCES RespuestaPregunta(respuestaExamenid, preguntaid) ON DELETE CASCADE,
  CONSTRAINT PRIMARY KEY(respuestaExamenid, preguntaid)
);
CREATE TABLE RespuestaString(
  respuestaExamenid int(11) NOT NULL,
  preguntaid int(11) NOT NULL,
  respuesta varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  CONSTRAINT fk_RespuestaString_RespuestaPregunta
      FOREIGN KEY(respuestaExamenid, preguntaid)
      REFERENCES RespuestaPregunta(respuestaExamenid, preguntaid) ON DELETE CASCADE,
  CONSTRAINT PRIMARY KEY(respuestaExamenid, preguntaid)
);

/* DATOS INICIALES */

/* admin pass = admin1 */
INSERT INTO Usuario(username, email, pwhash, rol, fullname) VALUES('admin', 'admin@mail.com', '$2a$14$C0gTluZGQVbau5vcsaB72e0iwiECRIJvCgwNk4cn7IFlEJEMFwuVC', 'administrador', 'admin');