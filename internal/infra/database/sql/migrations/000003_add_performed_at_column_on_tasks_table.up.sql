ALTER TABLE tasks
    ADD COLUMN performed_at DATETIME DEFAULT NULL after is_done;