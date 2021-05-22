insert into users (name, nick, email, password)
values 
("User 1", "user_1", "user1@gmail.com", "$2a$10$65Orz2VA/OV.yMsbm/V8WOsVl5vKmDOkIZUU2UFAm8i1Qg2tqtabu"),
("User 2", "user_2", "user2@gmail.com", "$2a$10$65Orz2VA/OV.yMsbm/V8WOsVl5vKmDOkIZUU2UFAm8i1Qg2tqtabu"),
("User 3", "user_3", "user3@gmail.com", "$2a$10$65Orz2VA/OV.yMsbm/V8WOsVl5vKmDOkIZUU2UFAm8i1Qg2tqtabu");

insert into followers (user_id, follower_id)
values
(1,2),
(3,1),
(1,3);