CREATE TABLE IF NOT EXISTS `admin` (
    `id` VARCHAR(36) PRIMARY KEY,
    `firstname` VARCHAR(255),
    `lastname` VARCHAR(255),
    `email` VARCHAR(255),
    `password` VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS `customer` (
    `id` VARCHAR(36) PRIMARY KEY,
    `firstname` VARCHAR(255),
    `lastname` VARCHAR(255),
    `email` VARCHAR(255),
    `password` VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS `hair_salon` (
    `id` VARCHAR(36) PRIMARY KEY,
    `name` VARCHAR(255),
    `address` VARCHAR(255),
    `description` VARCHAR(255),
    `is_accepted` BOOLEAN
);

CREATE TABLE IF NOT EXISTS `hairdresser` (
    `id` VARCHAR(36) PRIMARY KEY,
    `firstname` VARCHAR(255),
    `lastname` VARCHAR(255),
    `email` VARCHAR(255),
    `password` VARCHAR(255),
    `start_time` TIME,
    `end_time` TIME,
    `hair_salon_id` VARCHAR(36),
    CONSTRAINT fk_hairdresser_hair_salon FOREIGN KEY (hair_salon_id) REFERENCES hair_salon(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS `category` (
    `id` VARCHAR(36) PRIMARY KEY,
    `name` VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS `service` (
    `id` VARCHAR(36) PRIMARY KEY,
    `name` VARCHAR(255),
    `price` INT,
    `duration` INT,
    `category_id` VARCHAR(36),
    `hair_salon_id` VARCHAR(36),
    CONSTRAINT fk_service_category FOREIGN KEY (category_id) REFERENCES category(id) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT fk_service_hair_salon FOREIGN KEY (hair_salon_id) REFERENCES hair_salon(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS `slot` (
    `id` VARCHAR(36) PRIMARY KEY,
    `start_time` DATETIME,
    `end_time` DATETIME,
    `is_booked` BOOLEAN,
    `hairdresser_id` VARCHAR(36),
    CONSTRAINT fk_slot_hairdresser FOREIGN KEY (hairdresser_id) REFERENCES hairdresser(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS `booking` (
    `id` VARCHAR(36) PRIMARY KEY,
    `is_confirmed` BOOLEAN,
    `customer_id` VARCHAR(36),
    `service_id` VARCHAR(36),
    `slot_id` VARCHAR(36),
    CONSTRAINT fk_booking_customer FOREIGN KEY (customer_id) REFERENCES customer(id) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT fk_booking_service FOREIGN KEY (service_id) REFERENCES service(id) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT fk_booking_slot FOREIGN KEY (slot_id) REFERENCES slot(id) ON DELETE CASCADE ON UPDATE CASCADE
);