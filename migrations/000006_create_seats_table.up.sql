CREATE TABLE IF NOT EXISTS seats (
    id UUID PRIMARY KEY NOT NULL,
    hall_id UUID NOT NULL,
    seat_row INT NOT NULL,
    seat_number INT NOT NULL,
    FOREIGN KEY (hall_id) REFERENCES halls(id),
    UNIQUE (hall_id, seat_row, seat_number)
);