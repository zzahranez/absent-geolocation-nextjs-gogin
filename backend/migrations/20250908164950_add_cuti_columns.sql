-- +goose Up
-- +goose StatementBegin

CREATE TABLE status_cuti (
    id INT AUTO_INCREMENT PRIMARY KEY,
    status_name VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +goose StatementEnd
-- +goose StatementBegin

INSERT INTO status_cuti (status_name) VALUES
('Pending'),
('Approved'),
('Rejected');

-- +goose StatementEnd
-- +goose StatementBegin

CREATE TABLE cuti (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    submission_date DATE NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    information TEXT,
    status_cuti_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_status FOREIGN KEY (status_cuti_id) REFERENCES status_cuti(id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS cuti;

-- +goose StatementEnd
-- +goose StatementBegin

DROP TABLE IF EXISTS status_cuti;

-- +goose StatementEnd
