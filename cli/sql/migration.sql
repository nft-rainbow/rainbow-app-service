-- migrate Activity
Use rainbowApp

-- TODO:
-- create table activities
-- insert activites from poap_activity_configs
ALTER TABLE poap_activity_configs RENAME activities;
ALTER TABLE activities RENAME COLUMN contract_id TO contract_raw_id;

-- TODO:
-- create table activity_contracts
-- insert contract form poap_activity_configs
CREATE TABLE activity_contracts SELECT id,created_at,updated_at,deleted_at, contract_raw_id, contract_address, contract_type, chain_id, chain_type FROM activities;