-- +goose Up
CREATE TABLE weather_data (
    id BIGSERIAL PRIMARY KEY,
    houston_temperature FLOAT NOT NULL,
    nyc_temperature FLOAT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE weather_data;