-- Create users table
CREATE TABLE users (
    id VARCHAR(36) PRIMARY KEY COMMENT 'UUID format user identifier',
    line_user_id VARCHAR(255) NOT NULL UNIQUE COMMENT 'LINE user ID',
    display_name VARCHAR(255) NOT NULL COMMENT 'User display name',
    wallet_address VARCHAR(255) NULL COMMENT 'Blockchain wallet address',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Record creation timestamp',
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Record update timestamp',
    INDEX idx_line_user_id (line_user_id),
    INDEX idx_wallet_address (wallet_address)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='User account information';
