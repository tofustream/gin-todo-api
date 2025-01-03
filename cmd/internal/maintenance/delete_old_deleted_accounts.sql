-- 論理削除されているデータを、定期的に物理削除するためのSQL
DELETE FROM tasks
WHERE account_id IN (
    SELECT id FROM accounts
    WHERE is_deleted = TRUE
    AND created_at < NOW() - INTERVAL '1 days'
);

DELETE FROM accounts
WHERE is_deleted = TRUE
AND created_at < NOW() - INTERVAL '1 days';
