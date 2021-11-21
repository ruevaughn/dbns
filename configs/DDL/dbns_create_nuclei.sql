CREATE TABLE IF NOT EXISTS nuclei (
    id serial,
    templateID VARCHAR (512),
    host VARCHAR(512),
    severity VARCHAR(64),
    name TEXT,
    tags VARCHAR(256),
    matcher_name VARCHAR(512),
    type VARCHAR(64),
    matched VARCHAR(512),
    ip VARCHAR(16),
    first_discovered TIMESTAMP default now(),
    last_change TIMESTAMP DEFAULT now(),
    PRIMARY KEY(host, templateID)
)