CREATE TABLE roles
(
    id    INT AUTO_INCREMENT PRIMARY KEY,
    alias VARCHAR(255) NOT NULL
);

CREATE TABLE users
(
    id         INT AUTO_INCREMENT PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    email               VARCHAR(255) UNIQUE,
    password            VARCHAR(255) NOT NULL,
    role_id    INT          NOT NULL,
    deleted_at DATETIME DEFAULT NULL,
    FOREIGN KEY (role_id) REFERENCES roles (id)
        ON UPDATE CASCADE
);

CREATE TABLE tasks
(
    id         INT AUTO_INCREMENT PRIMARY KEY,
    user_id    INT           NOT NULL,
    summary    VARCHAR(2500) NOT NULL,
    created_at DATETIME      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    updated_by INT           NOT NULL,
    deleted_at DATETIME               DEFAULT NULL,
    deleted_by INT                    DEFAULT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id)
        ON DELETE CASCADE
        ON UPDATE CASCADE,
    FOREIGN KEY (updated_by) REFERENCES users (id)
        ON DELETE RESTRICT
        ON UPDATE CASCADE,
    FOREIGN KEY (deleted_by) REFERENCES users (id)
        ON DELETE SET NULL
        ON UPDATE CASCADE
);
