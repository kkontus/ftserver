DROP DATABASE IF EXISTS quiz;

CREATE DATABASE quiz;

-- select database with USE quiz;

DROP TABLE IF EXISTS questions;

CREATE TABLE questions
  (
     id       INT(6) UNSIGNED auto_increment PRIMARY KEY,
     question TEXT NOT NULL,
     correct  VARCHAR(255) NOT NULL,
     answers  TEXT NOT NULL
  );

INSERT INTO questions (question, correct, answers) VALUES ('What is the native name for Croatia?', 'Hrvatska', 'Slovenija, Croata, Hrvatska, Yugoslavia');
INSERT INTO questions (question, correct, answers) VALUES ('On which sea is Croatia located?', 'Adriatic', 'Adriatic, Baltic, Mediterranean, North');
INSERT INTO questions (question, correct, answers) VALUES ('Which Croatian city was formerly known as Ragusa?', 'Dubrovnik', 'Split, Zadar, Sibenik, Dubrovnik');
INSERT INTO questions (question, correct, answers) VALUES ('Croatia\'s biggest peninsula is what?', 'Istria', 'Istria, Peljesac, Pag, Korcula');
INSERT INTO questions (question, correct, answers) VALUES ('From which century have Croats lived in this place?', '7th', '5th, 6th, 7th, 8th');
INSERT INTO questions (question, correct, answers) VALUES ('What is the official currency in Croatia?', 'Kuna', 'Euro, Kuna, Dinar, Lira');
INSERT INTO questions (question, correct, answers) VALUES ('In which city in Croatia is there an ancient amphitheater?', 'Pula', 'Rijeka, Vodnjan, Pula, Umag');
INSERT INTO questions (question, correct, answers) VALUES ('Croatia had to get its freedom by war. When did that war begin?', '1991', '1990, 1991, 1992, 1995');
INSERT INTO questions (question, correct, answers) VALUES ('Who was the first King of Croatia?', 'Tomislav', 'Tomislav, Borna, Zvonimir, Trpimir');
INSERT INTO questions (question, correct, answers) VALUES ('What is Croatian for the Adriatic Sea?', 'Jadransko', 'Sjeverno, Balticko, Sredozemno, Jadransko');