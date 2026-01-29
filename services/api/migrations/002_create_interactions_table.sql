-- Create interactions table
CREATE TABLE interactions (
    id VARCHAR(36) PRIMARY KEY COMMENT 'UUID format interaction identifier',
    requester_id VARCHAR(36) NOT NULL COMMENT 'User ID who initiated the interaction',
    approver_id VARCHAR(36) NOT NULL COMMENT 'User ID who approves the interaction',
    status ENUM('pending', 'approved', 'rejected') NOT NULL DEFAULT 'pending' COMMENT 'Current interaction status',
    metadata JSON NULL COMMENT 'Additional metadata for the interaction',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Record creation timestamp',
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Record update timestamp',
    INDEX idx_requester_id (requester_id),
    INDEX idx_approver_id (approver_id),
    INDEX idx_status (status),
    INDEX idx_created_at (created_at),
    FOREIGN KEY (requester_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (approver_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='User interaction records';
