DROP TABLE IF EXISTS `incidents`;
CREATE TABLE `incidents`
(
  `id` int(6) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(30) ,
  `description` varchar(512) ,
  `acknowlege` enum('yes','no') DEFAULT 'no' ,
  `status`  enum('closed','open','inprogress') DEFAULT 'open',
  `priority` int(10) default 3,
  `severity` enum('low','medium','high') DEFAULT 'low',
  `reportedBy` varchar(30) ,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY(`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;
