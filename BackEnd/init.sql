CREATE DATABASE IF NOT EXISTS uva_devtest;
USE uva_devtest;

DROP TABLE IF EXISTS OpcionRespuesta;
DROP TABLE IF EXISTS RespuestaPregunta;
DROP TABLE IF EXISTS RespuestaExamen;
DROP TABLE IF EXISTS TestEtiqueta;
DROP TABLE IF EXISTS PreguntaEtiqueta;
DROP TABLE IF EXISTS Etiqueta;
DROP TABLE IF EXISTS Opcion;
DROP TABLE IF EXISTS PreguntaEquipo;
DROP TABLE IF EXISTS TestPregunta;
DROP TABLE IF EXISTS PreguntaFavorita;
DROP TABLE IF EXISTS Pregunta;
DROP TABLE IF EXISTS TestFavorito;
DROP TABLE IF EXISTS InvitacionTestUsuario;
DROP TABLE IF EXISTS InvitacionTestEquipo;
DROP TABLE IF EXISTS GestionTestEquipo;
DROP TABLE IF EXISTS Test;
DROP TABLE IF EXISTS EquipoUsuario;
DROP TABLE IF EXISTS Equipo;
DROP TABLE IF EXISTS TokenCorreo;
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

CREATE TABLE TokenCorreo(
  usuarioid int(11) NOT NULL,
  token varchar(40) COLLATE utf8_unicode_ci NOT NULL,
  caducidad DateTime NOT NULL,
  FOREIGN KEY(usuarioid) REFERENCES Usuario(id),
  PRIMARY KEY(token)
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
  maxMinutes int(11) NOT NULL,
  accesoPublico boolean NOT NULL,
  editable boolean NOT NULL,
  usuarioid int(11) NOT NULL,
  accesoPublicoNoPublicado boolean NOT NULL,
  horaCreacion DateTime NOT NULL,
  origenTestid int(11),
  notaMaxima int(11) NOT NULL,
  autoCorrect boolean NOT NULL,
  FOREIGN KEY(usuarioid) REFERENCES Usuario(id),
  FOREIGN KEY(origenTestid) REFERENCES Test(id),
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

CREATE TABLE TestFavorito(
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
  accesoPublicoNoPublicada boolean NOT NULL,
  penalizacion int(11) NOT NULL,
  CONSTRAINT CHK_penalizacion CHECK (penalizacion>=0 AND penalizacion<=100),
  FOREIGN KEY(usuarioid) REFERENCES Usuario(id) ON DELETE CASCADE,
  PRIMARY KEY(id)
);
CREATE TABLE PreguntaFavorita(
  usuarioid int(11) NOT NULL,
  preguntaid int(11) NOT NULL,
  FOREIGN KEY(usuarioid) REFERENCES Usuario(id) ON DELETE CASCADE,
  FOREIGN KEY(preguntaid) REFERENCES Pregunta(id) ON DELETE CASCADE,
  CONSTRAINT PRIMARY KEY(usuarioid, preguntaid)
);

CREATE TABLE TestPregunta(
  testid int(11) NOT NULL,
  preguntaid int(11) NOT NULL,
  valorFinal int(11) NOT NULL,
  CONSTRAINT CHK_valorFinal CHECK (valorFinal>=0),
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
  indice int(11) NOT NULL AUTO_INCREMENT,
  texto longtext COLLATE utf8_unicode_ci NOT NULL,
  correcta boolean NOT NULL,
  preguntaid int(11) NOT NULL,
  FOREIGN KEY(preguntaid) REFERENCES Pregunta(id) ON DELETE CASCADE,
  PRIMARY KEY(indice)
  /*CONSTRAINT PRIMARY KEY(preguntaid, indice)*/
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


CREATE TABLE TestEtiqueta(
  testid int(11) NOT NULL,
  etiquetanombre varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  FOREIGN KEY(testid) REFERENCES Test(id) ON DELETE CASCADE,
  FOREIGN KEY(etiquetanombre) REFERENCES Etiqueta(nombre) ON DELETE CASCADE,
  CONSTRAINT PRIMARY KEY(testid, etiquetanombre)
);

CREATE TABLE RespuestaExamen(
  id int(11) NOT NULL AUTO_INCREMENT,
  startTime DateTime NOT NULL,
  finishTime DateTime,
  entregado boolean NOT NULL,
  testid int(11) NOT NULL,
  usuarioid int(11) NOT NULL,
  puntuacion DEC(12,3),
  corregida boolean NOT NULL,
  FOREIGN KEY(testid) REFERENCES Test(id) ON DELETE CASCADE,
  FOREIGN KEY(usuarioid) REFERENCES Usuario(id) ON DELETE CASCADE,
  PRIMARY KEY(id)
);

CREATE TABLE RespuestaPregunta(
  respuestaExamenid int(11) NOT NULL,
  preguntaid int(11) NOT NULL,
  puntuacion int(11) NOT NULL, /*porcentaje*/
  corregida boolean NOT NULL,
  respuesta longtext COLLATE utf8_unicode_ci,
  CONSTRAINT CHK_puntuacion CHECK (puntuacion>=-100 AND puntuacion<=100),
  FOREIGN KEY(respuestaExamenid) REFERENCES RespuestaExamen(id) ON DELETE CASCADE,
  FOREIGN KEY(preguntaid) REFERENCES Pregunta(id) ON DELETE CASCADE,
  CONSTRAINT PRIMARY KEY(respuestaExamenid, preguntaid)
);

CREATE TABLE OpcionRespuesta(
  respuestaExamenid int(11) NOT NULL,
  preguntaid int(11) NOT NULL,
  opcionindice int(11) NOT NULL,
  CONSTRAINT fk_ORRespPreg
    FOREIGN KEY(respuestaExamenid, preguntaid)
    REFERENCES RespuestaPregunta(respuestaExamenid, preguntaid) ON DELETE CASCADE,
  CONSTRAINT fk_OROpcion
    FOREIGN KEY(preguntaid, opcionindice)
    REFERENCES Opcion(preguntaid, indice) ON DELETE CASCADE,
  CONSTRAINT PRIMARY KEY(respuestaExamenid, preguntaid, opcionindice)
);

/* DATOS INICIALES */

/* admin pass = admin1 */
INSERT INTO Usuario(username, email, pwhash, rol, fullname) VALUES('admin', 'admin@mail.com', '$2a$14$C0gTluZGQVbau5vcsaB72e0iwiECRIJvCgwNk4cn7IFlEJEMFwuVC', 'administrador', 'admin');