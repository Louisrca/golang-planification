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

INSERT INTO `admin` (`id`, `firstname`, `lastname`, `email`, `password`) VALUES
    ('197f586a-6402-4590-87c3-ceacd4558b22', 'Alex', 'Smith', 'alex.smith@example.com', '123');

INSERT INTO `customer` (`id`, `firstname`, `lastname`, `email`, `password`) VALUES
    ('6c4d2f89-e030-4e27-a184-2f0f0f2c33b5', 'Jamie', 'Doe', 'jamie.doe@example.com', '123');

INSERT INTO `hair_salon` (`id`, `name`, `address`, `description`, `is_accepted`) VALUES
    ('5f020c72-05b8-4777-85c6-a170a9faefed', 'Elegant Cuts', '123 Fashion St.', 'A trendy salon for modern styles.', TRUE);

INSERT INTO `hairdresser` (`id`, `firstname`, `lastname`, `email`, `password`, `start_time`, `end_time`, `hair_salon_id`) VALUES
    ('f283a4cc-65b8-4aa5-9a89-487636c20f7e', 'Chris', 'Johnson', 'chris.johnson@example.com', '123', '09:00:00', '17:00:00', '5f020c72-05b8-4777-85c6-a170a9faefed');

INSERT INTO `category` (`id`, `name`) VALUES
    ('c4ce1ccc-1db0-48da-8e95-50fd9c0185cd', 'Standard Cut');

INSERT INTO `service` (`id`, `name`, `price`, `duration`, `category_id`, `hair_salon_id`) VALUES
    ('7eea3cad-cd18-4547-9d9b-22450ede3b2a', 'Men’s Haircut', 30, 30, 'c4ce1ccc-1db0-48da-8e95-50fd9c0185cd', '5f020c72-05b8-4777-85c6-a170a9faefed');

INSERT INTO `slot` (`id`, `start_time`, `end_time`, `is_booked`, `hairdresser_id`) VALUES
    ('eacb1e57-99c4-4472-bee1-2908f2b6620c', '2024-01-20 10:00:00', '2024-01-20 10:30:00', FALSE, 'f283a4cc-65b8-4aa5-9a89-487636c20f7e');

INSERT INTO `booking` (`id`, `is_confirmed`, `customer_id`, `service_id`, `slot_id`) VALUES
    ('4ce15499-c7d9-4381-8ed2-dad9b8fe93c5', TRUE, '6c4d2f89-e030-4e27-a184-2f0f0f2c33b5', '7eea3cad-cd18-4547-9d9b-22450ede3b2a', 'eacb1e57-99c4-4472-bee1-2908f2b6620c');
