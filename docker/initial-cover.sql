create table local_db.cover
(
    id   int auto_increment
        primary key,
    name varchar(100) not null
);

INSERT INTO local_db.cover (id, name) VALUES (1, 'cover-lion');
INSERT INTO local_db.cover (id, name) VALUES (2, 'cover-elephant');
INSERT INTO local_db.cover (id, name) VALUES (3, 'cover-cheetah');
INSERT INTO local_db.cover (id, name) VALUES (4, 'cover-zebra');
INSERT INTO local_db.cover (id, name) VALUES (5, 'cover-buffalo');
INSERT INTO local_db.cover (id, name) VALUES (6, 'cover-giraffe');
INSERT INTO local_db.cover (id, name) VALUES (7, 'cover-zebra');
INSERT INTO local_db.cover (id, name) VALUES (8, 'test');
INSERT INTO local_db.cover (id, name) VALUES (9, 'test non');
