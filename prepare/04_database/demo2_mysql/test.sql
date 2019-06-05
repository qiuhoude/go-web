CREATE TABLE `userinfo`
(
  `uid`        INT(10) NOT NULL AUTO_INCREMENT,
  `username`   VARCHAR(64) NULL DEFAULT NULL,
  `departname` VARCHAR(64) NULL DEFAULT NULL,
  `created`    DATE NULL DEFAULT NULL,
  PRIMARY KEY (`uid`)
);

CREATE TABLE `userdetail`
(
  `uid`     INT(10) NOT NULL DEFAULT '0',
  `intro`   TEXT NULL,
  `profile` TEXT NULL,
  PRIMARY KEY (`uid`)
) COMMIT "用户详细";


CREATE TABLE `account`
(
  `id`    int(4) NOT NULL,
  `name`  varchar(30) DEFAULT NULL,
  `money` float(8,
  2
) DEFAULT NULL,
  PRIMARY KEY
(
  `id`
)
  );

INSERT
INTO
  `account`
VALUES
('1',
 'ruby',
 '3000.00');
INSERT INTO `account`

VALUES ('2', '王二狗', '1000.00');