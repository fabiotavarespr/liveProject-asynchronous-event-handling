-- Add UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Set timezone
-- For more information, please visit:
-- https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
SET TIMEZONE="UTC";

-- Create events schema
CREATE SCHEMA IF NOT EXISTS events;

-- Create events.processed_events table
CREATE TABLE IF NOT EXISTS events.processed_events (
    id uuid NOT NULL,
    processed_timestamp timestamp NOT NULL,
    event_name varchar(256) NOT NULL
);