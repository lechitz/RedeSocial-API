insert into usuarios (nome, nick, email, senha)
values ("Usuario 1", "usuario_1", "usuario1@gmail.com", "$2a$10$n9YMR2uvOkyuaRwFBYzSQOCa2nKn1iKKxmfLuGa/GuR4haC73Phm6"),
       ("Usuario 2", "usuario_2", "usuario2@gmail.com", "$2a$10$n9YMR2uvOkyuaRwFBYzSQOCa2nKn1iKKxmfLuGa/GuR4haC73Phm6"),
       ("Usuario 3", "usuario_3", "usuario3@gmail.com", "$2a$10$n9YMR2uvOkyuaRwFBYzSQOCa2nKn1iKKxmfLuGa/GuR4haC73Phm6");

insert into seguidores(usuario_id, seguidor_id)
values (1, 2), -- Usuário1 é SEGUIDO pelo Usuário2
       (3, 1),
       (1, 3);