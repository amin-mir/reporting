CREATE TABLE IF NOT EXISTS reports (
    report_id uuid,
    user_id uuid,
    resolver_id text,
    reviewer_ids list<text>,
    created_at timestamp,
    updated_at timestamp,
    status text,
    title text,
    PRIMARY KEY (report_id, created_at)
);

CREATE MATERIALIZED VIEW IF NOT EXISTS reports_by_user_id AS 
    SELECT * FROM reports
    WHERE user_id IS NOT NULL
    AND created_at IS NOT NULL
    AND report_id IS NOT NULL
    PRIMARY KEY (user_id, created_at, report_id);

CREATE MATERIALIZED VIEW IF NOT EXISTS reports_by_status AS
    SELECT * FROM reports
    WHERE status IS NOT NULL
    AND created_at IS NOT NULL
    AND report_id IS NOT NULL
    PRIMARY KEY (status, created_at, report_id);

CREATE TABLE IF NOT EXISTS reports_messages (
    report_id uuid,
    message_id uuid,
    user_id uuid,
    body text,
    created_at timestamp,
    PRIMARY KEY (report_id, message_id, created_at)
);