CREATE TABLE customer (
  uid INT(11) NOT NULL AUTO_INCREMENT,
  name VARCHAR(255) NOT NULL,
  password VARCHAR(255) NOT NULL,
  phone VARCHAR(20),
  address TEXT,
  PRIMARY KEY (uid)
) ENGINE=InnoDB;

CREATE TABLE merchant (
  merchant_id INT(11) NOT NULL AUTO_INCREMENT,
  name VARCHAR(255) NOT NULL,
  contact VARCHAR(20),
  address TEXT,
  PRIMARY KEY (merchant_id)
) ENGINE=InnoDB;

CREATE TABLE `order` (
  order_id INT(11) NOT NULL AUTO_INCREMENT,
  uid INT(11) NOT NULL,
  merchant_id INT(11) NOT NULL,
  status ENUM('pending', 'shipped', 'delivered', 'cancelled') NOT NULL DEFAULT 'pending',
  order_time DATETIME NOT NULL,
  delivery_address TEXT,
  PRIMARY KEY (order_id)
) ENGINE=InnoDB;