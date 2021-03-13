CREATE DATABASE IF NOT EXISTS uva_devtest;
USE uva_devtest;

DROP TABLE IF EXISTS Teamroles;
DROP TABLE IF EXISTS Teams;
DROP TABLE IF EXISTS Users;
CREATE TABLE Users (
  id int(11) NOT NULL AUTO_INCREMENT,
  username varchar(40) COLLATE utf8_unicode_ci NOT NULL,
  email varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  pwhash longtext COLLATE utf8_unicode_ci NOT NULL,
  type ENUM('admin', 'teacher', 'student') NOT NULL,
  fullname varchar(200) COLLATE utf8_unicode_ci,
  UNIQUE(username),
  UNIQUE(email),
  PRIMARY KEY (id)
);

CREATE TABLE Teams (
  id int(11) NOT NULL AUTO_INCREMENT,
  teamname varchar(40) COLLATE utf8_unicode_ci NOT NULL,
  description longtext COLLATE utf8_unicode_ci NOT NULL,
  UNIQUE(teamname),
  PRIMARY KEY(id)
);

CREATE TABLE Teamroles(
  userid int(11) NOT NULL,
  teamid int(11) NOT NULL,
  role ENUM('admin', 'member') NOT NULL,
  FOREIGN KEY(userid) REFERENCES Users(id) ON DELETE CASCADE,
  FOREIGN KEY(teamid) REFERENCES Teams(id) ON DELETE CASCADE,
  CONSTRAINT CompKey_ID_NAME PRIMARY KEY(userid, teamid)
);
/* admin pass = admin1 */
INSERT INTO Users(username, email, pwhash, type, fullname) VALUES('admin', 'admin@mail.com', '$2a$14$C0gTluZGQVbau5vcsaB72e0iwiECRIJvCgwNk4cn7IFlEJEMFwuVC', 'admin', 'admin')