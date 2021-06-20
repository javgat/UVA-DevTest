CREATE DATABASE IF NOT EXISTS uva_devtest;
USE uva_devtest;

DROP TABLE IF EXISTS Ejecucion;
DROP TABLE IF EXISTS Prueba;
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
DROP TABLE IF EXISTS TipoRol;
DROP TABLE IF EXISTS VistaPersonalizada;

CREATE TABLE VistaPersonalizada (
  rolBase ENUM('administrador', 'profesor', 'estudiante', 'noRegistrado') NOT NULL,
  mensajeInicio longtext COLLATE utf8_unicode_ci NOT NULL,
  PRIMARY KEY(rolBase)
);

INSERT INTO VistaPersonalizada(rolBase, mensajeInicio) VALUES('administrador', "Bienvenido Administrador");
INSERT INTO VistaPersonalizada(rolBase, mensajeInicio) VALUES('profesor', "Bienvenido Profesor");
INSERT INTO VistaPersonalizada(rolBase, mensajeInicio) VALUES('estudiante', "Bienvenido Estudiante");
INSERT INTO VistaPersonalizada(rolBase, mensajeInicio) VALUES('noRegistrado', "");

CREATE TABLE TipoRol (
  id int(11) NOT NULL AUTO_INCREMENT,
  rolBase ENUM('administrador', 'profesor', 'estudiante', 'noRegistrado') NOT NULL,
  nombre varchar(40) COLLATE utf8_unicode_ci NOT NULL,
  prioridad int(11) DEFAULT 1,
  verPTests boolean DEFAULT 1,
  verETests boolean DEFAULT 0,
  verEQuestions boolean DEFAULT 0,
  verPQuestions boolean DEFAULT 0,
  verAnswers boolean DEFAULT 0,
  changeRoles boolean DEFAULT 0,
  tenerTeams boolean DEFAULT 0,
  tenerEQuestions boolean DEFAULT 0,
  tenerETests boolean DEFAULT 0,
  tenerPTests boolean DEFAULT 0,
  adminPTests boolean DEFAULT 0,
  adminETests boolean DEFAULT 0,
  adminEQuestions boolean DEFAULT 0,
  adminAnswers boolean DEFAULT 0,
  adminUsers boolean DEFAULT 0,
  adminTeams boolean DEFAULT 0,
  adminConfiguration boolean DEFAULT 0,
  adminPermissions boolean DEFAULT 0,
  tipoInicial boolean DEFAULT 0,
  CONSTRAINT CHK_TipoRolAdminNoAdmin CHECK (rolBase='administrador' OR (adminPTests=0 AND
    adminETests=0 AND adminEQuestions=0 AND adminAnswers=0 AND adminUsers=0 AND adminTeams=0 AND
    adminConfiguration=0 AND adminPermissions=0)),
  CONSTRAINT CHK_PositivePrioridad CHECK (prioridad >=0),
  UNIQUE(nombre),
  PRIMARY KEY (id)
);

CREATE TABLE Usuario (
  id int(11) NOT NULL AUTO_INCREMENT,
  username varchar(40) COLLATE utf8_unicode_ci NOT NULL,
  email varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  pwhash longtext COLLATE utf8_unicode_ci NOT NULL,
  tipoRolId int(11) NOT NULL,
  fullname longtext COLLATE utf8_unicode_ci,
  FOREIGN KEY(tipoRolId) REFERENCES TipoRol(id),
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
  title varchar(100) COLLATE utf8_unicode_ci NOT NULL,
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
  visibilidad ENUM('alEntregar', 'alCorregir', 'manual') NOT NULL,
  cantidadFavoritos int(11) DEFAULT 0,
  tiempoEstricto boolean NOT NULL,
  maxIntentos int(11) NOT NULL, /*if <1, sin Limite*/
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
  FOREIGN KEY(usuarioid) REFERENCES Usuario(id),
  FOREIGN KEY(testid) REFERENCES Test(id),
  CONSTRAINT PRIMARY KEY(usuarioid, testid)
);

DELIMITER //
/* ON DELETE CASCADE de FOREIGN KEY Test */
CREATE TRIGGER testfavorito_testdelete AFTER DELETE ON Test
  FOR EACH ROW BEGIN
    DELETE FROM TestFavorito WHERE testid = OLD.id;
  END //
/* ON DELETE CASCADE de FOREIGN KEY Usuario */
CREATE TRIGGER testfavorito_usuariodelete AFTER DELETE ON Usuario
  FOR EACH ROW BEGIN
    DELETE FROM TestFavorito WHERE usuarioid = OLD.id;
  END //

/* Modifica cantidad favoritos al añadir y quitar*/
CREATE TRIGGER testfavorito_added AFTER INSERT ON TestFavorito
  FOR EACH ROW BEGIN
    UPDATE Test SET cantidadFavoritos=cantidadFavoritos+1 WHERE id=NEW.testid;
  END //

CREATE TRIGGER testfavorito_removed AFTER DELETE ON TestFavorito
  FOR EACH ROW BEGIN
    UPDATE Test SET cantidadFavoritos=cantidadFavoritos-1 WHERE id=OLD.testid;
  END //

DELIMITER ;


