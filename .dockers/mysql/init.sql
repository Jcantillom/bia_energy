SET @MYSQLDUMP_TEMP_LOG_BIN = @@SESSION.SQL_LOG_BIN;
SET @@SESSION.SQL_LOG_BIN= 0;

SET @@GLOBAL.GTID_PURGED='';

CREATE DATABASE IF NOT EXISTS `bia_energy` /*!40100 DEFAULT CHARACTER SET utf8 */;