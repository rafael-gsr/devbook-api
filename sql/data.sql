INSERT INTO users (name, nick, email, password) 
VALUES 
  ("USER 1", "USER_1_NICK", "USER_1@EMAIL.COM", "password"),
  ("USER 2", "USER_2_NICK", "USER_2@EMAIL.COM", "password"),
  ("USER 3", "USER_3_NICK", "USER_3@EMAIL.COM", "password")
;


INSERT INTO followers (user_id, follower_id)
VALUES
  (1 ,2),
  (3 ,1),
  (1 ,3)
;
