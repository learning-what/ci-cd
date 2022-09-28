CREATE TABLE user_info
(
    id int,
    username varchar,
    password varchar,
    primary key (id)
);

-- Warning: Do not store the password without any encryption.
INSERT INTO user_info
VALUES (1, 'zhangsan', '123456'),
       (2, 'xiaoming', '456789'),
       (3, 'xiaowang', 'abcdef'),
       (4, 'xiaohong', 'abc123');