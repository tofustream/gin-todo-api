CREATE TABLE IF NOT EXISTS tasks (
    id UUID PRIMARY KEY,                     -- タスクの一意識別子 (UUID)
    description TEXT NOT NULL,               -- タスクの説明
    created_at TIMESTAMP NOT NULL,           -- 作成日時
    updated_at TIMESTAMP NOT NULL,           -- 更新日時
    is_completed BOOLEAN NOT NULL DEFAULT FALSE, -- タスクの完了状態
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,   -- タスクの削除状態
    user_id UUID NOT NULL,                   -- ユーザーの一意識別子 (UUID)
    CONSTRAINT fk_user
        FOREIGN KEY(user_id)
        REFERENCES general_users(id)
        ON DELETE CASCADE
);
