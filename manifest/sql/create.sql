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

-- 用户表，支持邮箱登录
CREATE TABLE user (
  id INT(11) NOT NULL AUTO_INCREMENT,
  passport VARCHAR(64) NOT NULL,
  password VARCHAR(255) NOT NULL,
  nickname VARCHAR(64) NOT NULL,
  email VARCHAR(255) NOT NULL,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  UNIQUE KEY uk_passport (passport),
  UNIQUE KEY uk_email (email)
) ENGINE=InnoDB;

-- 邮件验证码表
CREATE TABLE email_code (
  id INT(11) NOT NULL AUTO_INCREMENT,
  email VARCHAR(255) NOT NULL,
  code VARCHAR(8) NOT NULL,
  purpose VARCHAR(32) NOT NULL,
  used TINYINT(1) NOT NULL DEFAULT 0,
  expires_at DATETIME NOT NULL,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  INDEX idx_email_purpose (email, purpose)
) ENGINE=InnoDB;

-- 物品表（RESTful 示例）
CREATE TABLE item (
  id INT(11) NOT NULL AUTO_INCREMENT,
  name VARCHAR(128) NOT NULL,
  description TEXT,
  price DECIMAL(10,2) NOT NULL DEFAULT 0,
  stock INT NOT NULL DEFAULT 0,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
) ENGINE=InnoDB;