CREATE TABLE Pregunta(
  id int(11) NOT NULL AUTO_INCREMENT,
  title varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  question longtext COLLATE utf8_unicode_ci NOT NULL,
  estimatedTime int(11),
  autoCorrect boolean NOT NULL,
  editable boolean NOT NULL,
  usuarioid int(11) NOT NULL,
  eleccionUnica boolean,
  solucion longtext COLLATE utf8_unicode_ci,
  accesoPublicoNoPublicada boolean NOT NULL,
  penalizacion int(11) NOT NULL,
  cantidadFavoritos int(11) DEFAULT 0,
  CONSTRAINT CHK_penalizacion CHECK (penalizacion>=0 AND penalizacion<=100),
  FOREIGN KEY(usuarioid) REFERENCES Usuario(id) ON DELETE CASCADE,
  PRIMARY KEY(id)
);
CREATE TABLE PreguntaFavorita(
  usuarioid int(11) NOT NULL,
  preguntaid int(11) NOT NULL,
  FOREIGN KEY(usuarioid) REFERENCES Usuario(id),
  FOREIGN KEY(preguntaid) REFERENCES Pregunta(id),
  CONSTRAINT PRIMARY KEY(usuarioid, preguntaid)
);

DELIMITER //
/* ON DELETE CASCADE de FOREIGN KEY Pregunta */
CREATE TRIGGER preguntafavorita_preguntadelete AFTER DELETE ON Pregunta
  FOR EACH ROW BEGIN
    DELETE FROM PreguntaFavorita WHERE preguntaid = OLD.id;
  END //
/* ON DELETE CASCADE de FOREIGN KEY Usuario */
CREATE TRIGGER preguntafavorita_usuariodelete AFTER DELETE ON Usuario
  FOR EACH ROW BEGIN
    DELETE FROM PreguntaFavorita WHERE usuarioid = OLD.id;
  END //

/* Modifica cantidad favoritos al añadir y quitar*/
CREATE TRIGGER preguntadfavorita_added AFTER INSERT ON PreguntaFavorita
  FOR EACH ROW BEGIN
    UPDATE Pregunta SET cantidadFavoritos=cantidadFavoritos+1 WHERE id=NEW.preguntaid;
  END //

CREATE TRIGGER preguntadfavorita_removed AFTER DELETE ON PreguntaFavorita
  FOR EACH ROW BEGIN
    UPDATE Pregunta SET cantidadFavoritos=cantidadFavoritos-1 WHERE id=OLD.preguntaid;
  END //

DELIMITER ;

CREATE TABLE TestPregunta(
  testid int(11) NOT NULL,
  preguntaid int(11) NOT NULL,
  valorFinal int(11) NOT NULL,
  posicion int(11) NOT NULL,
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
  visibleParaUsuario boolean NOT NULL,
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
  compila boolean,
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

CREATE TABLE Prueba(
  id int(11) NOT NULL,
  preguntaid int(11) NOT NULL,
  entrada longtext COLLATE utf8_unicode_ci NOT NULL,
  salida longtext COLLATE utf8_unicode_ci NOT NULL,
  contenidoVisible boolean NOT NULL,
  resultadoVisible boolean NOT NULL,
  FOREIGN KEY(preguntaid) REFERENCES Pregunta(id) ON DELETE CASCADE,
  PRIMARY KEY(id)
);

CREATE TABLE Ejecucion(
  pruebaid int(11) NOT NULL,
  respuestaExamenid int(11) NOT NULL,
  preguntaid int(11) NOT NULL,
  correcta boolean NOT NULL,
  error ENUM('tiempoExcedido', 'errorEjecucion', 'salidaIncorrecta'),
  FOREIGN KEY(pruebaid) REFERENCES Prueba(id) ON DELETE CASCADE,
  CONSTRAINT fk_EjRespPreg
    FOREIGN KEY(respuestaExamenid, preguntaid)
    REFERENCES RespuestaPregunta(respuestaExamenid, preguntaid) ON DELETE CASCADE,
);

/* DATOS INICIALES */

INSERT INTO TipoRol(id, rolBase, nombre, prioridad, verPTests, verETests, verEQuestions, verPQuestions,
    verAnswers, changeRoles, tenerTeams, tenerEQuestions, tenerETests, tenerPTests, adminPTests, 
    adminETests, adminEQuestions, adminAnswers, adminUsers, adminTeams, adminConfiguration, adminPermissions,
    tipoInicial) VALUES(1, 'administrador', 'administrador', 0, 1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,0);

INSERT INTO TipoRol(rolBase, nombre, prioridad, verPTests, verETests, verEQuestions, verPQuestions,
    verAnswers, tenerTeams, tenerEQuestions, tenerETests, tenerPTests) VALUES('profesor', 'profesor', 1, 1,1,1,1,1,1,1,1,1);

INSERT INTO TipoRol(rolBase, nombre, prioridad, verPTests, tipoInicial) VALUES('estudiante', 'estudiante', 10, 1, 1);

INSERT INTO TipoRol(rolBase, nombre, prioridad, verPTests) VALUES('noRegistrado', 'noRegistrado', 100, 1);

/* admin pass = admin1 */
INSERT INTO Usuario(username, email, pwhash, tipoRolId, fullname) VALUES('admin', 'admin@mail.com', '$2a$14$C0gTluZGQVbau5vcsaB72e0iwiECRIJvCgwNk4cn7IFlEJEMFwuVC', 1, 'admin');