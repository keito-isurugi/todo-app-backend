TRUNCATE TABLE todos RESTART IDENTITY CASCADE;
INSERT INTO todos (title, done_flag)
VALUES ('テストToDo1', 'false'),
       ('テストToDo2', 'false'),
       ('テストToDo3', 'false'),
       ('テストToDo4', 'false');