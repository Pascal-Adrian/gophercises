-- +goose up

INSERT INTO paths (path, url) VALUES
    ('/gophercises', 'https://gophercises.com/'),
    ('/quiz', 'https://github.com/gophercises/quiz'),
    ('/sqlc-docs', 'https://docs.sqlc.dev/en/latest/');


