CREATE TABLE trx_parking (
    plat_no VARCHAR(30) NOT NULL,
    slot_number INT NOT NULL,
    reg_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE hst_parking (
    hst_id INT AUTO_INCREMENT PRIMARY KEY,
    plat_no VARCHAR(30) NOT NULL,
    slot_number INT NOT NULL,
    reg_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE map_parking_lot (
    id INT AUTO_INCREMENT PRIMARY KEY,
    plat_no VARCHAR(10) DEFAULT ''
);


-- INSERT 10 ROWS in map_parking_lot table
INSERT INTO map_parking_lot (plat_no)
VALUES ('1234');
INSERT INTO map_parking_lot (plat_no)
VALUES ('');