CREATE TABLE IF NOT EXISTS shows (
    id UUID PRIMARY KEY NOT NULL,
    movie_id UUID NOT NULL,
    hall_id UUID NOT NULL,
    start_time TIMESTAMP NOT NULL,
    buffer_minutes INT DEFAULT 0, 
    FOREIGN KEY (movie_id) REFERENCES movies(id),
    FOREIGN KEY (hall_id) REFERENCES halls(id)
)