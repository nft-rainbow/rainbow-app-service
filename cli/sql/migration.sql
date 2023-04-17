# migrate Activity
Use rainbowApp
ALTER TABLE poap_activity_configs RENAME activities;
ALTER TABLE activities RENAME COLUMN contract_id TO contract_code;
CREATE TABLE activity_contracts SELECT id,created_at,updated_at,deleted_at, contract_id as contract_raw_id, contract_address, contract_type, chain_id, chain_type FROM activities;