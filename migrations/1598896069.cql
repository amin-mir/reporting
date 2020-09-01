CREATE TABLE IF NOT EXISTS reports (
    report_id uuid,
    user_id uuid,
    resolver_ids text,
    reviewer_ids list<text>,
    created_at timestamp,
    updated_at timestamp,
    status text,
    title text,
    PRIMARY KEY (report_id, created_at)
);

CREATE TABLE IF NOT EXISTS reports_by_user_id (
    report_id uuid,
    user_id uuid,
    resolver_ids text,
    reviewer_ids list<text>,
    created_at timestamp,
    updated_at timestamp,
    status text,
    title text,
    PRIMARY KEY (user_id, created_at, report_id)
) WITH CLUSTERING ORDER BY (created_at DESC);

CREATE TABLE IF NOT EXISTS reports_by_status (
    report_id uuid,
    user_id uuid,
    resolver_ids text,
    reviewer_ids list<text>,
    created_at timestamp,
    updated_at timestamp,
    status text,
    title text,
    PRIMARY KEY (status, created_at, report_id)
);

CREATE TABLE IF NOT EXISTS reports_messages (
    report_id uuid,
    message_id uuid,
    body text,
    created_at timestamp,
    PRIMARY KEY (report_id, message_id)
)