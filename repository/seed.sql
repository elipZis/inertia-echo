/* Seed a test user (johndoe@example.com/secret) */
INSERT INTO users("email", "password", "first_name", "last_name") VALUES
('johndoe@example.com', '$2y$10$i1jALNbm199VCCNpA0hNSe/3e34WOPdJy1yk7X.GrNAaJqGaUFgMK', "John", "Doe");