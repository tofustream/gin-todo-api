CREATE TABLE IF NOT EXISTS tasks (
    id UUID PRIMARY KEY,                     -- タスクの一意識別子 (UUID)
    description TEXT NOT NULL,              -- タスクの説明
    created_at TIMESTAMP NOT NULL,          -- 作成日時
    updated_at TIMESTAMP NOT NULL,          -- 更新日時
    is_completed BOOLEAN NOT NULL DEFAULT FALSE, -- タスクの完了状態
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE    -- タスクの削除状態
);
